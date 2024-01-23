package views

import "time"

type HistoricalExcel struct {
	ID             uint                   `json:"id"`
	Plate          string                 `json:"plate"`
	Imei           string                 `json:"imei"`
	Ip             string                 `json:"ip"`
	TimeStampEvent *time.Time             `json:"timeStampEvent"`
	Id_company     int                    `json:"id_company"`
	Company        string                 `json:"company"`
	Id_customer    int                    `json:"id_customer"`
	Customer       string                 `json:"customer"`
	Location       string                 `json:"location"`
	Latitude       float64                `json:"latitude"`
	Longitude      float64                `json:"longitude"`
	Altitude       int                    `json:"altitude"`
	Angle          int                    `json:"angle"`
	Satellites     int                    `json:"satellites"`
	Speed          int                    `json:"speed"`
	Hdop           int                    `json:"hdop"`
	Pdop           int                    `json:"pdop"`
	Event          string                 `json:"event"`
	TotalMileage   float64                `json:"Total Mileage"`
	TotalOdometer  float64                `json:"Total Odometer"`
	Properties     map[string]interface{} `json:"properties"`
}
