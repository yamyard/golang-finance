package yahoo

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

// ChartResponse 用于反序列化 Yahoo Finance 返回的 JSON 数据
type ChartResponse struct {
    Chart struct {
        Result []struct {
			// 时间戳（秒级）
            Timestamp []int64 `json:"timestamp"`
            Indicators struct {
                Quote []struct {
					// 收盘价
                    Close  []float64 `json:"close"`
					// 开盘价
                    Open   []float64 `json:"open"`
					// 最高价
                    High   []float64 `json:"high"`
					// 最低价
                    Low    []float64 `json:"low"`
					// 成交量
                    Volume []int64   `json:"volume"`
                } `json:"quote"`
            } `json:"indicators"`
        } `json:"result"`
    } `json:"chart"`
}

// GetHistory 从 Yahoo Finance 获取指定股票的历史数据
// symbol: 股票代码，如 "AAPL"
// interval: 间隔，如 "1d"（一天）
// rangeStr: 数据范围，如 "5d"（五天）
//
// 返回值：
// - *ChartResponse: 包含历史行情数据的结构体
// - error: 错误信息（如有）
func GetHistory(symbol, interval, rangeStr string) (*ChartResponse, error) {
	// 构建请求 URL
    url := fmt.Sprintf("https://query1.finance.yahoo.com/v8/finance/chart/%s?interval=%s&range=%s", symbol, interval, rangeStr)
    
	// 创建 HTTP 请求并添加 User-Agent，防止被反爬虫拦截
	req, _ := http.NewRequest("GET", url, nil)
    req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; MyGoClient/1.0)")
    
	// 发送 HTTP 请求
	resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    
	// 读取响应体
	body, _ := ioutil.ReadAll(resp.Body)

	// 检查 HTTP 状态码
    if resp.StatusCode != 200 {
        return nil, fmt.Errorf("HTTP Status: %d, body: %s", resp.StatusCode, string(body))
    }

	// 解析 JSON 数据到 ChartResponse 结构体
    var data ChartResponse
    err = json.Unmarshal(body, &data)
    if err != nil {
        return nil, fmt.Errorf("json decode error: %v, body: %s", err, string(body))
    }
    return &data, nil
}