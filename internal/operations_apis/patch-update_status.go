package operations_apis

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fluffy-bunny/echo_mini_mona_saas_app/models"
	"github.com/rs/zerolog/log"
)

var (
	// https://docs.microsoft.com/en-us/azure/marketplace/partner-center-portal/pc-saas-fulfillment-operations-api#update-the-status-of-an-operation
	operationsUpdateURL = "https://marketplaceapi.microsoft.com/api/saas/subscriptions/%s/operations/%s?api-version=2018-08-31"
)

func (s *Service) Update(subID, opID string, request *models.OperationsUpdateRequest) error {
	url := fmt.Sprintf(operationsUpdateURL, subID, opID)

	body, _ := json.Marshal(request)

	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := s.client.Do(req)

	if err != nil {
		log.Error().Err(err).Msg("Failed to update operations")
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return nil
	}
	log.Error().Int("StatusCode", resp.StatusCode).Msg("Failed to update operations")
	return fmt.Errorf("failed to activate subscription. status Code: %d", resp.StatusCode)
}
