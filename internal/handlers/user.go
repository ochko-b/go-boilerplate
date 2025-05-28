package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/ochko-b/goapp/internal/models"
	"github.com/ochko-b/goapp/internal/services"
	"github.com/ochko-b/goapp/internal/utils"
)

type UserHandler struct {
	userService *services.UserService
	validator   *validator.Validate
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
		validator:   validator.New(),
	}
}

func (h *UserHandler) GetProfile(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)

	user, err := h.userService.GetByID(c.Context(), userID)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusNotFound, "User Not Found")
	}

	return utils.SuccessResponse(c, user)
}

func (h *UserHandler) UpdateProfile(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)

	var req models.UpdateProfileRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid Request Body")
	}

	if err := h.validator.Struct(&req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	user, err := h.userService.UpdateProfile(c.Context(), userID, &req)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, user, "Profile updated successfully")
}

func (h *UserHandler) UpdateUserTransaction(c *fiber.Ctx) error {
	userID := c.Params("id")
	if _, err := uuid.Parse(userID); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid user ID")
	}

	var req models.UpdateProfileRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid Request Body")
	}

	if err := h.validator.Struct(&req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "here?")
	}

	user, err := h.userService.UpdateUserWithTransaction(c.Context(), userID, &req)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, user, "User updated successfully")
}

func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	userID := c.Params("id")
	if _, err := uuid.Parse(userID); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid user ID")
	}

	user, err := h.userService.GetByID(c.Context(), userID)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusNotFound, "User not found")
	}

	return utils.SuccessResponse(c, user)
}

func (h *UserHandler) ListUser(c *fiber.Ctx) error {
	limit := c.QueryInt("limit", 10)
	offset := c.QueryInt("offset", 0)

	if limit > 100 {
		limit = 100
	}

	users, err := h.userService.List(c.Context(), int32(limit), int32(offset))
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, fiber.Map{
		"users":  users,
		"limit":  limit,
		"offset": offset,
	})
}
