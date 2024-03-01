package swagger

type AvlRecordPoint struct {
	ID             uint    `json:"id"`
	Plate          string  `json:"plate"`
	TimeStampEvent string  `json:"timeStampEvent"`
	Location       string  `json:"location"`
	Latitude       float64 `json:"latitude"`
	Longitude      float64 `json:"longitude"`
}
