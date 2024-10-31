# RSSAggregator

A simple RSS feed reader with user management.

## Features

- User management via CLI
- PostgreSQL storage with UUID keys
- JSON config for settings
- Clean command system

## Setup

1. **Create PostgreSQL directories**
```bash
mkdir -p ~/postgres_data ~/postgres_run
```

2. **Initialize database**
```bash
/usr/lib/postgresql/17/bin/initdb -D ~/postgres_data
```

3. **Start PostgreSQL**
```bash
/usr/lib/postgresql/17/bin/pg_ctl -D ~/postgres_data \
  -o "-k /home/srinivas/postgres_run -p 5433" \
  -l ~/postgres_data/logfile start
```

4. **Setup database**
```bash
# Create DB
PGPORT=5433 PGHOST=/home/srinivas/postgres_run \
  /usr/lib/postgresql/17/bin/createdb gator

# Create user
PGPORT=5433 PGHOST=/home/srinivas/postgres_run \
  /usr/lib/postgresql/17/bin/psql -d gator \
  -c "CREATE USER postgres WITH PASSWORD 'postgres' SUPERUSER;"

# Grant privileges
PGPORT=5433 PGHOST=/home/srinivas/postgres_run \
  /usr/lib/postgresql/17/bin/psql -d gator \
  -c "GRANT ALL PRIVILEGES ON DATABASE gator TO postgres;"

# Run migrations
goose -dir sql/schema \
  postgres "postgres://postgres:postgres@localhost:5433/gator?sslmode=disable" up
```

## Project Structure

```
.
├── main.go           # Entry point
├── commands.go       # CLI command system
├── handler_user.go   # User operations
├── handler_reset.go  # DB reset
├── internal/
│   ├── config/      # Settings management
│   └── database/    # Generated DB code
└── sql/
    ├── schema/      # DB migrations
    └── queries/     # SQL queries
```

## Roadmap

- [x] Config system
- [x] CLI commands
- [x] Database setup
- [x] User management
- [ ] RSS handler
- [ ] API layer

## Troubleshooting

If connection fails:
```bash
# Check status
/usr/lib/postgresql/17/bin/pg_ctl -D ~/postgres_data status

# Start if needed
/usr/lib/postgresql/17/bin/pg_ctl -D ~/postgres_data \
  -o "-k /home/srinivas/postgres_run -p 5433" \
  -l ~/postgres_data/logfile start

# Test connection
PGPORT=5433 PGHOST=/home/srinivas/postgres_run \
  /usr/lib/postgresql/17/bin/psql -d gator

# Check logs
cat ~/postgres_data/logfile
```