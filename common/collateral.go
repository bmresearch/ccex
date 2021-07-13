package common

// Collateral represents a collateral type
type Collateral int

const (
	StableMargined = iota
	CoinMargined
)