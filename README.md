# RSSAggregator

A tool that helps you follow websites by collecting their updates in one place.

## Task
Build a command-line RSS feed aggregator that helps users follow multiple websites and get updates in one place.

## Spec
- Users can register and switch between accounts
- Settings persist between sessions
- Data stored securely in PostgreSQL
- Fetch and parse RSS feeds
- Clean command-line interface

## Plan
1. Set up basic project structure
2. Implement user management system
3. Add database integration
4. Create configuration system
5. Add RSS feed fetching and parsing
6. Add feed management commands
7. Build web API (upcoming)

## Code

### Project Structure
```
.
├── main.go           # Starting point
├── commands.go       # Handles CLI commands
├── handler_user.go   # User management
├── handler_reset.go  # Database cleanup
├── handler_agg.go    # RSS feed aggregation
├── handler_feed.go   # Feed management
├── rss_feed.go      # RSS feed parsing
├── internal/        
│   ├── config/      # Saves your settings
│   └── database/    # Talks to PostgreSQL
└── sql/
    ├── schema/      # Database structure
    └── queries/     # Database operations
```

### Available Commands
- `register <name>` - Create a new user
- `login <name>` - Switch to existing user
- `users` - List all users
- `reset` - Clear database
- `addfeed <name> <url>` - Add a new RSS feed
- `agg` - Test feed aggregation

### Setup Instructions

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

### Progress
- [x] Save user settings
- [x] Basic commands
- [x] Database setup
- [x] User system
- [x] RSS reading
- [x] Feed management
- [ ] Web API

### Troubleshooting

If something breaks, run these commands:
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