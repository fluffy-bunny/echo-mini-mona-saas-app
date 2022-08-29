package webhooks

import (
	"net/http"

	"github.com/fluffy-bunny/echo_mini_mona_saas_app/internal"
	"github.com/fluffy-bunny/echo_mini_mona_saas_app/models"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

// UnsubscribePost - POST
// 1. we got the unsubscribe request
// 2. we update the operation status to "Success" or "Failure"
// This request should be immediately offloaded to a workflow engine like temporal that would eventually make the operations callback
// using the operations api to set the status as "Success" or "Failure"
func (c *Container) UnsubscribePost(ctx echo.Context) error {
	log := log.With().Caller().Str("func", "UnsubscribePost").Logger()

	var err error
	record := &models.Unsubscribe{}
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
