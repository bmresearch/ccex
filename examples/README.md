# ccex examples

This package holds examples on how to use `ccex` in the following ways:

- exchange-agnostic using the builder package
- exchange-specific using any of the supported exchanges

## Preview

### Exchange agnostic

```go
package main

import (
	"fmt"
	"net/http"
	
	"github.com/murlokito/ccex"
	builder "github.com/murlokito/ccex/builder"
)

func main() {
	
	var (
		exchanges []ccex.Exchange		
    )
	
	clients := ccex.{
		ccex.FTX,
		ccex.Binance,
    }
	

	params := &ccex.Parameters{
		Debug:      false,
		HttpClient: &http.Client{},
		ProxyURL:   "",
		AccessKey:  "access-key",
		SecretKey:  "access-secret",
	}

	exchange, err := builder.NewExchangeFromParameters(ccex.FTX, params)
	if err != nil {
		fmt.Println(err)
	}
	
	exchange.
}

```

### FTX

```go
package main

import "github.com/murlokito/ccex"

func main(){
	
}

```