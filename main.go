package main

import (
	"bytes"
	"fmt"

	"github.com/fatih/structs"
	handlers_fullfillment "github.com/fluffy-bunny/echo_mini_mona_saas_app/handlers/fullfillment"
	handlers_landing "github.com/fluffy-bunny/echo_mini_mona_saas_app/handlers/landing"
	handlers_webhook "github.com/fluffy-bunny/echo_mini_mona_saas_app/handlers/webhooks"

	"github.com/fluffy-bunny/echo_mini_mona_saas_app/internal"
	"github.com/fluffy-bunny/viperEx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

// https://docs.microsoft.com/en-us/azure/marketplace/partner-center-portal/pc-saas-fulfillment-webhook
func main() {
	err := ReadViperConfig(internal.ConfigDefaultJSON, internal.AppConfig)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to read config")
	}
	err = internal.MakeAzureHttpClient()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to make azure http client")
	}

	// test call to get all our subscriptions
	internal.DumpGetSubscriptions()

	e := echo.New()

	//todo: handle the error!
	landingHandlers, _ := handlers_landing.NewContainer()
	webhooksHandlers, _ := handlers_webhook.NewContainer()
	fullfillmentHandlers, _ := handlers_fullfillment.NewContainer()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Webhooks
	//---------------------------------------------------------

	// SuspendPost - suspend - POST
	e.POST("/api/:version/saas/webhook", webhooksHandlers.WebhookPost)

	// ChangePlanPost - change-plan POST
	e.POST("/api/:version/saas/change-plan", webhooksHandlers.ChangePlanPost)
	// ChangeQuantityPost - change-quantity POST
	e.POST("/api/:version/saas/change-quantity", webhooksHandlers.ChangeQuantityPost)
	// RenewPost - renew POST
	e.POST("/api/:version/saas/renew", webhooksHandlers.RenewPost)
	// SuspendPost - suspend POST
	e.POST("/api/:version/saas/suspend", webhooksHandlers.SuspendPost)
	// UnsubscribePost - unsubscribe POST
	e.POST("/api/:version/saas/unsubscribe", webhooksHandlers.UnsubscribePost)
	// ReinstatePost - reinstate POST
	e.POST("/api/:version/saas/reinstate", webhooksHandlers.ReinstatePost)
	// SuspendPost - suspend - POST
	e.POST("/api/:version/saas/suspend", webhooksHandlers.SuspendPost)

	// Landing Pages
	//---------------------------------------------------------
	// HomeGet - home
	e.GET("/", landingHandlers.HomeGet)

	// SignupGet - purchased
	e.GET("/api/:version/saas/signup", landingHandlers.SignupGet)
	// PurchasedGet - purchased
	e.GET("/api/:version/saas/purchased/:subscription_id", landingHandlers.PurchasedGet)
	// ConfigureGet - configure
	e.GET("/api/:version/saas/configure/:subscription_id", landingHandlers.ConfigureGet)

	// Usefull SaaS APIs
	//---------------------------------------------------------
	// GetSubscription - get subscription
	e.GET("/api/:version/saas/subscription/:subscription_id", fullfillmentHandlers.GetSubscription)
	// GetSubscriptions - get subscriptions
	e.GET("/api/:version/saas/subscriptions", fullfillmentHandlers.GetSubscriptions)
	// Activate Subscription - POST
	e.POST("/api/:version/saas/subscription/:subscription_id/activate", fullfillmentHandlers.ActivateSubscriptionPost)

	// Start server
	//---------------------------------------------------------
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", internal.AppConfig.Port)))
}

// ReadViperConfig initial read
func ReadViperConfig(defaultConfig []byte, dst interface{}) error {
	v := viper.NewWithOptions(viper.KeyDelimiter("__"))
	var err error
	v.SetConfigType("json")
	// Environment Variables override everything.
	v.AutomaticEnv()

	// 1. Read in as buffer to set a default baseline.
	err = v.ReadConfig(bytes.NewBuffer(defaultConfig))
	if err != nil {
		log.Err(err).Msg("defaultConfig did not read in")
		return err
	}
	// we need to do a viper Unmarshal because that is the only way we get the
	// ENV variables to come in
	err = v.Unmarshal(dst)
	if err != nil {
		return err
	}
	// we do all settings here, becuase a v.AllSettings will NOT bring in the ENV variables
	structs.DefaultTagName = "mapstructure"
	allSettings := structs.Map(dst)

	// normal viper stuff
	myViperEx, err := viperEx.New(allSettings, func(ve *viperEx.ViperEx) error {
		ve.KeyDelimiter = "__"
		return nil
	})
	if err != nil {
		return err
	}
	myViperEx.UpdateFromEnv()
	err = myViperEx.Unmarshal(dst)
	return err
}
