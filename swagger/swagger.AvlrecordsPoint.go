package swagger

type AvlRecordPoint struct {
	ID             uint    `json:"id"`
	Plate          string  `json:"Plate"`
	TimeStampEvent string  `json:"TimeStampEvent"`
	Location       string  `json:"Location"`
	Latitude       float64 `json:"Latitude"`
	Longitude      float64 `json:"Longitude"`
}
