package ws

type (
	// Entry represents an entry in the order book.
	Entry struct {
		Price float64 `json:"price"`
		Size  float64 `json:"size"`
	}

	// OrderBookDelta defines which items of the order book need to be updated.
	OrderBookDelta struct {
		Delete []*Entry `json:"delete"`
		Update []*Entry `json:"update"`
		Insert []*Entry `json:"insert"`
	}

	// OrderBookSnapshot defines an order book snapshot.
	OrderBookSnapshot struct {
		Bids []*Entry `json:"bid"`
		Asks []*Entry `json:"asks"`
	}
)
