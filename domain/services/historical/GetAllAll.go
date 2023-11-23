package services

import (
	"SimonBK_Historical_Vehicles/api/views"
	"SimonBK_Historical_Vehicles/domain/models" // Reemplaza "tu_paquete" con el nombre correcto de tu paquete
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

// GetAllAvlRecords obtiene todos los registros Avl
func GetAllHistoricalExcel(db *gorm.DB, FkCompany *int, FkCustomer *int, Plate *string, Imei *string, fromDateStr string, toDateStr string) ([]views.HistoricalExcel, error) {
	// Caso 1: Ambas fechas están vacías
	var fromDate, toDate time.Time
	var err error
	if fromDateStr == "" && toDateStr == "" {
		toDate = time.Now()
		fromDate = toDate.AddDate(0, 0, -15)
	} else if fromDateStr != "" && toDateStr != "" {
		// Caso 2: Ninguna de las fechas está vacía
		fromDate, err = time.Parse(time.RFC3339, fromDateStr)
		if err != nil {
			return []views.HistoricalExcel{}, fmt.Errorf("fecha de inicio inválida: %w", err)
		}
		toDate, err = time.Parse(time.RFC3339, toDateStr)
		if err != nil {
			return []views.HistoricalExcel{}, fmt.Errorf("fecha final inválida: %w", err)
		}
	} else {
		// Caso 3: Solo una de las fechas está vacía
		return []views.HistoricalExcel{}, errors.New("por favor, especifique ambas fechas o deje ambas vacías para usar el rango predeterminado")
	}

	// Asegurarse de que toDate vaya hasta la última hora del día
	toDate = toDate.Add(time.Hour*23 + time.Minute*59 + time.Second*59)

	query := db.Order("time_stamp_event desc")

	// Filtro por Plate
	if Plate != nil && *Plate != "" {
		query = query.Where("plate LIKE ?", "%"+*Plate+"%")
	}

	// Filtro por Imei
	if Imei != nil && *Imei != "" {
		query = query.Where("imei LIKE ?", "%"+*Imei+"%")
	}

	// Filtro por rango de fechas de TimeStampEvent
	query = query.Where("time_stamp_event BETWEEN ? AND ?", fromDate, toDate)

	// Aplicar filtros contextuales basados en FkCompany y FkCustomer
	if FkCompany != nil && *FkCompany != 0 {
		query = query.Where("id_company = ?", *FkCompany)
	}
	if FkCustomer != nil && *FkCustomer != 0 {
		query = query.Where("id_customer = ?", *FkCustomer)
	}

	// Calcular el total de registros
	var total int64
	if err := query.Model(&models.AvlRecord{}).Count(&total).Error; err != nil {
		return []views.HistoricalExcel{}, fmt.Errorf("error al obtener el total de registros Avl: %w", err)
	}

	var records []models.AvlRecord
	if err := query.Find(&records).Error; err != nil {
		return []views.HistoricalExcel{}, err
	}

	var responseRecords []views.HistoricalExcel
	for _, record := range records {
		var properties map[string]interface{}
		err = json.Unmarshal([]byte(record.Properties), &properties)
		if err != nil {
			return nil, fmt.Errorf("error al deserializar Properties: %w", err)
		}

		responseRecord := views.HistoricalExcel{
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
			Properties:     properties,
		}

		responseRecords = append(responseRecords, responseRecord)
	}

	return responseRecords, nil
}
