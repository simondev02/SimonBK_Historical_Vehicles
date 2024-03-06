package services

import (
	"SimonBK_Historical_Vehicles/api/views"
	"SimonBK_Historical_Vehicles/api/views/inputs"
	"SimonBK_Historical_Vehicles/domain/models"
	"encoding/json"
	"fmt"
	"log"
)

func FindRecordsHistorical(tour inputs.Params) (views.Return, error) {

	var records []models.AvlRecord
	query := tour.Db.Order("time_stamp_event desc")

	// Filtro por Plate
	if tour.Plate != nil {
		query = query.Where("plate ILIKE ?", "%"+*tour.Plate+"%")
	}

	// Filtro por Imei
	if tour.Imei != nil {
		query = query.Where("imei LIKE ?", "%"+*tour.Imei+"%")
	}

	// Filtro por rango de fechas de TimeStampEvent
	query = query.Where("time_stamp_event BETWEEN ? AND ?", tour.FromDate, tour.ToDate)

	// Aplicar filtros contextuales basados en FkCompany y FkCustomer
	if tour.FkCompany != nil && *tour.FkCompany != 0 {
		query = query.Where("id_company = ?", *tour.FkCompany)
	}
	if tour.FkCustomer != nil && *tour.FkCustomer != 0 {
		query = query.Where("id_customer = ?", *tour.FkCustomer)
	}

	// Calcular el total de registros
	var total int64
	if err := query.Model(&records).Count(&total).Error; err != nil {
		return views.Return{}, fmt.Errorf("error al obtener el total de registros Avl: %w", err)
	}

	if err := query.Find(&records).Error; err != nil {
		return views.Return{}, err
	}
	// Aplicar Offset y Limit a la consulta original
	query = query.Debug().Offset((tour.Page - 1) * tour.PageSize).
		Limit(tour.PageSize).
		Find(&records)

	if query.Error != nil {
		log.Println(query.Error)
		return views.Return{}, fmt.Errorf("error al obtener registros Avl: %w", query.Error)
	}

	if len(records) == 0 {
		return views.Return{}, fmt.Errorf("registros no encontrados")
	}

	var responseRecords []interface{}
	for _, record := range records {
		var properties views.Properties
		err := json.Unmarshal([]byte(*record.Properties), &properties)
		if err != nil {
			return views.Return{}, fmt.Errorf("error al deserializar las propiedades: %w", err)
		}

		var totalMileage *float64
		if properties.TotalMileage != nil {
			totalMileage = properties.TotalMileage
		}

		responseRecord := views.Historical{
			ID:             record.ID,
			Plate:          record.Plate,
			TimeStampEvent: record.TimeStampEvent,
			Location:       record.Location,
			Speed:          record.Speed,
			Event:          record.Event,
			TotalMileage:   totalMileage,
			TotalOdometer:  properties.TotalOdometer,
		}

		responseRecords = append(responseRecords, responseRecord)
	}

	return views.Return{
		Page:     tour.Page,
		PageSize: tour.PageSize,
		Total:    int(total),
		Result:   responseRecords,
	}, nil
}
