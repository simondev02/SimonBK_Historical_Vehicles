package views

type AvlRecordPointResponse struct {
	ID             uint    `json:"id"`
	Plate          string  `json:"plate"`
	Imei           string  `json:"iemi"`
	TimeStampEvent string  `json:"timeStampEvent"`
	Location       string  `json:"location"`
	Latitude       float64 `json:"latitude"`
	Longitude      float64 `json:"longitude"`
}
