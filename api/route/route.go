package route

import (
	"time"

	"github.com/MrBooi/go_chat_backend/api/middleware"
	"github.com/MrBooi/go_chat_backend/bootstrap"
	"github.com/MrBooi/go_chat_backend/mongo"
	"github.com/gin-gonic/gin"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db mongo.Database, gin *gin.Engine) {
	publicRouter := gin.Group("")
	// All Public APIs
	healthCheckRouter(publicRouter)
	NewRegisterRouter(env, timeout, db, publicRouter)
	NewLoginRouter(env, timeout, db, publicRouter)
	NewRefreshTokenRouter(env, timeout, db, publicRouter)

	protectedRouter := gin.Group("")
	// middleware to verify if user has a valid accessToken
	protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))
	// private routes
	NewUserRouter(env, timeout, db, protectedRouter)
	NewMessageRouter(env, timeout, db, protectedRouter)

}

func healthCheckRouter(group *gin.RouterGroup) {
	// @Summary check the api health for our app
	// @Description Get the help of the endpoints
	// @Success 200 {string} string  "ok"
	// @Router /healthcheck/
	group.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "All systems work fine.",
		})
	})

}
