package repository

import (
	"context"

	"github.com/MrBooi/go_chat_backend/domain"
	"github.com/MrBooi/go_chat_backend/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type userRepository struct {
	database   mongo.Database
	collection string
}

func (u *userRepository) GetUser(c context.Context, uuid string) (domain.User, error) {
	collection := u.database.Collection(u.collection)

	var user domain.User

	err := collection.FindOne(c, bson.M{"uuid": uuid}).Decode(&user)
	return user, err
}

func (u *userRepository) Create(c context.Context, user *domain.User) error {
	collection := u.database.Collection(u.collection)

	_, err := collection.InsertOne(c, user)

	return err
}

func (u *userRepository) GetByID(c context.Context, id string) (domain.User, error) {
	collection := u.database.Collection(u.collection)

	var user domain.User

	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return user, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&user)
	return user, err
}

func NewUserRepository(db mongo.Database, collection string) domain.UserRepository {
	return &userRepository{
		database:   db,
		collection: collection,
	}
}
