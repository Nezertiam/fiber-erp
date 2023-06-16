package repositories

import (
	"github.com/nezertiam/fiber-erp/internals/core/domain"
	"github.com/nezertiam/fiber-erp/internals/core/ports"

	"gorm.io/gorm"
)

type UserRepository struct {
	client *gorm.DB
}

// Will ensure that UserRepository implements ports.UserRepository
var _ ports.UserRepository = (*UserRepository)(nil)

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		client: db,
	}
}

// ------- FIND BY EMAIL -------

func (r *UserRepository) FindByEmail(email string) (*domain.User, error) {
	var user domain.User
	if err := r.client.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// ------- FIND BY ID -------

func (r *UserRepository) FindByID(id string) (*domain.User, error) {
	var user domain.User
	if err := r.client.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
