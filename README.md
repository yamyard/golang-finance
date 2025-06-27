# Golang Finance

`golang-finance` is a simple financial data fetching package written in Go.
Now supporting retrieval of historical stock price data from Yahoo Finance.

## Example
```
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/yamyard/golang-finance/yahoo"
)

func main() {
	// Fetch the last 5 days of daily historical data for AAPL
	data, err := yahoo.GetHistory("AAPL", "1d", "5d")
	if err != nil {
		log.Fatal(err)
	}

	// Print date, open, high, low, close, and volume
	for i, ts := range data.Chart.Result[0].Timestamp {
		quote := data.Chart.Result[0].Indicators.Quote[0]
		t := time.Unix(ts, 0)
		fmt.Printf(
			"Date: %s | Open: %.2f | High: %.2f | Low: %.2f | Close: %.2f | Volume: %d\n",
			t.Format("2006-01-02"),
			quote.Open[i],
			quote.High[i],
			quote.Low[i],
			quote.Close[i],
			quote.Volume[i],
		)
	}
}
```
