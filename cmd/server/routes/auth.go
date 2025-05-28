package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ochko-b/goapp/internal/handlers"
)

func setupAuthRoutes(api fiber.Router, authHandler *handlers.AuthHandler) {
	auth := api.Group("/auth")

	auth.Post("/register", authHandler.Register)
	auth.Post("/login", authHandler.Login)
}

func setupProtectedAuthRoutes(protected fiber.Router, authHandler *handlers.AuthHandler) {
	protected.Post("/auth/refresh", authHandler.Register)
}
