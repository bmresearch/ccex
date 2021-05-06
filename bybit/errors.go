package bybit

import "fmt"

var (
	ErrMethodNotImplemented = fmt.Errorf("bybit: method not implemented")
	ErrNotAuthenticated = fmt.Errorf("bybit: not authenticated")
)
