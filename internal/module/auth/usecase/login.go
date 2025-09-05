package usecase

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/fiqriardiansyah/user-shopping-api-golang/internal/helper"
	"github.com/fiqriardiansyah/user-shopping-api-golang/internal/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (u *AuthUseCase) Login(ctx context.Context, request model.AuthLoginRequest) (*model.AuthLoginResponse, error) {
	response := &model.AuthLoginResponse{}

	userExists, err := u.AuthRepository.GetByEmail(u.db, request.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, helper.Internal(err.Error())
	}

	if userExists == nil {
		return nil, helper.NotFound(fmt.Sprintf("User with email %s not found", request.Email))
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userExists.Password), []byte(request.Password)); err != nil {
		return nil, helper.BadRequest("Email or password wrong!")
	}

	roles, err := u.AuthRepository.GetRoles(u.db, userExists.ID)
	if err != nil {
		return nil, helper.Internal(err.Error())
	}
	rolesStr := []string{}

	for _, r := range *roles {
		rolesStr = append(rolesStr, r.Name)
	}

	accessTokenParam := helper.GenerateTokenParam{
		UserId:   userExists.ID,
		Email:    userExists.Email,
		Duration: time.Now().Add((60 * 24) * time.Minute),
		Secret:   os.Getenv("JWT_SECRET"),
		Roles:    rolesStr,
	}

	refreshTokenParam := helper.GenerateTokenParam{
		UserId:   userExists.ID,
		Email:    userExists.Email,
		Duration: time.Now().Add((24 * time.Hour) * 30),
		Secret:   os.Getenv("JWT_REFRESH_TOKEN_SECRET"),
		Roles:    rolesStr,
	}

	accessToken, err := helper.GenerateToken(accessTokenParam)
	if err != nil {
		return nil, helper.Internal("Could not generate token")
	}
	refreshToken, err := helper.GenerateToken(refreshTokenParam) // 1 month
	if err != nil {
		return nil, helper.Internal("Could not generate refresh token")
	}

	response.Email = request.Email
	response.RefreshToken = refreshToken
	response.AccessToken = accessToken

	return response, nil
}
