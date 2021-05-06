package ftx

import (
	"github.com/murlokito/ccex/common"
	"github.com/murlokito/ccex/ftx/rest"
	"time"
)

type (

	// OrdersClient represents the client for the FTX Orders API.
	OrdersClient struct {
		client *rest.Client
	}
)

func (o OrdersClient) GetOpenOrders(market string) (common.Response, error) {
	panic("implement me")
}

func (o OrdersClient) GetOrderHistory(market string, limit int, start, end time.Time) (common.Response, error) {
	panic("implement me")
}

func (o OrdersClient) GetOpenTriggerOrders(market, triggerOrderType string) (common.Response, error) {
	panic("implement me")
}

func (o OrdersClient) GetTriggerOrderHistory(market, side, triggerOrderType, orderType string, limit int, start, end time.Time) (common.Response, error) {
	panic("implement me")
}

func (o OrdersClient) GetTriggerOrderTriggers() (common.Response, error) {
	panic("implement me")
}

func (o OrdersClient) PostOrder(market, side, orderType, clientId string, price, size float32, reduceOnly, postOnly, ioc bool) (common.Response, error) {
	panic("implement me")
}

func (o OrdersClient) PostTriggerOrder(market, side, triggerOrderType string, size, triggerPrice, orderPrice, trailValue float32, reduceOnly, retryUntilFilled bool) (common.Response, error) {
	panic("implement me")
}

func (o OrdersClient) PostModifyOrder(orderId int, price, size float32, clientId string, byClientId bool) (common.Response, error) {
	panic("implement me")
}

func (o OrdersClient) PostModifyTriggerOrder(orderId int, size, triggerPrice, orderPrice, trailValue float32) (common.Response, error) {
	panic("implement me")
}

func (o OrdersClient) GetOrderStatus(orderId, clientId int, byClientId bool) (common.Response, error) {
	panic("implement me")
}

func (o OrdersClient) DeleteOrder(orderId, clientId int, byClientId bool) (common.Response, error) {
	panic("implement me")
}

func (o OrdersClient) DeleteTriggerOrder(orderId int) (common.Response, error) {
	panic("implement me")
}

func (o OrdersClient) DeleteAllOrders(market string, conditionalOrdersOnly, limitOrdersOnly bool) (common.Response, error) {
	panic("implement me")
}

func NewOrdersClient(client *rest.Client) (*OrdersClient, error) {
	return &OrdersClient{
		client: client,
	}, nil
}
