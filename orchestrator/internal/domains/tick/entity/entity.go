package tickentity

type Bar struct {
	Symbol     string                        `json:"symbol"`
	Timestamp  int64                         `json:"timestamp"`
	Open       float64                       `json:"open"`
	High       float64                       `json:"high"`
	Low        float64                       `json:"low"`
	Close      float64                       `json:"close"`
	Volume     float64                       `json:"volume"`
	Indicators map[string]map[string]float64 `json:"indicators"`
}
