// Package stocks is defining a Go package named `stocks`. This package can contain functions,
// variables, and other Go code that can be used to work with stocks and financial data.
package stocks

// This is defining a constant variable named `yahooURL` with the value of the Yahoo Finance API URL.
// The `%s` is a placeholder for the stock symbol that will be added to the URL when making a request
// to the API.
const (
	yahooURL = "https://query1.finance.yahoo.com/v7/finance/quote?symbols=%s"
)
