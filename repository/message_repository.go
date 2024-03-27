package repository

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/MrBooi/go_chat_backend/domain"
	"github.com/MrBooi/go_chat_backend/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type messageRepository struct {
	database   mongo.Database
	collection string
}

func (mr *messageRepository) PrivateMessageList(c context.Context, id string, options domain.PaginationOptions) ([]domain.Message, error) {
	collection := mr.database.Collection(mr.collection)

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("Invalid user Id")
	}

	// Default options
	opt := domain.NewPaginationOptions(
		options.Page,
		options.PerPage,
	)

	skip := opt.ApplyOptions(opt).CalculateSkip()

	match := bson.M{
		"type": "private",
		"$or": bson.A{
			bson.M{"sender_id": _id},
			bson.M{"receiver_id": _id},
		},
	}
	fmt.Println(skip, match)

	pipeline := bson.M{}
	ctx := context.TODO()
	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var messages []domain.Message
	if err := cursor.All(ctx, &messages); err != nil {
		return nil, err
	}

	return messages, nil
}

func (mr *messageRepository) ConversationMessages(c context.Context, id string, conversationId string, options domain.PaginationOptions) ([]domain.Message, error) {
	collection := mr.database.Collection(mr.collection)
	if id == "" || conversationId == "" {
		return nil, errors.New("Invalid ids")
	}
	_, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("Invalid user Id")
	}
	// Default options
	opt := domain.NewPaginationOptions(
		options.Page,
		options.PerPage,
	)

	skip := opt.ApplyOptions(opt).CalculateSkip()
	match := bson.M{"conversation_id": conversationId}
	pipeline := bson.A{
		bson.D{{Key: "$match", Value: match}},
		bson.D{{Key: "$skip", Value: skip}},
		bson.D{{Key: "$limit", Value: opt.PerPage}},
		bson.D{{Key: "$sort", Value: bson.D{{Key: "updatedAt", Value: -1}}}},
	}
	ctx := context.TODO() // Use context.Background() if you're not using Go 1.16+
	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var messages []domain.Message
	if err := cursor.All(ctx, &messages); err != nil {
		return nil, err
	}

	return messages, nil

}

func (mr *messageRepository) Create(c context.Context, message *domain.Message) (domain.Message, error) {
	collection := mr.database.Collection(mr.collection)

	var conversationID string
	if message.MessageType == "private" {
		ids := []string{message.SenderId.String(), message.ReceiverId.String()}
		sort.Strings(ids)
		conversationID = strings.Join(ids, "_")
	} else {
		conversationID = message.ReceiverId.String()
	}

	data := &domain.Message{
		SenderId:       message.SenderId,
		ReceiverId:     message.ReceiverId,
		Content:        message.Content,
		MessageType:    message.MessageType,
		ParentId:       message.ParentId,
		ConversationId: conversationID,
	}

	_, err := collection.InsertOne(c, data)

	return *data, err
}

func NewMessageRepository(db mongo.Database, collection string) domain.MessageRepository {
	return &messageRepository{
		database:   db,
		collection: collection,
	}
}
