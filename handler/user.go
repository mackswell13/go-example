package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/mackswell13/go-example/views/user"
)

type UserHandler struct {

}

func(h UserHandler) HandleUserRender(c echo.Context) error {
	return render(c, user.Show())
}
