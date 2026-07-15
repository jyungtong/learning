import { afterAll, beforeAll, expect, test } from "bun:test";
import { rmSync } from "node:fs";
import { createServer } from "node:net";

import { POST } from "../src/app/api/greet/route";

let server: ReturnType<typeof Bun.spawn>;
const serverBinary = `/tmp/go-nextjs-stack-server-${process.pid}`;
let previousServiceUrl: string | undefined;

async function reserveAddress() {
  const listener = createServer();

  await new Promise<void>((resolve, reject) => {
    listener.once("error", reject);
    listener.listen(0, "127.0.0.1", resolve);
  });

  const address = listener.address();
  if (address === null || typeof address === "string") {
    throw new Error("Could not reserve test port");
  }

  await new Promise<void>((resolve, reject) => {
    listener.close((error) => (error ? reject(error) : resolve()));
  });

  return `127.0.0.1:${address.port}`;
}

beforeAll(async () => {
  const address = await reserveAddress();
  previousServiceUrl = process.env.GREET_SERVICE_URL;
  process.env.GREET_SERVICE_URL = `http://${address}`;

  const build = Bun.spawn(["go", "build", "-o", serverBinary, "./server"], {
    cwd: new URL("../../", import.meta.url).pathname,
    stderr: "ignore",
    stdout: "ignore",
  });
  if ((await build.exited) !== 0) {
    throw new Error("Go service did not build");
  }

  server = Bun.spawn([serverBinary], {
    env: { ...process.env, ADDRESS: address },
    stderr: "ignore",
    stdout: "ignore",
  });

  for (let attempt = 0; attempt < 30; attempt += 1) {
    try {
      const response = await fetch(
        `${process.env.GREET_SERVICE_URL}/greet.v1.GreetService/SayHello`,
        {
          method: "POST",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify({ name: "healthcheck" }),
        },
      );
      const body = (await response.json()) as { message?: string };
      if (response.ok && body.message === "Hello, healthcheck!") {
        return;
      }
    } catch {
      await Bun.sleep(100);
    }
  }

  throw new Error("Go service did not start");
});

afterAll(async () => {
  server.kill();
  await server.exited;
  rmSync(serverBinary, { force: true });
  if (previousServiceUrl === undefined) {
    delete process.env.GREET_SERVICE_URL;
  } else {
    process.env.GREET_SERVICE_URL = previousServiceUrl;
  }
});

test("proxies SayHello through the generated Connect client", async () => {
  const response = await POST(
    new Request("http://localhost:3000/api/greet", {
      method: "POST",
      body: JSON.stringify({ name: "Ada" }),
    }),
  );

  expect(response.status).toBe(200);
  expect(await response.json()).toEqual({ message: "Hello, Ada!" });
});

test("rejects a request without a name", async () => {
  const response = await POST(
    new Request("http://localhost:3000/api/greet", {
      method: "POST",
      body: JSON.stringify({}),
    }),
  );

  expect(response.status).toBe(400);
  expect(await response.json()).toEqual({ error: "Name is required." });
});

test("returns a gateway error when Go service is unavailable", async () => {
  const previousServiceUrl = process.env.GREET_SERVICE_URL;
  process.env.GREET_SERVICE_URL = "http://localhost:1";

  try {
    const response = await POST(
      new Request("http://localhost:3000/api/greet", {
        method: "POST",
        body: JSON.stringify({ name: "Ada" }),
      }),
    );

    expect(response.status).toBe(502);
    expect(await response.json()).toEqual({ error: "Greeting service is unavailable." });
  } finally {
    process.env.GREET_SERVICE_URL = previousServiceUrl;
  }
});
