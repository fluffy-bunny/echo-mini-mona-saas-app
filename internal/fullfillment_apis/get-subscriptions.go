package fullfillment_apis

import (
	"github.com/fluffy-bunny/echo_mini_mona_saas_app/internal"
	"github.com/fluffy-bunny/echo_mini_mona_saas_app/models"
	"github.com/rs/zerolog/log"
)

var (
	// https://docs.microsoft.com/en-us/azure/marketplace/partner-center-portal/pc-saas-fulfillment-subscription-api#get-httpsmarketplaceapimicrosoftcomapisaassubscriptionssubscriptionidapi-versionapiversion
	GetSubscriptionsURL = "https://marketplaceapi.microsoft.com/api/saas/subscriptions?api-version=2018-08-31"
)

func (s *FullfillmentAPIS) GetSubscriptions() ([]models.Subscription, error) {
	httpResponse, err := s.client.Get(GetSubscriptionsURL)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get subscriptions")
		return nil, err
	}
	log.Info().Int("StatusCode", httpResponse.StatusCode).Msg("Get Subscriptions")
	record := &models.Subscriptions{}
	err = internal.UnmarshalFromRequestBody(httpResponse.Body, record)
	if err != nil {
		log.Error().Err(err).Msg("Failed to unmarshal body")
		return nil, err
	}
	return record.Subscriptions, nil
}
