package controllers

import (
	"SimonBK_Historical_Vehicles/domain/services"
	"SimonBK_Historical_Vehicles/infra/db"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetAllAvlRecordsHandler maneja la solicitud GET para obtener todos los registros Avl
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

	records, err := services.GetAllAvlRecords(db.DBConn, fkCompany, fkCustomer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener registros Avl"})
		return
	}

	c.JSON(http.StatusOK, records)
}

// GetAvlRecordByIDHandler maneja la solicitud GET para obtener un registro Avl por ID
func GetAvlRecordByIDHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	record, err := services.GetAvlRecordByID(db.DBConn, uint(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, record)
}
