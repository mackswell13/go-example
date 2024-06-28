package main

import (
	"github.com/labstack/echo/v4"
	"github.com/mackswell13/go-example/handler"
)

func main() {
	app := echo.New()

	userHandler := handler.UserHandler{}

	app.GET("/user", userHandler.HandleUserShow)

	app.Start(":3000")
}
