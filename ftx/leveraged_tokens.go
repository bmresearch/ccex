package ftx

import "github.com/murlokito/ccex/ftx/rest"

type (

	// LeveragedTokensClient represents the client for the FTX Leveraged Tokens API.
	LeveragedTokensClient struct {
		client *rest.Client
	}
)

// NewLeveragedTokensClient returns a new configured account client
func NewLeveragedTokensClient(client *rest.Client) (*LeveragedTokensClient, error) {
	return &LeveragedTokensClient{
		client: client,
	}, nil
}

