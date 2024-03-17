package domain

import "golang.org/x/net/context"

type LoginRequest struct {
	Uuid string `json:"uuid" binding:"required"`
}

type LoginUsecase interface {
	GetUserByUuid(c context.Context, uuid string) (User, error)
}
