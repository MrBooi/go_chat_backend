package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/net/context"
)

type Message struct {
	ID             primitive.ObjectID `bson:"_id" json:"id" `
	SenderId       primitive.ObjectID `json:"sender_id" binding:"required"`
	ReceiverId     primitive.ObjectID `json:"receiver_id"`
	Content        string             `json:"content" binding:"required"`
	MessageType    string             `json:"message_type" binding:"required"`
	ParentId       primitive.ObjectID `json:"parent_id"`
	ConversationId string             `json:"conversation_id" binding:"required"`
}

type MessageRepository interface {
	Create(c context.Context, body *Message) (Message, error)
	ConversationMessages(c context.Context, id string, conversationId string, options PaginationOptions) ([]Message, error)
	PrivateMessageList(c context.Context, id string, options PaginationOptions) ([]Message, error)
}

type MessageUsecase interface {
	Create(c context.Context, body *Message) (Message, error)
	ConversationMessages(c context.Context, id string, conversationId string) ([]Message, error)
	PrivateMessageList(c context.Context, id string) ([]Message, error)
}
