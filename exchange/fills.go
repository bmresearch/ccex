package exchange

import (
	"time"

	"github.com/murlokito/ccex/common"
)

type (
	Fills interface {
		GetFills(market, order string, orderId, limit int, start, end time.Time) (common.Response, error)
	}
)
