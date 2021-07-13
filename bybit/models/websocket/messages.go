package websocket

import (
	"github.com/murlokito/ccex/bybit/websocket"
	"github.com/murlokito/ccex/models/ws"
	"time"
)

type (
	// BaseDataResponse holds the base message attributes.
	BaseDataResponse struct {
		Topic     string      `json:"topic"`
		Data      interface{} `json:"data"`
		Type      string      `json:"type,omitempty"`
		Timestamp int64       `json:"timestamp_e6,omitempty"`
	}


	// Trade represents a trade in the real-time trades feed.
	Trade struct {
		Symbol    string  `json:"symbol"`
		Side      string  `json:"side"`
		Size      float64 `json:"size"`
		Price     float64 `json:"price"`
		Timestamp string  `json:"timestamp"`
	}

	// Trades is an helper type to represent a slice of trades.
	Trades []Trade

	// Item represents an item in the order book data.
	Item struct {
		Price  float64 `json:"price"`
		Symbol string  `json:"symbol"`
		Id     int64   `json:"id"`
		Side   string  `json:"side"`
		Size   float64   `json:"size,omitempty"`
	}

	// Items is an helper type to represent a slice of trades.
	Items []Item

	// OrderBookDelta represents an order book delta, where items must be deleted, updated or inserted in the original snapshot.
	OrderBookDelta struct {
		Delete Items `json:"delete"`
		Update Items `json:"update"`
		Insert Items `json:"insert"`
	}

	// OrderBookSnapshot represents an order book snapshot, where all data is new.
	OrderBookSnapshot struct {
		Data Items `json:"data"`
	}

	// Instrument represents an instrument information message
	Instrument struct {
		LastPriceE4  int       `json:"last_price_e4,omitempty"`
		Bid1PriceE4  int       `json:"bid1_price_e4,omitempty"`
		Ask1PriceE4  int       `json:"ask1_price_e4,omitempty"`
		MarkPriceE4  int       `json:"mark_price_e4,omitempty"`
		IndexPriceE4 int       `json:"index_price_e4,omitempty"`
		UpdatedAt    time.Time `json:"updated_at,omitempty"`
	}

	// Instruments is an helper type to represent a slice of instruments.
	Instruments []Instrument

	// InstrumentDelta represents a snapshot of the instrument data
	InstrumentDelta struct {
		Delete Instruments `json:"delete"`
		Update Instruments `json:"update"`
		Insert Instruments `json:"insert"`
	}

	// InstrumentSnapshot represents a snapshot of the instrument data
	InstrumentSnapshot struct {
		Data Instrument `json:"data"`
	}
)

// Standard converts the trade information specific to Bybit into the unified ws.Trade model.
func (t *Trades) Standard() ws.Trades {
	var trades []ws.Trade

	for _, trade := range *t {
		trades = append(trades, trade.Standard())
	}

	return trades
}

// Standard converts the trade information specific to Bybit into the unified ws.Trade model.
func (t *Trade) Standard() ws.Trade {
	return ws.Trade{
		Side:                t.Side,
		Size:                t.Size,
		Price:               t.Price,
		ProvidesLiquidation: false,
		Liquidation:         false,
	}
}

// Standard converts the order book entry specific to Bybit into the unified ws.Entry model.
func (i *Item) Standard() *ws.Entry {
	return &ws.Entry{
		Price: i.Price,
		Size:  i.Size,
	}
}

// Standard converts the order book snapshot specific to Bybit into the unified ws.OrderBookSnapshot model.
func (obs *OrderBookSnapshot) Standard() ws.OrderBookSnapshot {
	var (
		bids []*ws.Entry
		asks []*ws.Entry
	)

	for _, entry := range obs.Data {
		switch entry.Side {
		case websocket.Buy:
			bids = append(bids, entry.Standard())
			break
		case websocket.Sell:
			asks = append(asks, entry.Standard())
			break
		}
	}

	return ws.OrderBookSnapshot{
		Bids: bids,
		Asks: asks,
	}
}

// Standard converts the trade information specific to Bybit into the unified ws.OrderBookDelta model.
func (obd *OrderBookDelta) Standard() ws.OrderBookDelta {
	var (
		deleteEntries []*ws.Entry
		updateEntries []*ws.Entry
		insertEntries []*ws.Entry
	)

	for _, entry := range obd.Delete {
		deleteEntries = append(deleteEntries, entry.Standard())
	}

	for _, entry := range obd.Update {
		updateEntries = append(updateEntries, entry.Standard())
	}

	for _, entry := range obd.Insert {
		insertEntries = append(insertEntries, entry.Standard())
	}

	return ws.OrderBookDelta{
		Delete: deleteEntries,
		Update: updateEntries,
		Insert: insertEntries,
	}
}

// Update updates the instrument information specific to Bybit.
func (i *Instrument) Update(new *InstrumentDelta) {
	i.UpdatedAt = new.Update[0].UpdatedAt
	i.Bid1PriceE4 = new.Update[0].Bid1PriceE4
	i.Ask1PriceE4 = new.Update[0].Ask1PriceE4
	i.LastPriceE4 = new.Update[0].LastPriceE4
	i.MarkPriceE4 = new.Update[0].MarkPriceE4
	i.IndexPriceE4 = new.Update[0].IndexPriceE4
}

// Instrument gets a pointer to the inner Instrument data structure.
func (i *InstrumentSnapshot) Instrument() *Instrument {
	return &i.Data
}

// Standard converts the instrument snapshot information specific to Bybit into the unified ws.Ticker model.
func (i *InstrumentSnapshot) Standard() ws.Ticker {
	return ws.Ticker{
		Bid:        float64(i.Data.Bid1PriceE4 /  10000),
		Ask:        float64(i.Data.Ask1PriceE4 /  10000),
		LastPrice:  float64(i.Data.LastPriceE4 /  10000),
		MarkPrice:  float64(i.Data.MarkPriceE4 /  10000),
		IndexPrice: float64(i.Data.IndexPriceE4 / 10000),
		Timestamp:  i.Data.UpdatedAt,
	}
}

// Standard converts the instrument information specific to Bybit into the unified ws.Ticker model.
func (i *Instrument) Standard() ws.Ticker {
	return ws.Ticker{
		Bid:        float64(i.Bid1PriceE4 /  10000),
		Ask:        float64(i.Ask1PriceE4 /  10000),
		LastPrice:  float64(i.LastPriceE4 /  10000),
		MarkPrice:  float64(i.MarkPriceE4 /  10000),
		IndexPrice: float64(i.IndexPriceE4 / 10000),
		Timestamp:  i.UpdatedAt,
	}
}
