package services

import (
	"SimonBK_Historical_Vehicles/api/views/inputs"
	tours "SimonBK_Historical_Vehicles/domain/services/tours/utilities"
	"fmt"

	"github.com/xuri/excelize/v2"
)

func DownloadHistoricalToursExcel(tourIn inputs.ToursInputs) (string, error) {

	fromDate, toDate, err := tours.ValidateDates(tourIn)
	if err != nil {
		return "", fmt.Errorf("error al validar fechas: %w", err)
	}

	data, err := tours.FindRecordsExcel(tourIn)
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
	// Añadir datos
	for i, record := range data {
		file.SetCellValue("Sheet1", fmt.Sprintf("A%d", i+2), record.ID)
		file.SetCellValue("Sheet1", fmt.Sprintf("B%d", i+2), *record.Plate)
		file.SetCellValue("Sheet1", fmt.Sprintf("C%d", i+2), *record.Imei)
		file.SetCellValue("Sheet1", fmt.Sprintf("D%d", i+2), record.TimeStampEvent)
		file.SetCellValue("Sheet1", fmt.Sprintf("E%d", i+2), *record.Location)
		file.SetCellValue("Sheet1", fmt.Sprintf("F%d", i+2), *record.Latitude)
		file.SetCellValue("Sheet1", fmt.Sprintf("G%d", i+2), *record.Longitude)
	}

	var identifier string
	if tourIn.Plate != nil {
		identifier = *tourIn.Plate
	} else {
		identifier = *tourIn.Imei
	}

	filename := fmt.Sprintf("Recorrido_%s_%s_%s.xlsx", identifier, fromDate, toDate)
	if err := file.SaveAs(filename); err != nil {
		return "", err
	}

	return filename, nil
}
