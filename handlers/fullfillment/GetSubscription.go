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

type getSubscriptionParams struct {
	SubscriptionID string `param:"subscription_id" query:"subscription_id" header:"subscription_id" form:"subscription_id" json:"subscription_id" xml:"subscription_id"`
}

// GetSubscriptions - purchased
func (c *Container) GetSubscription(ctx echo.Context) error {
	log := log.With().Caller().Str("func", "GetSubscription").Logger()
	u := new(getSubscriptionParams)
	if err := ctx.Bind(u); err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
	}
	log.Debug().Interface("params", u).Msg("params")

	fullfillmentAPIs := fullfillment_apis.New(internal.AzureHttpClient)
	records, err := fullfillmentAPIs.GetSubscriptions()
	if err != nil {
		log.Error().Err(err).Msg("Failed to get subscriptions")
		return ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
	}
	return ctx.JSON(http.StatusOK, records)

}
