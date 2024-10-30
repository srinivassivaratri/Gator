# RSSAggregator

## Task
Build an RSS feed reader

## Spec
- Store user settings in JSON config file (~/.gatorconfig.json)
- Save DB connection URL with SSL disabled for local development
- Save current username
- Read/write config operations
- CLI commands for user management
- PostgreSQL database with UUID primary keys

## Plan
1. Config system ✓
   - JSON file storage
   - User settings management
   - Database connection info
2. CLI command system ✓
   - Command registration
   - Argument handling
   - Error management
3. Database setup ✓
   - Users table with UUID, timestamps, and unique names
   - Local PostgreSQL instance
   - Goose migrations
4. User management ✓
   - Register new users
   - Login existing users
   - Config file user tracking
5. RSS handler (TODO)
6. API layer (TODO)

## Code
- `main.go`: Entry point with CLI command handling
- `commands.go`: Command registration and execution system
- `handler_user.go`: User-related command handlers
- `internal/config/config.go`: Config file management
- `internal/database/`: Generated database code
- `sql/schema/`: Database migrations
- `sql/queries/`: SQL queries for SQLC

## Setup
```bash
# Database setup
sudo -u postgres psql -c "DROP DATABASE IF EXISTS gator;"
sudo -u postgres psql -c "CREATE DATABASE gator;"
sudo -u postgres psql -d gator -c "GRANT ALL PRIVILEGES ON DATABASE gator TO postgres;"
sudo -u postgres psql -d gator -c "GRANT ALL PRIVILEGES ON SCHEMA public TO postgres;"
sudo -u postgres psql -d gator -c "ALTER SCHEMA public OWNER TO postgres;"

# Run migrations
goose -dir sql/schema postgres "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable" up