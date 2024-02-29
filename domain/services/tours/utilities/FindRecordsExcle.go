package tours

import (
	"SimonBK_Historical_Vehicles/api/views/inputs"
	"SimonBK_Historical_Vehicles/api/views/outputs"
	"SimonBK_Historical_Vehicles/domain/models"
)

func FindRecordsExcel(tourIn inputs.ToursInputs) ([]outputs.ToursOutputs, error) {
	var records []models.AvlRecord
	db := tourIn.Db.Debug() // Habilitar la impresión de consultas SQL

	// Crear consulta para obtener registros Avl
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

	// Ejecutar la consulta y asignar el resultado a records
	if err := query.Find(&records).Error; err != nil {
		return nil, err
	}

	var responseRecordsPoint []outputs.ToursOutputs
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

	// Devolver responseRecordsPoint en lugar de un objeto ToursOutputs vacío
	return responseRecordsPoint, nil
}
