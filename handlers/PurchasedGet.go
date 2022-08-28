package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fluffy-bunny/echo_mini_mona_saas_app/internal"
	"github.com/fluffy-bunny/echo_mini_mona_saas_app/models"
	"github.com/go-resty/resty/v2"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

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
	subscriptionInfoFmt = "https://%s.blob.core.windows.net%s"
)

type params struct {
	SubscriptionID string `param:"subscription_id" query:"subscription_id" header:"subscription_id" form:"subscription_id" json:"subscription_id" xml:"subscription_id"`
	SubQuery       string `param:"_sub" query:"_sub" header:"_sub" form:"_sub" json:"_sub" xml:"_sub"`
}

// PurchasedGet - purchased
func (c *Container) PurchasedGet(ctx echo.Context) error {
	log := log.With().Caller().Str("func", "PurchasedGet").Logger()
	u := new(params)
	if err := ctx.Bind(u); err != nil {
		return err
	}
	log = log.With().Str("subscription_id", u.SubscriptionID).Logger()
	log.Info().Send()
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})

	subInfoUrl := fmt.Sprintf(subscriptionInfoFmt, internal.AppConfig.MonaStorageAccountName, u.SubQuery)
	log.Info().Str("subInfoUrl", subInfoUrl).Send()

	client := resty.New()
	resp, err := client.R().
		EnableTrace().
		Get(subInfoUrl)
	if err != nil {
		log.Error().Err(err).Send()
		return ctx.JSON(http.StatusBadRequest, models.HelloWorld{
			Message: "failed_to_get_subscription_info",
		})
	}
	statusCode := resp.StatusCode()
	if statusCode != http.StatusOK {
		log.Info().Int("status_code", resp.StatusCode()).Str("body", string(resp.Body())).Send()
		return ctx.JSON(http.StatusBadRequest, models.HelloWorld{
			Message: "failed_to_get_subscription_info - status_code:" + string(statusCode),
		})
	}

	subInfo := &models.SubscriptionInfo{}
	err = json.Unmarshal(resp.Body(), &subInfo)
	if err != nil {
		log.Error().Err(err).Send()
		return ctx.JSON(http.StatusBadRequest, models.HelloWorld{
			Message: "failed_to_unmarshal_subscription_info",
		})

		return c.Redirect(http.StatusFound, wellknown.ErrorPath+"?error=failed_to_unmarshal_subscription_info")
	}
	log.Info().Interface("subInfo", subInfo).Send()
	subInfoJson, _ := json.Marshal(subInfo)

}
