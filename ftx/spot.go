package ftx

import "github.com/murlokito/ccex/ftx/rest"

type (

	// SpotClient represents the client for the FTX Spot API.
	SpotClient struct {
		client *rest.Client
	}
)

// NewSpotClient returns a new configured account client
func NewSpotClient(client *rest.Client) (*SpotClient, error) {
	return &SpotClient{
		client: client,
	}, nil
}

