# TanStack Query Todo POC

## Goal

Build a small todo application that demonstrates TanStack Query data fetching, mutations, cache invalidation, and optimistic updates in Next.js 16.

## Scope

Users can create todos, mark them complete or active, delete them, and filter the visible list by status. The interface includes loading, empty, and recoverable error states plus active and completed counts.

Editing titles, authentication, pagination, drag-and-drop, and durable persistence are excluded.

## Architecture

- Next.js Route Handlers expose JSON endpoints under `/api/todos`.
- A server-only module owns a seeded, process-local array of todos. Data resets when the Next.js process restarts and is not shared across instances.
- A client component uses TanStack Query for reads and mutations.
- The root layout mounts one browser-stable `QueryClientProvider`.

## API

- `GET /api/todos` returns all todos in creation order.
- `POST /api/todos` accepts `{ "title": string }`, validates a non-empty trimmed title, and returns the created todo.
- `PATCH /api/todos/:id` accepts `{ "completed": boolean }` and returns the updated todo.
- `DELETE /api/todos/:id` removes the todo and returns HTTP 204.
- Invalid input returns HTTP 400. Unknown IDs return HTTP 404.

Each todo has `id`, `title`, `completed`, and `createdAt` fields.

## Client Data Flow

The list query uses key `['todos']`. Create invalidates that key after success. Toggle and delete update cached todos optimistically, preserve a snapshot, roll back on failure, and invalidate after settlement so server state remains authoritative.

Filters are local presentation state because the complete dataset is already cached. Mutation failures appear in a compact alert without discarding existing data.

## Interface

The page is a focused task workspace: compact header and counts, prominent add input, segmented status filter, and stable todo rows with checkbox and icon-only delete action. Layout remains usable on narrow mobile and desktop screens. Controls include accessible labels, focus styles, and disabled states while relevant work is pending.

## Verification

- ESLint and production build must pass.
- Browser verification covers initial fetch, create, toggle, filtering, delete, empty-state behavior, mobile layout, and console errors.
