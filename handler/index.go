package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/mackswell13/go-example/views/user"
)

type IndexHandler struct {

}

func(h IndexHandler) HandleIndexRender(c echo.Context) error {
	return render(c, )
}
