package middlewares

import (
	"github.com/nezertiam/fiber-erp/internals/core/ports"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

type Middlewares struct{}

var _ ports.Middlewares = (*Middlewares)(nil)

func NewMiddlewares() *Middlewares {
	return &Middlewares{}
}

func (*Middlewares) NewAuthMiddleware(secret string) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(secret),
	})
}
