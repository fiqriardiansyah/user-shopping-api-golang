package usecase

import (
	"os"
	"time"

	"github.com/fiqriardiansyah/user-shopping-api-golang/internal/helper"
	"github.com/fiqriardiansyah/user-shopping-api-golang/internal/model"
)

func (u *AuthUseCase) Refresh(refreshToken string) (*model.AuthLoginResponse, error) {
	claims, err := helper.ValidateToken(refreshToken, os.Getenv("JWT_REFRESH_TOKEN_SECRET"))
	if err != nil {
		return nil, helper.BadRequest(err.Error())
	}

	roles, err := u.AuthRepository.GetRoles(u.db, claims.UserId)
	if err != nil {
		return nil, helper.Internal(err.Error())
	}
	rolesStr := []string{}

	for _, r := range *roles {
		rolesStr = append(rolesStr, r.Name)
	}

	accessTokenParam := helper.GenerateTokenParam{
		UserId:   claims.UserId,
		Email:    claims.Email,
		Roles:    rolesStr,
		Duration: time.Now().Add((60 * 24) * time.Minute),
		Secret:   os.Getenv("JWT_SECRET"),
	}

	accessToken, errAt := helper.GenerateToken(accessTokenParam)
	if errAt != nil {
		return nil, helper.Internal("Could not generate access token")
	}

	return &model.AuthLoginResponse{
		Email:        claims.Email,
		RefreshToken: refreshToken,
		AccessToken:  accessToken,
	}, nil
}
