package server

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/swagger"
	_ "github.com/nezertiam/fiber-erp/docs"
	"github.com/nezertiam/fiber-erp/internals/core/ports"

	fiber "github.com/gofiber/fiber/v2"
)

type Server struct {
	//We will add every new Handler here
	userHandlers ports.UserHandlers
	middlewares  ports.Middlewares
	//paymentHandlers ports.IPaymentHandlers
}

func NewServer(uHandlers ports.UserHandlers, middlewares ports.Middlewares) *Server {
	return &Server{
		userHandlers: uHandlers,
		middlewares:  middlewares,
	}
}

func (s *Server) Initialize() {
	app := fiber.New()

	// Swagger
	app.Get("/swagger/*", swagger.HandlerDefault) // default

	// ----- v1 api -----
	v1 := app.Group("/v1/api")
	v1.Get("/health", HealthCheck)

	// PUBLIC ROUTES
	public := v1.Group("public")                          // /v1/api/public
	userRoutes := public.Group("/auth")                   // /v1/api/public/auth
	userRoutes.Post("/login", s.userHandlers.Login)       // /v1/api/public/auth/login
	userRoutes.Post("/register", s.userHandlers.Register) // /v1/api/public/auth/register

	// PROTECTED ROUTES
	jwt := s.middlewares.NewAuthMiddleware(os.Getenv("JWT_SECRET"))
	protected := v1.Group("/protected", jwt)       // /v1/api/protected
	userRoutes = protected.Group("/users")         // /v1/api/protected/users
	userRoutes.Get("/me", s.userHandlers.GetMe)    // /v1/api/protected/users/me
	userRoutes.Get("/:id", s.userHandlers.GetUser) // /v1/api/protected/users/:id

	server_port := os.Getenv("SERVER_PORT")
	if err := app.Listen(fmt.Sprintf(":%s", server_port)); err != nil {
		log.Fatal(err)
	}
}

// Swagger documentation
func HealthCheck(c *fiber.Ctx) error {
	res := map[string]interface{}{
		"data": "Server is up and running",
	}

	if err := c.JSON(res); err != nil {
		return err
	}

	return nil
}
