package utilities

import (
	"SimonBK_Historical_Vehicles/api/views/inputs"
	"fmt"
	"time"
)

func ValidateDates(param inputs.Params) (*time.Time, *time.Time, error) {

	// 1. si no se dan fechas, se establece un rango predeterminado desde la hora actual hasta 8 horas atrás
	if param.FromDate == nil && param.ToDate == nil {

		toDate, err := GetLastRecordDateByPlateOrImei(param.Db, param.Plate, param.Imei)
		if err != nil {
			return nil, nil, err
		}
		fromDate := toDate.AddDate(0, 0, -1)
		param.FromDate = &fromDate
		param.ToDate = toDate
		// 2 . se aplican las dos fechasa
	} else {
		if param.FromDate.After(*param.ToDate) {
			return nil, nil, fmt.Errorf("la fecha de inicio no puede ser posterior a la fecha final")
		}
	}

	return param.FromDate, param.ToDate, nil
}
