package common

// Market represents a market type
type Market int

const (
	Spot = iota
	Futures
	Options
)