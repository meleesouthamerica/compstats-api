package config

import (
	"database/sql"
	"errors"
	"log"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"

	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/redis"
	"github.com/splorg/compstats-api/internal/database"
)

type ApiConfig struct {
	DB *database.Queries
	Store *session.Store
}

func NewApiConfig() (*ApiConfig, error) {
	redisHost := os.Getenv("REDIS_HOST")
	if redisHost == "" {
		return nil, errors.New("REDIS_HOST is not defined in .env")
	}

	redisPort := os.Getenv("REDIS_PORT")
	if redisPort == "" {
		return nil, errors.New("REDIS_PORT is not defined in .env")
	}

	redisPortNumber, err := strconv.Atoi(redisPort)
	if err != nil {
		return nil, err
	}

	redisPassword := os.Getenv("REDIS_PASSWORD")
	if redisPassword == "" {
		return nil, errors.New("REDIS_PASSWORD is not defined in .env")
	}

	// redisClient := redis.NewClient(&redis.Options{
	// 	Addr:     redisAddr,
	// 	Password: redisPassword,
	// 	DB:       0,
	// })
	storage := redis.New(redis.Config{
		Host:     redisHost,
		Port:     redisPortNumber,
		Password: redisPassword,
	})

	store := session.New(session.Config{
		Storage:    storage,
		Expiration: 7 * 24 * time.Hour,
		KeyLookup:  "cookie:session_id",
	})

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		return nil, errors.New("DB_URL is not defined in .env")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, errors.New("cannot connect to database")
	}

	var testQuery int

	err = conn.QueryRow("SELECT 1").Scan(&testQuery)
	if err != nil {
		return nil, errors.New("database connection test failed")
	} else {
		log.Print("connection test query executed successfully")
	}

	apiConfig := &ApiConfig{
		DB:    database.New(conn),
		Store: store,
	}

	return apiConfig, nil
}
