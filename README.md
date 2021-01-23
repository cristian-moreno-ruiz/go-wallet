This is a wallet API made in GO.

# How to Run in Development mode

(Need air installed, use `go get -u github.com/cosmtrek/air`)
1. Install packages:
````
go get github.com/go-shadow/moment
````
2. Run `docker-compose up -d mongo` to start the database
3. Run `air main.go`.

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
	- [X] First, it will parse a CSV with buy/sell operations
	- [X] It should be able to calculate the net profit
	- [X] Should accept a CSV in a POST request
	- [ ] Accept some config parameters 
- [ ] Add wallet feature to store how much I have and track buys/sells
	- [ ] Add feature to calculate profit and taxes from the wallet
	- [ ] Add feature to have multiple wallets
- [ ] Add middlewares to handle auth/multiple users
- [ ] How to handle multiple exchanges
- [ ] Compatibility with CSVs exported from exchanges directly

## Engineering
- [X] Refactor vendor modules (use new Go Modules feature)
- [ ] Refactor Project structure, maybe follow https://github.com/golang-standards/project-layout
- [ ] Move fetcher to helpers
- [X] Use go modules
- [ ] JSON responses
- [ ] Unit tests
	- [ ] Taxes
	- [ ] Accounting
- [ ] Documentation
- [ ] Environment variables
- [X] Split Controller and Service
- [ ] It was maybe a bad idea to mix buy/sell operations in the same table but with different types...
