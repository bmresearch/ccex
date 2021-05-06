package models

import "time"

// Position holds the necessary information to represent an open position
type Position struct {
	Cost                         float64 `json:"cost"`
	EntryPrice                   float64 `json:"entryPrice"`
	EstimatedLiquidationPrice    float64 `json:"estimatedLiquidationPrice,omitempty"`
	Future                       string  `json:"future"`
	InitialMarginRequirement     float64 `json:"initialMarginRequirement"`
	LongOrderSize                float64 `json:"longOrderSize"`
	MaintenanceMarginRequirement float64 `json:"maintenanceMarginRequirement"`
	NetSize                      float64 `json:"netSize"`
	OpenSize                     float64 `json:"openSize"`
	RealizedPnl                  float64 `json:"realizedPnl"`
	ShortOrderSize               float64 `json:"shortOrderSize"`
	Side                         string  `json:"side"`
	Size                         float64 `json:"size"`
	UnrealizedPnl                int     `json:"unrealizedPnl"`
	CollateralUsed               float64 `json:"collateralUsed,omitempty"`
}

// Account holds the necessary information to represent the account information
type Account struct {
	BackstopProvider             bool       `json:"backstopProvider"`
	Collateral                   float64    `json:"collateral"`
	FreeCollateral               float64    `json:"freeCollateral"`
	InitialMarginRequirement     float64    `json:"initialMarginRequirement"`
	Leverage                     int        `json:"leverage"`
	Liquidating                  bool       `json:"liquidating"`
	MaintenanceMarginRequirement float64    `json:"maintenanceMarginRequirement"`
	MakerFee                     float64    `json:"makerFee"`
	MarginFraction               float64    `json:"marginFraction"`
	OpenMarginFraction           float64    `json:"openMarginFraction"`
	TakerFee                     float64    `json:"takerFee"`
	TotalAccountValue            float64    `json:"totalAccountValue"`
	TotalPositionSize            float64    `json:"totalPositionSize"`
	Username                     string     `json:"username"`
	Positions                    []Position `json:"positions"`
}

// SubAccount holds the necessary information to represent a sub account
type SubAccount struct {
	Nickname    string `json:"nickname"`
	Deletable   bool   `json:"deletable"`
	Editable    bool   `json:"editable"`
	Competition bool   `json:"competition,omitempty"`
}

// SubAccountTransfer holds the necessary information to represent a sub account transfer
type SubAccountTransfer struct {
	ID     int       `json:"id"`
	Coin   string    `json:"coin"`
	Size   int       `json:"size"`
	Time   time.Time `json:"time"`
	Notes  string    `json:"notes"`
	Status string    `json:"status"`
}

// SubAccountBalance holds the necessary information to represent a sub account balance
type SubAccountBalance struct {
	Coin                   string  `json:"coin"`
	Free                   float64 `json:"free"`
	Total                  float64 `json:"total"`
	SpotBorrow             int     `json:"spotBorrow"`
	AvailableWithoutBorrow float64 `json:"availableWithoutBorrow"`
}
