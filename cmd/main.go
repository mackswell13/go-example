package main

import (
	"github.com/labstack/echo/v4"
	"github.com/mackswell13/go-example/handler"
)

func main() {
	app := echo.New()

	userHandler := handler.UserHandler{}
	indexHandler := handler.IndexHandler{}

	app.GET("/user", userHandler.HandleUserRender)
	app.GET("/", indexHandler.HandleIndexRender)

	app.Start(":3000")
}
