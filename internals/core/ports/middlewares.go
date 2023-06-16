package ports

import (
	"github.com/gofiber/fiber/v2"
)

type Middlewares interface {
	NewAuthMiddleware(secret string) fiber.Handler
}
