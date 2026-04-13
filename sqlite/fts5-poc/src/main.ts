/**
 * SQLite FTS5 proof-of-concept
 *
 * Covers:
 *  - virtual table creation
 *  - basic MATCH query
 *  - column-scoped search
 *  - phrase search
 *  - prefix search
 *  - boolean operators (AND / OR / NOT)
 *  - BM25 relevance ranking
 *  - snippet() for highlighted excerpts
 *  - content table + triggers for incremental sync
 *  - integrity-check
 */
import { DatabaseSync } from "node:sqlite";

const db = new DatabaseSync(":memory:");

// ── 1. Content table (source of truth) ───────────────────────────────────────
db.exec(`
  CREATE TABLE articles (
    id    INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    body  TEXT NOT NULL
  )
`);

// ── 2. FTS5 virtual table ─────────────────────────────────────────────────────
//  content=  → reads from articles; keeps the index small (no data duplication)
//  content_rowid= → maps fts rowid to articles.id
db.exec(`
  CREATE VIRTUAL TABLE articles_fts USING fts5(
    title,
    body,
    content='articles',
    content_rowid='id'
  )
`);

// ── 3. Sync triggers ──────────────────────────────────────────────────────────
db.exec(`
  CREATE TRIGGER articles_ai AFTER INSERT ON articles BEGIN
    INSERT INTO articles_fts(rowid, title, body)
    VALUES (new.id, new.title, new.body);
  END;

  CREATE TRIGGER articles_ad AFTER DELETE ON articles BEGIN
    INSERT INTO articles_fts(articles_fts, rowid, title, body)
    VALUES ('delete', old.id, old.title, old.body);
  END;

  CREATE TRIGGER articles_au AFTER UPDATE ON articles BEGIN
    INSERT INTO articles_fts(articles_fts, rowid, title, body)
    VALUES ('delete', old.id, old.title, old.body);
    INSERT INTO articles_fts(rowid, title, body)
    VALUES (new.id, new.title, new.body);
  END;
`);

// ── 4. Seed data ──────────────────────────────────────────────────────────────
const insert = db.prepare(`INSERT INTO articles (title, body) VALUES (?, ?)`);

const docs: [string, string][] = [
  [
    "Introduction to SQLite",
    "SQLite is a lightweight, serverless, self-contained SQL database engine widely used in mobile and embedded applications.",
  ],
  [
    "Full-Text Search with FTS5",
    "FTS5 is SQLite's latest full-text search extension. It supports BM25 ranking, prefix queries, and phrase matching.",
  ],
  [
    "PostgreSQL vs SQLite",
    "PostgreSQL is a powerful open-source relational database. SQLite trades advanced features for simplicity and zero configuration.",
  ],
  [
    "Building REST APIs with Node.js",
    "Node.js is a popular JavaScript runtime for building scalable server-side applications and REST APIs.",
  ],
  [
    "Database Indexing Strategies",
    "Proper indexing dramatically improves query performance. B-tree indexes suit range queries; hash indexes excel at equality lookups.",
  ],
  [
    "SQLite WAL Mode",
    "Write-Ahead Logging allows concurrent reads during writes, making SQLite suitable for higher-concurrency workloads.",
  ],
];

for (const [title, body] of docs) insert.run(title, body);

// ── helpers ───────────────────────────────────────────────────────────────────
const sep = "─".repeat(60);
function header(label: string) {
  console.log(`\n${sep}\n  ${label}\n${sep}`);
}

function search(label: string, query: string) {
  header(label);
  const rows = db
    .prepare(
      `SELECT a.id, a.title
       FROM articles_fts f
       JOIN articles a ON a.id = f.rowid
       WHERE articles_fts MATCH ?`
    )
    .all(query);

  if (rows.length === 0) console.log("  (no results)");
  else rows.forEach((r) => console.log(" ", r));
}

// ── 5. Basic MATCH ────────────────────────────────────────────────────────────
search("Basic MATCH: 'SQLite'", "SQLite");

// ── 6. Column-scoped search ───────────────────────────────────────────────────
search("Column filter — title only: 'SQLite'", "title:SQLite");

// ── 7. Phrase search ──────────────────────────────────────────────────────────
search('Phrase search: "full-text search"', '"full-text search"');

// ── 8. Prefix search ─────────────────────────────────────────────────────────
search("Prefix search: 'dat*'", "dat*");

// ── 9. Boolean OR ─────────────────────────────────────────────────────────────
// FTS5 treats '.' as a token separator, so wrap dotted terms in double-quotes
search(`Boolean OR: '"Node.js" OR PostgreSQL'`, `"Node.js" OR PostgreSQL`);

// ── 10. Boolean NOT ───────────────────────────────────────────────────────────
search(`Boolean NOT: 'SQLite NOT WAL'`, `SQLite NOT WAL`);

// ── 11. BM25 ranking ─────────────────────────────────────────────────────────
header("BM25 ranked: 'database'  (lower score = more relevant)");
db.prepare(
  `SELECT a.id, a.title, round(bm25(articles_fts), 4) AS score
   FROM articles_fts f
   JOIN articles a ON a.id = f.rowid
   WHERE articles_fts MATCH 'database'
   ORDER BY score`
)
  .all()
  .forEach((r) => console.log(" ", r));

// ── 12. snippet() ─────────────────────────────────────────────────────────────
//  snippet(table, column_index, open_tag, close_tag, ellipsis, max_tokens)
header("snippet() highlight: 'SQLite'");
db.prepare(
  `SELECT snippet(articles_fts, 1, '[', ']', '...', 10) AS excerpt
   FROM articles_fts
   WHERE articles_fts MATCH 'SQLite'`
)
  .all()
  .forEach((r) => console.log(" ", r));

// ── 13. Live insert picked up by trigger ──────────────────────────────────────
header("Insert new doc then search: 'vector'");
db.prepare(`INSERT INTO articles (title, body) VALUES (?, ?)`).run(
  "Vector Search in SQLite",
  "sqlite-vec adds vector embeddings and approximate nearest-neighbor search to SQLite."
);
db.prepare(
  `SELECT a.id, a.title
   FROM articles_fts f
   JOIN articles a ON a.id = f.rowid
   WHERE articles_fts MATCH 'vector'`
)
  .all()
  .forEach((r) => console.log(" ", r));

// ── 14. integrity-check ───────────────────────────────────────────────────────
header("FTS5 integrity-check");
db.prepare(`INSERT INTO articles_fts(articles_fts) VALUES ('integrity-check')`).run();
console.log("  passed (no error thrown)");

db.close();
