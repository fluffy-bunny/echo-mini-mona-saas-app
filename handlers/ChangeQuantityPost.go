package handlers

import (
	"net/http"

	"github.com/fluffy-bunny/echo_mini_mona_saas_app/models"
	"github.com/labstack/echo/v4"
)

// ChangeQuantityPost - suspend - POST
func (c *Container) ChangeQuantityPost(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}
