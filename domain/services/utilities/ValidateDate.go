package utilities

import (
	"SimonBK_Historical_Vehicles/api/views/inputs"
	"fmt"
	"time"
)

func ValidateDates(param inputs.Params) (*time.Time, *time.Time, error) {

	// Si solo una de las fechas está presente, devuelve un error
	if (param.FromDate == nil && param.ToDate != nil) || (param.FromDate != nil && param.ToDate == nil) {
		return nil, nil, fmt.Errorf("ambas fechas deben estar presentes")
	}
	// 1. si no se dan fechas, se establece un rango predeterminado desde la hora actual hasta 8 horas atrás
	if param.FromDate == nil && param.ToDate == nil {

		toDate, err := GetLastRecordDateByPlateOrImei(param.Db, param.Plate, param.Imei)
		if err != nil {
			now := time.Now()
			fromDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
			toDate := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location())
			param.FromDate = &fromDate
			param.ToDate = &toDate
		} else {
			fromDate := toDate.AddDate(0, 0, -1)
			fromDate = time.Date(fromDate.Year(), fromDate.Month(), fromDate.Day(), 0, 0, 0, 0, fromDate.Location())
			*toDate = time.Date(toDate.Year(), toDate.Month(), toDate.Day(), 23, 59, 59, 0, toDate.Location())
			param.FromDate = &fromDate
			param.ToDate =
				toDate
		}

		// 2 . se aplican las dos fechasa
	} else {
		if param.FromDate.After(*param.ToDate) {
			return nil, nil, fmt.Errorf("la fecha de inicio no puede ser posterior a la fecha final")
		}
		fromDate := time.Date(param.FromDate.Year(), param.FromDate.Month(), param.FromDate.Day(), 0, 0, 0, 0, param.FromDate.Location())
		toDate := time.Date(param.ToDate.Year(), param.ToDate.Month(), param.ToDate.Day(), 23, 59, 59, 0, param.ToDate.Location())
		param.FromDate = &fromDate
		param.ToDate = &toDate
	}

	return param.FromDate, param.ToDate, nil
}
