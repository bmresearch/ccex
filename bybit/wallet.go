package bybit

import (
	"github.com/murlokito/ccex/bybit/rest"
	models "github.com/murlokito/ccex/models/rest"
)

type (
	WalletClient struct {
		client *rest.Client
	}
)

func (w WalletClient) GetWalletBalances(req *models.RequestForWalletBalances) (models.ResponseForWalletBalances, error) {
	panic("implement me")
}

func (w WalletClient) GetDepositAddress(req *models.RequestForDepositAddress) (models.ResponseForDepositAddress, error) {
	panic("implement me")
}

func (w WalletClient) GetWalletDepositHistory(req *models.RequestForWalletDepositHistory) (models.ResponseForWalletDepositHistory, error) {
	panic("implement me")
}

func (w WalletClient) GetWalletWithdrawalHistory(req *models.RequestForWalletWithdrawalHistory) (models.ResponseForWalletWithdrawalHistory, error) {
	panic("implement me")
}

func (w WalletClient) GetSavedAddresses(req *models.RequestForSavedAddresses) (models.ResponseForSavedAddresses, error) {
	panic("implement me")
}

func (w WalletClient) PostCreateSavedAddress(req *models.RequestForSavedAddressCreation) (models.ResponseForSavedAddressCreation, error) {
	panic("implement me")
}

func (w WalletClient) DeleteSavedAddress(req *models.RequestForSavedAddressDeletion) (models.ResponseForSavedAddressDeletion, error) {
	panic("implement me")
}

func (w WalletClient) PostWalletWithdrawal(req *models.RequestForWalletWithdrawal) (models.ResponseForWalletWithdrawal, error) {
	panic("implement me")
}

// NewWalletClient returns a new configured account client
func NewWalletClient(client *rest.Client) (*WalletClient, error) {
	return &WalletClient{
		client: client,
	}, nil
}