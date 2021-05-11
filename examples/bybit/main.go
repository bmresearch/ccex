package main

import (
	"fmt"
	"github.com/murlokito/ccex/bybit"
	"github.com/murlokito/ccex/examples"
	"github.com/murlokito/ccex/ftx/websocket"
	"github.com/murlokito/ccex/models/ws"
	"time"
)

func main(){

	markets := []string{
		"BTC-PERP", "ETH-PERP",
	}

	data := examples.Data{
		Trades: map[string][]ws.TradeData{},
		Ticker: map[string]ws.TickerData{},
	}

	tradeHandler := func(message ws.TradeMessage) {
		for _, trade := range message.Data {
			data.Trades[message.Market] = append(data.Trades[message.Market], trade)
			tradeVol := trade.Size * trade.Price
			if tradeVol > 500000 {
				fmt.Printf("{%v} {%v} Volume: $%v Price: $%v Liquidation: %v\n", message.Market, trade.Side, examples.PrettyFormat(tradeVol), trade.Price, trade.Liquidation)
			}
		}
	}

	client, err := bybit.NewClient(nil, nil, nil, tradeHandler, nil)
	if err != nil {
		fmt.Printf("err: %v", err)
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
		for k, v := range data.Trades {
			fmt.Println(fmt.Sprintf("trades (inception) for %v - %v", k, len(v)))
			fmt.Println(fmt.Sprintf("trades (Δ 1m) for %v - %v", k, examples.GetDeltaTrades(v, 1*time.Minute, time.Now())))
			fmt.Println(fmt.Sprintf("trades (Δ 15m) for %v - %v", k, examples.GetDeltaTrades(v, 15*time.Minute, time.Now())))
			fmt.Println(fmt.Sprintf("volume (Δ 1m) for %v - $%v", k, examples.PrettyFormat(examples.GetDeltaVol(v, 1*time.Minute, time.Now()))))
			fmt.Println(fmt.Sprintf("volume (Δ 15m) for %v - $%v", k, examples.PrettyFormat(examples.GetDeltaVol(v, 15*time.Minute, time.Now()))))
		}
		for k, v := range data.Ticker {
			fmt.Println(fmt.Sprintf("latest ticker for %v - $%.2f", k, v.Last))
		}
		time.Sleep(1 * time.Minute)
	}
}
