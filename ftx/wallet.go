package ftx

import (
	"github.com/murlokito/ccex/ftx/rest"
	rest3 "github.com/murlokito/ccex/models/rest"
)

type (

	// WalletClient represents the client for the FTX Wallet API.
	WalletClient struct {
		client *rest.Client
	}
)

func (w WalletClient) GetWalletBalances(req *rest3.RequestForWalletBalances) (rest3.ResponseForWalletBalances, error) {
	panic("implement me")
}

func (w WalletClient) GetDepositAddress(req *rest3.RequestForDepositAddress) (rest3.ResponseForDepositAddress, error) {
	panic("implement me")
}

func (w WalletClient) GetWalletDepositHistory(req *rest3.RequestForWalletDepositHistory) (rest3.ResponseForWalletDepositHistory, error) {
	panic("implement me")
}

func (w WalletClient) GetWalletWithdrawalHistory(req *rest3.RequestForWalletWithdrawalHistory) (rest3.ResponseForWalletWithdrawalHistory, error) {
	panic("implement me")
}

func (w WalletClient) GetSavedAddresses(req *rest3.RequestForSavedAddresses) (rest3.ResponseForSavedAddresses, error) {
	panic("implement me")
}

func (w WalletClient) PostCreateSavedAddress(req *rest3.RequestForSavedAddressCreation) (rest3.ResponseForSavedAddressCreation, error) {
	panic("implement me")
}

func (w WalletClient) DeleteSavedAddress(req *rest3.RequestForSavedAddressDeletion) (rest3.ResponseForSavedAddressDeletion, error) {
	panic("implement me")
}

func (w WalletClient) PostWalletWithdrawal(req *rest3.RequestForWalletWithdrawal) (rest3.ResponseForWalletWithdrawal, error) {
	panic("implement me")
}

func NewWalletClient(client *rest.Client) (*WalletClient, error) {
	return &WalletClient{
		client: client,
	}, nil
}
