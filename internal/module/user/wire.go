//go:build wireinject
// +build wireinject

package user

import (
	"github.com/fiqriardiansyah/user-shopping-api-golang/internal/helper"
	"github.com/fiqriardiansyah/user-shopping-api-golang/internal/module/user/repository"
	"github.com/fiqriardiansyah/user-shopping-api-golang/internal/module/user/usecase"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"gorm.io/gorm"
)

type UserHandlerSet struct {
	Controller *UserController
	UseCase    *usecase.UserUseCase
}

func InitializeUserHandler(db *gorm.DB, validate *validator.Validate, config *helper.Config) *UserHandlerSet {
	wire.Build(
		repository.NewUserRepository,
		usecase.NewUserUseCase,
		NewUserController,
		wire.Struct(new(UserHandlerSet), "*"))
	return nil
}
