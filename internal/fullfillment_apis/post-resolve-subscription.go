package fullfillment_apis

import (
	"net/http"

	"github.com/fluffy-bunny/echo_mini_mona_saas_app/internal"
	"github.com/fluffy-bunny/echo_mini_mona_saas_app/models"
	"github.com/gogo/status"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
)

var (
	// https://learn.microsoft.com/en-us/azure/marketplace/partner-center-portal/pc-saas-fulfillment-subscription-api#resolve-a-purchased-subscription
	ResolveSubscriptionURL = "https://marketplaceapi.microsoft.com/api/saas/subscriptions/resolve?api-version=2018-08-31"
)

func (s *Service) ReloveSubscription(token string) (*models.ResolvedSubsription, error) {

	req, err := http.NewRequest("POST", ResolveSubscriptionURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header = http.Header{
		"content-type":           {"application/json"},
		"x-ms-marketplace-token": {token},
	}

	httpResponse, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	log.Info().Int("StatusCode", httpResponse.StatusCode).Msg("Resolve Subscription")
	if httpResponse.StatusCode != http.StatusOK {
		return nil, status.Errorf(codes.NotFound, "Subscription not found, httpStatusCode: %d", httpResponse.StatusCode)
	}
	record := &models.ResolvedSubsription{}
	err = internal.SafeUnmarshalFromHttpResponse(httpResponse, record)
	if err != nil {
		log.Error().Err(err).Msg("Failed to unmarshal body")
		return nil, err
	}
	return record, nil
}
