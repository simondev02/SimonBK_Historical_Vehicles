package controllers

import (
	services "SimonBK_Historical_Vehicles/domain/services/historical"
	"SimonBK_Historical_Vehicles/domain/services/utilities"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Obtiene todos los puntos de registros Avl
// @Description Recupera todos los puntos de registros Avl con opciones de paginación y filtrado por FkCompany y FkCustomer si están presentes en el contexto.
// @Tags AvlRecords
// @Accept  json
// @Produce  json
// @Param Plate query string false "Placa del vehículo"
// @Param Imei query string false "Imei del dispositivo"
// @Param fromDate query string false "Fecha de inicio para filtrar los registros Avl"
// @Param toDate query string false "Fecha de fin para filtrar los registros Avl"
// @Success 200 {array} swagger.AvlRecordPoint "Lista de puntos de registros Avl"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Security ApiKeyAuth
// @Router /avlrecords/excel/ [get]
func GetAllExcelHistoricalHandler(c *gin.Context) {

	params, err := utilities.GetParams(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	filename, err := services.DownloadHistoricalExcel(params)
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
