package services

import (
	"SimonBK_Historical_Vehicles/api/views"
	"SimonBK_Historical_Vehicles/api/views/inputs" // Reemplaza "tu_paquete" con el nombre correcto de tu paquete
	"SimonBK_Historical_Vehicles/domain/services/utilities"
	"fmt"
)

// GetAllAvlRecords obtiene todos los registros Avl
func GetAllHistorical(params inputs.Params) (views.Return, error) {

	// 1.1 Validar fechas
	fromDate, toDate, err := utilities.ValidateDates(params)
	if err != nil {
		return views.Return{}, fmt.Errorf("error al validar fechas: %w", err)
	}
	params.FromDate = fromDate
	params.ToDate = toDate

	// 2. Buscar registros Avl
	records, err := FindRecordsHistorical(params)
	if err != nil {
		return views.Return{}, fmt.Errorf("error al obtener registros Avl: %w", err)
	}
	return records, nil
}
