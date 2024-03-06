package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionUser = "users"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id"`
	Uuid     string             `bson:"uuid"`
	Name     string             `bson:"name"`
	Email    string             `bson:"email"`
	PhotoUrl string             `bson:"photo_url"`
}

type UserRepository interface {
	Create(c context.Context, user *User) error
	GetByID(c context.Context, id string) (User, error)
	GetUser(c context.Context, uuid string) (User, error)
}
