package controller

import (
	"net/http"

	"github.com/MrBooi/go_chat_backend/bootstrap"
	"github.com/MrBooi/go_chat_backend/domain"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MessageController struct {
	MessageUsecase domain.MessageUsecase
	Env            *bootstrap.Env
}

func (mc *MessageController) AddPrivateMessage(c *gin.Context) {
	var request domain.Message

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	message := domain.Message{
		ID:             primitive.NewObjectID(),
		ConversationId: request.ConversationId,
		ReceiverId:     request.ReceiverId,
		SenderId:       request.SenderId,
		Content:        request.Content,
		MessageType:    request.MessageType,
	}

	_, err = mc.MessageUsecase.Create(c, &message)

	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, "created")
}

func (mc *MessageController) UserConversations(c *gin.Context) {

}

func (mc *MessageController) PrivateMessages(c *gin.Context) {

}
