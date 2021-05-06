package ftx

import "github.com/murlokito/ccex/ftx/rest"

type (

	// OptionsClient represents the client for the FTX Options API.
	OptionsClient struct {
		client *rest.Client
	}
)

// NewOptionsClient returns a new configured account client
func NewOptionsClient(client *rest.Client) (*OptionsClient, error) {
	return &OptionsClient{
		client: client,
	}, nil
}

