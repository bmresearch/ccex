package rest

type (
	// RequestForPositions represents a request for the account's positions.
	RequestForPositions struct {
		ShowAveragePrice bool
	}

	// Position holds the necessary information to represent an open position.
	Position struct {
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

	// ResponseForPositions holds the necessary information to represent the response for positions data.
	ResponseForPositions struct {
		BaseResponse
		Result []Position `json:"result"`
	}

	// Account
	// TODO: Remove
	Account struct {
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

	// ResponseForAccount holds the necessary information to represent the response the account data
	ResponseForAccount struct {
		BaseResponse
		Result Account `json:"result"`
	}

	RequestForAccountLeverageChange struct {
		Leverage int
	}

	RequestForFuturesAccountLeverageChange struct {
		Future   string
		Leverage int
	}

	// ResponseForAccountLeverageChange represents the response for account leverage change request.
	ResponseForAccountLeverageChange struct {
		BaseResponse
	}

	// 	ResponseForFuturesAccountLeverageChange represents the response for futures account leverage change request.
	ResponseForFuturesAccountLeverageChange struct {
		BaseResponse
	}
)
