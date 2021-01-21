This is a wallet API made in GO.

# How to Run in Development mode

(Need air installed, use `go get -u github.com/cosmtrek/air`)
1. Install packages:
````
go get github.com/go-shadow/moment
````
2. Run `air main.go`.

# Endpoints

# Next Steps

Featues to be implemented, refactors and dev tools to add to the project:

## Features
- [ ] Select which currencies you want to show (maybe  default config and overridable via query string overrides?)
- [ ] Add main Crypto currencies
	- [X] BTC current value
	- [ ] BTC History
	- [ ] Add more cryptos
- [ ] Add crypto trading benefit calculation for Taxes declaration
	- [ ] First, it will parse a CSV with buy/sell operations
	- [ ] It should be able to calculate the net profit
	- [ ] Should accept a CSV in a POST request
	- [ ] Accept some config parameters 
- [ ] Add wallet feature to store how much I have and track buys/sells
	- [ ] Add feature to calculate profit and taxes from the wallet

## Engineering
- [ ] Refactor vendor modules (use new Go Modules feature)
- [ ] Refactor Project structure, maybe follow https://github.com/golang-standards/project-layout
- [ ] Move fetcher to helpers
- [ ] Use go modules
- [ ] JSON responses
- [ ] Unit tests
	- [ ] Taxes
	- [ ] Accounting
- [ ] Documentation
