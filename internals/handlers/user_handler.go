package handlers

import (
	"github.com/nezertiam/fiber-erp/internals/core/ports"

	"github.com/gofiber/fiber/v2"
)

type UserHandlers struct {
	userService ports.UserService
}

var _ ports.UserHandlers = (*UserHandlers)(nil)

func NewUserHandlers(userService ports.UserService) *UserHandlers {
	return &UserHandlers{
		userService: userService,
	}
}

// ------- LOGIN -------

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *UserHandlers) Login(c *fiber.Ctx) error {
	// Parse body
	credentials := new(LoginRequest)
	if err := c.BodyParser(credentials); err != nil {
		return err
	}
	// Call service
	status, token, err := h.userService.Login(credentials.Email, credentials.Password)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	// Return token
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
	})
}

// ------- REGISTER -------

type RegisterRequest struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

func (h *UserHandlers) Register(c *fiber.Ctx) error {
	// Parse body
	credentials := new(RegisterRequest)
	if err := c.BodyParser(credentials); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	// Call service
	status, err := h.userService.Register(credentials.Email, credentials.Password, credentials.ConfirmPassword)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	// Return created status
	return c.SendStatus(fiber.StatusCreated)
}

// ------- GET USER -------
func (h *UserHandlers) GetUser(c *fiber.Ctx) error {
	// Get id from params
	id := c.Params("id")
	// Call service
	status, user, err := h.userService.GetUser(id)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	// Return user
	return c.Status(status).JSON(fiber.Map{
		"user": user,
	})
}
