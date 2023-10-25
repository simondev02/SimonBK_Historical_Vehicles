package views

type APIResponse struct {
	Status  string            `json:"status"`
	Message string            `json:"message"`
	Data    AvlRecordresponse `json:"data"`
}
