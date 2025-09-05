package usecase

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/fiqriardiansyah/user-shopping-api-golang/internal/constant"
	"github.com/fiqriardiansyah/user-shopping-api-golang/internal/entity"
	"github.com/fiqriardiansyah/user-shopping-api-golang/internal/helper"
	"github.com/fiqriardiansyah/user-shopping-api-golang/internal/model"
	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (u *AuthUseCase) Register(ctx context.Context, request model.AuthRegisterRequest) (*model.AuthRegisterResponse, error) {
	tx := u.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	response := &model.AuthRegisterResponse{}
	newUser := &entity.User{
		Email: request.Email,
		Name:  request.Name,
	}

	userExists, err := u.AuthRepository.GetByEmail(tx, request.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, helper.Internal(err.Error())
	}

	if userExists != nil {
		return nil, helper.BadRequest(fmt.Sprintf("User with this email %s has already registered", request.Email))
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, helper.Internal("Could not generate password")
	}

	newUser.Password = string(passwordHash)

	if err := u.AuthRepository.Create(tx, newUser); err != nil {
		return nil, helper.Internal(err.Error())
	}

	role := &entity.Role{}
	roleResult := tx.Where("name = ? ", constant.RoleBuyer).Find(role)
	if roleResult.Error != nil || errors.Is(roleResult.Error, gorm.ErrRecordNotFound) {
		return nil, helper.Internal(roleResult.Error.Error())
	}

	userRole := &entity.UserRole{
		RoleID: role.ID,
		UserID: newUser.ID,
	}

	if err := tx.Create(userRole).Error; err != nil {
		return nil, helper.Internal(err.Error())
	}

	if err := copier.Copy(response, newUser); err != nil {
		tx.Rollback()
		return nil, helper.Internal(err.Error())
	}

	accessTokenParam := helper.GenerateTokenParam{
		UserId:   newUser.ID,
		Email:    newUser.Email,
		Duration: time.Now().Add((60 * 24) * time.Minute),
		Secret:   os.Getenv("JWT_SECRET"),
		Roles: []string{
			string(constant.RoleBuyer),
		},
	}

	refreshTokenParam := helper.GenerateTokenParam{
		UserId:   newUser.ID,
		Email:    newUser.Email,
		Duration: time.Now().Add((24 * time.Hour) * 30),
		Secret:   os.Getenv("JWT_REFRESH_TOKEN_SECRET"),
		Roles: []string{
			string(constant.RoleBuyer),
		},
	}

	accessToken, errAt := helper.GenerateToken(accessTokenParam)
	if errAt != nil {
		tx.Rollback()
		return nil, helper.Internal("Could not generate access token")
	}

	refreshToken, errRt := helper.GenerateToken(refreshTokenParam) // 1 month
	if errRt != nil {
		tx.Rollback()
		return nil, helper.Internal("Could not generate refresh token")
	}

	if err := tx.Commit().Error; err != nil {
		return nil, helper.Internal(err.Error())
	}

	response.AccessToken = accessToken
	response.RefreshToken = refreshToken

	return response, nil
}
