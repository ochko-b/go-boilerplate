package utils

import "github.com/gofiber/fiber/v2"

type APIResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}

func SuccessResponse(c *fiber.Ctx, data any, message ...string) error {
	response := APIResponse{
		Success: true,
		Data:    data,
	}

	if len(message) > 0 {
		response.Message = message[0]
	}
	return c.JSON(response)
}

func ErrorResponse(c *fiber.Ctx, status int, err string) error {
	response := APIResponse{
		Success: false,
		Message: err,
	}
	return c.Status(status).JSON(response)
}
