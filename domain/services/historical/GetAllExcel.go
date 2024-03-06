package services

import (
	"SimonBK_Historical_Vehicles/api/views/inputs"
	"SimonBK_Historical_Vehicles/api/views/outputs"
	"SimonBK_Historical_Vehicles/domain/models" // Reemplaza "tu_paquete" con el nombre correcto de tu paquete
	"encoding/json"
	"fmt"
)

// GetAllAvlRecords obtiene todos los registros Avl
func GetAllHistoricalExcel(params inputs.Params) ([]outputs.HistoricalExcel, error) {

	query := params.Db.Order("time_stamp_event desc")

	// Filtro por Plate
	if params.Plate != nil {
		query = query.Where("plate LIKE ?", "%"+*params.Plate+"%")
	}

	// Filtro por Imei
	if params.Imei != nil && *params.Imei != "" {
		query = query.Where("imei LIKE ?", "%"+*params.Imei+"%")
	}

	// Filtro por rango de fechas de TimeStampEvent
	query = query.Where("time_stamp_event BETWEEN ? AND ?", params.FromDate, params.ToDate)

	// Aplicar filtros contextuales basados en FkCompany y FkCustomer
	if params.FkCompany != nil && *params.FkCompany != 0 {
		query = query.Where("id_company = ?", *params.FkCompany)
	}
	if params.FkCustomer != nil && *params.FkCustomer != 0 {
		query = query.Where("id_customer = ?", *params.FkCustomer)
	}

	// Calcular el total de registros
	var total int64
	if err := query.Model(&models.AvlRecord{}).Count(&total).Error; err != nil {
		return []outputs.HistoricalExcel{}, fmt.Errorf("error al obtener el total de registros Avl: %w", err)
	}

	var records []models.AvlRecord
	if err := query.Find(&records).Error; err != nil {
		return []outputs.HistoricalExcel{}, err
	}

	var responseRecords []outputs.HistoricalExcel
	for _, record := range records {
		var properties map[string]interface{}
		err := json.Unmarshal([]byte(*record.Properties), &properties)
		if err != nil {
			return nil, fmt.Errorf("error al deserializar Properties: %w", err)
		}

		totalMileage, ok := properties["Total Mileage"].(float64)
		if !ok {
			totalMileage = 0 // o cualquier valor predeterminado
		}

		totalOdometer, ok := properties["Total Odometer"].(float64)
		if !ok {
			totalOdometer = 0 // o cualquier valor predeterminado
		}

		responseRecord := outputs.HistoricalExcel{
			ID:             record.ID,
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
			TotalMileage:   &totalMileage,
			TotalOdometer:  &totalOdometer,
			Properties:     properties,
		}

		responseRecords = append(responseRecords, responseRecord)
	}

	return responseRecords, nil
}
