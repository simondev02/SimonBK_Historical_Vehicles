package outputs

import "time"

type HistoricalExcel struct {
	ID             uint                   `json:"id"`
	Plate          *string                `json:"plate"`
	Imei           *string                `json:"imei"`
	Ip             *string                `json:"ip"`
	TimeStampEvent *time.Time             `json:"timeStampEvent"`
	Id_company     *uint                  `json:"id_company"`
	Company        *string                `json:"company"`
	Id_customer    *uint                  `json:"id_customer"`
	Customer       *string                `json:"customer"`
	Location       *string                `json:"location"`
	Latitude       *float64               `json:"latitude"`
	Longitude      *float64               `json:"longitude"`
	Altitude       *uint                  `json:"altitude"`
	Angle          *uint                  `json:"angle"`
	Satellites     *uint                  `json:"satellites"`
	Speed          *uint                  `json:"speed"`
	Hdop           *uint                  `json:"hdop"`
	Pdop           *uint                  `json:"pdop"`
	Event          *string                `json:"event"`
	TotalMileage   *float64               `json:"Total Mileage"`
	TotalOdometer  *float64               `json:"Total Odometer"`
	Properties     map[string]interface{} `json:"properties"`
}
