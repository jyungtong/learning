import { DatabaseSync } from "node:sqlite";

const db = new DatabaseSync("users.db");

// db.exec(`PRAGMA journal_mode=WAL`);
// db.exec(`PRAGMA synchronous=NORMAL`);

db.exec(`
  CREATE TABLE IF NOT EXISTS users (
    id         INTEGER PRIMARY KEY AUTOINCREMENT,
    name       TEXT    NOT NULL,
    email      TEXT    NOT NULL UNIQUE,
    created_at TEXT    NOT NULL DEFAULT (datetime('now'))
  )
`);

export default db;
