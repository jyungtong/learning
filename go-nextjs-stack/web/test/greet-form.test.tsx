import { afterEach, expect, mock, test } from "bun:test";
import { GlobalRegistrator } from "@happy-dom/global-registrator";

GlobalRegistrator.register();

const { cleanup, render, screen } = await import("@testing-library/react");
const { default: userEvent } = await import("@testing-library/user-event");
const { GreetForm } = await import("../src/components/greet-form");

const originalFetch = globalThis.fetch;

afterEach(() => {
  cleanup();
  globalThis.fetch = originalFetch;
});

test("submits a name and displays the greeting", async () => {
  const fetchMock = mock(() =>
    Promise.resolve(
      new Response(JSON.stringify({ message: "Hello, Ada!" }), { status: 200 }),
    ),
  );
  globalThis.fetch = fetchMock as typeof fetch;

  const user = userEvent.setup();
  render(<GreetForm />);

  await user.type(screen.getByLabelText("Name"), "Ada");
  await user.click(screen.getByRole("button", { name: "Send greeting" }));

  expect(await screen.findByText("Hello, Ada!")).toBeTruthy();
  expect(fetchMock).toHaveBeenCalledWith("/api/greet", {
    body: JSON.stringify({ name: "Ada" }),
    headers: { "Content-Type": "application/json" },
    method: "POST",
  });
});

test("displays an API error", async () => {
  const fetchMock = mock(() =>
    Promise.resolve(
      new Response(JSON.stringify({ error: "Name is required." }), { status: 400 }),
    ),
  );
  globalThis.fetch = fetchMock as typeof fetch;

  const user = userEvent.setup();
  render(<GreetForm />);

  await user.type(screen.getByLabelText("Name"), "Ada");
  await user.click(screen.getByRole("button", { name: "Send greeting" }));

  expect(await screen.findByText("Name is required.")).toBeTruthy();
});

test("displays a fallback error when the request cannot be completed", async () => {
  const fetchMock = mock(() => Promise.reject(new Error("Network error")));
  globalThis.fetch = fetchMock as typeof fetch;

  const user = userEvent.setup();
  render(<GreetForm />);

  await user.type(screen.getByLabelText("Name"), "Ada");
  await user.click(screen.getByRole("button", { name: "Send greeting" }));

  expect(await screen.findByText("Greeting request failed.")).toBeTruthy();
});
