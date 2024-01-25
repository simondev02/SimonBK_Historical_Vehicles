package services

import (
	"SimonBK_Historical_Vehicles/api/views"
	"fmt"

	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
)

func DownloadHistoricalToursExcel(db *gorm.DB, FkCompany *int, FkCustomer *int, Plate *string, Imei *string, fromDateStr string, toDateStr string, page int, pageSize int) (string, error) {
	data, err := GetAllHistoricalTours(db, FkCompany, FkCustomer, Plate, Imei, fromDateStr, toDateStr, page, pageSize)
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
		record, ok := intf.(views.Tours)
		if !ok {
			return "", fmt.Errorf("no se pudo convertir el registro %d a Historical", i)
		}
		row := i + 2 // Comenzar en la segunda fila
		file.SetCellValue("Sheet1", fmt.Sprintf("A%d", row), record.ID)
		file.SetCellValue("Sheet1", fmt.Sprintf("B%d", row), record.Plate)
		file.SetCellValue("Sheet1", fmt.Sprintf("C%d", row), record.Imei)
		file.SetCellValue("Sheet1", fmt.Sprintf("D%d", row), record.TimeStampEvent)
		file.SetCellValue("Sheet1", fmt.Sprintf("E%d", row), record.Location)
		file.SetCellValue("Sheet1", fmt.Sprintf("F%d", row), record.Latitude)
		file.SetCellValue("Sheet1", fmt.Sprintf("G%d", row), record.Longitude)
	}

	filename := fmt.Sprintf("Recorrido_%s_%s_%s.xlsx", *Plate, fromDateStr, toDateStr)
	if err := file.SaveAs(filename); err != nil {
		return "", err
	}

	return filename, nil
}
