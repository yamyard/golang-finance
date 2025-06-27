package yahoo

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

type ChartResponse struct {
    Chart struct {
        Result []struct {
            Timestamp []int64 `json:"timestamp"`
            Indicators struct {
                Quote []struct {
                    Close  []float64 `json:"close"`
                    Open   []float64 `json:"open"`
                    High   []float64 `json:"high"`
                    Low    []float64 `json:"low"`
                    Volume []int64   `json:"volume"`
                } `json:"quote"`
            } `json:"indicators"`
        } `json:"result"`
    } `json:"chart"`
}

func GetHistory(symbol, interval, rangeStr string) (*ChartResponse, error) {
    url := fmt.Sprintf("https://query1.finance.yahoo.com/v8/finance/chart/%s?interval=%s&range=%s", symbol, interval, rangeStr)
    req, _ := http.NewRequest("GET", url, nil)
    req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; MyGoClient/1.0)")
    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)

    if resp.StatusCode != 200 {
        return nil, fmt.Errorf("HTTP Status: %d, body: %s", resp.StatusCode, string(body))
    }

    var data ChartResponse
    err = json.Unmarshal(body, &data)
    if err != nil {
        return nil, fmt.Errorf("json decode error: %v, body: %s", err, string(body))
    }
    return &data, nil
}