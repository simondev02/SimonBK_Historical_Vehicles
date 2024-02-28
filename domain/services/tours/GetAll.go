package services

import (
	"SimonBK_Historical_Vehicles/api/views"
	"SimonBK_Historical_Vehicles/api/views/inputs"
	tours "SimonBK_Historical_Vehicles/domain/services/tours/utilities"
	"fmt"
)

func GetAllHistoricalTours(tourIn inputs.ToursInputs) (views.Return, error) {

	// 1.1 Validar fechas
	fromDate, toDate, err := tours.ValidateDates(tourIn)
	if err != nil {
		return views.Return{}, fmt.Errorf("error al validar fechas: %w", err)
	}
	tourIn.FromDate = fromDate
	tourIn.ToDate = toDate

	// 2. Buscar registros Avl
	records, err := tours.FindRecords(tourIn)
	if err != nil {
		return views.Return{}, fmt.Errorf("error al obtener registros Avl: %w", err)
	}
	return records, nil

}
