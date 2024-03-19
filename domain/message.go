package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/net/context"
)

type Message struct {
	SenderId       primitive.ObjectID `json:"sender_id" binding:"required"`
	ReceiverId     primitive.ObjectID `json:"receiver_id"`
	Content        string             `json:"content" binding:"required"`
	MessageType    string             `json:"message_type" binding:"required"`
	ParentId       primitive.ObjectID `json:"parent_id" binding:"required"`
	ConversationId string             `json:"conversation_id" binding:"required"`
}

type MessageUsecase interface {
	Create(c context.Context, body Message) (Message, error)
	ConversationMessages(c context.Context, id string, conversationId string)
	PrivateMessageList(c context.Context, id string)
}
