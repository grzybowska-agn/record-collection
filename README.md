# Record Collection

A vinyl collection API built while learning Go and PostgreSQL.

## Endpoints

| Method | Path | Description | Authentication |
|---|---|---|---|
| GET | /health | Check that the API is running | Public |
| GET | /records | List records, newest first | Public |
| POST | /records | Create a record | Bearer token |
| PATCH | /records/{id} | Update selected fields | Bearer token |
| DELETE | /records/{id} | Delete a record | Bearer token |

Records persist in PostgreSQL. IDs and creation timestamps are generated automatically.

PATCH preserves omitted fields. Optional text fields can be cleared with an empty string; pressing year and reissue status can be cleared with null.

Requests are validated before saving. Automated tests cover validation, partial updates, and authentication.

## Run locally

Requires Go, PostgreSQL 17, and the golang-migrate CLI.

1. Start PostgreSQL:

   ```bash
   brew services start postgresql@17
   ```

2. Copy the example configuration:

   ```bash
   cp .env.example .env
   ```

3. Edit `.env` with your settings:

   ```dotenv
   DATABASE_URL=postgres://YOUR_USERNAME@localhost:5432/record_collection?sslmode=disable
   PORT=8080
   API_TOKEN=YOUR_RANDOM_TOKEN
   ```

   Generate a token with:

   ```bash
   openssl rand -hex 32
   ```

   Keep `.env` private; it is excluded from Git.

4. For a fresh setup, create the database:

   ```bash
   /opt/homebrew/opt/postgresql@17/bin/createdb record_collection
   ```

5. Load your trusted local configuration into the terminal environment:

   ```bash
   set -a
   source .env
   set +a
   ```

6. Apply pending migrations:

   ```bash
   migrate -path migrations -database "$DATABASE_URL" up
   ```

   Applied migrations are tracked in the database. Repeated runs report
   `no change` when nothing is pending.

7. Start the API from the project root:

   ```bash
   go run ./cmd/api
   ```

The default port is 8080. Change PORT to use another port.

## Authentication

Write requests require this header:

```http
Authorization: Bearer YOUR_API_TOKEN
```

Missing or invalid credentials return 401 Unauthorized.

GET endpoints are public. Use HTTPS when deploying the API.

## Checks

Run from the project root:

```bash
go fmt ./...
go test ./...
go vet ./...
```

## Migrations

Schema changes live in numbered `.up.sql` files under `migrations/`.

- Add a new migration for each schema change.
- Do not edit migrations that have already been applied.
- Run migrations separately from API startup.
- Down migrations are not configured.

## Next steps

- Add HTTP and database-operation timeouts.
- Add graceful shutdown.
- Tighten JSON request handling and standardise error responses.
- Deploy the API and PostgreSQL database.