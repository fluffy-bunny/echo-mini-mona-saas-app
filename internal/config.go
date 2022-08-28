package internal

/*
NOTE:
ENV overrides follow azure double underscore convention
i.e.
	AZURE_AD__CREDENTIALS__CLIENT_ID=<some value>
	AZURE_AD__CREDENTIALS__CLIENT_SECRET=<some value>
*/
type (
	OAuth2ClientCredentials struct {
		ClientID     string `json:"client_id" mapstructure:"CLIENT_ID"`
		ClientSecret string `json:"client_secret" mapstructure:"CLIENT_SECRET"`
	}
	// https://docs.microsoft.com/en-us/azure/marketplace/marketplace-metering-service-authentication
	AzureAD struct {
		Tenant      string                  `json:"tenant" mapstructure:"TENANT"`
		Credentials OAuth2ClientCredentials `json:"credentials" mapstructure:"CREDENTIALS"`
	}
	Config struct {
		Port    int     `json:"port" mapstructure:"PORT"`
		AzureAD AzureAD `json:"azure_ad" mapstructure:"AZURE_AD"`
	}
)

const (
	// https://login.microsoftonline.com/{{tenant_id}}/oauth2/token
	TokenEndpointFormat = "https://login.microsoftonline.com/%s/oauth2/v2.0/token"
)

var (
	ConfigDefaultJSON = []byte(`
	{
		"PORT": 1111,
		"AZURE_AD": {
			"TENANT": "<in-environemnt>",
			"CREDENTIALS": {
				"CLIENT_ID": "<in-environemnt>",
				"CLIENT_SECRET": "<in-environemnt>"
			}
		}
	}
	`)
)
