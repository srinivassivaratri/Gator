# 🐊 Gator - Your Terminal RSS Feed Aggregator

Gator is a powerful command-line RSS feed aggregator that brings all your favorite web content into one place. Built in Go, it helps you follow websites, collect their updates, and read them right in your terminal.

## Why Gator?

RSS feeds are scattered across the web, making it hard to keep up with your favorite content. Most RSS readers are bloated web apps or desktop applications. I wanted something:

- Lightning fast and lightweight
- Usable entirely from the terminal
- Multi-user friendly
- With smart feed management

Gator solves this by providing a simple CLI that just works. No browser needed, no fancy UI - just your content, when you want it.

## 🚀 Quick Start

1. **Install Gator**
```bash
go install github.com/yourusername/RSSAggregator@latest
```

2. **Create your account**
```bash
gator register myusername
```

3. **Follow some feeds**
```bash
gator follow "https://blog.boot.dev/index.xml"
gator follow "https://news.ycombinator.com/rss"
```

4. **Start reading!**
```bash
gator browse
```

## 📖 Usage

### Core Commands

- `register <name>` - Create your account
- `login <name>` - Switch between accounts
- `follow <url>` - Follow a new RSS feed
- `unfollow <url>` - Stop following a feed
- `following` - List your followed feeds
- `browse [limit] [page]` - Read posts with pagination
  - `browse` - Show 10 posts from page 1
  - `browse 5` - Show 5 posts from page 1
  - `browse 5 2` - Show 5 posts from page 2
- `agg <interval>` - Start collecting posts (e.g., `agg 1m`)

### Feed Collection Options

The `agg` command supports various time intervals:
- `agg 30s` - Collect every 30 seconds
- `agg 1m` - Collect every minute
- `agg 1h` - Collect every hour

## 🛠️ Development Setup

### Prerequisites
- Go 1.22+
- PostgreSQL 17+

### Database Setup
```bash
# Create database directories
mkdir -p ~/postgres_data ~/postgres_run

# Initialize and start PostgreSQL
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

### Configuration
Create `~/.gatorconfig.json`:
```json
{
    "db_url": "postgres://postgres:postgres@localhost:5433/gator?sslmode=disable",
    "current_user_name": ""
}
```

## 🤝 Contributing

1. **Get the code**
```bash
git clone https://github.com/yourusername/RSSAggregator.git
cd RSSAggregator
```

2. **Build it**
```bash
go build
```

3. **Run tests**
```bash
go test ./...
```

4. Submit a PR with your changes!

