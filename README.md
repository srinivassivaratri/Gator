# RSSAggregator (Gator)

A simple command-line tool that helps you follow your favorite websites and get all their updates in one place. Think of it as your personal news feed aggregator!

## What does it do?
- Follows RSS feeds from websites you like
- Collects all posts in one place
- Shows you updates right in your terminal
- Keeps track of what you've read
- Works with multiple user accounts

## Prerequisites
You'll need:
- Go (1.22 or newer)
- PostgreSQL (17 or newer)

## Installation

1. **Get the code**
```bash
go install github.com/yourusername/RSSAggregator@latest
```

2. **Set up config file**
Create `~/.gatorconfig.json`:
```json
{
    "db_url": "postgres://postgres:postgres@localhost:5433/gator?sslmode=disable",
    "current_user_name": ""
}
```

3. **Set up database**
```bash
# Create database folders
mkdir -p ~/postgres_data ~/postgres_run

# Start PostgreSQL
/usr/lib/postgresql/17/bin/initdb -D ~/postgres_data
/usr/lib/postgresql/17/bin/pg_ctl -D ~/postgres_data \
  -o "-k /home/srinivas/postgres_run -p 5433" \
  -l ~/postgres_data/logfile start

# Create database
PGPORT=5433 PGHOST=/home/srinivas/postgres_run createdb gator

# Set up permissions
PGPORT=5433 PGHOST=/home/srinivas/postgres_run psql -d gator \
  -c "CREATE USER postgres WITH PASSWORD 'postgres' SUPERUSER;"

# Run migrations
goose -dir sql/schema postgres "postgres://postgres:postgres@localhost:5433/gator?sslmode=disable" up
```

## Quick Start

1. **Create your account**
```bash
gator register myusername
```

2. **Follow some feeds**
```bash
# Follow some popular tech blogs
gator follow "https://blog.boot.dev/index.xml"
gator follow "https://news.ycombinator.com/rss"
```

3. **Start collecting posts**
```bash
# Fetch posts every minute
gator agg 1m
```

4. **Browse your posts**
```bash
# Show latest 5 posts
gator browse 5
```

## Available Commands
- `register <name>` - Create your account
- `login <name>` - Switch accounts
- `follow <url>` - Follow a website's RSS feed
- `unfollow <url>` - Stop following a feed
- `following` - List feeds you follow
- `agg <interval>` - Start collecting posts (e.g., agg 1m)
- `browse [limit] [page]` - Read posts with pagination
  - `browse` - Show 10 posts from page 1
  - `browse 5` - Show 5 posts from page 1
  - `browse 5 2` - Show 5 posts from page 2

## Features
✅ User Management
  - Multiple accounts
  - Easy switching
  - Settings persist

✅ Feed Management
  - Follow/unfollow feeds
  - List your follows
  - Smart feed fetching

✅ Post Collection
  - Automatic fetching
  - Deduplication
  - Chronological order
  - Smart pagination
    - Configurable page size
    - Next page hints
    - Page navigation

✅ Data Storage
  - PostgreSQL backend
  - Clean migrations
  - Data persistence

## Troubleshooting

If something's not working:
```bash
# Check if database is running
/usr/lib/postgresql/17/bin/pg_ctl -D ~/postgres_data status

# Start database if needed
/usr/lib/postgresql/17/bin/pg_ctl -D ~/postgres_data \
  -o "-k /home/srinivas/postgres_run -p 5433" \
  -l ~/postgres_data/logfile start

# Check database logs
cat ~/postgres_data/logfile
```

## Development

Want to hack on Gator? Here's how:
```bash
# Get the code
git clone https://github.com/yourusername/RSSAggregator.git

# Build it
go build

# Run it (for development)
go run .

# Install it (for production)
go install
```

## Contributing
Pull requests welcome! Check out our issues page or add your own features.

## License
MIT - Do whatever you want with it!