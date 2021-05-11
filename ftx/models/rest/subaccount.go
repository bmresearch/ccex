package rest

import "time"

type (

	// RequestForSubAccounts represents a request for sub accounts information.
	RequestForSubAccounts struct {
	}

	// SubAccount holds the necessary information to represent a sub account.
	SubAccount struct {
		Nickname    string `json:"nickname"`
		Deletable   bool   `json:"deletable"`
		Editable    bool   `json:"editable"`
		Competition bool   `json:"competition,omitempty"`
	}

	// ResponseForSubAccounts holds the necessary information to represent the response for sub account data
	ResponseForSubAccounts struct {
		BaseResponse
		Result []SubAccount `json:"result"`
	}

	// RequestForSubAccountChange represents a request for sub account change.
	RequestForSubAccountChange struct {
		Name string
		New  string
	}

	// ResponseForSubAccountChange holds the necessary information to represent the response for a sub account change
	ResponseForSubAccountChange struct {
		BaseResponse
		Result SubAccount `json:"result"`
	}

	// RequestForSubAccountTransfer represents a request for sub account transfer.
	RequestForSubAccountTransfer struct {
		Coin   string
		Source string
		Target string
		Size   float64
	}

	// SubAccountTransfer holds the necessary information to represent a sub account transfer.
	SubAccountTransfer struct {
		ID     int       `json:"id"`
		Coin   string    `json:"coin"`
		Size   int       `json:"size"`
		Time   time.Time `json:"time"`
		Notes  string    `json:"notes"`
		Status string    `json:"status"`
	}

	// RequestForSubAccountBalance represents a request for a sub account balance.
	RequestForSubAccountBalance struct {
		Name string
	}

	// ResponseForSubAccountCreation holds the necessary information to represent the response for a sub account creation
	ResponseForSubAccountCreation struct {
		BaseResponse
		Result SubAccount `json:"result"`
	}

	// ResponseForSubAccountTransfer holds the necessary information to represent the response for a sub account transfer
	ResponseForSubAccountTransfer struct {
		BaseResponse
		Result SubAccountTransfer `json:"result"`
	}

	// SubAccountBalance holds the necessary information to represent a sub account balance.
	SubAccountBalance struct {
		Coin                   string  `json:"coin"`
		Free                   float64 `json:"free"`
		Total                  float64 `json:"total"`
		SpotBorrow             int     `json:"spotBorrow"`
		AvailableWithoutBorrow float64 `json:"availableWithoutBorrow"`
	}

	// ResponseForSubAccountBalance holds the necessary information to represent the response for sub account balances
	ResponseForSubAccountBalance struct {
		BaseResponse
		Result []SubAccountBalance `json:"result"`
	}

)
