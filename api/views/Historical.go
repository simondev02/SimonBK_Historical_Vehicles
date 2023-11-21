package views

import "time"

type Historical struct {
	ID             uint       `json:"id"`
	Plate          string     `json:"plate"`
	TimeStampEvent *time.Time `json:"timeStampEvent"`
	Location       string     `json:"location"`
	Speed          int        `json:"speed"`
	Event          string     `json:"event"`
}
