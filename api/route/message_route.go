package route

import (
	"time"

	"github.com/MrBooi/go_chat_backend/api/controller"
	"github.com/MrBooi/go_chat_backend/bootstrap"
	"github.com/MrBooi/go_chat_backend/domain"
	"github.com/MrBooi/go_chat_backend/mongo"
	"github.com/MrBooi/go_chat_backend/repository"
	"github.com/MrBooi/go_chat_backend/usecase"
	"github.com/gin-gonic/gin"
)

func NewMessageRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	mr := repository.NewMessageRepository(db, domain.CollectionUser)
	mc := controller.MessageController{
		MessageUsecase: usecase.NewMessageUsecase(mr, timeout),
		Env:            env,
	}

	group.POST("/users/:id/messages", mc.AddPrivateMessage)
	group.GET("/users/:id/messages", mc.UserConversations)
	group.GET("/messages/private", mc.PrivateMessages)

}
