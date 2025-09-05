package model

type AuthLoginRequest struct {
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"min=6"`
}

type AuthLoginResponse struct {
	Email        string `json:"email"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type AuthRegisterRequest struct {
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"min=6"`
	Name     string `json:"name" validate:"min=2"`
}

type AuthRegisterResponse struct {
	Email        string `json:"email"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
