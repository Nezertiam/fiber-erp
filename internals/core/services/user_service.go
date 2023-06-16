package services

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/nezertiam/fiber-erp/internals/core/domain"
	"github.com/nezertiam/fiber-erp/internals/core/ports"
)

type UserService struct {
	userRepository ports.UserRepository
}

var _ ports.UserService = (*UserService)(nil)

func NewUserService(repository ports.UserRepository) *UserService {
	return &UserService{
		userRepository: repository,
	}
}

// ------- LOGIN -------

func (s *UserService) Login(email string, password string) (status int, token *string, err error) {
	// Validate credentials
	if email == "" || password == "" {
		return fiber.StatusBadRequest, nil, errors.New("email and password are required")
	}

	// Check if user exists
	user, err := s.userRepository.FindByEmail(email)
	if err != nil {
		return fiber.StatusNotFound, nil, err
	}

	// TODO: Check if password is correct
	log.Println(user.ID)

	// Create JWT token

	jwt := "token here pliz"
	return fiber.StatusOK, &jwt, nil
}

// ------- REGISTER -------

func (s *UserService) Register(email string, password string, confirmPass string) (status int, err error) {
	// Validate not empty
	if email == "" || password == "" || confirmPass == "" {
		return fiber.StatusBadRequest, errors.New("email and password are required")
	}
	// Validate password match
	if password != confirmPass {
		return fiber.StatusBadRequest, errors.New("passwords do not match")
	}
	// Check if user exists
	if _, err = s.userRepository.FindByEmail(email); err == nil {
		return fiber.StatusBadRequest, errors.New("user already exists")
	}
	// Create user
	// user := domain.User{
	// 	Email:    email,
	// 	Password: password,
	// }
	// Save user
	// TODO: Implement save user in repository
	// if err = s.userRepository.Save(&user); err != nil {
	// 	return fiber.StatusInternalServerError, nil, err
	// }
	return fiber.StatusCreated, nil
}

// ------- GET USER -------

func (s *UserService) GetUser(id string) (status int, user *domain.User, err error) {
	// Validate not empty
	if id == "" {
		return fiber.StatusBadRequest, nil, errors.New("id is required")
	}

	// Check if user exists
	user, err = s.userRepository.FindByID(id)
	if err != nil {
		return fiber.StatusNotFound, nil, err
	}

	return fiber.StatusOK, user, nil
}
