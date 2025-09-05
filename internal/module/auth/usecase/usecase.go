package usecase

import (
	"github.com/fiqriardiansyah/user-shopping-api-golang/internal/module/auth/repository"
	"gorm.io/gorm"
)

type AuthUseCase struct {
	*repository.AuthRepository
	db *gorm.DB
}

func NewAuthUseCase(authRepo *repository.AuthRepository, db *gorm.DB) *AuthUseCase {
	return &AuthUseCase{
		AuthRepository: authRepo,
		db:             db,
	}
}
