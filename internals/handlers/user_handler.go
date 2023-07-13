package handlers

import (
	"github.com/nezertiam/fiber-erp/internals/core/domain"
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
type LoginResponseSuccess struct {
	Token *string      `json:"token"`
	User  *domain.User `json:"user"`
}
type LoginResponseError struct {
	Errors interface{} `json:"message"`
}

// Login ... Generate token after providing good credentials
// @Summary Generate token after providing good credentials
// @Description Generate token after providing good credentials
// @Tags Users
// @Param body body LoginRequest true "Credentials"
// @Success 200 {object} LoginResponseSuccess
// @Failure 400 {object} LoginResponseError
// @Failure 404 {object} LoginResponseError
// @Router /v1/api/public/auth/login [post]
func (h *UserHandlers) Login(c *fiber.Ctx) error {

	// Parse body
	credentials := new(LoginRequest)
	if err := c.BodyParser(credentials); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(LoginResponseError{
			Errors: err.Error(),
		})
	}
	// Call service
	status, token, user, err := h.userService.Login(credentials.Email, credentials.Password)
	if err != nil {
		return c.Status(status).JSON(LoginResponseError{
			Errors: err,
		})
	}
	// Return token
	return c.Status(fiber.StatusOK).JSON(LoginResponseSuccess{
		Token: token,
		User:  user,
	})
}

// ------- REGISTER -------
type RegisterRequest struct {
	Email           string `json:"email"`
	Name            string `json:"name"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}
type RegisterResponseError struct {
	Errors interface{} `json:"errors"`
}

// Register ... Create a new user
// @Summary Create a new user
// @Description Create a new user
// @Tags Users
// @Param body body RegisterRequest true "Credentials"
// @Success 201
// @Failure 400 {object} RegisterResponseError
// @Router /v1/api/public/auth/register [post]
func (h *UserHandlers) Register(c *fiber.Ctx) error {

	// Parse body
	credentials := new(RegisterRequest)
	if err := c.BodyParser(credentials); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(RegisterResponseError{
			Errors: []string{err.Error()},
		})
	}
	// Call service
	if status, err := h.userService.Register(credentials.Email, credentials.Name, credentials.Password, credentials.ConfirmPassword); err != nil {
		return c.Status(status).JSON(RegisterResponseError{
			Errors: err,
		})
	}
	// Return created status
	return c.SendStatus(fiber.StatusCreated)
}

// ------- GET USER -------
type GetUserResponseSuccess struct {
	Data *domain.User `json:"user"`
}
type GetUserResponseError struct {
	Message string `json:"message"`
}

// Get User ... Retrieve a user
// @Summary Retrieve a user
// @Description Retrieve a user
// @Tags Users
// @Param id path string true "User ID"
// @Success 200 {object} GetUserResponseSuccess
// @Failure 400 {object} RegisterResponseError
// @Failure 404 {object} RegisterResponseError
// @Router /v1/api/protected/users/:id [get]
func (h *UserHandlers) GetUser(c *fiber.Ctx) error {
	// Get id from params
	id := c.Params("id")
	// Call service
	status, user, err := h.userService.GetUser(id)
	if err != nil {
		return c.Status(status).JSON(GetUserResponseError{
			Message: err.Error(),
		})
	}
	// Return user
	return c.Status(status).JSON(GetUserResponseSuccess{
		Data: user,
	})
}
