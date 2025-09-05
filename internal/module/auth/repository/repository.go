package repository

import (
	"github.com/fiqriardiansyah/user-shopping-api-golang/internal/entity"
	"github.com/fiqriardiansyah/user-shopping-api-golang/internal/helper"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthRepository struct {
	helper.Repository[entity.User]
}

func NewAuthRepository() *AuthRepository {
	return &AuthRepository{}
}

func (r *AuthRepository) GetByEmail(tx *gorm.DB, email string) (*entity.User, error) {
	user := entity.User{}
	if err := tx.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *AuthRepository) GetRoles(tx *gorm.DB, id uuid.UUID) (*[]entity.Role, error) {
	user := entity.User{}

	if err := tx.Preload("Roles").Where("id = ?", id).Find(&user).Error; err != nil {
		return nil, err
	}

	return &user.Roles, nil
}
