package tickdto

import "github.com/goccy/go-json"

type BarRequestType string

const (
	Subscribe   BarRequestType = "subscribe"
	Unsubscribe BarRequestType = "unsubscribe"
)

type BarRequest struct {
	RequestType BarRequestType `json:"request_type"`
	Timeframe   uint8          `json:"timeframe"`
	Symbol      string         `json:"symbol"`
	Indicators  []string       `json:"indicators"`
}

func NewBarRequest(requestType BarRequestType, symbol string, timeframe uint8, indicators []string) ([]byte, error) {
	req := BarRequest{
		RequestType: requestType,
		Timeframe:   timeframe,
		Symbol:      symbol,
		Indicators:  indicators,
	}
	return json.Marshal(req)
}
