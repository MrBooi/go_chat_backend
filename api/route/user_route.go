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

func NewUserRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	uc := controller.UserController{
		UserUsecase: usecase.NewUserUsecase(ur, timeout),
		Env:         env,
	}
	// profile
	group.GET("user/profile", uc.Profile)
	group.PUT("user/update", uc.UpdateUser)

}
