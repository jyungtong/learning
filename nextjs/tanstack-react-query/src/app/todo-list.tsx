"use client";

import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import {
  Check,
  CircleAlert,
  ListTodo,
  LoaderCircle,
  Plus,
  Trash2,
} from "lucide-react";
import { FormEvent, useState } from "react";

type Todo = {
  id: string;
  title: string;
  completed: boolean;
  createdAt: string;
};

type Filter = "all" | "active" | "completed";
const queryKey = ["todos"] as const;

async function request<T>(url: string, init?: RequestInit): Promise<T> {
  const response = await fetch(url, init);
  if (!response.ok) {
    const body = (await response.json().catch(() => null)) as {
      message?: string;
    } | null;
    throw new Error(body?.message ?? "Request failed");
  }
  return response.status === 204
    ? (undefined as T)
    : ((await response.json()) as T);
}

export function TodoList() {
  const queryClient = useQueryClient();
  const [title, setTitle] = useState("");
  const [filter, setFilter] = useState<Filter>("all");

  const todosQuery = useQuery({
    queryKey,
    queryFn: () => request<Todo[]>("/api/todos"),
  });

  const createMutation = useMutation({
    mutationFn: (newTitle: string) =>
      request<Todo>("/api/todos", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ title: newTitle }),
      }),
    onSuccess: () => {
      setTitle("");
      void queryClient.invalidateQueries({ queryKey });
    },
  });

  const toggleMutation = useMutation({
    mutationFn: ({ id, completed }: Pick<Todo, "id" | "completed">) =>
      request<Todo>(`/api/todos/${id}`, {
        method: "PATCH",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ completed }),
      }),
    onMutate: async ({ id, completed }) => {
      await queryClient.cancelQueries({ queryKey });
      const previous = queryClient.getQueryData<Todo[]>(queryKey);
      queryClient.setQueryData<Todo[]>(queryKey, (current = []) =>
        current.map((todo) =>
          todo.id === id ? { ...todo, completed } : todo,
        ),
      );
      return { previous };
    },
    onError: (_error, _variables, context) => {
      queryClient.setQueryData(queryKey, context?.previous);
    },
    onSettled: () => queryClient.invalidateQueries({ queryKey }),
  });

  const deleteMutation = useMutation({
    mutationFn: (id: string) =>
      request<void>(`/api/todos/${id}`, { method: "DELETE" }),
    onMutate: async (id) => {
      await queryClient.cancelQueries({ queryKey });
      const previous = queryClient.getQueryData<Todo[]>(queryKey);
      queryClient.setQueryData<Todo[]>(queryKey, (current = []) =>
        current.filter((todo) => todo.id !== id),
      );
      return { previous };
    },
    onError: (_error, _variables, context) => {
      queryClient.setQueryData(queryKey, context?.previous);
    },
    onSettled: () => queryClient.invalidateQueries({ queryKey }),
  });

  function handleSubmit(event: FormEvent<HTMLFormElement>) {
    event.preventDefault();
    const trimmedTitle = title.trim();
    if (trimmedTitle) createMutation.mutate(trimmedTitle);
  }

  const todos = todosQuery.data ?? [];
  const activeCount = todos.filter((todo) => !todo.completed).length;
  const completedCount = todos.length - activeCount;
  const visibleTodos = todos.filter((todo) => {
    if (filter === "active") return !todo.completed;
    if (filter === "completed") return todo.completed;
    return true;
  });
  const mutationError =
    createMutation.error ?? toggleMutation.error ?? deleteMutation.error;
  const listMutationPending =
    toggleMutation.isPending || deleteMutation.isPending;

  return (
    <main className="mx-auto flex min-h-screen w-full max-w-3xl flex-col px-5 py-10 sm:px-8 sm:py-16">
      <header className="mb-10 flex items-start justify-between gap-6">
        <div>
          <div className="mb-3 flex items-center gap-2 text-sm font-semibold text-emerald-700">
            <ListTodo aria-hidden="true" size={18} /> DAILY QUEUE
          </div>
          <h1 className="text-4xl font-bold text-zinc-950 sm:text-5xl">Tasks</h1>
          <p className="mt-3 text-zinc-600">
            {activeCount} active, {completedCount} completed
          </p>
        </div>
        {todosQuery.isFetching && !todosQuery.isPending ? (
          <LoaderCircle
            className="mt-1 animate-spin text-zinc-400"
            aria-label="Refreshing"
          />
        ) : null}
      </header>

      <form className="flex gap-2" onSubmit={handleSubmit}>
        <label className="sr-only" htmlFor="new-todo">
          New task
        </label>
        <input
          id="new-todo"
          className="min-w-0 flex-1 border border-zinc-300 bg-white px-4 py-3 text-base text-zinc-950 outline-none transition focus:border-emerald-600 focus:ring-2 focus:ring-emerald-600/20"
          value={title}
          onChange={(event) => setTitle(event.target.value)}
          placeholder="What needs doing?"
          maxLength={120}
          disabled={createMutation.isPending}
        />
        <button
          className="inline-flex h-12 w-12 shrink-0 items-center justify-center bg-zinc-950 text-white transition hover:bg-emerald-700 focus:outline-none focus:ring-2 focus:ring-emerald-600 focus:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50"
          type="submit"
          disabled={!title.trim() || createMutation.isPending}
          title="Add task"
        >
          {createMutation.isPending ? (
            <LoaderCircle
              className="animate-spin"
              aria-label="Adding task"
              size={20}
            />
          ) : (
            <Plus aria-hidden="true" size={21} />
          )}
        </button>
      </form>

      <div className="mt-8 flex border-b border-zinc-200" aria-label="Task filters">
        {(["all", "active", "completed"] as const).map((option) => (
          <button
            key={option}
            className={`relative px-4 py-3 text-sm font-semibold capitalize transition focus:outline-none focus:ring-2 focus:ring-inset focus:ring-emerald-600 ${
              filter === option
                ? "text-zinc-950 after:absolute after:inset-x-0 after:bottom-0 after:h-0.5 after:bg-emerald-600"
                : "text-zinc-500 hover:text-zinc-900"
            }`}
            type="button"
            onClick={() => setFilter(option)}
            aria-pressed={filter === option}
          >
            {option}
          </button>
        ))}
      </div>

      {mutationError ? (
        <div
          className="mt-5 flex items-center gap-2 border-l-2 border-red-600 bg-red-50 px-4 py-3 text-sm text-red-800"
          role="alert"
        >
          <CircleAlert aria-hidden="true" size={18} />
          {mutationError.message}
        </div>
      ) : null}

      {todosQuery.isPending ? (
        <div className="flex flex-1 items-center justify-center py-24 text-zinc-500">
          <LoaderCircle
            className="mr-2 animate-spin"
            aria-hidden="true"
            size={20}
          />
          Loading tasks
        </div>
      ) : todosQuery.isError ? (
        <div className="py-20 text-center">
          <CircleAlert
            className="mx-auto mb-3 text-red-600"
            aria-hidden="true"
          />
          <p className="font-semibold text-zinc-900">Could not load tasks</p>
          <button
            className="mt-4 text-sm font-semibold text-emerald-700 underline underline-offset-4"
            onClick={() => void todosQuery.refetch()}
            type="button"
          >
            Try again
          </button>
        </div>
      ) : visibleTodos.length === 0 ? (
        <div className="py-20 text-center text-zinc-500">
          <Check className="mx-auto mb-3 text-emerald-600" aria-hidden="true" />
          <p>{filter === "all" ? "No tasks yet" : `No ${filter} tasks`}</p>
        </div>
      ) : (
        <ul className="divide-y divide-zinc-200">
          {visibleTodos.map((todo) => (
            <li
              key={todo.id}
              className="group flex min-h-16 items-center gap-3 py-3"
            >
              <label className="flex min-w-0 flex-1 cursor-pointer items-center gap-3">
                <input
                  className="peer sr-only"
                  type="checkbox"
                  checked={todo.completed}
                  onChange={() =>
                    toggleMutation.mutate({
                      id: todo.id,
                      completed: !todo.completed,
                    })
                  }
                  disabled={listMutationPending}
                />
                <span className="flex h-6 w-6 shrink-0 items-center justify-center border-2 border-zinc-300 text-white transition peer-checked:border-emerald-600 peer-checked:bg-emerald-600 peer-focus-visible:ring-2 peer-focus-visible:ring-emerald-600 peer-focus-visible:ring-offset-2">
                  {todo.completed ? (
                    <Check aria-hidden="true" size={16} strokeWidth={3} />
                  ) : null}
                </span>
                <span
                  className={`truncate text-base ${
                    todo.completed
                      ? "text-zinc-400 line-through"
                      : "text-zinc-800"
                  }`}
                >
                  {todo.title}
                </span>
              </label>
              <button
                className="flex h-10 w-10 shrink-0 items-center justify-center text-zinc-400 transition hover:bg-red-50 hover:text-red-700 focus:outline-none focus:ring-2 focus:ring-red-600 disabled:cursor-not-allowed disabled:opacity-40 sm:opacity-0 sm:group-hover:opacity-100 sm:focus:opacity-100"
                type="button"
                onClick={() => deleteMutation.mutate(todo.id)}
                disabled={listMutationPending}
                title={`Delete ${todo.title}`}
              >
                <Trash2 aria-hidden="true" size={18} />
              </button>
            </li>
          ))}
        </ul>
      )}
    </main>
  );
}
