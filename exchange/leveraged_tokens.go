package exchange

import "github.com/murlokito/ccex/common"

type (
	LeveragedTokens interface {
		GetLeveragedTokens() (common.Response, error)

		GetLeveragedTokenInfo(token string) (common.Response, error)

		GetLeveragedTokenBalances() (common.Response, error)

		PostLeveragedTokenCreationRequest(token string, size float32) (common.Response, error)

		GetLeveragedTokenRedemptionRequests() (common.Response, error)

		PostLeveragedTokenRedemptionRequest(token string, size float32) (common.Response, error)
	}
)
