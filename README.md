
A REST API built with Go, Gin, and GORM, backed by PostgreSQL.

## Prerequisites

- Go 1.26.5+
- PostgreSQL (local install)

## Setup

1. Clone the repo and enter the directory:
   git clone <https://github.com/patelabhinav-angelone/hello>
   cd hello

2. Start PostgreSQL if you don't already have it running:

3. Create the database:

   createdb hello
   

4. Copy the example env file and fill in your values:
   cp .env.example .env
   
   Set `DB_HOST`, `DB_PORT`, `DB_USER`, `DB_PASSWORD`, `DB_NAME` to match your database, and set `JWT_SECRET` to a real secret.

5. Install dependencies:
   ```
   go mod download
   ```

6. Run the server:
   ```
   go run main.go
   ```

The server starts on `http://localhost:8080`.

## API Endpoints

| Method | Path               | Description         |
|--------|--------------------|----------------------|
| GET    | `/status`          | Health check         |
| POST   | `/api/auth/register`| Register a new user  |
| POST   | `/api/auth/login`   | Log in               |
| GET    | `/api/users`        | List all users       |
| GET    | `/api/users/:id`    | Get a user by ID     |
