package services

import (
	"fmt"
	"time"

	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
)

func DownloadHistoricalExcel(db *gorm.DB, FkCompany *int, FkCustomer *int, page int, pageSize int, Plate *string, Imei *string, fromDateStr string, toDateStr string) (string, error) {
	data, err := GetAllHistoricalExcel(db, FkCompany, FkCustomer, Plate, Imei, fromDateStr, toDateStr)
	if err != nil {
		return "", err
	}

	file := excelize.NewFile()

	// Añadir encabezados
	headers := []string{"ID", "Plate", "Imei", "Ip", "TimeStampEvent", "Id_company", "Company", "Id_customer", "Customer", "Location", "Latitude", "Longitude", "Altitude", "Angle", "Satellites", "Speed", "Hdop", "Pdop", "Event"}
	for i, header := range headers {
		file.SetCellValue("Sheet1", fmt.Sprintf("%c1", 'A'+i), header)
	}

	// Añadir datos
	for i, record := range data {

		row := i + 2 // Comenzar en la segunda fila
		file.SetCellValue("Sheet1", fmt.Sprintf("A%d", row), record.ID)
		file.SetCellValue("Sheet1", fmt.Sprintf("B%d", row), record.Plate)
		file.SetCellValue("Sheet1", fmt.Sprintf("C%d", row), record.Imei)
		file.SetCellValue("Sheet1", fmt.Sprintf("D%d", row), record.Ip)
		if record.TimeStampEvent != nil {
			file.SetCellValue("Sheet1", fmt.Sprintf("E%d", row), record.TimeStampEvent.Format(time.RFC3339))
		}
		file.SetCellValue("Sheet1", fmt.Sprintf("F%d", row), record.Id_company)
		file.SetCellValue("Sheet1", fmt.Sprintf("G%d", row), record.Company)
		file.SetCellValue("Sheet1", fmt.Sprintf("H%d", row), record.Id_customer)
		file.SetCellValue("Sheet1", fmt.Sprintf("I%d", row), record.Customer)
		file.SetCellValue("Sheet1", fmt.Sprintf("J%d", row), record.Location)
		file.SetCellValue("Sheet1", fmt.Sprintf("K%d", row), record.Latitude)
		file.SetCellValue("Sheet1", fmt.Sprintf("L%d", row), record.Longitude)
		file.SetCellValue("Sheet1", fmt.Sprintf("M%d", row), record.Altitude)
		file.SetCellValue("Sheet1", fmt.Sprintf("N%d", row), record.Angle)
		file.SetCellValue("Sheet1", fmt.Sprintf("O%d", row), record.Satellites)
		file.SetCellValue("Sheet1", fmt.Sprintf("P%d", row), record.Speed)
		file.SetCellValue("Sheet1", fmt.Sprintf("Q%d", row), record.Hdop)
		file.SetCellValue("Sheet1", fmt.Sprintf("R%d", row), record.Pdop)
		file.SetCellValue("Sheet1", fmt.Sprintf("S%d", row), record.Event)
		file.SetCellValue("Sheet1", fmt.Sprintf("S%d", row), record.Properties)

		// Para las propiedades adicionales, puedes decidir cómo manejarlas según tus necesidades
	}

	filename := "Historical.xlsx"
	if err := file.SaveAs(filename); err != nil {
		return "", err
	}

	return filename, nil
}
