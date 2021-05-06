package common

// Response represents a response from an exchange
type Response interface {
	WasSuccessful() bool
	GetResult() interface{}
}
