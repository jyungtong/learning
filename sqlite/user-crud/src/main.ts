import { Worker } from "node:worker_threads";
import { fileURLToPath } from "node:url";
import { createUser, updateUser, deleteUser } from "./user";

const ts = () => new Date().toISOString().slice(11, 23);
const sleep = (ms: number) => new Promise((r) => setTimeout(r, ms));

const CYCLES = 10;
const READER_ITERATIONS = 20;

// Spawn reader worker with tsx support for .ts files
const worker = new Worker(fileURLToPath(new URL("./worker.ts", import.meta.url)), {
  execArgv: ["--import", "tsx"],
  workerData: { iterations: READER_ITERATIONS },
});

let reads = 0;
let errors = 0;

worker.on("message", ({ ok, msg }: { ok: boolean; msg: string }) => {
  console.log(msg);
  ok ? reads++ : errors++;
});

worker.on("error", (err) => console.error("Worker error:", err));

const workerDone = new Promise<void>((resolve) => worker.on("exit", resolve));

// Main thread: hammer writes
console.log(`[${ts()}] --- start: ${CYCLES} write cycles, ${READER_ITERATIONS} concurrent reads ---`);

for (let i = 0; i < CYCLES; i++) {
  const user = createUser(`User_${i}`, `user${i}_${Date.now()}@example.com`);
  console.log(`[${ts()}] writer CREATE id=${user.id}`);

  await sleep(20);

  updateUser(user.id, { email: `user${i}_updated_${Date.now()}@example.com` });
  console.log(`[${ts()}] writer UPDATE id=${user.id}`);

  await sleep(20);

  deleteUser(user.id);
  console.log(`[${ts()}] writer DELETE id=${user.id}`);

  await sleep(20);
}

await workerDone;
console.log(`\n[${ts()}] --- done ---`);
console.log(`reads ok=${reads} errors=${errors}`);
