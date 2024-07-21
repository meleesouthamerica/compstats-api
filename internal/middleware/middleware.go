package middleware

import (
	"crypto/sha256"
	"crypto/subtle"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/keyauth"
	"github.com/splorg/compstats-api/internal/config"
)

func newApiKeyValidator(apiKey string) func(c *fiber.Ctx, key string) (bool, error) {
	return func(c *fiber.Ctx, key string) (bool, error) {
		hashedAPIKey := sha256.Sum256([]byte(apiKey))
		hashedKey := sha256.Sum256([]byte(key))

		if subtle.ConstantTimeCompare(hashedAPIKey[:], hashedKey[:]) == 1 {
			return true, nil
		}

		return false, keyauth.ErrMissingOrMalformedAPIKey
	}
}

type middleware struct {
	*config.ApiConfig
}

func NewMiddleware(cfg *config.ApiConfig) *middleware {
	return &middleware{ApiConfig: cfg}
}

func (m *middleware) SessionAuthentication(c *fiber.Ctx) error {
	s, err := m.Store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}

	userID := s.Get("uid")
	if userID == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}

	return c.Next()
}

func (m *middleware) NewApiKeyAuthentication(key string) func(*fiber.Ctx) error {
	if key == "" {
		log.Fatal("API_KEY not defined in .env")
	}

	return keyauth.New(keyauth.Config{
		KeyLookup: "header:X-API-KEY",
		Validator: newApiKeyValidator(key),
	})
}
