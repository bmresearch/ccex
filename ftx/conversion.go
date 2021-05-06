package ftx

import (
	"github.com/murlokito/ccex/common"
	"github.com/murlokito/ccex/ftx/rest"
)

type (

	// ConversionClient represents the client for the FTX Conversion API.
	ConversionClient struct {
		client *rest.Client
	}
)

func (c ConversionClient) PostConversionQuoteRequest(fromCoin, toCoin string, size float32) (common.Response, error) {
	panic("implement me")
}

func (c ConversionClient) GetConversionQuoteStatus(quoteId int, market string) (common.Response, error) {
	panic("implement me")
}

func (c ConversionClient) PostConversionQuoteAcceptance(quoteId int) (common.Response, error) {
	panic("implement me")
}

// NewConversionClient returns a new configured account client
func NewConversionClient(client *rest.Client) (*ConversionClient, error) {
	return &ConversionClient{
		client: client,
	}, nil
}
