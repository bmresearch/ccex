package ftx

import (
	"github.com/murlokito/ccex/common"
	"github.com/murlokito/ccex/ftx/rest"
	"time"
)

type (

	// WalletClient represents the client for the FTX Wallet API.
	WalletClient struct {
		client *rest.Client
	}
)

func (w WalletClient) GetWalletCoins() (common.Response, error) {
	panic("implement me")
}

func (w WalletClient) GetWalletBalances() (common.Response, error) {
	panic("implement me")
}

func (w WalletClient) GetAllWalletBalances() (common.Response, error) {
	panic("implement me")
}

func (w WalletClient) GetDepositAddress(coin, method string) (common.Response, error) {
	panic("implement me")
}

func (w WalletClient) GetWalletDepositHistory(limit int, start, end time.Time) (common.Response, error) {
	panic("implement me")
}

func (w WalletClient) GetWalletWithdrawalHistory(limit int, start, end time.Time) (common.Response, error) {
	panic("implement me")
}

func (w WalletClient) GetWalletAirdropHistory(limit int, start, end time.Time) (common.Response, error) {
	panic("implement me")
}

func (w WalletClient) GetSavedAddresses(coin string) (common.Response, error) {
	panic("implement me")
}

func (w WalletClient) PostCreateSavedAddress(coin, address, addressName, tag string, isPrimeTrust bool) (common.Response, error) {
	panic("implement me")
}

func (w WalletClient) DeleteSavedAddress(addressId int) (common.Response, error) {
	panic("implement me")
}

func (w WalletClient) PostWalletWithdrawal(coin, address, tag, password, code string, size int) (common.Response, error) {
	panic("implement me")
}

func NewWalletClient(client *rest.Client) (*WalletClient, error) {
	return &WalletClient{
		client: client,
	}, nil
}
