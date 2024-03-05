package main

import (
	"fmt"
	"time"

	"github.com/MrBooi/go_chat_backend/api/route"
	"github.com/MrBooi/go_chat_backend/bootstrap"
	"github.com/gin-gonic/gin"
)

func main() {
	app := bootstrap.App()

	env := app.Env
	db := app.Mongo.Database(env.DBName)
	defer app.CloseDBConnection()

	timeout := time.Duration(env.ContextTimeout) * time.Second

	gin := gin.Default()

	// TODO for testing set trusted proxies to [127.0.0.1]
	gin.ForwardedByClientIP = true
	err := gin.SetTrustedProxies([]string{"127.0.0.1"})

	// Setup swagger doc

	route.Setup(env, timeout, db, gin)
	route.SetupSwagger(env, timeout, gin)

	err = gin.Run(env.ServerAddress)

	if err != nil {
		fmt.Println(err)
	}
}
