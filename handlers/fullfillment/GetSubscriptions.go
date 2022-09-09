package fullfillment

import (
	"net/http"

	"github.com/fluffy-bunny/echo_mini_mona_saas_app/internal"
	"github.com/fluffy-bunny/echo_mini_mona_saas_app/internal/fullfillment_apis"
	"github.com/fluffy-bunny/echo_mini_mona_saas_app/models"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

// References:
//--------------------------------------------------

// GetSubscriptions -
func (c *Container) GetSubscriptions(ctx echo.Context) error {
	log := log.With().Caller().Str("func", "GetSubscriptions").Logger()
	fullfillmentAPIs := fullfillment_apis.New(internal.AzureHttpClient)
	records, err := fullfillmentAPIs.GetSubscriptions()
	if err != nil {
		log.Error().Err(err).Msg("Failed to get subscriptions")
		return ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
	}
	return ctx.JSON(http.StatusOK, records)

}
