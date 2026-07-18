import { createTodo, getTodos } from "@/lib/todos";

export const dynamic = "force-dynamic";

export function GET() {
  return Response.json(getTodos());
}

export async function POST(request: Request) {
  let body: unknown;
  try {
    body = await request.json();
  } catch {
    return Response.json({ message: "Invalid JSON body" }, { status: 400 });
  }

  const title =
    typeof body === "object" && body !== null && "title" in body
      ? body.title
      : undefined;
  if (typeof title !== "string" || !title.trim()) {
    return Response.json(
      { message: "Title must not be empty" },
      { status: 400 },
    );
  }

  return Response.json(createTodo(title.trim()), { status: 201 });
}
