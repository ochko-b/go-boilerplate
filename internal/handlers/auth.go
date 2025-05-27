package handlers

import (
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/ochko-b/goapp/internal/models"
	"github.com/ochko-b/goapp/internal/services"
	"github.com/ochko-b/goapp/internal/utils"
)

type AuthHandler struct {
	authService *services.AuthService
	validator   *validator.Validate
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		validator:   validator.New(),
	}
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var req models.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body")
	}

	if err := h.validator.Struct(&req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	user, token, err := h.authService.Register(c.Context(), &req)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Email already exists")
		}
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	response := models.AuthResponse{
		User:  *user,
		Token: token,
	}

	return utils.SuccessResponse(c, response, "User registered successfully")
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req models.LoginRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body")
	}

	if err := h.validator.Struct(&req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	user, token, err := h.authService.Login(c.Context(), &req)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Invalid credentials")
	}

	response := models.AuthResponse{
		User:  *user,
		Token: token,
	}

	return utils.SuccessResponse(c, response, "Login successful")
}

func (h *AuthHandler) Refresh(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)
	email := c.Locals("email").(string)

	token, err := h.authService.RefreshToken(userID, email)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, fiber.Map{"token": token}, "Token refreshed successfully")
}
