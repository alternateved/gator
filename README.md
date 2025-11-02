# gator

A minimal command-line RSS feed aggregator.

## Installation

This aggregator requires PostgreSQL and Go to be installed. To install the program, simply run `make install` in the root of the project.

## Configuration

Configuration is stored in `$XDG_CONFIG_HOME/gator/config.json`. Provide the appropriate `db_url`:

```json
{"db_url": "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable"}
```

## Commands

- `register <name>` - Register a user and save it in the database
- `login <name>` - Log in to your user account
- `reset` - Reset all users in the database
- `users` - List all users and mark the current user
- `agg` - Aggregate feeds (should be run in the background)
- `addfeed <name> <url>` - Add a feed and follow it with the current user
- `feeds` - List all feeds
- `follow <url>` - Follow a feed with the current user
- `unfollow <url>` - Unfollow a feed with the current user
- `following` - List followed feeds
- `browse <limit>` - Browse aggregated posts (displays 2 posts by default)
