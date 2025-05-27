package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func CORS(origins string) fiber.Handler {
	return cors.New(cors.Config{
		AllowOrigins:     origins,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
		AllowHeaders:     "Origin,Content-type,Accept,Authorization,X-Requested-With",
		AllowCredentials: true,
	})
}
