package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

// References:
//--------------------------------------------------
// https://github.com/microsoft/mona-saas/tree/main/docs#what-is-the-subscription-configuration-page

type configParams struct {
	SubscriptionID string `param:"subscription_id" query:"subscription_id" header:"subscription_id" form:"subscription_id" json:"subscription_id" xml:"subscription_id"`
}

// ConfigureGet - purchased
func (c *Container) ConfigureGet(ctx echo.Context) error {
	log := log.With().Caller().Str("func", "PurchasedGet").Logger()
	u := new(configParams)
	if err := ctx.Bind(u); err != nil {
		return err
	}
	log.Debug().Interface("params", u).Msg("params")

	return ctx.JSON(http.StatusOK, u)
}
