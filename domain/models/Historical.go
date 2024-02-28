package models

import (
	"time"

	"gorm.io/gorm"
)

type AvlRecord struct {
	gorm.Model
	Plate          *string    `json:"plate"`
	Imei           *string    `json:"imei"`
	Ip             *string    `json:"ip"`
	TimeStampEvent *time.Time `json:"time_stamp_event"`
	Id_company     *uint      `json:"id_company"`
	Company        *string    `json:"company"`
	Id_customer    *uint      `json:"id_customer"`
	Customer       *string    `json:"customer"`
	Location       *string    `json:"location"`
	Latitude       *float64   `json:"latitude"`
	Longitude      *float64   `json:"longitude"`
	Altitude       *uint      `json:"altitude"`
	Angle          *uint      `json:"angle"`
	Satellites     *uint      `json:"satellites"`
	Speed          *uint      `json:"speed"`
	Hdop           *uint      `json:"hdop"`
	Pdop           *uint      `json:"pdop"`
	Event          *string    `json:"event"`
	Properties     *string    `json:"properties"`
}
