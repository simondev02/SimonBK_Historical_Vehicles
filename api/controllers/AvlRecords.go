package controllers

import (
	"SimonBK_Historical_Vehicles/domain/services"
	"SimonBK_Historical_Vehicles/infra/db"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// @Summary Obtiene todos los registros Avl
// @Description Recupera todos los registros Avl con opciones de paginación y filtrado por FkCompany y FkCustomer si están presentes en el contexto.
// @Tags AvlRecords
// @Accept  json
// @Produce  json
// @Param page query int false "Número de página para la paginación" default(1)
// @Param pageSize query int false "Tamaño de página para la paginación" default(10)
// @Param Plate query string false "Placa del vehículo"
// @Param Imei query string false "Imei del dispositivo"
// @Param fromDate query string false "Fecha de inicio para filtrar los registros Avl"
// @Param toDate query string false "Fecha de fin para filtrar los registros Avl"
// @Success 200 {array} swagger.AvlRecord "Lista de registros Avl"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Security ApiKeyAuth
// @Router /avlrecords/ [get]
func GetAllAvlRecordsHandler(c *gin.Context) {

	var fkCompany *int
	var fkCustomer *int

	// Intentar obtener FkCompany y FkCustomer del contexto de Gin
	fkCompaniaValue, exists := c.Get("FkCompany")
	if exists {
		val, ok := fkCompaniaValue.(int)
		if ok {
			fkCompany = &val
		}
	}

	fkClienteValue, exists := c.Get("FkCustomer")
	if exists {
		val, ok := fkClienteValue.(int)
		if ok {
			fkCustomer = &val
		}
	}

	// Obtener parámetros de paginación de la solicitud
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	Plate := c.DefaultQuery("Plate", "")
	Imei := c.DefaultQuery("Imei", "")
	fromDateStr := c.DefaultQuery("fromDate", "")
	toDateStr := c.DefaultQuery("toDate", "")

	// Convertir las fechas de string a time.Time
	fromDate, err := time.Parse("2006-01-02", fromDateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Fecha de inicio inválida"})
		return
	}
	toDate, err := time.Parse("2006-01-02", toDateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Fecha final inválida"})
		return
	}

	// Verificar que la fecha final no sea menor que la fecha de inicio
	if toDate.Before(fromDate) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "La fecha final no puede ser menor que la fecha de inicio"})
		return
	}

	records, err := services.GetAllAvlRecords(db.DBConn, fkCompany, fkCustomer, page, pageSize, &Plate, &Imei, fromDateStr, toDateStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener registros Avl"})
		return
	}

	c.JSON(http.StatusOK, records)
}
