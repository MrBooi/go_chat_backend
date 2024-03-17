package domain

import "context"

type RegisterRequest struct {
	Uuid     string `json:"uuid" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Name     string `json:"name" binding:"required"`
	PhotoUrl string `json:"photo_url" binding:"required"`
}

type RegisterResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type RegisterUsecase interface {
	Create(c context.Context, user *User) error
	GetUserByUuidOrEmail(c context.Context, uuid string, email string) (User, error)
	CreateAccessToken(user *User, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *User, secret string, expiry int) (refreshToken string, err error)
}
