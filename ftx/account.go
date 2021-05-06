package ftx

import (
	"github.com/murlokito/ccex/common"
	"github.com/murlokito/ccex/ftx/rest"
)

type (

	// AccountClient represents the client for the FTX Account API.
	AccountClient struct {
		client *rest.Client
	}
)

func (a AccountClient) GetAccount() (common.Response, error) {
	panic("implement me")
}

func (a AccountClient) GetPositions(showAvgPrice bool) (common.Response, error) {
	panic("implement me")
}

func (a AccountClient) GetSubAccounts() (common.Response, error) {
	panic("implement me")
}

func (a AccountClient) GetSubAccountBalance(subAccount string) (common.Response, error) {
	panic("implement me")
}

func (a AccountClient) PostSubAccountNameChange(subAccountName, newSubAccountName string) (common.Response, error) {
	panic("implement me")
}

func (a AccountClient) PostSubAccountAssetTransfer(coin, sourceSubAccount, targetSubAccount string, size int) (common.Response, error) {
	panic("implement me")
}

func (a AccountClient) PostAccountLeverageChange(leverage int) (common.Response, error) {
	panic("implement me")
}

func (a AccountClient) PostFuturesAccountLeverageChange(symbol string, leverage int) (common.Response, error) {
	panic("implement me")
}

// NewAccountClient returns a new configured account client
func NewAccountClient(client *rest.Client) (*AccountClient, error) {
	return &AccountClient{
		client: client,
	}, nil
}
