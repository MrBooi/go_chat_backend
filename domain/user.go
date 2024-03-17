package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionUser = "users"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id" json:"id" `
	Uuid     string             `bson:"uuid" json:"uuid" `
	Name     string             `bson:"name" json:"name" `
	Email    string             `bson:"email" json:"email" `
	PhotoUrl string             `bson:"photo_url" json:"photo_url" `
}

type UserRepository interface {
	Create(c context.Context, user *User) error
	GetUserByUuidOrEmail(c context.Context, uuid string, email string) (User, error)
	GetUserByUuid(c context.Context, uuid string) (User, error)
}
