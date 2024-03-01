package routers

import (
	historical "SimonBK_Historical_Vehicles/api/controllers/historical"
	tours "SimonBK_Historical_Vehicles/api/controllers/tours"
	"SimonBK_Historical_Vehicles/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRouter configura las rutas de la aplicación
func SetupRouter(r *gin.Engine) {
	// Validacion de acces Token
	r.Use(middleware.ValidateTokenMiddleware())

	// Grupo de rutas para vehiculos
	avlRecordsGroup := r.Group("/avlrecords")
	{
		avlRecordsGroup.GET("/", historical.GetAllAvlRecordsHandler)
		avlRecordsGroup.GET("/:id", historical.GetAvlRecordByIDHandler)
		avlRecordsGroup.GET("/excel", historical.GetAllExcelHistoricalHandler)

	}

	// Grupo de rutas para vehiculos
	toursGroup := r.Group("/tours")
	{
		toursGroup.GET("/", tours.GetAllHistoricalToursHandler)
		toursGroup.GET("/:id", tours.GetHistoricalToursByIdHandler)
		toursGroup.GET("/excel/", tours.GetAllExcelToursHandler)
	}
}
