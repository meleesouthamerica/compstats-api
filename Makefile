create-migration:
	@goose --dir ./internal/config/migrations sqlite3 ./compstats.db create MIGRATION_NAME sql

migrate-up:
	@goose --dir ./internal/config/migrations sqlite3 ./compstats.db up

migrate-down:
	@goose --dir ./internal/config/migrations sqlite3 ./compstats.db down

run:
	@air