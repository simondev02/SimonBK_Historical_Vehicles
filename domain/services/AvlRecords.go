package services

import (
	"SimonBK_Historical_Vehicles/api/views"
	"SimonBK_Historical_Vehicles/domain/models" // Reemplaza "tu_paquete" con el nombre correcto de tu paquete
	"encoding/json"
	"fmt"
	"time"

	"gorm.io/gorm"
)

// GetAllAvlRecords obtiene todos los registros Avl
func GetAllAvlRecords(db *gorm.DB, FkCompany *int, FkCustomer *int, page int, pageSize int, Plate *string, Imei *string, fromDateStr string, toDateStr string) (views.AvlRecord, error) {
	// Convertir las fechas de string a time.Time
	fromDate, err := time.Parse("2006-01-02", fromDateStr)
	if err != nil {
		return views.AvlRecord{}, err
	}
	toDate, err := time.Parse("2006-01-02", toDateStr)
	if err != nil {
		return views.AvlRecord{}, err
	}

	// Asegurarse de que toDate vaya hasta la última hora del día
	toDate = toDate.Add(time.Hour*23 + time.Minute*59 + time.Second*59)

	// Calcular el offset basado en la página y el tamaño de página
	offset := (page - 1) * pageSize

	query := db.Order("time_stamp_event desc").Offset(offset).Limit(pageSize)

	// Filtro por Plate
	if Plate != nil && *Plate != "" {
		query = query.Where("plate = ?", *Plate)
	}

	// Filtro por Imei
	if Imei != nil && *Imei != "" {
		query = query.Where("imei = ?", *Imei)
	}

	// Filtro por rango de fechas de TimeStampEvent
	if !fromDate.IsZero() && !toDate.IsZero() {
		query = query.Where("time_stamp_event BETWEEN ? AND ?", fromDate, toDate)
	}

	// Apply contextual filters based on FkCompany and FkCustomer
	if *FkCompany != 0 && *FkCustomer == 0 {
		query = query.Where("id_company = ?", *FkCompany)
	} else if *FkCompany != 0 && *FkCustomer != 0 {
		query = query.Where("id_company = ?", *FkCompany).Where("id_customer = ?", *FkCustomer)
	}

	// Calcular el total de registros
	var total int64
	if err := query.Model(&models.AvlRecord{}).Count(&total).Error; err != nil {
		return views.AvlRecord{}, fmt.Errorf("error al obtener el total de registros Avl: %w", err)
	}
	// Aplicar paginación
	query = query.Offset((page - 1) * pageSize).Limit(pageSize)

	var records []models.AvlRecord
	if err := query.Find(&records).Error; err != nil {
		return views.AvlRecord{}, err
	}

	var responseRecords []views.AvlRecordresponse
	for _, record := range records {
		var properties map[string]interface{}
		if err := json.Unmarshal([]byte(record.Properties), &properties); err != nil {
			return views.AvlRecord{}, err
		}
		formattedTimeStamp, err := FormatTimestamp(record.TimeStampEvent)
		if err != nil {
			fmt.Println("Error al formatear la fecha:", err)
			return views.AvlRecord{}, err
		}
		responseRecord := views.AvlRecordresponse{
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

		responseRecords = append(responseRecords, responseRecord)
	}

	return views.AvlRecord{
		Page:     page,
		PageSize: pageSize,
		Total:    int(total),
		Result:   responseRecords,
	}, nil
}
