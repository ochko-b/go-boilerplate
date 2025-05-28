package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ochko-b/goapp/internal/handlers"
	"github.com/ochko-b/goapp/internal/middleware"
)

func setupUserRoutes(api fiber.Router, userHandler *handlers.UserHandler, jwtSecret string) {
	protected := api.Group("/", middleware.JWTAuth(jwtSecret))

	// Profile routes
	protected.Get("/users/me", userHandler.GetProfile)
	protected.Put("/users/me", userHandler.UpdateProfile)

	// Management routes
	protected.Get("/users/:id", userHandler.GetUser)
	protected.Put("/users/:id", userHandler.UpdateUserTransaction)
	protected.Get("/users", userHandler.ListUser)
}
