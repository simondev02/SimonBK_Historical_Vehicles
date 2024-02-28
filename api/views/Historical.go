package views

import "time"

type Historical struct {
	ID             uint       `json:"id"`
	Plate          *string    `json:"plate"`
	TimeStampEvent *time.Time `json:"timeStampEvent"`
	Location       *string    `json:"location"`
	Speed          *uint      `json:"speed"`
	Event          *string    `json:"event"`
	TotalMileage   *float64   `json:"totalMileage"`
	TotalOdometer  *float64   `json:"totalOdometer"`
}

type Properties struct {
	TotalMileage  *float64 `json:"Total Mileage"`
	TotalOdometer *float64 `json:"Total Odometer"`
}
