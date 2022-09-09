package landing

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (c *Container) HomeGet(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "HOME: echo-mini-mona-saas-app")
}
