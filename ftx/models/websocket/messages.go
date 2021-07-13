package websocket

import (
	"fmt"
	"github.com/murlokito/ccex/models/ws"
	"time"
)

type (
	// BaseMessage holds the common attributes of the websocket responses.
	BaseMessage struct {
		Type    string `json:"type"`
		Channel string `json:"channel"`
		Market  string `json:"market"`
	}

	// Trade holds the data from the markets message
	Trade struct {
		Side        string  `json:"side"`
		Size        float64 `json:"size"`
		Price       float64 `json:"price"`
		Liquidation bool    `json:"liquidation"`
		Timestamp   string  `json:"time"`
	}

	// Trades is an helper type to represent a slice of trades.
	Trades []Trade

	// TradeMessage holds a message from the trades channel
	TradeMessage struct {
		BaseMessage
		Data Trades `json:"data"`
	}

	// Item represents an item in the order book data.
	Item []float64

	// Items is an helper type to represent a slice of trades.
	Items []Item

	// OrderBookData represents an order book delta, where items must be deleted, updated or inserted in the original snapshot.
	OrderBookData struct {
		Action    string  `json:"action"`
		Bids      Items   `json:"bids"`
		Asks      Items   `json:"asks"`
		Timestamp float64 `json:"timestamp"`
	}

	// OrderBookMessage holds a message from the orderbook channel
	OrderBookMessage struct {
		BaseMessage
		Data OrderBookData `json:"data"`
	}

	// Ticker holds the data from the markets message
	Ticker struct {
		Bid       float64 `json:"bid"`
		Ask       float64 `json:"ask"`
		Last      float64 `json:"last"`
		Timestamp float64 `json:"time"`
	}

	// TickerMessage holds a message from the ticker channel
	TickerMessage struct {
		BaseMessage
		Data Ticker `json:"data"`
	}
)

// Standard converts the trade information specific to FTX into the unified ws.Trade model.
func (t *Trades) Standard() ws.Trades {
	var trades []ws.Trade

	for _, trade := range *t {
		trades = append(trades, trade.Standard())
	}

	return trades
}

// Standard converts the trade information specific to FTX into the unified ws.Trade model.
func (t *Trade) Standard() ws.Trade {
	return ws.Trade{
		Side:                t.Side,
		Size:                t.Size,
		Price:               t.Price,
		ProvidesLiquidation: true,
		Liquidation:         t.Liquidation,
	}
}

// Standard converts the order book entry specific to FTX into the unified ws.Entry model.
func (i Item) Standard() *ws.Entry {
	return &ws.Entry{
		Price: i[0],
		Size:  i[1],
	}
}

// Snapshot converts the order book snapshot specific to FTX into the unified ws.OrderBookSnapshot model.
func (obm *OrderBookMessage) Snapshot() ws.OrderBookSnapshot {
	var (
		bids []*ws.Entry
		asks []*ws.Entry
	)

	for _, entry := range obm.Data.Bids {
		bids = append(bids, entry.Standard())
	}

	for _, entry := range obm.Data.Bids {
		bids = append(bids, entry.Standard())
	}

	return ws.OrderBookSnapshot{
		Bids: bids,
		Asks: asks,
	}
}

// Standard converts the instrument information specific to FTX into the unified ws.Ticker model.
func (i *Ticker) Standard() ws.Ticker {
	t, err := time.Parse(time.RFC3339, fmt.Sprintf("%v", i.Timestamp))
	if err != nil {
		fmt.Println(err.Error())
	}
	return ws.Ticker{
		Bid:        i.Bid,
		Ask:        i.Ask,
		LastPrice:  i.Last,
		MarkPrice:  0,
		IndexPrice: 0,
		Timestamp:  t,
	}
}

// Delta converts the order book delta specific to FTX into the unified ws.OrderBookDelta model.
func (obm *OrderBookMessage) Delta() ws.OrderBookDelta {
	var (
		deleteEntries []*ws.Entry
		updateEntries []*ws.Entry
		insertEntries []*ws.Entry
	)

	for _, item := range obm.Data.Bids {
		it := item.Standard()
		if it.Size == 0 {
			deleteEntries = append(deleteEntries, it)
			continue
		}
		updateEntries = append(updateEntries, it)
	}

	for _, item := range obm.Data.Asks {
		it := item.Standard()
		if it.Size == 0 {
			deleteEntries = append(deleteEntries, it)
			continue
		}
		updateEntries = append(updateEntries, it)
	}

	return ws.OrderBookDelta{
		Delete: deleteEntries,
		Update: updateEntries,
		Insert: insertEntries,
	}
}
