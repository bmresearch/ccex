package exchange

import (
	"time"

	"github.com/murlokito/ccex/common"
)

type (
	Funding interface {
		GetFundingPayments(future string, start, end time.Time) (common.Response, error)
	}
)
