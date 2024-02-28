package tours

import (
	"SimonBK_Historical_Vehicles/api/views"
	"SimonBK_Historical_Vehicles/api/views/inputs"
	"SimonBK_Historical_Vehicles/api/views/outputs"
	"SimonBK_Historical_Vehicles/domain/models"
	"fmt"
)

func FindRecords(tourIn inputs.ToursInputs) (views.Return, error) {
	var records []models.AvlRecord
	db := tourIn.Db.Debug() // Habilitar la impresión de consultas SQL

	// 1 .Verificar si Page y PageSize son nulos y asignar valores predeterminados
	page := uint(1)
	pageSize := uint(300)
	if tourIn.Page != nil {
		page = *tourIn.Page
	}
	if tourIn.PageSize != nil {
		pageSize = *tourIn.PageSize
	}

	// 2. Crear consulta para obtener registros Avl
	query := db.Order("time_stamp_event desc").
		Select("id, plate, imei, time_stamp_event, location, latitude, longitude, altitude").
		Where("time_stamp_event BETWEEN ? AND ?", tourIn.FromDate, tourIn.ToDate)

	if tourIn.FkCompany != nil {
		query = query.Where("id_company = ?", tourIn.FkCompany)
	}

	if tourIn.FkCustomer != nil {
		query = query.Where("id_customer = ?", tourIn.FkCustomer)
	}

	if tourIn.Imei != nil {
		query = query.Where("imei ILIKE ?", "%"+*tourIn.Imei+"%")
	}

	if tourIn.Plate != nil {
		query = query.Where("plate ILIKE ?", "%"+*tourIn.Plate+"%")
	}

	// 3 . Crear una copia de la consulta para contar el total de registros
	countQuery := *query
	var total int64
	if err := countQuery.Model(&models.AvlRecord{}).Count(&total).Error; err != nil {
		return views.Return{}, fmt.Errorf("error al contar registros Avl: %w", err)
	}

	// Aplicar Offset y Limit a la consulta original
	query = query.Offset(int((page - 1) * pageSize)).
		Limit(int(pageSize)).
		Find(&records)

	if query.Error != nil {
		return views.Return{}, fmt.Errorf("error al obtener registros Avl: %w", query.Error)
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
		Page:     int(page),
		PageSize: int(pageSize),
		Total:    int(total),
		Result:   responseRecordsPoint,
	}, nil
}
