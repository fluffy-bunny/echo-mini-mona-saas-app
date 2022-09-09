package fullfillment_apis

import (
	"net/http"
)

type (
	Service struct {
		client *http.Client
	}
)

func New(client *http.Client) *Service {
	return &Service{
		client: client,
	}
}
