package controllers

import (
	inputs "SimonBK_Historical_Vehicles/api/views/inputs"
	services "SimonBK_Historical_Vehicles/domain/services/tours"
	"SimonBK_Historical_Vehicles/infra/db"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// @Summary Obtiene todos los puntos de registros Avl
// @Description Recupera todos los puntos de registros Avl con opciones de paginación y filtrado por FkCompany y FkCustomer si están presentes en el contexto.
// @Tags Tours
// @Accept  json
// @Produce  json
// @Param page query int false "Número de página para la paginación" default(1)
// @Param pageSize query int false "Tamaño de página para la paginación" default(10)
// @Param Plate query string false "Placa del vehículo"
// @Param Imei query string false "Imei del dispositivo"
// @Param fromDate query string false "Fecha de inicio para filtrar los registros Avl"
// @Param toDate query string false "Fecha de fin para filtrar los registros Avl"
// @Success 200 {array} swagger.AvlRecordPoint "Lista de puntos de registros Avl"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Security ApiKeyAuth
// @Router /tours/excel/ [get]
func GetAllExcelToursHandler(c *gin.Context) {

	var fkCompany *uint
	if temp, exists := c.Get("FkCompany"); exists {
		val := uint(temp.(int))
		fkCompany = &val
	}

	var fkCustomer *uint
	if temp, exists := c.Get("FkCustomer"); exists {
		val := uint(temp.(int))
		fkCustomer = &val
	}

	var plate *string
	if temp := c.Query("Plate"); temp != "" {
		plate = &temp
	}
	fmt.Println(plate)

	var imei *string
	if temp := c.Query("Imei"); temp != "" {
		imei = &temp
	}

	var fromDate *time.Time
	if temp := c.Query("fromDate"); temp != "" {
		val, _ := time.Parse(time.RFC3339, temp)
		fromDate = &val
	}

	var toDate *time.Time
	if temp := c.Query("toDate"); temp != "" {
		val, _ := time.Parse(time.RFC3339, temp)
		toDate = &val
	}

	var page *uint
	if temp := c.Query("page"); temp != "" {
		val, _ := strconv.Atoi(temp)
		valUint := uint(val)
		page = &valUint
	}

	var pageSize *uint
	if temp := c.Query("pageSize"); temp != "" {
		val, _ := strconv.Atoi(temp)
		valUint := uint(val)
		pageSize = &valUint
	}

	tourIn := inputs.ToursInputs{
		Db:         db.DBConn,
		FkCompany:  fkCompany,
		FkCustomer: fkCustomer,
		Plate:      plate,
		Imei:       imei,
		FromDate:   fromDate,
		ToDate:     toDate,
		Page:       page,
		PageSize:   pageSize,
	}
	filename, err := services.DownloadHistoricalToursExcel(tourIn)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error al obtener registros Avl: %v", err)})
		return
	}

	// Leer el contenido del archivo
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error al leer el archivo: %v", err)})
		return
	}

	// Establecer las cabeceras para la descarga
	c.Writer.Header().Set("Content-type", "application/octet-stream")
	c.Writer.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))

	// Escribir el contenido en la respuesta
	c.Writer.Write(content)
}
