package route

import (
	"time"

	"github.com/MrBooi/go_chat_backend/bootstrap"
	"github.com/MrBooi/go_chat_backend/mongo"
	"github.com/gin-gonic/gin"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db mongo.Database, gin *gin.Engine) {
	publicRouter := gin.Group("")
	// All Public APIs
	healthCheckRouter(publicRouter)

}

func healthCheckRouter(group *gin.RouterGroup) {
	group.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "All systems work fine.",
		})
	})
}
