"use client";

import { FormEvent, useState } from "react";

export function GreetForm() {
  const [name, setName] = useState("");
  const [error, setError] = useState("");
  const [message, setMessage] = useState("");

  async function handleSubmit(event: FormEvent<HTMLFormElement>) {
    event.preventDefault();

    try {
      const response = await fetch("/api/greet", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ name }),
      });
      const body = (await response.json()) as { error?: string; message?: string };

      if (!response.ok || typeof body.message !== "string") {
        setMessage("");
        setError(body.error ?? "Greeting request failed.");
        return;
      }

      setError("");
      setMessage(body.message);
    } catch {
      setMessage("");
      setError("Greeting request failed.");
    }
  }

  return (
    <form className="greet-form" onSubmit={handleSubmit}>
      <label htmlFor="name">Name</label>
      <div className="greet-controls">
        <input
          id="name"
          name="name"
          onChange={(event) => setName(event.target.value)}
          placeholder="Ada"
          required
          value={name}
        />
        <button type="submit">Send greeting</button>
      </div>
      {error ? <p className="greet-error">{error}</p> : null}
      {message ? <output className="greet-result">{message}</output> : null}
    </form>
  );
}
