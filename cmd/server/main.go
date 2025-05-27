package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	fiber_recover "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	"github.com/ochko-b/goapp/internal/config"
	"github.com/ochko-b/goapp/internal/database"
	"github.com/ochko-b/goapp/internal/handlers"
	"github.com/ochko-b/goapp/internal/middleware"
	"github.com/ochko-b/goapp/internal/repository"
	"github.com/ochko-b/goapp/internal/services"
)

func main() {
	// Load Env variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	// Load Configuration
	cfg := config.Load()

	// Initialize DB
	db, err := database.Connect(cfg.Database)
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	defer db.Close()

	// Initize Repository
	repo := repository.New(db)

	// Initialize Services
	authService := services.NewAuthService(repo, cfg.JWT)
	userService := services.NewUserService(repo)

	// Initialize  handlers
	authHandler := handlers.NewAuthHandler(authService)
	userHandler := handlers.NewUserHandler(userService)
	healthHandler := handlers.NewHealthHandler()

	// Initialize Fiber app
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return c.Status(code).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

	// Global Middleware
	app.Use(fiber_recover.New())
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: cfg.CORS.Origins,
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
	}))

	setupRoutes(app, authHandler, userHandler, healthHandler, cfg.JWT.Secret)

	log.Printf("Server starting on %s:%s", cfg.Server.Host, cfg.Server.Port)
	log.Fatal(app.Listen(cfg.Server.Host + ":" + cfg.Server.Port))
}

func setupRoutes(app *fiber.App, authHandler *handlers.AuthHandler, userHandler *handlers.UserHandler, healthHandler *handlers.HealthHandler, jwtSecret string) {
	app.Get("/health", healthHandler.Check)

	api := app.Group("/api/v1")

	auth := api.Group("/auth")
	auth.Post("/register", authHandler.Register)
	auth.Post("/login", authHandler.Login)

	protected := api.Group("/", middleware.JWTAuth(jwtSecret))
	protected.Post("/auth/refresh", authHandler.Refresh)
	protected.Get("/users/me", userHandler.GetProfile)
	protected.Put("/users/me", userHandler.UpdateProfile)
	protected.Get("/users", userHandler.ListUser)
	protected.Get("/users/:id", userHandler.GetUser)
}
