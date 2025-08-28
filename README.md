# Gator Blog Aggregator

A basic RSS aggregator as a [Boot.dev](https://boot.dev) guided project.

## Requirements

- Go - To build / install
- Postgres - Active DB while running
- Goose - Setup DB

## Building

```bash
go build

# Optionally, have Go install to GOPATH
go install
```

## Setup

### Postgres

How to setup a postgres instance as a user, with specific file store location

```bash
# Initialize Postgres
initdb -D <dblocation>

# Start
mkdir -p /run/user/<user-id>/postgresql/
postgres -D <dblocation> 2> /dev/null &

# Setup DB
goose postgres "postgresql://[userspec@][hostspec][/dbname]" --dir sql/schema up
```

### Config File

In `.gatorconfig.json`, set `db_url` to the same URI used to Setup DB.
(i.e. `postgresql://[userspec@][hostspec][/dbname]`)

`current_user_name` will be set by the program

## Running

1. Ensure Postgres is running (See Start of Postgres Setup)
2. Create a user with `gator register <username>`
3. Run `gator agg <interval>`

### Commands

- `login <username>` - Set to user
- `register <username>` - Create (and set to) user
- `users` - List users
- `reset` - Delete all users and feeds
- `agg <duration>` - Fetch a feed every `<duration>`
- `addfeed <name> <url>` - Create feed listener
- `feeds` - List all feeds
- `follow <url>` - Add feed to current user's browse page
- `unfollow <url>` - Remove feed from current user's browse page
- `following` - List all feeds followed by current user
- `browse [limit]` - Show last `[limit]` posts from followed feeds (default 2)
