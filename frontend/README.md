# TweetApp Frontend

A SvelteKit frontend for the Go/Gin Twitter-clone backend in this repo.

## Tech Stack

- **SvelteKit** (Svelte 5, runes mode) + TypeScript
- **Vite** dev server
- Plain `fetch` for API calls (no state library) — `localStorage` for the JWT token
- No CSS framework — hand-written component styles

## Running

```bash
cd frontend
npm install
npm run dev
```

Runs on `http://localhost:5173`. Requires the Go backend running on `http://localhost:8080` (see repo root `README.md`).

`src/routes/+layout.ts` sets `ssr = false` — this is a client-only SPA. `localStorage` (used for the auth token) isn't available during SSR, and the app doesn't need SEO, so SSR is disabled outright instead of guarding every access.

## Project Structure

```
src/
├── lib/
│   ├── api.ts       # All backend API calls + auth token storage
│   ├── stores.ts    # Svelte stores: currentUser, isLoggedIn
│   └── avatar.ts     # getInitials()/getAvatarColor() for the colored avatar badges
├── routes/
│   ├── +layout.svelte   # Navbar (gradient, search bar, avatar, logout) + auth-gate redirect
│   ├── +layout.ts       # ssr = false
│   ├── +page.svelte     # Home: current user's own tweets + compose modal
│   ├── login/+page.svelte
│   ├── register/+page.svelte
│   ├── profile/+page.svelte       # Own profile: edit/delete own tweets
│   ├── users/[id]/+page.svelte    # Read-only view of another user's profile + tweets
│   └── search/+page.svelte        # Search users by name/email, expand to view their tweets
```

## Auth Flow

- `login()` / `register()` (`src/lib/api.ts`) store the JWT in `localStorage` under `token`, and the user id under `user_id`.
- `register` hits `POST /api/auth/register`, which does **not** return a token — so `register()` immediately calls `login()` with the same credentials afterward to establish a session.
- `+layout.svelte` checks `isAuthenticated()` on mount; if there's no token and the current route isn't `/login` or `/register`, it redirects to `/login`. This is a UX redirect only — the actual security boundary is enforced server-side (see below).
- `logout()` clears `localStorage` and the Svelte stores, then redirects to `/login`.

## Backend Endpoints Used

| Frontend function | Method | Endpoint | Auth |
|---|---|---|---|
| `login` | POST | `/api/auth/login` | — |
| `register` | POST | `/api/auth/register` | — |
| `getCurrentUser` | GET | `/api/users/me` | required |
| `getUserById` | GET | `/api/users/:id` | public |
| `getAllUsers` | GET | `/api/users` | public |
| `searchUsers` | GET | `/api/users/search?q=` | public |
| `getAllTweets` | GET | `/api/tweets` | public |
| `getTweetsByUserId` | GET | `/api/tweets/user/:id` | public |
| `createTweet` | POST | `/api/tweets/user/:id` | required |
| `updateTweet` | PUT | `/api/users/:id/tweets/:tweetId` | required |
| `deleteTweet` | DELETE | `/api/users/:id/tweets/:tweetId` | required |

Reads are intentionally public (like a public profile); only create/update/delete require a bearer token. `getAuthHeaders()` in `api.ts` attaches `Authorization: Bearer <token>` when one exists.

`API_BASE` in `src/lib/api.ts` is hardcoded to `http://localhost:8080/api` — update this if the backend runs elsewhere.

## Pages

- **`/login`, `/register`** — gradient auth cards. Register auto-logs-in on success.
- **`/` (Home)** — shows only the logged-in user's own tweets (`getTweetsByUserId`, not the global feed). A "Tweet" button sits below the navbar, right-aligned; clicking it opens a modal to compose and post.
- **`/profile`** — same tweet list as Home, but with inline edit/delete per tweet (checked against the tweet's `user_id` server-side).
- **`/users/[id]`** — read-only view of another user's profile and tweets.
- **`/search`** — search box that calls `GET /api/users/search?q=`; each result is a collapsible row that lazy-loads that user's tweets (`getTweetsByUserId`) on first expand and caches them in memory for the rest of the session.

## Styling

- Global font: Inter (loaded via Google Fonts in `app.html`).
- Shared visual language: `#1da1f2 → #6a5cf5` gradient for primary actions/headers, white cards with `16px` radius and soft shadows, pill-shaped buttons.
- `lib/avatar.ts` derives a deterministic color + initials from a user's name so the same user always gets the same colored avatar badge across pages, without storing an actual avatar image.

## Known Gaps / Not Implemented

- No pagination anywhere (tweets/users lists load in full).
- No like/follow/retweet features.
- `API_BASE` is hardcoded rather than driven by an env var.
- No automated frontend tests (backend has none either); `svelte-check` is used ad hoc for type/a11y validation.
