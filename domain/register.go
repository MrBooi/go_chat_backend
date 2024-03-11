package domain

import "context"

type RegisterRequest struct {
	Uuid     string `json:"uuid" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Name     string `json:"name" binding:"required"`
	PhotoUrl string `json:"photo_url" binding:"required"`
}

type RegisterUsecase interface {
	Create(c context.Context, user *User) error
	GetByID(c context.Context, email string) (User, error)
}
