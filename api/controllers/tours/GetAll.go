package controllers

import (
	services "SimonBK_Historical_Vehicles/domain/services/tours"
	"SimonBK_Historical_Vehicles/domain/services/utilities"
	"fmt"
	"net/http"
	"strconv"

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

	params, err := utilities.GetParams(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if temp := c.Query("page"); temp != "" {
		val, err := strconv.Atoi(temp)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("error parsing page: %v", err)})
			return
		}
		params.Page = val
	} else {

	}

	if temp := c.Query("pageSize"); temp != "" {
		val, err := strconv.Atoi(temp)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("error parsing pageSize: %v", err)})
			return
		}
		params.PageSize = val
	} else {
		params.PageSize = 10
	}

	records, err := services.GetAllHistoricalTours(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error al obtener registros Avl: %v", err)})
		return
	}

	c.JSON(http.StatusOK, records)
}
