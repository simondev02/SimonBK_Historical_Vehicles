package tours

import (
	"SimonBK_Historical_Vehicles/api/views/inputs"
	"fmt"
	"time"
)

func ValidateDates(tourIn inputs.ToursInputs) (*time.Time, *time.Time, error) {

	// 1. si no se dan fechas, se establece un rango predeterminado desde la hora actual hasta 8 horas atrás
	if tourIn.FromDate == nil && tourIn.ToDate == nil {
		toDate, err := GetLastRecordDateByPlateOrImei(tourIn.Db, tourIn.Plate, tourIn.Imei)
		if err != nil {
			return nil, nil, err
		}
		fromDate := toDate.AddDate(0, 0, -1)
		tourIn.FromDate = &fromDate
		tourIn.ToDate = toDate
		// 2 . se aplican las dos fechasa
	} else {
		if tourIn.FromDate.After(*tourIn.ToDate) {
			return nil, nil, fmt.Errorf("la fecha de inicio no puede ser posterior a la fecha final")
		}
	}

	return tourIn.FromDate, tourIn.ToDate, nil
}
