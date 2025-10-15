//go:build wireinject
// +build wireinject

package auth

import (
	"github.com/fiqriardiansyah/user-shopping-api-golang/internal/helper"
	"github.com/fiqriardiansyah/user-shopping-api-golang/internal/module/auth/repository"
	"github.com/fiqriardiansyah/user-shopping-api-golang/internal/module/auth/usecase"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"gorm.io/gorm"
)

type AuthHandlerSet struct {
	Controller *AuthController
	UseCase    *usecase.AuthUseCase
}

func InitializeAuthHandler(db *gorm.DB, validator *validator.Validate, config *helper.Config) *AuthHandlerSet {
	wire.Build(
		usecase.NewAuthUseCase,
		repository.NewAuthRepository,
		NewAuthController,
		wire.Struct(new(AuthHandlerSet), "*"))
	return nil
}
