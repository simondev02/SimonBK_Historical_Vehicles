package services

import (
	"SimonBK_Historical_Vehicles/api/views"
	"SimonBK_Historical_Vehicles/api/views/inputs"
	"SimonBK_Historical_Vehicles/api/views/outputs"
	"SimonBK_Historical_Vehicles/domain/models"
	"fmt"
)

func FindRecordsTours(tour inputs.Params) (views.Return, error) {

	var records []models.AvlRecord
	db := tour.Db.Debug() // Habilitar la impresión de consultas SQL

	// 2. Crear consulta para obtener registros Avl
	query := db.Order("time_stamp_event desc").
		Select("id, plate, imei, time_stamp_event, location, latitude, longitude, altitude").
		Where("time_stamp_event BETWEEN ? AND ?", tour.FromDate, tour.ToDate)

	if tour.FkCompany != nil && *tour.FkCompany > 0 {
		query = query.Where("id_company = ?", tour.FkCompany)
	}

	if tour.FkCustomer != nil && *tour.FkCustomer > 0 {
		query = query.Where("id_customer = ?", tour.FkCustomer)
	}

	if tour.Imei != nil {
		query = query.Where("imei ILIKE ?", "%"+*tour.Imei+"%")
	}

	if tour.Plate != nil {
		query = query.Where("plate ILIKE ?", "%"+*tour.Plate+"%")
	}

	// 3 . Crear una copia de la consulta para contar el total de registros
	countQuery := *query
	var total int64
	if err := countQuery.Model(&records).Count(&total).Error; err != nil {
		return views.Return{}, fmt.Errorf("error al contar registros Avl: %w", err)
	}

	// Aplicar Offset y Limit a la consulta original solo si Page y PageSize no son 0
	if tour.Page != 0 && tour.PageSize != 0 {
		query = query.Offset((tour.Page - 1) * tour.PageSize).
			Limit(tour.PageSize)
	}

	query = query.Find(&records)

	if query.Error != nil {
		return views.Return{}, fmt.Errorf("error al obtener registros Avl: %w", query.Error)
	}

	if len(records) == 0 {
		return views.Return{}, fmt.Errorf("este vehículo no tiene históricos")
	}

	var responseRecordsPoint []interface{}
	for _, record := range records {
		responseRecordPoint := outputs.ToursOutputs{
			ID:             record.ID,
			Plate:          record.Plate,
			Imei:           record.Imei,
			TimeStampEvent: record.TimeStampEvent,
			Location:       record.Location,
			Latitude:       record.Latitude,
			Longitude:      record.Longitude,
		}

		responseRecordsPoint = append(responseRecordsPoint, responseRecordPoint)
	}

	return views.Return{
		Page:     tour.Page,
		PageSize: tour.PageSize,
		Total:    int(total),
		Result:   responseRecordsPoint,
	}, nil
}
