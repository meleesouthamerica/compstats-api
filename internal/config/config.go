package config

import (
	"database/sql"
	"embed"
	"log"
	"time"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"

	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/sqlite3"
	"github.com/splorg/compstats-api/internal/database"
)

type ApiConfig struct {
	DB    *database.Queries
	Store *session.Store
}

//go:embed migrations/*.sql
var embedMigrations embed.FS

func NewApiConfig() (*ApiConfig, error) {
	db, err := sql.Open("sqlite3", "compstats.db")
	if err != nil {
		return nil, err
	}

	goose.SetDialect("sqlite3")
	goose.SetBaseFS(embedMigrations)

	if err := goose.Up(db, "migrations"); err != nil {
		log.Fatal(err)
	}

	if err := goose.Version(db, "migrations"); err != nil {
		log.Fatal(err)
	}

	_, err = db.Query("SELECT 1")
	if err != nil {
		log.Fatal(err)
	}

	storage := sqlite3.New(sqlite3.Config{
		Database:        "./compstats.db",
		Table:           "sessions",
		Reset:           false,
		GCInterval:      10 * time.Second,
		MaxOpenConns:    100,
		MaxIdleConns:    100,
		ConnMaxLifetime: 1 * time.Second,
	})

	store := session.New(session.Config{
		Storage:    storage,
		Expiration: 7 * 24 * time.Hour,
		KeyLookup:  "cookie:session_id",
	})

	apiConfig := &ApiConfig{
		DB:    database.New(db),
		Store: store,
	}

	return apiConfig, nil
}
