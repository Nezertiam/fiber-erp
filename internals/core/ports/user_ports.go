package ports

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nezertiam/fiber-erp/internals/core/domain"
)

type UserHandlers interface {
	Login(c *fiber.Ctx) error
	Register(c *fiber.Ctx) error
	GetUser(c *fiber.Ctx) error
}

type UserService interface {
	Login(email string, password string) (status int, token *string, user *domain.User, err interface{})
	Register(email string, name string, password string, confirmPassword string) (status int, err interface{})
	GetUser(id string) (status int, user *domain.User, err error)
}

type UserRepository interface {
	FindByEmail(email string) (*domain.User, error)
	FindByID(id string) (*domain.User, error)
	Create(user *domain.User) error
}
