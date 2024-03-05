package route

import (
	"time"

	"github.com/MrBooi/go_chat_backend/bootstrap"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs"
)

func SetupSwagger(env *bootstrap.Env, timeout time.Duration, gin *gin.Engine) {
	docs.SwaggerInfo.BasePath = "/api/v1"
	ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	gin.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
