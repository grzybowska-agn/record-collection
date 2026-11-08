# Record Collection

A vinyl collection API built while learning Go and PostgreSQL.

## Current state

- GET /health — checks that the API is running.
- POST /records — creates a record in PostgreSQL.
- GET /records — lists records, newest first.
- Records persist across application restarts.
- Each record includes an automatically generated ID and creation timestamp.
- Database schema changes are stored in migrations/.
- Configuration is read from environment variables, with local .env support.

## Run locally

1. Start PostgreSQL:
   brew services start postgresql@17

2. Copy the example configuration:
   cp .env.example .env

3. Edit .env with your local PostgreSQL connection details.

4. Start the API from the project root:
   go run ./cmd/api

The API listens on http://localhost:8080.

## Next steps

- Read the server port from PORT.
- Add request validation.
- Add update and delete endpoints.
- Protect write endpoints before public deployment.
- Deploy the API and PostgreSQL database.