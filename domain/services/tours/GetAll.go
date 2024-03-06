package services

import (
	"SimonBK_Historical_Vehicles/api/views"
	"SimonBK_Historical_Vehicles/api/views/inputs"
	services "SimonBK_Historical_Vehicles/domain/services/utilities"
	"fmt"
)

func GetAllHistoricalTours(params inputs.Params) (views.Return, error) {

	// 1.1 Validar fechas
	fromDate, toDate, err := services.ValidateDates(params)
	if err != nil {
		return views.Return{}, fmt.Errorf("error al validar fechas: %w", err)
	}
	params.FromDate = fromDate
	params.ToDate = toDate

	// 2. Buscar registros Avl
	records, err := FindRecordsTours(params)
	if err != nil {
		return views.Return{}, fmt.Errorf("error al obtener registros Avl: %w", err)
	}
	return records, nil

}
