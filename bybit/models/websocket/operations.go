package websocket

type (
	// BaseOperation is the base of a message sent, all attributes are common.
	BaseOperation struct {
		Op string `json:"op"`
	}

	// Message is a common message sent.
	Message struct {
		BaseOperation
		Args []string `json:"args"`
	}

	// SubscriptionResponse represents a response to a subscription.
	SubscriptionResponse struct {
		Success      bool    `json:"success"`
		Message      string  `json:"ret_msg"`
		ConnectionId string  `json:"conn_id"`
		Request      Message `json:"request"`
	}
)
