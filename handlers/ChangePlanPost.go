package handlers

import (
	"net/http"

	"github.com/fluffy-bunny/echo_mini_mona_saas_app/models"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

// ChangePlanPost - suspend - POST
func (c *Container) ChangePlanPost(ctx echo.Context) error {
	log := log.With().Caller().Str("func", "ChangePlanPost").Logger()
	log.Info().Msg("ChangePlanPost")
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}
