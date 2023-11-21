package controllers

import (
	services "SimonBK_Historical_Vehicles/domain/services/tours"
	"SimonBK_Historical_Vehicles/infra/db"
	"fmt"
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
// @Router /tours/ [get]
func GetAllHistoricalToursHandler(c *gin.Context) {
	var fkCompany *int
	var fkCustomer *int

	fkCompanyValue, exists := c.Get("FkCompany")
	if exists {
		val, ok := fkCompanyValue.(int)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": "FkCompany debe ser un número entero."})
			return
		}
		fkCompany = &val
	}

	fkCustomerValue, exists := c.Get("FkCustomer")
	if exists {
		val, ok := fkCustomerValue.(int)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": "FkCustomer debe ser un número entero."})
			return
		}
		fkCustomer = &val
	}

	Plate := c.DefaultQuery("Plate", "")
	Imei := c.DefaultQuery("Imei", "")
	fromDateStr := c.DefaultQuery("fromDate", "")
	toDateStr := c.DefaultQuery("toDate", "")
	pageStr := c.DefaultQuery("page", "")
	pageSizeStr := c.DefaultQuery("pageSize", "")

	var page, pageSize int
	var err error

	if pageStr != "" {
		page, err = strconv.Atoi(pageStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "La página debe ser un número entero."})
			return
		}
	}
	if pageSizeStr != "" {
		pageSize, err = strconv.Atoi(pageSizeStr)
		fmt.Println(pageSize)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "El tamaño de la página debe ser un número entero."})
			return
		}
	}

	// Manejar la lógica de fechas
	var fromDate, toDate time.Time
	if fromDateStr != "" && toDateStr != "" {
		fromDate, err = time.Parse(time.RFC3339, fromDateStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "La fecha de inicio no es válida."})
			return
		}

		toDate, err = time.Parse(time.RFC3339, toDateStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "La fecha final no es válida."})
			return
		}

		if fromDate.After(toDate) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "La fecha de inicio no puede ser posterior a la fecha final."})
			return
		}
	}

	records, err := services.GetAllHistoricalTours(db.DBConn, fkCompany, fkCustomer, &Plate, &Imei, fromDateStr, toDateStr, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error al obtener registros Avl: %v", err)})
		return
	}

	c.JSON(http.StatusOK, records)
}
