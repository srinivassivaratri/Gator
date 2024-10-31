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
   - JSON file storage ✓
   - User settings management ✓
   - Database connection info ✓
2. CLI command system ✓
   - Command registration ✓
   - Argument handling ✓
   - Error management ✓
3. Database setup ✓
   - Users table with UUID, timestamps, and unique names ✓
   - Local PostgreSQL instance ✓
   - Goose migrations ✓
4. User management ✓
   - Register new users ✓
   - Login existing users ✓
   - Config file user tracking ✓
   - List users command ✓
   - Reset database command ✓
5. RSS handler (TODO)
   - Feed table setup
   - Feed post storage
   - Feed fetching logic
6. API layer (TODO)
   - RESTful endpoints
   - Feed management
   - User authentication

## Code
- `main.go`: Entry point with CLI command handling
- `commands.go`: Command registration and execution system
- `handler_user.go`: User-related command handlers
- `handler_reset.go`: Database reset functionality
- `internal/config/config.go`: Config file management
- `internal/database/`: Generated database code
- `sql/schema/`: Database migrations
- `sql/queries/`: SQL queries for SQLC

## Setup (Local PostgreSQL)
```bash
# Create local PostgreSQL directories
mkdir -p ~/postgres_data ~/postgres_run

# Initialize PostgreSQL database
/usr/lib/postgresql/17/bin/initdb -D ~/postgres_data

# Start PostgreSQL on port 5433
/usr/lib/postgresql/17/bin/pg_ctl -D ~/postgres_data -o "-k /home/srinivas/postgres_run -p 5433" -l ~/postgres_data/logfile start

# Create database and set up permissions
PGPORT=5433 PGHOST=/home/srinivas/postgres_run /usr/lib/postgresql/17/bin/createdb gator
PGPORT=5433 PGHOST=/home/srinivas/postgres_run /usr/lib/postgresql/17/bin/psql -d gator -c "CREATE USER postgres WITH PASSWORD 'postgres' SUPERUSER;"
PGPORT=5433 PGHOST=/home/srinivas/postgres_run /usr/lib/postgresql/17/bin/psql -d gator -c "GRANT ALL PRIVILEGES ON DATABASE gator TO postgres;"

# Run migrations
goose -dir sql/schema postgres "postgres://postgres:postgres@localhost:5433/gator?sslmode=disable" up
```

## Troubleshooting
If you get connection errors:
1. Check PostgreSQL status: `/usr/lib/postgresql/17/bin/pg_ctl -D ~/postgres_data status`
2. Start PostgreSQL if needed: `/usr/lib/postgresql/17/bin/pg_ctl -D ~/postgres_data -o "-k /home/srinivas/postgres_run -p 5433" -l ~/postgres_data/logfile start`
3. Verify connection: `PGPORT=5433 PGHOST=/home/srinivas/postgres_run /usr/lib/postgresql/17/bin/psql -d gator`
4. Check logs: `cat ~/postgres_data/logfile`