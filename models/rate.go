package models

// Rate ...
type Rate struct {
	Base      string             `json:"base"`
	Date      string             `json:"date"` // Not using anything with Time as this serves only to get the latest rate
	Rates     map[string]float64 `json:"rates"`
	ShouldBuy bool               `json:"should_buy"`
}

// HistoricRates ...
type HistoricRates struct {
	Base      string                        `json:"base"`
	StartDate string                        `json:"start_at"`
	EndDate   string                        `json:"end_at"`
	Rates     map[string]map[string]float64 `json:"rates"`
}
