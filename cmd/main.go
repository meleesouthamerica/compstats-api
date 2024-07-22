package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"

	"github.com/splorg/compstats-api/internal/auth"
	"github.com/splorg/compstats-api/internal/config"
	"github.com/splorg/compstats-api/internal/half"
	"github.com/splorg/compstats-api/internal/middleware"
	"github.com/splorg/compstats-api/internal/player"
	"github.com/splorg/compstats-api/internal/tournament"
	"github.com/splorg/compstats-api/internal/validator"
	_ "github.com/splorg/compstats-api/docs"
)

var apiKey = os.Getenv("API_KEY")

//	@title			CompStats API
//	@version		1.0
//	@description	REST API for CompStats - stat tracking for online melee game tournaments.
//	@BasePath	/
func main() {
	validator.Setup()

	apiConfig, err := config.NewApiConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	app := fiber.New()
	app.Use(cors.New(cors.Config{}))
	app.Use("/swagger/*", swagger.HandlerDefault)

	middleware := middleware.NewMiddleware(apiConfig)

	authHandler := auth.NewAuthHandler(apiConfig)
	tournamentsHandler := tournament.NewTournamentHandler(apiConfig)
	halfsHandler := half.NewHalfHandler(apiConfig)
	playersHandler := player.NewPlayerHandler(apiConfig)

	publicGroup := app.Group("/auth")

	publicGroup.Post("/register", authHandler.Register)
	publicGroup.Post("/login", authHandler.Login)
	publicGroup.Post("/logout", authHandler.Logout)

	// regular CRUD routes - for manual use by admins
	sessionOnly := app.Group("/admin")
	sessionOnly.Use(middleware.SessionAuthentication)

	sessionOnly.Get("/test", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to an admin-only route!")
	})
	sessionOnly.Get("/tournaments", tournamentsHandler.GetAllTournaments)
	sessionOnly.Get("/tournaments/:id", tournamentsHandler.GetTournamentByID)
	sessionOnly.Post("/tournaments", tournamentsHandler.CreateTournament)
	sessionOnly.Patch("/tournaments/:id", tournamentsHandler.UpdateTournament)
	sessionOnly.Delete("/tournaments/:id", tournamentsHandler.DeleteTournament)

	sessionOnly.Get("/players", playersHandler.GetAllPlayers)
	sessionOnly.Get("/players/:id", playersHandler.GetPlayerByID)
	sessionOnly.Post("/players", playersHandler.CreatePlayer)
	sessionOnly.Patch("/players/:id", playersHandler.UpdatePlayer)
	sessionOnly.Delete("/players/:id", playersHandler.DeletePlayer)

	sessionOnly.Get("/halfs", halfsHandler.GetAllHalfs)
	sessionOnly.Get("/halfs/:id", halfsHandler.GetHalfByID)
	sessionOnly.Post("/halfs", halfsHandler.CreateHalf)
	sessionOnly.Patch("/halfs/:id", halfsHandler.UpdateHalf)
	sessionOnly.Delete("/halfs/:id", halfsHandler.DeleteHalf)

	// to be accessed from Bannerlord game server/Discord bots/etc
	apiKeyAuthentication := middleware.NewApiKeyAuthentication(apiKey)

	apiKeyOnly := app.Group("/")
	apiKeyOnly.Use(apiKeyAuthentication)

	apiKeyOnly.Get("/test", func(c *fiber.Ctx) error {
		return c.SendString("API KEY only route")
	})
	apiKeyOnly.Post("/stats", playersHandler.UpdatePlayerStats)

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT is not defined in .env")
	}

	app.Listen(port)
}
