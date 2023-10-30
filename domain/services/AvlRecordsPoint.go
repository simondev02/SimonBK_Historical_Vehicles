package services

import (
	"SimonBK_Historical_Vehicles/api/views"
	"SimonBK_Historical_Vehicles/domain/models"
	"fmt"
	"time"

	"gorm.io/gorm"
)

// GetAllAvlRecords obtiene todos los registros Avl
func GetAllAvlRecordsPoint(db *gorm.DB, FkCompany *int, FkCustomer *int, Plate *string, Imei *string, dateStr string) ([]views.AvlRecordPointResponse, error) {

	// Comprobar si se proporcionó un parámetro Imei o Plate
	if (Imei == nil || *Imei == "") && (Plate == nil || *Plate == "") {
		return nil, fmt.Errorf("debe proporcionar un parámetro Imei o Plate")
	}

	// Convertir la fecha de string a time.Time
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return nil, err
	}

	// Asegurarse de que date vaya desde el inicio hasta la última hora del día
	fromDate := date
	toDate := date.Add(time.Hour*23 + time.Minute*59 + time.Second*59)

	selectedColumns := []string{
		"id",
		"plate",
		"imei",
		"time_stamp_event",
		"location",
		"latitude",
		"longitude",
		"altitude",
	}

	query := db.Select(selectedColumns).Order("time_stamp_event desc")

	if Plate != nil && *Plate != "" {
		query = query.Where("plate = ?", *Plate)
	}

	if Imei != nil && *Imei != "" {
		query = query.Where("imei = ?", *Imei)
	}

	// Filtro por rango de fechas de TimeStampEvent
	if !fromDate.IsZero() && !toDate.IsZero() {
		query = query.Where("time_stamp_event BETWEEN ? AND ?", fromDate, toDate)
	}

	if *FkCompany != 0 && *FkCustomer == 0 {
		query = query.Where("id_company = ?", *FkCompany)
	} else if *FkCompany != 0 && *FkCustomer != 0 {
		query = query.Where("id_company = ?", *FkCompany).Where("id_customer = ?", *FkCustomer)
	}

	// Activa el modo de depuración para ver la consulta SQL
	query = query.Debug()

	var records []models.AvlRecord
	if err := query.Find(&records).Error; err != nil {
		return nil, fmt.Errorf("error al obtener registros Avl: %w", err)
	}

	var responseRecordsPoint []views.AvlRecordPointResponse
	for _, record := range records {

		formattedTimeStamp, err := FormatTimestamp(record.TimeStampEvent)
		if err != nil {
			return nil, fmt.Errorf("error al formatear la fecha: %w", err)
		}
		responseRecordPoint := views.AvlRecordPointResponse{
			ID:             record.ID,
			Plate:          record.Plate,
			Imei:           record.Imei,
			TimeStampEvent: formattedTimeStamp,
			Location:       record.Location,
			Latitude:       record.Latitude,
			Longitude:      record.Longitude,
		}

		responseRecordsPoint = append(responseRecordsPoint, responseRecordPoint)
	}

	return responseRecordsPoint, nil
}
