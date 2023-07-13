package services

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nezertiam/fiber-erp/internals/core/domain"
	"github.com/nezertiam/fiber-erp/internals/core/ports"

	jtoken "github.com/golang-jwt/jwt/v4"
	"github.com/golodash/galidator"
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

func (s *UserService) Login(email string, password string) (status int, token *string, user *domain.User, err interface{}) {
	// Validate credentials
	g := galidator.G()
	validator := g.ComplexValidator(galidator.Rules{
		"Email":    g.RuleSet("email").Required().Email(),
		"Password": g.RuleSet("password").Required().String(),
	})
	if err := validator.Validate(map[string]string{
		"Email":    email,
		"Password": password,
	}); err != nil {
		return fiber.StatusBadRequest, nil, nil, err
	}

	// Check if user exists
	user, err = s.userRepository.FindByEmail(email)
	if err != nil {
		var err [1]string
		err[0] = "wrong credentials"
		return fiber.StatusNotFound, nil, nil, map[string]any{
			"email":    err,
			"password": err,
		}
	}

	// Check if password is correct
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		var err [1]string
		err[0] = "wrong credentials"
		return fiber.StatusNotFound, nil, nil, map[string]any{
			"email":    err,
			"password": err,
		}
	}

	// Create the JWT claims, which includes the user ID and expiry time
	day := time.Hour * 24
	secret := os.Getenv("JWT_SECRET")
	claims := jtoken.MapClaims{
		"ID":    user.ID,
		"email": user.Email,
		"exp":   time.Now().Add(day * 1).Unix(),
	}
	// Create token
	jwt := jtoken.NewWithClaims(jtoken.SigningMethodHS256, claims)
	// Generate encoded token and send it as response.
	t, err := jwt.SignedString([]byte(secret))
	if err != nil {
		return fiber.StatusInternalServerError, nil, nil, err
	}
	return fiber.StatusOK, &t, user, nil
}

// ------- REGISTER -------

func (s *UserService) Register(email string, name string, password string, confirmPassword string) (status int, errors interface{}) {

	// Validate fields
	g := galidator.G()
	validator := g.ComplexValidator(galidator.Rules{
		"Email":           g.RuleSet("email").Required().Email(),
		"Name":            g.RuleSet("name").Required().Min(3).Max(20).String().NonEmpty(),
		"Password":        g.RuleSet("password").Required().String(),
		"ConfirmPassword": g.RuleSet("confirmPassword").Required().String(),
	})
	if err := validator.Validate(map[string]string{
		"Email":           strings.Trim(email, " "),
		"Name":            strings.Trim(name, " "),
		"Password":        strings.Trim(password, " "),
		"ConfirmPassword": strings.Trim(confirmPassword, " "),
	}); err != nil {
		return fiber.StatusBadRequest, err
	}
	// Check if passwords match
	if password != confirmPassword {
		var err [1]string
		err[0] = "Passwords do not match"
		return fiber.StatusBadRequest, map[string]any{
			"confirmPassword": err,
		}
	}
	// Check if password is strong enough
	const minEntropyBits = 60
	if err := passwordvalidator.Validate(password, minEntropyBits); err != nil {
		var err [1]string
		err[0] = "insecure password, try including more special characters, using uppercase letters, using numbers or using a longer password"
		return fiber.StatusBadRequest, map[string]any{
			"password": err,
		}
	}

	// Check if user exists
	if _, err := s.userRepository.FindByEmail(email); err == nil {
		var err [1]string
		err[0] = "email already exists"
		return fiber.StatusBadRequest, map[string]any{
			"email": err,
		}
	}

	// Hash password
	var hash string
	if bcryptPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost); err != nil {
		return fiber.StatusInternalServerError, err
	} else {
		hash = string(bcryptPassword)
	}
	// Create user
	user := domain.User{
		Email:    email,
		Password: hash,
	}
	fmt.Println(user)
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
