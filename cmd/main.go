package main

import (
	"fmt"

	"github.com/MrBooi/go_chat_backend/bootstrap"
)

func main() {
	app := bootstrap.App()

	env := app.Env

	fmt.Print(env.AppEnv)
}
