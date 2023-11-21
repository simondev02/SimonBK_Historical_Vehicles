package controllers

import (
	services "SimonBK_Historical_Vehicles/domain/services/historical"
	"SimonBK_Historical_Vehicles/infra/db"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @Security ApiKeyAuth
// @Summary Obtiene un historico específico
// @Description Obtiene un historico por su ID específico
// @Tags AvlRecords
// @Accept json
// @Produce json
// @Param id path int true "ID del Historico"
// @Success 200 {object} swagger.AvlRecord "Detalles del vehículo"
// @Failure 400 {object} map[string]string "Error: ID inválido"
// @Failure 404 {object} map[string]string "Error: Vehículo no encontrado"
// @Router /avlrecords/{id} [get]
func GetAvlRecordByIDHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido. Debe ser un número entero."})
		return
	}
	record, err := services.GetHistoricalByID(db.DBConn, uint(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Registro con ID %d no encontrado.", id)})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error al obtener el registro con ID %d: %v", id, err)})
		return
	}
	c.JSON(http.StatusOK, record)
}
