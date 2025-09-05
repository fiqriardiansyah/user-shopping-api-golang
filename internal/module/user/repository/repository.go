package repository

import (
	"github.com/fiqriardiansyah/user-shopping-api-golang/internal/entity"
	"github.com/fiqriardiansyah/user-shopping-api-golang/internal/helper"
)

type UserRepository struct {
	helper.Repository[entity.User]
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}
