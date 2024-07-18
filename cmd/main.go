package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/splorg/compstats-api/internal/auth"
	"github.com/splorg/compstats-api/internal/config"
	"github.com/splorg/compstats-api/internal/middleware"
	"github.com/splorg/compstats-api/internal/tournament"
	"github.com/splorg/compstats-api/internal/validator"
)

func main() {
	validator.Setup()

	apiConfig, err := config.NewApiConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	app := fiber.New()
	app.Use(cors.New(cors.Config{}))

	middleware := middleware.NewMiddleware(apiConfig)

	authHandler := auth.NewAuthHandler(apiConfig)

	app.Post("/register", authHandler.Register)
	app.Post("/login", authHandler.Login)
	app.Post("/logout", authHandler.Logout)

	sessionOnly := app.Group("/auth")
	sessionOnly.Use(middleware.SessionAuthentication)

	sessionOnly.Get("/protected", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to the protected route!")
	})

	tournamentsHandler := tournament.NewTournamentHandler(apiConfig)

	sessionOnly.Get("/tournaments", tournamentsHandler.GetAllTournaments)
	sessionOnly.Get("/tournaments/:id", tournamentsHandler.GetTournamentByID)
	sessionOnly.Post("/tournaments", tournamentsHandler.CreateTournament)
	sessionOnly.Patch("/tournaments/:id", tournamentsHandler.UpdateTournament)
	sessionOnly.Delete("/tournaments/:id", tournamentsHandler.DeleteTournament)

	// to be accessed from Bannerlord game server/Discord bots/etc
	apiKeyOnly := app.Group("/server")
	apiKeyOnly.Use(middleware.ApiKeyAuthentication())

	apiKeyOnly.Get("/players", func(c *fiber.Ctx) error {
		return c.SendString("API KEY only route")
	})

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT is not defined in .env")
	}

	app.Listen(port)
}
