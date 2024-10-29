# RSSAggregator

## Task
Build an RSS feed reader

## Spec
- Store user settings in JSON config file (~/.gatorconfig.json)
- Save DB connection URL
- Save current username
- Read/write config operations

## Plan
1. Config system ✓
2. Database setup
3. RSS handler
4. API layer

## Code
- `main.go`: Entry point that demonstrates config operations
- `internal/config/config.go`: Handles reading/writing config file
  - Uses JSON for storage
  - Stores in ~/.gatorconfig.json
  - Currently manages DB URL and username