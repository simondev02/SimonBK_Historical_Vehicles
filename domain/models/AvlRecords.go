package models

import (
	"gorm.io/gorm"
)

type AvlRecord struct {
	gorm.Model
	Plate          string  `json:"plate"`
	Imei           string  `json:"imei"`
	Ip             string  `json:"ip"`
	TimeStampEvent string  `json:"timeStampEvent"`
	Id_company     int     `json:"id_company"`
	Company        string  `json:"company"`
	Id_customer    int     `json:"id_customer"`
	Customer       string  `json:"customer"`
	Location       string  `json:"location"`
	Latitude       float64 `json:"latitude"`
	Longitude      float64 `json:"longitude"`
	Altitude       int     `json:"altitude"`
	Angle          int     `json:"angle"`
	Satellites     int     `json:"satellites"`
	Speed          int     `json:"speed"`
	Hdop           int     `json:"hdop"`
	Pdop           int     `json:"pdop"`
	Event          string  `json:"event"`
	Properties     string  `json:"properties"`
}
