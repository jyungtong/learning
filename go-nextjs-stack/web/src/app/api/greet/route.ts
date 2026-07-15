import { createClient } from "@connectrpc/connect";
import { createConnectTransport } from "@connectrpc/connect-node";

import { GreetService } from "@/gen/greet/v1/greet_pb";

export const runtime = "nodejs";

export async function POST(request: Request) {
  let body: unknown;

  try {
    body = await request.json();
  } catch {
    return Response.json({ error: "Name is required." }, { status: 400 });
  }

  if (
    typeof body !== "object" ||
    body === null ||
    !("name" in body) ||
    typeof body.name !== "string" ||
    body.name.trim() === ""
  ) {
    return Response.json({ error: "Name is required." }, { status: 400 });
  }

  const client = createClient(
    GreetService,
    createConnectTransport({
      baseUrl: process.env.GREET_SERVICE_URL ?? "http://localhost:8080",
      httpVersion: "1.1",
    }),
  );

  try {
    const response = await client.sayHello({ name: body.name.trim() });
    return Response.json({ message: response.message });
  } catch {
    return Response.json(
      { error: "Greeting service is unavailable." },
      { status: 502 },
    );
  }
}
