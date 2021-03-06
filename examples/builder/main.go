package main

import (
	"fmt"
	"github.com/murlokito/ccex/builder"
	"github.com/murlokito/ccex/examples"
	"github.com/murlokito/ccex/ftx/websocket"
	"github.com/murlokito/ccex/models/rest"
	"github.com/murlokito/ccex/models/ws"
	"time"
)

func main() {

	markets := []string{
		"BTC-PERP", "ETH-PERP",
	}

	data := examples.Data{
		Trades: map[string][]ws.Trade{},
		Ticker: map[string]ws.Ticker{},
	}

	tickerHandler := func(symbol string, message ws.Ticker) {
		//fmt.Printf("bid: %v ask: %v last: %v\n", message.Data.Bid, message.Data.Ask, message.Data.Last)
		data.Ticker[symbol] = message
	}

	tradeHandler := func(symbol string, message ws.Trades) {
		//complete := fmt.Sprintf("num trades: %v", len(message.Data))
		for _, trade := range message {
			data.Trades[symbol] = append(data.Trades[symbol], trade)
			//str := fmt.Sprintf("price: %v size: %v side: %v liq: %v\n", trade.Price, trade.Price, trade.Side, trade.Liquidation)
			//complete += str
			tradeVol := trade.Size * trade.Price
			if tradeVol > 100000 {
				fmt.Printf("{%v} {%v} Volume: $%.2f Price: $%v Liquidation: %v\n", symbol, trade.Side, tradeVol, trade.Price, trade.Liquidation)
			}
		}
		//fmt.Println(complete)
	}


	client, err := builder.NewExchangeWith(builder.FTX, nil, tickerHandler, tradeHandler, nil)
	if err != nil {
		fmt.Printf("err: %v", err)
	}

	/* List all futures */
	futures, err := client.Futures.GetFutures(&rest.RequestForFutures{})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(futures)
	}

	/* Funding Rates */
	fundingRates, err := client.Futures.GetFundingRate(&rest.RequestForFundingRates{
		Future: "BTC-PERP",
		Start:  time.Time{},
		End:    time.Time{},
	})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(fundingRates)
	}

	client.Streaming.Connect()

	for _, market := range markets {
		err = client.Streaming.Subscribe(websocket.Ticker, market)
		if err != nil {
			fmt.Printf("err: %v", err)
		}

		err = client.Streaming.Subscribe(websocket.Trades, market)
		if err != nil {
			fmt.Printf("err: %v", err)
		}
	}

	for {
		if client.Streaming.Connected() {
			fmt.Println(fmt.Sprintf("client is connected - active subs %v", client.Streaming.Subscriptions()))
		}
		for k, v := range data.Trades {
			fmt.Println(fmt.Sprintf("number of trades for %v - %v", k, len(v)))
		}
		for k, v := range data.Ticker {
			fmt.Println(fmt.Sprintf("latest ticker for %v - %v", k, v.LastPrice))
		}
		time.Sleep(15 * time.Second)
	}
}
