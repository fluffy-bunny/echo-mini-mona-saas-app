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
	// TEST: Unmarshal the body into a generic map to see what we are getting
	//-----------------------------------------------------------------------
	record := make(map[string]interface{})
	err = internal.SafeUnmarshalFromHttpRequest(ctx.Request(), &record)
	if err != nil {
		log.Error().Err(err).Msg("Failed to unmarshal body")
		return ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
	}
	log.Info().Interface("unsubscribe_record", record).Msg("UnsubscribePost")

	// Unmarshal the body into a UnsubscribeInto struct
	//-----------------------------------------------------------------------
	unsubscribeInfo := &models.UnsubscribeInfo{}
	err = internal.SafeUnmarshalFromHttpRequest(ctx.Request(), unsubscribeInfo)
	if err != nil {
		log.Error().Err(err).Msg("Failed to unmarshal body")
		return ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
	}
	log.Info().Interface("unsubscribeInfo", unsubscribeInfo).Msg("UnsubscribePost")

	// TODO: Kick off a workflow to clean up the subscription on our end.

	/*
		Getting a 400 with this
			err = internal.SendOperationUpdateSuccess(unsubscribeInfo.Subscription.SubscriptionID, unsubscribeInfo.OperationId)
			if err != nil {
				log.Error().Err(err).Msg("Failed to update operation")
				return ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
			}
	*/
	return ctx.JSON(http.StatusOK, record)
}
