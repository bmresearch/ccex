package main

import (
	"fmt"
	"github.com/murlokito/ccex/examples"
	"github.com/murlokito/ccex/ftx"
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
		OrderBook: map[string]ws.OrderBookSnapshot{},
	}

	tickerHandler := func(symbol string, ticker ws.Ticker) {
		fmt.Printf("bid: %v ask: %v last: %v\n", ticker.Bid, ticker.Ask, ticker.LastPrice)
		data.Ticker[symbol] = ticker
	}

	tradesHandler := func(symbol string, trades ws.Trades) {
		complete := fmt.Sprintf("num trades: %v", len(trades))
		for _, trade := range trades {
			data.Trades[symbol] = append(data.Trades[symbol], trade)
			str := fmt.Sprintf("price: %v size: %v side: %v liq: %v\n", trade.Price, trade.Price, trade.Side, trade.Liquidation)
			complete += str
			tradeVol := trade.Size * trade.Price
			if tradeVol > 500000 {
				fmt.Printf("{%v} {%v} Volume: $%v Price: $%v Liquidation: %v\n", symbol, trade.Side, examples.PrettyFormat(tradeVol), trade.Price, trade.Liquidation)
			}
		}
		fmt.Println(complete)
	}

	obsSnapshotHandler := func(symbol string, obs ws.OrderBookSnapshot) {
		fmt.Printf("{%v} bids: %v asks: %v \n", symbol, len(obs.Bids),len(obs.Asks))
		data.OrderBook[symbol] = obs
	}

	obdSnapshotHandler := func(symbol string, obd ws.OrderBookDelta) {
		fmt.Printf("{%v} deleted: %v updated: %v inserted: %v \n", symbol, len(obd.Delete), len(obd.Update), len(obd.Insert))
		// TODO: process order book deltas and update snapshot
	}

	ftxClient, err := ftx.NewClientWith(nil, tickerHandler, tradesHandler, obsSnapshotHandler, obdSnapshotHandler)
	if err != nil {
		fmt.Printf("err: %v", err)
	}

	/* List all futures
	futures, err := ftxClient.Futures.GetFutures(&rest.RequestForFutures{})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Sprintf("futures available: %v", len(futures.Result))
	}
	*/

	/* Funding Rates
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
	*/

	/* Historical OHLCV of the FTX indexes

	start, err := time.Parse(time.RFC3339, "2020-10-02T15:04:05+07:00")
	end, err := time.Parse(time.RFC3339, "2020-10-03T15:04:05+07:00")

	historical, err := ftxClient.Futures.GetHistoricalIndex(&rest.RequestForHistoricalIndex{
		Index: "DEFI",
		Resolution: 60, // 60 seconds per candle
		Limit: 2000,
		Start: start,
		End: end,
	})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(historical)
	}
	*/

	ftxClient.Streaming.Connect()

	for _, market := range markets {
/*		err = ftxClient.Streaming.Subscribe(websocket.Ticker, market)
		if err != nil {
			fmt.Printf("err: %v", err)
		}

		err = ftxClient.Streaming.Subscribe(websocket.Trades, market)
		if err != nil {
			fmt.Printf("err: %v", err)
		}
*/
		err = ftxClient.Streaming.Subscribe(websocket.Orderbook, market)
		if err != nil {
			fmt.Printf("err: %v", err)
		}
	}

	go func() {
		for {
			/* Open interest data */
			for _, market := range markets {
				openInterest, err := ftxClient.Futures.GetOpenInterest(&rest.RequestForOpenInterest{
					Future: market,
				})
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println(
						fmt.Sprintf(
							"%v: open interest for %v - OI: %.2f - 24h VOL: %.2f - Next FR: %.6f",
							time.Now().Format(time.RFC3339),
							market,
							openInterest.Result.OpenInterest,
							openInterest.Result.Volume,
							openInterest.Result.FundingRate,
						),
					)
				}
			}
			time.Sleep(1 * time.Minute)
		}
	}()

	for {
		for k, v := range data.Trades {
			fmt.Println(fmt.Sprintf("trades (inception) for %v - %v", k, len(v)))
			fmt.Println(fmt.Sprintf("trades (Δ 1m) for %v - %v", k, examples.GetDeltaTrades(v, 1*time.Minute, time.Now())))
			fmt.Println(fmt.Sprintf("trades (Δ 15m) for %v - %v", k, examples.GetDeltaTrades(v, 15*time.Minute, time.Now())))
			fmt.Println(fmt.Sprintf("volume (Δ 1m) for %v - $%v", k, examples.PrettyFormat(examples.GetDeltaVol(v, 1*time.Minute, time.Now()))))
			fmt.Println(fmt.Sprintf("volume (Δ 15m) for %v - $%v", k, examples.PrettyFormat(examples.GetDeltaVol(v, 15*time.Minute, time.Now()))))
		}
		for k, v := range data.Ticker {
			fmt.Println(fmt.Sprintf("latest ticker for %v - $%.2f", k, v.LastPrice))
		}
		time.Sleep(1 * time.Minute)
	}

}
