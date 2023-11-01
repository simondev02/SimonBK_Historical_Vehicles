package controllers

import (
	"SimonBK_Historical_Vehicles/domain/services"
	"SimonBK_Historical_Vehicles/infra/db"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Obtiene todos los puntos de registros Avl
// @Description Recupera todos los puntos de registros Avl con opciones de paginación y filtrado por FkCompany y FkCustomer si están presentes en el contexto.
// @Tags AvlRecordsPoints
// @Accept  json
// @Produce  json
// @Param Plate query string false "Placa del vehículo"
// @Param Imei query string false "Imei del dispositivo"
// @Param fromDate query string false "Fecha de inicio para filtrar los registros Avl"
// @Param toDate query string false "Fecha de fin para filtrar los registros Avl"
// @Success 200 {array} swagger.AvlRecordPoint "Lista de puntos de registros Avl"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Security ApiKeyAuth
// @Router /avlrecords/point/ [get]
func GetAllAvlRecordsPointHandler(c *gin.Context) {
	var fkCompany, fkCustomer *int

	fkCompany = tryGetContextValueAsInt(c, "FkCompany")
	fkCustomer = tryGetContextValueAsInt(c, "FkCustomer")

	Plate := c.DefaultQuery("Plate", "")
	Imei := c.DefaultQuery("Imei", "")
	fromDateStr := c.DefaultQuery("fromDate", "")
	toDateStr := c.DefaultQuery("toDate", "")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	records, err := services.GetAllAvlRecordsPoint(db.DBConn, fkCompany, fkCustomer, &Plate, &Imei, fromDateStr, toDateStr, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error al obtener registros Avl: %v", err)})
		return
	}

	c.JSON(http.StatusOK, records)
}

func tryGetContextValueAsInt(c *gin.Context, key string) *int {
	value, exists := c.Get(key)
	if !exists {
		return nil
	}
	val, ok := value.(int)
	if !ok {
		return nil
	}
	return &val
}
