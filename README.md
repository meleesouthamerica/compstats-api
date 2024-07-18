# compstats-api

REST API for compstats
- Players
- Teams
- Tournaments
- Halfs (1 mission = 1 half)
- Player stats per half

This project uses:

- [Fiber](https://gofiber.io/) for routing/middleware
- [Goose](https://pressly.github.io/goose/) for database migrations
- [SQLC](https://sqlc.dev/) for generating type-safe code for SQL queries