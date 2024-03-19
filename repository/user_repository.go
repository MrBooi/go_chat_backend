package repository

import (
	"context"

	"github.com/MrBooi/go_chat_backend/domain"
	"github.com/MrBooi/go_chat_backend/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type userRepository struct {
	database   mongo.Database
	collection string
}

func (u *userRepository) UpdateUser(c context.Context, id string, body domain.UpdateUserRequest) (domain.User, error) {
	collection := u.database.Collection(u.collection)

	var user domain.User
	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return user, err
	}
	filter := bson.M{"_id": idHex}

	update := bson.D{{Key: "$set", Value: body}}

	_, err = collection.UpdateOne(c, filter, update, options.Update().SetUpsert(true))

	if err != nil {
		return user, err
	}
	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&user)
	if err != nil {
		return user, err
	}

	return user, nil
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
