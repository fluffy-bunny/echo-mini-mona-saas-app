package landing

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
// https://github.com/microsoft/mona-saas/tree/main/docs#what-is-the-subscription-purchase-confirmation-page

const (
// https://e9e3-172-90-216-122.ngrok.io/saas/purchased/426159d1-821c-4fe0-c981-7b7033bb0895?_sub=%2Fstage-subscriptions%2F426159d1-821c-4fe0-c981-7b7033bb0895%3Fsv%3D2020-04-08%26se%3D2022-08-26T17%253A15%253A13Z%26sr%3Db%26sp%3Dr%26sig%3DqZ57djvRdaPgKIw5kBia2CJi3FJNVSAXh%252F6NYrRFpII%253D
// our endpoint gets called with an url that looks like this.
// 1. get the query param _sub, that is then mapped to ${SUBSCRIPTION_INFO_PATH}
// when we make that GET call we get a JSON response that looks like the following
/*
	{
		"subscriptionId": "426159d1-821c-4fe0-c981-7b7033bb0895",
		"subscriptionName": "monatest",
		"offerId": "herb-offer-mona-2-preview",
		"planId": "test-plan-2",
		"isTest": false,
		"isFreeTrial": false,
		"seatQuantity": null,
		"status": 2,
		"term": {
			"termUnit": "P1M",
			"startDate": null,
			"endDate": null
		},
		"beneficiary": {
			"userId": "100320011A1F1D3F",
			"userEmail": "herb@mapped.com",
			"aadObjectId": "273b6470-e9ca-456e-837f-866ad2014fcc",
			"aadTenantId": "5905aed6-175d-4ef4-914d-3ee92a398d55"
		},
		"purchaser": {
			"userId": "100320011A1F1D3F",
			"userEmail": "herb@mapped.com",
			"aadObjectId": "273b6470-e9ca-456e-837f-866ad2014fcc",
			"aadTenantId": "5905aed6-175d-4ef4-914d-3ee92a398d55"
		}
	}
*/
// We have enough here to create a user in our system.
//"https://${STORAGE_ACCOUNT_NAME}.blob.core.windows.net${SUBSCRIPTION_INFO_PATH}"
//subscriptionInfoFmt = "https://%s.blob.core.windows.net%s"
)

type sigunupParams struct {
	Version string `param:"version" query:"version" header:"version" form:"version" json:"version" xml:"version"`
	Token   string `param:"token" query:"token" header:"token" form:"token" json:"token" xml:"token"`
}

// SignupGet - signup
func (c *Container) SignupGet(ctx echo.Context) error {
	log := log.With().Caller().Str("func", "PurchasedGet").Logger()
	u := new(sigunupParams)
	if err := ctx.Bind(u); err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
	}
	log = log.With().Interface("params", u).Logger()
	if len(u.Token) == 0 {
		return ctx.JSON(http.StatusBadRequest, "missing_params")
	}
	log.Info().Msg("got token")

	fa := fullfillment_apis.New(internal.AzureHttpClient)
	resolvedSubscription, err := fa.ReloveSubscription(u.Token)
	if err != nil {
		log.Error().Err(err).Msg("error resolving subscription")
		return ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
	}
	// Make a call back to the market place to activate the subscription
	//----------------------------------------

	err = fa.ActivateSubscription(resolvedSubscription.Subscription.ID, &models.ActivateRequest{
		PlanID:   resolvedSubscription.PlanID,
		Quantity: internal.StringPointer("1"),
	})
	if err != nil {
		log.Error().Err(err).Msg("error activating subscription")
		return ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
	}

	return ctx.JSON(http.StatusOK, u)

}
