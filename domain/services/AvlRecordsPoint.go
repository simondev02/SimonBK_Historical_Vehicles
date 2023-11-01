package services

import (
	"SimonBK_Historical_Vehicles/api/views"
	"SimonBK_Historical_Vehicles/domain/models"
	"fmt"
	"time"

	"gorm.io/gorm"
)

func GetAllAvlRecordsPoint(db *gorm.DB, FkCompany *int, FkCustomer *int, Plate *string, Imei *string, fromDateStr string, toDateStr string, page int, pageSize int) (views.SalidaPoint, error) {

	// Comprobar si se proporcionó un parámetro Imei o Plate
	if (Imei == nil || *Imei == "") && (Plate == nil || *Plate == "") {
		return views.SalidaPoint{}, fmt.Errorf("debe proporcionar un parámetro Imei o Plate")
	}

	// Convertir la fecha de string a time.Time
	var fromDate, toDate time.Time
	var err error

	if fromDateStr == "" && toDateStr == "" {
		// Si no se proporcionan las fechas, establecer un rango predeterminado desde la hora actual hasta 8 horas atrás
		toDate = time.Now()
		fromDate = toDate.Add(-8 * time.Hour)
	} else {
		// Si se proporcionan las fechas, convertirlas de string a time.Time
		fromDate, err = time.Parse("2006-01-02 15:04:05", fromDateStr)
		if err != nil {
			return views.SalidaPoint{}, err
		}
		toDate, err = time.Parse("2006-01-02 15:04:05", toDateStr)
		if err != nil {
			return views.SalidaPoint{}, err
		}
	}

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

	// Calcular el total de registros
	var total int64
	if err := query.Model(&models.AvlRecord{}).Count(&total).Error; err != nil {
		return views.SalidaPoint{}, fmt.Errorf("error al obtener el total de registros Avl: %w", err)
	}

	// Aplicar paginación
	query = query.Offset((page - 1) * pageSize).Limit(pageSize)

	// Activa el modo de depuración para ver la consulta SQL
	query = query.Debug()

	var records []models.AvlRecord
	if err := query.Find(&records).Error; err != nil {
		return views.SalidaPoint{}, fmt.Errorf("error al obtener registros Avl: %w", err)
	}

	var responseRecordsPoint []views.AvlRecordPointResponse
	for _, record := range records {

		formattedTimeStamp, err := FormatTimestamp(record.TimeStampEvent)
		if err != nil {
			return views.SalidaPoint{}, fmt.Errorf("error al formatear la fecha: %w", err)
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

	return views.SalidaPoint{
		Page:     page,
		PageSize: pageSize,
		Total:    int(total),
		Result:   responseRecordsPoint,
	}, nil
}
