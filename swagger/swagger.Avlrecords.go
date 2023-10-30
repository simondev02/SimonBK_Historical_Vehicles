package swagger

type AvlRecord struct {
	Plate          string  `json:"Plate"`
	Imei           string  `json:"Imei"`
	Ip             string  `json:"Ip"`
	TimeStampEvent string  `json:"TimeStampEvent"`
	Id_company     int     `json:"id_company"`
	Company        string  `json:"Company"`
	Id_customer    int     `json:"id_customer"`
	Customer       string  `json:"Customer"`
	Location       string  `json:"Location"`
	Latitude       float64 `json:"Latitude"`
	Longitude      float64 `json:"Longitude"`
	Altitude       int     `json:"Altitude"`
	Angle          int     `json:"Angle"`
	Satellites     int     `json:"Satellites"`
	Speed          int     `json:"Speed"`
	Hdop           int     `json:"Hdop"`
	Pdop           int     `json:"Pdop"`
	Event          string  `json:"Event"`
	Properties     string  `json:"Properties"`
}
