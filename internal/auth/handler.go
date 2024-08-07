package auth

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/splorg/compstats-api/internal/config"
	"github.com/splorg/compstats-api/internal/database"
	"github.com/splorg/compstats-api/internal/util"
	"github.com/splorg/compstats-api/internal/validator"
)

type AuthHandler struct {
	*config.ApiConfig
}

func NewAuthHandler(config *config.ApiConfig) *AuthHandler {
	return &AuthHandler{ApiConfig: config}
}

// Register godoc
//
//	@Summary		Register
//	@Description	user registration
//	@Accept			json
//	@Produce		json
//
// @Param dto body auth.registerDTO true "register json"
//
//	@Success		201	{object}	database.User
//	@Router			/auth/register [post]
func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var req registerDTO

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": err.Error()})
	}

	if err := validator.ValidateStruct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	password, err := util.HashPassword([]byte(req.Password))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to encrypt password"})
	}

	newUser, err := h.DB.CreateUser(c.Context(), database.CreateUserParams{
		Name:     req.Name,
		Password: string(password),
		Email:    req.Email,
	})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "credentials already in use"})
	}

	return c.Status(fiber.StatusCreated).JSON(newUser)
}

// Login godoc
//
//	@Summary		Login
//	@Description	user login
//	@Accept			json
//	@Produce		json
//
// @Param dto body auth.loginDTO true "login json"
//
//	@Success		200	{object}	database.User
//	@Router			/auth/login [post]
func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req loginDTO

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": err.Error()})
	}

	if err := validator.ValidateStruct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	foundUser, err := h.DB.FindUserByEmail(c.Context(), req.Email)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "no user found"})
	}

	if err := util.ComparePassword([]byte(foundUser.Password), []byte(req.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid credentials"})
	}

	s, _ := h.Store.Get(c)

	if s.Fresh() {
		sid := s.ID()

		uid := foundUser.ID

		s.Set("uid", uid)
		s.Set("sid", sid)
		s.Set("ip", c.Context().RemoteIP().String())
		s.Set("login", time.Unix(time.Now().Unix(), 0).UTC().String())
		s.Set("ua", string(c.Request().Header.UserAgent()))

		err = s.Save()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
	}

	return c.Status(fiber.StatusOK).JSON(foundUser)
}

// Logout godoc
//
//	@Summary		Logout
//	@Description	User logout route
//	@Accept			json
//	@Produce		json
//	@Router			/auth/logout [post]
func (h *AuthHandler) Logout(c *fiber.Ctx) error {
	s, _ := h.Store.Get(c)

	s.Destroy()

	return c.Status(fiber.StatusOK).JSON(authResponse{Message: "logged out successfully"})
}
