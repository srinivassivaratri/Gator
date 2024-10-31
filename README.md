# RSSAggregator

A tool that helps you follow websites by collecting their updates in one place.

## What it does now

- Create users
- Switch between users
- Store everything safely in PostgreSQL
- Remember your settings between uses

## Quick start

1. **Set up storage**
```bash
# Make folders for database
mkdir -p ~/postgres_data ~/postgres_run

# Start up PostgreSQL
/usr/lib/postgresql/17/bin/initdb -D ~/postgres_data
/usr/lib/postgresql/17/bin/pg_ctl -D ~/postgres_data \
  -o "-k /home/srinivas/postgres_run -p 5433" \
  -l ~/postgres_data/logfile start
```

2. **Create database**
```bash
# Make a new database called 'gator'
PGPORT=5433 PGHOST=/home/srinivas/postgres_run \
  /usr/lib/postgresql/17/bin/createdb gator

# Set up permissions
PGPORT=5433 PGHOST=/home/srinivas/postgres_run \
  /usr/lib/postgresql/17/bin/psql -d gator \
  -c "CREATE USER postgres WITH PASSWORD 'postgres' SUPERUSER;"

# Let our user access it
PGPORT=5433 PGHOST=/home/srinivas/postgres_run \
  /usr/lib/postgresql/17/bin/psql -d gator \
  -c "GRANT ALL PRIVILEGES ON DATABASE gator TO postgres;"

# Set up tables
goose -dir sql/schema \
  postgres "postgres://postgres:postgres@localhost:5433/gator?sslmode=disable" up
```

## How it's built

```
.
├── main.go           # Starting point
├── commands.go       # Handles CLI commands
├── handler_user.go   # User management
├── handler_reset.go  # Database cleanup
├── internal/        
│   ├── config/      # Saves your settings
│   └── database/    # Talks to PostgreSQL
└── sql/
    ├── schema/      # Database structure
    └── queries/     # Database operations
```

## What's next

- [x] Save user settings
- [x] Basic commands
- [x] Database setup
- [x] User system
- [ ] RSS reading
- [ ] Web API

## If something breaks

Run these commands:
```bash
# Is database running?
/usr/lib/postgresql/17/bin/pg_ctl -D ~/postgres_data status

# Start it if needed
/usr/lib/postgresql/17/bin/pg_ctl -D ~/postgres_data \
  -o "-k /home/srinivas/postgres_run -p 5433" \
  -l ~/postgres_data/logfile start

# Test connection
PGPORT=5433 PGHOST=/home/srinivas/postgres_run \
  /usr/lib/postgresql/17/bin/psql -d gator

# Check error logs
cat ~/postgres_data/logfile
```