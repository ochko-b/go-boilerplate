package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ochko-b/goapp/internal/handlers"
	"github.com/ochko-b/goapp/internal/middleware"
)

type Handlers struct {
	Auth   *handlers.AuthHandler
	User   *handlers.UserHandler
	Health *handlers.HealthHandler
}

func Setup(app *fiber.App, h *Handlers, jwtSecret string) {
	app.Get("/health", h.Health.Check)

	api := app.Group("/api/v1")

	setupAuthRoutes(api, h.Auth)
	setupUserRoutes(api, h.User, jwtSecret)

	protected := api.Group("/", middleware.JWTAuth(jwtSecret))
	setupProtectedAuthRoutes(protected, h.Auth)
}
