package services

import (
	"SimonBK_Historical_Vehicles/api/views"
	"SimonBK_Historical_Vehicles/domain/models"
	"encoding/json"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

// GetAvlRecordByID obtiene un registro Avl por ID
func GetAvlRecordByID(db *gorm.DB, id uint) (*views.AvlRecordresponse, error) {
	var record models.AvlRecord
	if err := db.First(&record, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, err
	}

	// Deserializar las propiedades JSON
	var properties map[string]interface{}
	if err := json.Unmarshal([]byte(record.Properties), &properties); err != nil {
		return nil, err
	}

	// Formatear la marca de tiempo
	formattedTimeStamp, err := FormatTimestamp(record.TimeStampEvent)
	if err != nil {
		fmt.Println("Error al formatear la fecha:", err)
		return nil, err
	}

	// Convertir el registro al formato de respuesta
	responseRecord := &views.AvlRecordresponse{
		ID:             record.ID,
		CreatedAt:      record.CreatedAt.String(),
		Plate:          record.Plate,
		Imei:           record.Imei,
		Ip:             record.Ip,
		TimeStampEvent: formattedTimeStamp,
		Id_company:     record.Id_company,
		Company:        record.Company,
		Id_customer:    record.Id_customer,
		Customer:       record.Customer,
		Location:       record.Location,
		Latitude:       record.Latitude,
		Longitude:      record.Longitude,
		Altitude:       record.Altitude,
		Angle:          record.Angle,
		Satellites:     record.Satellites,
		Speed:          record.Speed,
		Hdop:           record.Hdop,
		Pdop:           record.Pdop,
		Event:          record.Event,
		Properties:     properties,
	}

	return responseRecord, nil
}
