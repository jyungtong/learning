import { deleteTodo, updateTodo } from "@/lib/todos";

type TodoRouteContext = RouteContext<"/api/todos/[id]">;

export async function PATCH(request: Request, context: TodoRouteContext) {
  let body: unknown;
  try {
    body = await request.json();
  } catch {
    return Response.json({ message: "Invalid JSON body" }, { status: 400 });
  }

  const completed =
    typeof body === "object" && body !== null && "completed" in body
      ? body.completed
      : undefined;
  if (typeof completed !== "boolean") {
    return Response.json(
      { message: "Completed must be a boolean" },
      { status: 400 },
    );
  }

  const { id } = await context.params;
  const todo = updateTodo(id, completed);
  if (!todo) {
    return Response.json({ message: "Todo not found" }, { status: 404 });
  }
  return Response.json(todo);
}

export async function DELETE(_request: Request, context: TodoRouteContext) {
  const { id } = await context.params;
  if (!deleteTodo(id)) {
    return Response.json({ message: "Todo not found" }, { status: 404 });
  }
  return new Response(null, { status: 204 });
}
