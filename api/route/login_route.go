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

func NewLoginRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	lc := &controller.LoginController{
		LoginUsecase: usecase.NewLoginUsecase(ur, timeout),
		Env:          env,
	}
	group.POST("/login", lc.Login)
}
