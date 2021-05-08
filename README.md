
# ccex - cryptocurrency exchange library
A golang library for cryptocurrency trading with support for many bitcoin, ethereum and altcoin exchange markets, from spot to derivatives.

The ccex library is used to connect and trade with exchanges. It provides quick access to market information for storage, analysis, algorithmic trading etc.

This library is intended to be used by everyone.

## Supported Cryptocurrency Exchange Markets

| logo                                                                                                                                                                                             | id                 | name                                                                                    | ver | doc                                                                                         |
|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|--------------------|-----------------------------------------------------------------------------------------|:---:|:-------------------------------------------------------------------------------------------:|
## Getting Started

These instructions will get you a copy of the project up and running on your local machine for various purposes.

### Prerequisites

In order to be able to install the library in your machine.

```
go > 1.14
```

### Installing

To install this package simply run the following command

```shell
go get -u github.com/bmresearch/ccex
```

## Usage

### Intro

The library consists of public and private methods.
Anyone can use the public methods without needing anything besides the library, while the private ones require an API key from the exchange.

Public APIs include the following:

- market data
- trading pairs
- price feeds
- order books
- trade history
- tickers
- OHLC(V) for charting

In order to use the private APIs you need to obtain API keys from an exchange.
This means signing up to the exchange and creating API keys for your account.

Private APIs allow the following:

- manage accounts
- query account balances
- trade by submitting orders
- deposit and withdraw funds
- query personal orders
- get deposit and withdrawal history
- transfer funds between sub-accounts

### Examples

We have gone through great lengths in order to provide a unified API for multiple exchanges.
Check the package [examples](https://github.com/bmresearch/ccex/examples) to find examples on how to use the librar accordingly.

## Contributing

Please read the [CONTRIBUTING](https://github.com/bmresearch/ccex/blob/master/CONTRIBUTING.md) document before making changes that you would like adopted in the code.

## Support Developer Team

We are investing a significant amount of time into the development of this library. If ccex made your life easier and you want to help us improve it further, or if you want to speed up development of new features and exchanges, please support us with a tip. We appreciate all contributions!

## Contributors

* **Hugo Carvalho** - *Maintainer* - [murlokito](https://github.com/murlokito)

See also the list of [contributors](https://github.com/bmresearch/ccex/contributors) who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE.md](https://github.com/bmresearch/ccex/LICENSE) file for details

