package repository

import (
	"context"

	"github.com/MrBooi/go_chat_backend/domain"
	"github.com/MrBooi/go_chat_backend/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

type userRepository struct {
	database   mongo.Database
	collection string
}

func (u *userRepository) GetUserByUuid(c context.Context, uuid string) (domain.User, error) {
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

func (u *userRepository) GetUserByUuidOrEmail(c context.Context, uuid string, email string) (domain.User, error) {
	collection := u.database.Collection(u.collection)

	var user domain.User

	pipeline := bson.M{
		"uuid": uuid,
		"$or": []interface{}{
			bson.M{"email": email},
		},
	}

	err := collection.FindOne(c, pipeline).Decode(&user)
	if err != nil {
		return user, err
	}
	return user, err
}

func NewUserRepository(db mongo.Database, collection string) domain.UserRepository {
	return &userRepository{
		database:   db,
		collection: collection,
	}
}
