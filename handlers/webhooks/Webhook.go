package webhooks

import (
	"encoding/json"
	"net/http"

	"github.com/fluffy-bunny/echo_mini_mona_saas_app/internal"
	"github.com/fluffy-bunny/echo_mini_mona_saas_app/models"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

func (c *Container) getWebhookFromRequest(ctx echo.Context) (*models.Webhook, error) {
	log := log.With().Caller().Str("func", "getWebhookFromRequest").Logger()

	var err error
	bodyBytes, err := internal.SafeGetBodyFromHttpRequest(ctx.Request())
	if err != nil {
		log.Error().Err(err).Msg("Failed to get body")
		return nil, err
	}
	metadata := &models.WebhookMetadata{}
	err = json.Unmarshal(bodyBytes, metadata)
	if err != nil {
		log.Error().Err(err).Msg("Failed to unmarshal body")
		return nil, err
	}
	webhook := &models.Webhook{
		Metadata: metadata,
	}
	switch metadata.Action {
	case "ChangePlan":
		changePlanWebhook := &models.ChangePlanWebhook{}
		err = json.Unmarshal(bodyBytes, changePlanWebhook)
		if err != nil {
			log.Error().Err(err).Msg("Failed to unmarshal body")
			return nil, err
		}
		webhook.SetChangePlanWebhook(changePlanWebhook)
	case "ChangeQuantity":
		changeQuantityWebhook := &models.ChangeQuantityWebhook{}
		err = json.Unmarshal(bodyBytes, changeQuantityWebhook)
		if err != nil {
			log.Error().Err(err).Msg("Failed to unmarshal body")
			return nil, err
		}
		webhook.SetChangeQuantityWebhook(changeQuantityWebhook)
	case "Suspend":
		suspendWebhook := &models.SuspendWebhook{}
		err = json.Unmarshal(bodyBytes, suspendWebhook)
		if err != nil {
			log.Error().Err(err).Msg("Failed to unmarshal body")
			return nil, err
		}
		webhook.SetSuspendWebhook(suspendWebhook)
	case "Reinstate":
		reinstateWebhook := &models.ReinstateWebhook{}
		err = json.Unmarshal(bodyBytes, reinstateWebhook)
		if err != nil {
			log.Error().Err(err).Msg("Failed to unmarshal body")
			return nil, err
		}
		webhook.SetReinstateWebhook(reinstateWebhook)
	case "Renew":
		renewWebhook := &models.RenewWebhook{}
		err = json.Unmarshal(bodyBytes, renewWebhook)
		if err != nil {
			log.Error().Err(err).Msg("Failed to unmarshal body")
			return nil, err
		}
		webhook.SetRenewWebhook(renewWebhook)
	case "Unsubscribe":
		unsubscribeWebhook := &models.UnsubscribeWebhook{}
		err = json.Unmarshal(bodyBytes, unsubscribeWebhook)
		if err != nil {
			log.Error().Err(err).Msg("Failed to unmarshal body")
			return nil, err
		}
		webhook.SetUnsubscribeWebhook(unsubscribeWebhook)
	}

	return webhook, nil
}

// WebhookPost - POST
func (c *Container) WebhookPost(ctx echo.Context) error {
	log := log.With().Caller().Str("func", "WebhookPost").Logger()

	var err error
	webhook, err := c.getWebhookFromRequest(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get webhook from request")
		return ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
	}
	log.Info().Interface("webhook", webhook).Msg("webhook")

	return ctx.JSON(http.StatusOK, webhook)
}
