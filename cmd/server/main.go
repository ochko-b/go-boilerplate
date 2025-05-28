package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	fiber_recover "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	"github.com/ochko-b/goapp/cmd/server/routes"
	"github.com/ochko-b/goapp/internal/config"
	"github.com/ochko-b/goapp/internal/database"
	"github.com/ochko-b/goapp/internal/handlers"
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

	routes.Setup(app, &routes.Handlers{
		Auth:   authHandler,
		User:   userHandler,
		Health: healthHandler,
	}, cfg.JWT.Secret)

	log.Printf("Server starting on %s:%s", cfg.Server.Host, cfg.Server.Port)
	log.Fatal(app.Listen(cfg.Server.Host + ":" + cfg.Server.Port))
}
