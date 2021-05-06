package ftx

import "github.com/murlokito/ccex/ftx/rest"

type (

	// WalletClient represents the client for the FTX Wallet API.
	WalletClient struct {
		client *rest.Client
	}
)


func NewWalletClient(client *rest.Client) (*WalletClient, error){
	return &WalletClient{
		client: client,
	}, nil
}