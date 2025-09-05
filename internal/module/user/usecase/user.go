package usecase

import (
	"errors"

	"github.com/fiqriardiansyah/user-shopping-api-golang/internal/entity"
	"github.com/fiqriardiansyah/user-shopping-api-golang/internal/helper"
	"github.com/fiqriardiansyah/user-shopping-api-golang/internal/model"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (u *UserUseCase) User(userId uuid.UUID) (*model.User, error) {
	user := entity.User{}
	userModel := model.User{}

	result := u.db.Where("id = ?", userId).Preload(clause.Associations).First(&user)
	if result.Error != nil || errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, helper.NotFound("User not found")
	}

	if err := copier.Copy(&userModel, &user); err != nil {
		return nil, helper.Internal(err.Error())
	}

	return &userModel, nil
}
