// @API Historicos Avl
// @Se encuentra el hostoricos de registro generados por los vehiculo y dispositivos
// @version 1
// @BasePath /Vehicle
// @SecurityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
package main

import (
	"SimonBK_Historical_Vehicles/docs"
	"SimonBK_Historical_Vehicles/infra/db"
	"log"

	"SimonBK_Historical_Vehicles/routers"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	// Establecer la conexión con la base de datos
	err := db.ConnectDB()

	if err != nil {
		fmt.Println("Error al conectar con la base de datos:", err)
		return
	}
	// Configurar CORS
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
		// Esto no puede aceptar todo debe restringirse a los encabezados que se necesitan
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Configurar Swagger
	docs.SwaggerInfo.Title = "Mi API"
	docs.SwaggerInfo.Description = "Esta es mi API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = os.Getenv("SERVICE_HOST") + ":" + os.Getenv("SERVICE_PORT")
	docs.SwaggerInfo.BasePath = "/"

	if err != nil {
		fmt.Println("Error al configurar CORS:", err)
		return
	}

	// Configurar e iniciar el enrutador
	routers.SetupRouter(r)

	// Agregar la ruta de Swagger sin el middleware de validación de token
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Configurar la señal de captura
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		// Código de limpieza: cierra la conexión a la base de datos
		db.CloseDB()
		os.Exit(0)
	}()
	certFile := os.Getenv("TLS_CERT")
	certKey := os.Getenv("TLS_KEY")
	if certFile == "" || certKey == "" {
		log.Println("Error al leer las variables de entorno.")
		db.CloseDB()
		os.Exit(1)
	}

	err = r.RunTLS(":"+os.Getenv("SERVICE_PORT"), certFile, certKey) // escucha y sirve en 0.0.0.0:60060  (por defecto)

	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
		return
	}
}
