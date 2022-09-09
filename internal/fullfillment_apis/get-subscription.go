package fullfillment_apis

import (
	"fmt"

	"github.com/fluffy-bunny/echo_mini_mona_saas_app/internal"
	"github.com/fluffy-bunny/echo_mini_mona_saas_app/models"
	"github.com/rs/zerolog/log"
)

var (
	// https://docs.microsoft.com/en-us/azure/marketplace/partner-center-portal/pc-saas-fulfillment-subscription-api#get-httpsmarketplaceapimicrosoftcomapisaassubscriptionssubscriptionidapi-versionapiversion
	GetSubscriptionURL = "https://marketplaceapi.microsoft.com/api/saas/subscriptions/%s?api-version=2018-08-31"
)

func (s *Service) GetSubscription(id string) (*models.Subscription, error) {
	url := fmt.Sprintf(GetSubscriptionURL, id)
	httpResponse, err := s.client.Get(url)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get subscriptions")
		return nil, err
	}

	log.Info().Int("StatusCode", httpResponse.StatusCode).Msg("Get Subscriptions")
	record := &models.Subscription{}
	err = internal.SafeUnmarshalFromHttpResponse(httpResponse, record)
	if err != nil {
		log.Error().Err(err).Msg("Failed to unmarshal body")
		return nil, err
	}
	return record, nil
}
