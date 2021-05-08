package rest

import "time"

type (
	RequestForOpenOrders struct {
		Market string
	}

	OpenOrder struct {

	}

	ResponseForOpenOrders struct {
		BaseResponse
	}

	RequestForOrderHistory struct {
		Market string
		Limit int
		Start time.Time
		End time.Time
	}

	Order struct {

	}

	ResponseForOrderHistory struct {
		BaseResponse
	}

	RequestForOpenTriggerOrders struct {
		Market string
		TriggerOrderType string
	}

	OpenTriggerOrder struct {

	}

	ResponseForOpenTriggerOrders struct {
		BaseResponse
	}



)