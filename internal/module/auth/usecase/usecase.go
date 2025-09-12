package usecase

import (
	"github.com/fiqriardiansyah/user-shopping-api-golang/internal/helper"
	"github.com/fiqriardiansyah/user-shopping-api-golang/internal/module/auth/repository"
	"gorm.io/gorm"
)

type AuthUseCase struct {
	*helper.Config
	*repository.AuthRepository
	db *gorm.DB
}

func NewAuthUseCase(authRepo *repository.AuthRepository, db *gorm.DB, config *helper.Config) *AuthUseCase {
	return &AuthUseCase{
		AuthRepository: authRepo,
		db:             db,
		Config:         config,
	}
}
