/*
References:
https://docs.microsoft.com/en-us/azure/marketplace/partner-center-portal/pc-saas-registration
this tells you how to make a AD app, I used my Mono app that got deployed to azure when I created my clientid/secret

Also, this is there that resource = "20e940b3-4c77-4b0b-9a53-9e16a1b010a7" is documented.
*/
package internal

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/rs/zerolog/log"
	cc "golang.org/x/oauth2/clientcredentials"
)

const (
	// https://login.microsoftonline.com/{{tenant_id}}/oauth2/token
	TokenEndpointFormat = "https://login.microsoftonline.com/%s/oauth2/token"
)

var (
	AzureHttpClient *http.Client
)

func MakeAzureHttpClient() error {
	tokenEndpoint := fmt.Sprintf(TokenEndpointFormat, AppConfig.AzureAD.Tenant)
	_, err := url.ParseRequestURI(tokenEndpoint)
	if err != nil {
		log.Error().Err(err).Str("tokenEndpoint", tokenEndpoint).Msg("Failed to parse token endpoint")
		return err
	}
	config := &cc.Config{
		ClientID:     AppConfig.AzureAD.Credentials.ClientID,
		ClientSecret: AppConfig.AzureAD.Credentials.ClientSecret,
		TokenURL:     tokenEndpoint,
		EndpointParams: url.Values{
			// https://docs.microsoft.com/en-us/azure/marketplace/partner-center-portal/pc-saas-registration
			// "resource": []string{"20e940b3-4c77-4b0b-9a53-9e16a1b010a7"},
			//-----------------------------------------------
			"resource": []string{"20e940b3-4c77-4b0b-9a53-9e16a1b010a7"},
		},
	}
	ctx := context.Background()
	AzureHttpClient = config.Client(ctx)
	return nil
}
