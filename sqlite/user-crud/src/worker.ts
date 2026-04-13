import { DatabaseSync } from "node:sqlite";
import { workerData, parentPort } from "node:worker_threads";

const db = new DatabaseSync("users.db");

// db.exec(`PRAGMA journal_mode=WAL`);
// db.exec(`PRAGMA synchronous=NORMAL`);

const ts = () => new Date().toISOString().slice(11, 23);

const { iterations } = workerData as { iterations: number };

for (let i = 0; i < iterations; i++) {
  try {
    const users = db.prepare("SELECT * FROM users").all();
    const msg = `[${ts()}] reader READ_ALL count=${users.length}`;
    parentPort!.postMessage({ ok: true, msg });
  } catch (err: any) {
    parentPort!.postMessage({ ok: false, msg: `[${ts()}] reader ERROR: ${err.message}` });
  }
  await new Promise((r) => setImmediate(r));
}
