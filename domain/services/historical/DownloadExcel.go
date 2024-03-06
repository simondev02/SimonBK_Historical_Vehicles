package services

import (
	"SimonBK_Historical_Vehicles/api/views/inputs"
	"SimonBK_Historical_Vehicles/domain/services/utilities"
	"fmt"
	"math"
	"time"

	"github.com/xuri/excelize/v2"
)

func DownloadHistoricalExcel(params inputs.Params) (string, error) {

	// 1.1 Validar fechas
	fromDate, toDate, err := utilities.ValidateDates(params)
	if err != nil {
		return "", fmt.Errorf("error al validar fechas: %w", err)
	}
	params.FromDate = fromDate
	params.ToDate = toDate

	data, err := GetAllHistoricalExcel(params)
	if err != nil {
		return "", err
	}

	file := excelize.NewFile()

	// Añadir encabezado
	headers := []string{"ID", "Placa", "Imei", "Ip", "Fecha del evento", "Id compañia", "Compañia", "Id del cliente", "Cliente", "Ubicación", "Latitud", "Longitud", "Altitud", "Angulo", "Satelite", "Velocidad", "Hdop", "Pdop", "Evento", "Kilometraje Can", "Kilometraje Gps"}
	for i, header := range headers {
		file.SetCellValue("Sheet1", fmt.Sprintf("%c1", 'A'+i), header)
	}

	// Añadir datos
	for i, record := range data {

		row := i + 2 // Comenzar en la segunda fila
		file.SetCellValue("Sheet1", fmt.Sprintf("A%d", row), record.ID)
		file.SetCellValue("Sheet1", fmt.Sprintf("B%d", row), *record.Plate)
		file.SetCellValue("Sheet1", fmt.Sprintf("C%d", row), *record.Imei)
		file.SetCellValue("Sheet1", fmt.Sprintf("D%d", row), *record.Ip)
		if record.TimeStampEvent != nil {
			file.SetCellValue("Sheet1", fmt.Sprintf("E%d", row), record.TimeStampEvent.Format(time.RFC3339))
		}
		file.SetCellValue("Sheet1", fmt.Sprintf("F%d", row), *record.Id_company)
		file.SetCellValue("Sheet1", fmt.Sprintf("G%d", row), *record.Company)
		file.SetCellValue("Sheet1", fmt.Sprintf("H%d", row), *record.Id_customer)
		file.SetCellValue("Sheet1", fmt.Sprintf("I%d", row), *record.Customer)
		file.SetCellValue("Sheet1", fmt.Sprintf("J%d", row), *record.Location)
		file.SetCellValue("Sheet1", fmt.Sprintf("K%d", row), *record.Latitude)
		file.SetCellValue("Sheet1", fmt.Sprintf("L%d", row), *record.Longitude)
		file.SetCellValue("Sheet1", fmt.Sprintf("M%d", row), *record.Altitude)
		file.SetCellValue("Sheet1", fmt.Sprintf("N%d", row), *record.Angle)
		file.SetCellValue("Sheet1", fmt.Sprintf("O%d", row), *record.Satellites)
		file.SetCellValue("Sheet1", fmt.Sprintf("P%d", row), *record.Speed)
		file.SetCellValue("Sheet1", fmt.Sprintf("Q%d", row), *record.Hdop)
		file.SetCellValue("Sheet1", fmt.Sprintf("R%d", row), *record.Pdop)
		file.SetCellValue("Sheet1", fmt.Sprintf("S%d", row), *record.Event)
		file.SetCellValue("Sheet1", fmt.Sprintf("T%d", row), math.Round(float64(*record.TotalMileage)/1000))
		file.SetCellValue("Sheet1", fmt.Sprintf("U%d", row), math.Round(float64(*record.TotalOdometer)/1000))
		file.SetCellValue("Sheet1", fmt.Sprintf("V%d", row), record.Properties)

		// Para las propiedades adicionales, puedes decidir cómo manejarlas según tus necesidades
	}

	filename := fmt.Sprintf("Recorrido_%s_%s_%s.xlsx", *params.Plate, params.FromDate, params.ToDate)
	if err := file.SaveAs(filename); err != nil {
		return "", err
	}

	return filename, nil
}
