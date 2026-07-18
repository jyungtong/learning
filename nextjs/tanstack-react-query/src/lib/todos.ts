import "server-only";

export type Todo = {
  id: string;
  title: string;
  completed: boolean;
  createdAt: string;
};

type TodoStore = { todos: Todo[] };

const globalStore = globalThis as typeof globalThis & { todoStore?: TodoStore };
const store =
  globalStore.todoStore ??
  (globalStore.todoStore = {
    todos: [
      {
        id: crypto.randomUUID(),
        title: "Learn query invalidation",
        completed: true,
        createdAt: new Date(Date.now() - 60_000).toISOString(),
      },
      {
        id: crypto.randomUUID(),
        title: "Try an optimistic update",
        completed: false,
        createdAt: new Date().toISOString(),
      },
    ],
  });

export function getTodos() {
  return [...store.todos];
}

export function createTodo(title: string) {
  const todo: Todo = {
    id: crypto.randomUUID(),
    title,
    completed: false,
    createdAt: new Date().toISOString(),
  };
  store.todos.push(todo);
  return todo;
}

export function updateTodo(id: string, completed: boolean) {
  const todo = store.todos.find((item) => item.id === id);
  if (!todo) return null;
  todo.completed = completed;
  return { ...todo };
}

export function deleteTodo(id: string) {
  const index = store.todos.findIndex((item) => item.id === id);
  if (index === -1) return false;
  store.todos.splice(index, 1);
  return true;
}
