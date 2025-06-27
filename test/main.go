package main

import (
	"fmt"
	"log"
	"time"

	"github.com/yamyard/golang-finance/yahoo"
)

func main() {
	// 调用 yahoo 包的 GetHistory 方法，获取 AAPL 最近 5 天的历史行情数据
	data, err := yahoo.GetHistory("AAPL", "1d", "5d")
	if err != nil {
		// 如果发生错误，打印并退出程序
		log.Fatal(err)
	}
	
	// 遍历所有时间戳，输出对应的收盘价
	for i, ts := range data.Chart.Result[0].Timestamp {
		close := data.Chart.Result[0].Indicators.Quote[0].Close[i]
		// 将 Unix 时间戳转换为可读日期
		t := time.Unix(ts, 0)
		fmt.Printf("Date: %s, Close Price: %.2f\n", t.Format("2006-01-02"), close)
	}
}