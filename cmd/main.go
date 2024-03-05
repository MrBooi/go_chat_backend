package main

import (
	"time"

	"github.com/MrBooi/go_chat_backend/api/route"
	"github.com/MrBooi/go_chat_backend/bootstrap"
	"github.com/gin-gonic/gin"
)

func main() {
	app := bootstrap.App()

	env := app.Env

	timeout := time.Duration(env.ContextTimeout) * time.Second

	gin := gin.Default()

	route.Setup(env, timeout, gin)

	gin.Run(env.ServerAddress)
}
