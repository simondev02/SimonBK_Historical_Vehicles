package controllers

import (
	services "SimonBK_Historical_Vehicles/domain/services/historical"
	"SimonBK_Historical_Vehicles/domain/services/utilities"
	"net/http"

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
// @Param fromDate query string false "Fecha de inicio para filtrar los registros Avl" 2004-09-30
// @Param toDate query string false "Fecha de fin para filtrar los registros Avl" 2004-09-30
// @Success 200 {array} swagger.AvlRecord "Lista de registros Avl"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Security ApiKeyAuth
// @Router /avlrecords/ [get]
func GetAllAvlRecordsHandler(c *gin.Context) {

	params, err := utilities.GetParams(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	records, err := services.GetAllHistorical(params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, records)
}
