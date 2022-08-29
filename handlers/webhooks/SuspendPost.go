package webhooks

import (
	"net/http"

	"github.com/fluffy-bunny/echo_mini_mona_saas_app/internal"
	"github.com/fluffy-bunny/echo_mini_mona_saas_app/models"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

// SuspendPost - POST
func (c *Container) SuspendPost(ctx echo.Context) error {
	log := log.With().Caller().Str("func", "SuspendPost").Logger()

	var err error
	record := &models.Suspend{}
	err = internal.UnmarshalFromRequestBody(ctx.Request().Body, record)
	if err != nil {
		log.Error().Err(err).Msg("Failed to unmarshal body")
		return ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
	}
	err = internal.SendOperationUpdateSuccess(record.SubscriptionID, record.ID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to update operation")
		return ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
	}
	return ctx.JSON(http.StatusOK, record)
}
