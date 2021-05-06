package exchange

import (
	"github.com/murlokito/ccex/common"
	"time"
)

type (
	// Orders specifies functionality for the orders API
	Orders interface {
		// GetOpenOrders fetches the open orders. Optionally pass `market` to only fetch orders from a single market.
		GetOpenOrders(market string) (common.Response, error)

		/*
			GetOrderHistory fetches the orders history.
			Optionally pass `market` to only fetch orders from a single market and `start` and `end` for orders within a period.
		*/
		GetOrderHistory(market string, limit int, start, end time.Time) (common.Response, error)

		GetOpenTriggerOrders(market, triggerOrderType string) (common.Response, error)

		GetTriggerOrderHistory(market, side, triggerOrderType, orderType string, limit int, start, end time.Time) (common.Response, error)

		GetTriggerOrderTriggers() (common.Response, error)

		PostOrder(market, side, orderType, clientId string, price, size float32, reduceOnly, postOnly, ioc bool) (common.Response, error)

		PostTriggerOrder(market, side, triggerOrderType string, size, triggerPrice, orderPrice, trailValue float32, reduceOnly, retryUntilFilled bool) (common.Response, error)

		PostModifyOrder(orderId int, price, size float32, clientId string, byClientId bool) (common.Response, error)

		PostModifyTriggerOrder(orderId int, size, triggerPrice, orderPrice, trailValue float32) (common.Response, error)

		GetOrderStatus(orderId, clientId int, byClientId bool) (common.Response, error)

		DeleteOrder(orderId, clientId int, byClientId bool) (common.Response, error)

		DeleteTriggerOrder(orderId int) (common.Response, error)

		DeleteAllOrders(market string, conditionalOrdersOnly, limitOrdersOnly bool) (common.Response, error)
	}
)
