# Record Collection

A vinyl collection API built while learning Go and PostgreSQL.

## Current state

- Health endpoint: GET /health
- Create record: POST /records
- List records: GET /records
- Records are currently stored in memory and lost on restart.
- PostgreSQL schema is saved in migrations/001_create_records.sql.
- The application connects to PostgreSQL at startup.
- Database connection settings are currently local to the developer's machine.

## Run

Start PostgreSQL:
brew services start postgresql@17

Start the API:
go run .

## Next session

1. Replace in-memory record creation with a PostgreSQL INSERT.
2. Replace in-memory listing with a SELECT.
3. Verify that records survive an application restart.
4. Add input validation, then update and delete endpoints.