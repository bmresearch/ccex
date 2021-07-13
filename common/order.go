package common

// Order represents an order type
type Order int

const (
	LimitOrder = iota
	MarketOrder
	TriggerOrder
)
