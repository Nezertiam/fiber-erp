package services

import (
	"errors"
	"log"
	"net/mail"

	"github.com/gofiber/fiber/v2"
	"github.com/nezertiam/fiber-erp/internals/core/domain"
	"github.com/nezertiam/fiber-erp/internals/core/ports"

	passwordvalidator "github.com/wagslane/go-password-validator"
	"golang.org/x/crypto/bcrypt"
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

func (s *UserService) Register(email string, name string, password string, confirmPass string) (status int, err []error) {
	errorsArray := []error{}
	// Validate fields
	if email == "" {
		errorsArray = append(errorsArray, errors.New("email is required"))
	}
	if name == "" {
		errorsArray = append(errorsArray, errors.New("name is required"))
	}
	if password == "" {
		errorsArray = append(errorsArray, errors.New("password is required"))
	}
	if confirmPass == "" {
		errorsArray = append(errorsArray, errors.New("confirmPass is required"))
	}
	if _, er := mail.ParseAddress(email); er != nil {
		errorsArray = append(errorsArray, errors.New("email is not valid")) // Is Email
	}
	if password != confirmPass {
		errorsArray = append(errorsArray, errors.New("passwords do not match"))
	}
	if len(errorsArray) > 0 {
		return fiber.StatusBadRequest, errorsArray
	}
	// Check if user exists
	if _, er := s.userRepository.FindByEmail(email); er == nil {
		errorsArray = append(errorsArray, errors.New("email already exists"))
		return fiber.StatusBadRequest, errorsArray
	}
	// Check if password is strong enough
	const minEntropyBits = 60
	if err := passwordvalidator.Validate("some password", minEntropyBits); err != nil {
		errorsArray = append(errorsArray, errors.New("password is not strong enough"))
		return fiber.StatusBadRequest, errorsArray
	}
	// Hash password
	var hash string
	if bcryptPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost); err != nil {
		errorsArray = append(errorsArray, errors.New("error hashing password"))
		return fiber.StatusInternalServerError, errorsArray
	} else {
		hash = string(bcryptPassword)
	}
	// Create user
	user := domain.User{
		Email:    email,
		Password: hash,
	}
	// Save user
	if err := s.userRepository.Create(&user); err != nil {
		return fiber.StatusInternalServerError, nil
	}
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
