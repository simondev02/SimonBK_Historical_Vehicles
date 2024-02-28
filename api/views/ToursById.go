package views

import "time"

type ToursById struct {
	ID             uint       `json:"id"`
	TimeStampEvent *time.Time `json:"timeStampEvent"`
	Location       *string    `json:"location"`
	Event          *string    `json:"event"`
}
