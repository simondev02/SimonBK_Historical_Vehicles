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
func GetAllHistorical(db *gorm.DB, FkCompany *int, FkCustomer *int, page int, pageSize int, Plate *string, Imei *string, fromDateStr string, toDateStr string) (views.Return, error) {
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
			return views.Return{}, fmt.Errorf("fecha de inicio inválida: %w", err)
		}
		toDate, err = time.Parse(time.RFC3339, toDateStr)
		if err != nil {
			return views.Return{}, fmt.Errorf("fecha final inválida: %w", err)
		}
	} else {
		// Caso 3: Solo una de las fechas está vacía
		return views.Return{}, errors.New("por favor, especifique ambas fechas o deje ambas vacías para usar el rango predeterminado")
	}

	// Asegurarse de que toDate vaya hasta la última hora del día
	toDate = toDate.Add(time.Hour*23 + time.Minute*59 + time.Second*59)

	// Calcular el offset basado en la página y el tamaño de página
	offset := (page - 1) * pageSize

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
		return views.Return{}, fmt.Errorf("error al obtener el total de registros Avl: %w", err)
	}

	// Aplicar paginación
	query = query.Offset(offset).Limit(pageSize)

	var records []models.AvlRecord
	if err := query.Find(&records).Error; err != nil {
		return views.Return{}, err
	}

	var responseRecords []interface{}
	for _, record := range records {
		var properties views.Properties
		err := json.Unmarshal([]byte(record.Properties), &properties)
		if err != nil {
			return views.Return{}, fmt.Errorf("error al deserializar las propiedades: %w", err)
		}

		var totalMileage *float64
		if properties.TotalMileage != 0 {
			totalMileage = &properties.TotalMileage
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
		Page:     page,
		PageSize: pageSize,
		Total:    int(total),
		Result:   responseRecords,
	}, nil
}
