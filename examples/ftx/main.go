package main

import (
	"fmt"
	"github.com/murlokito/ccex/ftx"
	"github.com/murlokito/ccex/ftx/websocket"
	"github.com/murlokito/ccex/models/rest"
	"github.com/murlokito/ccex/models/ws"
	"time"
)

type Data struct {
	Trades map[string][]ws.TradeData
	Ticker map[string]ws.TickerData
}

func main() {

	markets := []string{
		"BTC-PERP", "ETH-PERP",
	}

	data := Data{
		Trades: map[string][]ws.TradeData{},
		Ticker: map[string]ws.TickerData{},
	}

	tickerHandler := func(message ws.TickerMessage) {
		//fmt.Printf("bid: %v ask: %v last: %v\n", message.Data.Bid, message.Data.Ask, message.Data.Last)
		data.Ticker[message.Market] = message.Data
	}

	tradeHandler := func(message ws.TradeMessage) {
		//complete := fmt.Sprintf("num trades: %v", len(message.Data))
		for _, trade := range message.Data {
			data.Trades[message.Market] = append(data.Trades[message.Market], trade)
			//str := fmt.Sprintf("price: %v size: %v side: %v liq: %v\n", trade.Price, trade.Price, trade.Side, trade.Liquidation)
			//complete += str
			tradeVol := trade.Size * trade.Price
			if tradeVol > 100000 {
				fmt.Printf("{%v} {%v} Volume: $%.2f Price: $%v Liquidation: %v\n", message.Market, trade.Side, tradeVol, trade.Price, trade.Liquidation)
			}
		}
		//fmt.Println(complete)
	}


	ftxClient, err := ftx.NewFTXClient(nil, nil, tickerHandler, tradeHandler, nil)
	if err != nil {
		fmt.Printf("err: %v", err)
	}

	/* List all futures */
	futures, err := ftxClient.Futures.GetFutures(&rest.RequestForFutures{})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(futures)
	}

	/* Funding Rates */
	fundingRates, err := ftxClient.Futures.GetFundingRate(&rest.RequestForFundingRates{
		Future: "BTC-PERP",
		Start:  time.Time{},
		End:    time.Time{},
	})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(fundingRates)
	}

	ftxClient.Streaming.Connect()

	for _, market := range markets {
		err = ftxClient.Streaming.Subscribe(websocket.Ticker, market)
		if err != nil {
			fmt.Printf("err: %v", err)
		}

		err = ftxClient.Streaming.Subscribe(websocket.Trades, market)
		if err != nil {
			fmt.Printf("err: %v", err)
		}
	}

	for {
		if ftxClient.Streaming.Connected() {
			fmt.Println(fmt.Sprintf("client is connected - active subs %v", ftxClient.Streaming.Subscriptions()))
		}
		for k, v := range data.Trades {
			fmt.Println(fmt.Sprintf("number of trades for %v - %v", k, len(v)))
		}
		for k, v := range data.Ticker {
			fmt.Println(fmt.Sprintf("latest ticker for %v - %v", k, v.Last))
		}
		time.Sleep(15 * time.Second)
	}

}
