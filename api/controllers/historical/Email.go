package controllers

import (
	services "SimonBK_Historical_Vehicles/domain/services/historical"
	"SimonBK_Historical_Vehicles/infra/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendRecordsByEmailController(c *gin.Context) {
	// Obtener la conexión a la base de datos directamente desde el paquete donde está definida
	sqlDB, err := db.DBConn.DB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener la base de datos SQL"})
		return
	}

	err = services.SendRecordsByEmail(sqlDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Email enviado"})
}
