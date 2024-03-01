package views

type Return struct {
	Page     int           `json:"page"`
	PageSize int           `json:"pageSize"`
	Total    int           `json:"total"`
	Result   []interface{} `json:"result"`
}
