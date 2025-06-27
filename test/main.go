package main

import (
	"fmt"
	"log"
	"time"

	"github.com/yamyard/golang-finance/yahoo"
)

func main() {
	data, err := yahoo.GetHistory("AAPL", "1d", "5d")
	if err != nil {
		log.Fatal(err)
	}
	for i, ts := range data.Chart.Result[0].Timestamp {
		close := data.Chart.Result[0].Indicators.Quote[0].Close[i]
		t := time.Unix(ts, 0)
		fmt.Printf("Date: %s, Close Price: %.2f\n", t.Format("2006-01-02"), close)
	}
}