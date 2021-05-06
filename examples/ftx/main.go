package ftx

import (
	"fmt"
	"github.com/murlokito/ccex/auth"
	"github.com/murlokito/ccex/builder"
	"github.com/murlokito/ccex/config"
	"github.com/murlokito/ccex/ftx/websocket"
	"github.com/murlokito/ccex/ftx/websocket/models"
)

func main() {

	cfg := &config.Configuration{
		Auth:       &auth.Authentication{
			Key:    "some-key",
			Secret: "some-secret",
		},
		SubAccount: "some-sub-account",
	}

	ftxClient, err := builder.NewFTXClient(cfg)
	if err != nil {
		fmt.Printf("err: %v", err)
	}

	handler := func(message models.TickerMessage) {
		fmt.Printf("bid: %v ask: %v last: %v", message.Data.Bid, message.Data.Ask, message.Data.Last)
	}

	err = ftxClient.Websocket.TickerHandler(handler)
	if err != nil {
		fmt.Printf("err: %v", err)
	}

	err = ftxClient.Websocket.Subscribe(websocket.Ticker, "BTC-PERP")
	if err != nil {
		fmt.Printf("err: %v", err)
	}

}
