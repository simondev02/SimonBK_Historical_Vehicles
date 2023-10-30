package controllers

import (
	"SimonBK_Historical_Vehicles/domain/services"
	"SimonBK_Historical_Vehicles/infra/db"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

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
