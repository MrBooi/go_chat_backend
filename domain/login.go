package domain

import "golang.org/x/net/context"

type LoginRequest struct {
	Uuid string `json:"uuid" binding:"required"`
}

type LoginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type LoginUsecase interface {
	GetByID(c context.Context, id string) (User, error)
	GetUserByUuid(c context.Context, uuid string) (User, error)
	CreateAccessToken(user *User, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *User, secret string, expiry int) (refreshToken string, err error)
}
