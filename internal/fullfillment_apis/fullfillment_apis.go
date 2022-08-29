package fullfillment_apis

import (
	"net/http"
)

type (
	FullfillmentAPIS struct {
		client *http.Client
	}
)

func New(client *http.Client) *FullfillmentAPIS {
	return &FullfillmentAPIS{
		client: client,
	}
}
