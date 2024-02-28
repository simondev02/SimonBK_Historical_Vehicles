package services

import (
	"SimonBK_Historical_Vehicles/api/views/inputs"
	"SimonBK_Historical_Vehicles/api/views/outputs"
	tours "SimonBK_Historical_Vehicles/domain/services/tours/utilities"
	"fmt"

	"github.com/xuri/excelize/v2"
)

func DownloadHistoricalToursExcel(tourIn inputs.ToursInputs) (string, error) {

	fromDate, toDate, err := tours.ValidateDates(tourIn)
	if err != nil {
		return "", fmt.Errorf("error al validar fechas: %w", err)
	}

	data, err := GetAllHistoricalTours(tourIn)
	if err != nil {
		return "", err
	}
	file := excelize.NewFile()
	// Añadir encabezados
	file.SetCellValue("Sheet1", "A1", "ID")
	file.SetCellValue("Sheet1", "B1", "Placa")
	file.SetCellValue("Sheet1", "C1", "Imei")
	file.SetCellValue("Sheet1", "D1", "Fecha de Evento")
	file.SetCellValue("Sheet1", "E1", "Ubicación")
	file.SetCellValue("Sheet1", "F1", "Latitud")
	file.SetCellValue("Sheet1", "G1", "Longitud")

	// Añadir datos
	for i, intf := range data.Result {
		record, ok := intf.(outputs.ToursOutputs)
		if !ok {
			return "", fmt.Errorf("no se pudo convertir el registro %d a Historical", i)
		}
		row := i + 2
		file.SetCellValue("Sheet1", fmt.Sprintf("A%d", row), record.ID)
		file.SetCellValue("Sheet1", fmt.Sprintf("B%d", row), *record.Plate)
		file.SetCellValue("Sheet1", fmt.Sprintf("C%d", row), *record.Imei)
		file.SetCellValue("Sheet1", fmt.Sprintf("D%d", row), record.TimeStampEvent)
		file.SetCellValue("Sheet1", fmt.Sprintf("E%d", row), *record.Location)
		file.SetCellValue("Sheet1", fmt.Sprintf("F%d", row), *record.Latitude)
		file.SetCellValue("Sheet1", fmt.Sprintf("G%d", row), *record.Longitude)
	}

	filename := fmt.Sprintf("Recorrido_%s_%s_%s.xlsx", *tourIn.Plate, fromDate, toDate)
	if err := file.SaveAs(filename); err != nil {
		return "", err
	}

	return filename, nil
}
