package views

type Salida struct {
	Page     int                 `json:"page"`
	PageSize int                 `json:"pageSize"`
	Total    int                 `json:"total"`
	Result   []AvlRecordresponse `json:"result"`
}
type SalidaPoint struct {
	Page     int                      `json:"page"`
	PageSize int                      `json:"pageSize"`
	Total    int                      `json:"total"`
	Result   []AvlRecordPointResponse `json:"result"`
}
