package routers

import (
	"SimonBK_Historical_Vehicles/api/controllers"
	"SimonBK_Historical_Vehicles/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRouter configura las rutas de la aplicación
func SetupRouter(r *gin.Engine) {
	// Validacion de acces Token
	r.Use(middleware.ValidateTokenMiddleware())

	// Grupo de rutas para vehiculos
	historicalGroup := r.Group("/avlrecords")
	{
		historicalGroup.GET("/", controllers.GetAllAvlRecordsHandler)
		historicalGroup.GET("/:id", controllers.GetAvlRecordByIDHandler)
		historicalGroup.GET("/point/", controllers.GetAllAvlRecordsPointHandler)

	}
}
