package usecase

import (
	"github.com/fiqriardiansyah/user-shopping-api-golang/internal/module/user/repository"
	"gorm.io/gorm"
)

type UserUseCase struct {
	*repository.UserRepository
	db *gorm.DB
}

func NewUserUseCase(userRepository *repository.UserRepository, db *gorm.DB) *UserUseCase {
	return &UserUseCase{
		UserRepository: userRepository,
		db:             db,
	}
}
