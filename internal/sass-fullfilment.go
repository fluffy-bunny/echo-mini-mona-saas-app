package internal

import (
	"fmt"
	"net/http/httputil"

	"github.com/rs/zerolog/log"
)

// https://docs.microsoft.com/en-us/azure/marketplace/partner-center-portal/pc-saas-fulfillment-subscription-api
var (
	ActivateURLFormat   = "https://marketplaceapi.microsoft.com/api/saas/subscriptions/%s/activate?api-version=2018-08-31"
	GetSubscriptionsURL = "https://marketplaceapi.microsoft.com/api/saas/subscriptions?api-version=2018-08-31"
)

func DumpGetSubscriptions() {
	httpResponse, err := AzureHttpClient.Get(GetSubscriptionsURL)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to get subscriptions")
	}
	log.Info().Int("StatusCode", httpResponse.StatusCode).Msg("Get Subscriptions")
	respDump, err := httputil.DumpResponse(httpResponse, true)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to dump response")
	}
	fmt.Printf("RESPONSE:\n%s", string(respDump))
}
