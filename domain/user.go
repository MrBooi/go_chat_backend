package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionUser = "users"
)

type User struct {
	ID              primitive.ObjectID `bson:"_id" json:"id" `
	Uuid            string             `bson:"uuid" json:"uuid" `
	Name            string             `bson:"name" json:"name" `
	Email           string             `bson:"email" json:"email" `
	PhotoUrl        string             `bson:"photo_url" json:"photo_url"`
	IsActive        bool               `bson:"is_active" json:"is_active" default:"true"`
	IsTermsAccepted bool               `bson:"is_terms_accepted" json:"is_terms_accepted"`
}

type UpdateUserRequest struct {
	Name            string `bson:"name" json:"name" `
	PhotoUrl        string `bson:"photo_url" json:"photo_url"`
	IsActive        bool   `bson:"is_active" json:"is_active" default:"true"`
	IsTermsAccepted bool   `bson:"is_terms_accepted" json:"is_terms_accepted"`
}

type UserRepository interface {
	Create(c context.Context, user *User) error
	GetByID(c context.Context, id string) (User, error)
	GetUserByUuidOrEmail(c context.Context, uuid string, email string) (User, error)
	GetUserByUuid(c context.Context, uuid string) (User, error)
	UpdateUser(c context.Context, id string, body UpdateUserRequest) (User, error)
}

type UserUsecase interface {
	UpdateUser(c context.Context, id string, body UpdateUserRequest) (User, error)
	GetProfile(c context.Context, id string) (User, error)
}
