package fullfillment_apis

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fluffy-bunny/echo_mini_mona_saas_app/models"
	"github.com/rs/zerolog/log"
)

var (
	// https://docs.microsoft.com/en-us/azure/marketplace/partner-center-portal/pc-saas-fulfillment-subscription-api#post-httpsmarketplaceapimicrosoftcomapisaassubscriptionssubscriptionidactivateapi-versionapiversion
	ActivateSubscriptionURL = "https://marketplaceapi.microsoft.com/api/saas/subscriptions/%s/activate?api-version=2018-08-31"
)

func (s *FullfillmentAPIS) ActivateSubscription(id string, request *models.ActivateRequest) error {
	url := fmt.Sprintf(ActivateSubscriptionURL, id)

	body, _ := json.Marshal(request)

	resp, err := s.client.Post(url, "application/json", bytes.NewBuffer(body))

	if err != nil {
		log.Error().Err(err).Msg("Failed to activate subscription")
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return nil
	}
	return fmt.Errorf("failed to activate subscription. status Code: %d", resp.StatusCode)

}
