package views

type HistoricalById struct {
	ID         uint                   `json:"id"`
	Properties map[string]interface{} `json:"properties"`
}
