package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/ochko-b/goapp/internal/utils"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Check(c *fiber.Ctx) error {
	return utils.SuccessResponse(c, fiber.Map{
		"status":    "ok",
		"timestamp": time.Now(),
	})
}
