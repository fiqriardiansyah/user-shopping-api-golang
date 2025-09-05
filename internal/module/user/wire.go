//go:build wireinject
// +build wireinject

package user

import (
	"github.com/fiqriardiansyah/user-shopping-api-golang/internal/module/user/repository"
	"github.com/fiqriardiansyah/user-shopping-api-golang/internal/module/user/usecase"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializeUserHandler(db *gorm.DB, validate *validator.Validate) *UserController {
	wire.Build(repository.NewUserRepository, usecase.NewUserUseCase, NewUserController)
	return nil
}
