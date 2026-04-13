import db from "./db";

export interface User {
  id: number;
  name: string;
  email: string;
  created_at: string;
}

export function createUser(name: string, email: string): User {
  return db
    .prepare("INSERT INTO users (name, email) VALUES (?, ?) RETURNING *")
    .get(name, email) as User;
}

export function getUserById(id: number): User | undefined {
  return db.prepare("SELECT * FROM users WHERE id = ?").get(id) as
    | User
    | undefined;
}

export function getAllUsers(): User[] {
  return db.prepare("SELECT * FROM users").all() as User[];
}

export function updateUser(
  id: number,
  fields: Partial<Pick<User, "name" | "email">>
): User | undefined {
  const entries = Object.entries(fields).filter(([, v]) => v !== undefined);
  if (entries.length === 0) return getUserById(id);

  const setClauses = entries.map(([k]) => `${k} = ?`).join(", ");
  const values = entries.map(([, v]) => v);

  db.prepare(`UPDATE users SET ${setClauses} WHERE id = ?`).run(...values, id);
  return getUserById(id);
}

export function deleteUser(id: number): boolean {
  const result = db.prepare("DELETE FROM users WHERE id = ?").run(id);
  return (result.changes as number) > 0;
}
