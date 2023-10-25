package services

import (
	"SimonBK_Historical_Vehicles/api/views"
	"SimonBK_Historical_Vehicles/domain/models" // Reemplaza "tu_paquete" con el nombre correcto de tu paquete
	"encoding/json"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

// GetAllAvlRecords obtiene todos los registros Avl
func GetAllAvlRecords(db *gorm.DB, FkCompany *int, FkCustomer *int) ([]views.AvlRecordresponse, error) {
	query := db.Order("time_stamp_event desc").Limit(10)

	fmt.Println("FkCompany: ", *FkCompany)
	fmt.Println("FkCustomer: ", *FkCustomer)

	// Apply contextual filters based on FkCompany and FkCustomer
	if FkCompany != nil && *FkCompany != 0 && (FkCustomer == nil || *FkCustomer == 0) {
		query = query.Where("id_company = ?", *FkCompany)
	} else if FkCompany != nil && *FkCompany != 0 && FkCustomer != nil && *FkCustomer != 0 {
		query = query.Where("id_company = ?", *FkCompany).Where("id_customer = ?", *FkCustomer)
	}

	var records []models.AvlRecord
	if err := query.Find(&records).Error; err != nil {
		return nil, err
	}

	var responseRecords []views.AvlRecordresponse
	for _, record := range records {
		var properties map[string]interface{}
		if err := json.Unmarshal([]byte(record.Properties), &properties); err != nil {
			return nil, err
		}
		responseRecord := views.AvlRecordresponse{
			ID:             record.ID,
			CreatedAt:      record.CreatedAt.String(),
			UpdatedAt:      record.UpdatedAt.String(),
			DeletedAt:      record.DeletedAt.Time.String(),
			Plate:          record.Plate,
			Imei:           record.Imei,
			Ip:             record.Ip,
			TimeStampEvent: record.TimeStampEvent,
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

		responseRecords = append(responseRecords, responseRecord)
	}

	return responseRecords, nil
}

// GetAvlRecordByID obtiene un registro Avl por ID
func GetAvlRecordByID(db *gorm.DB, id uint) (*models.AvlRecord, error) {
	var record models.AvlRecord
	if err := db.First(&record, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, err
	}
	return &record, nil
}
