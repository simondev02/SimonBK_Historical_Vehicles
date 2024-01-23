package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // Importa el driver de PostgreSQL
)

var SQLDB *sql.DB

func ConnectSQLDB() error {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error al leer variables de entorno", err)
		return err
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=require", dbHost, dbPort, dbUser, dbPass, dbName)

	// Crear conexión con la base de datos
	SQLDB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Error al conectar con la base de datos:", err)
		return err
	}

	// Verificar la conexión
	if err = SQLDB.Ping(); err != nil {
		log.Fatal("Error al verificar la conexión con la base de datos:", err)
		return err
	}

	log.Println("SQL DB CONNECTED")
	return nil
}

func CloseSQLDB() error {
	if err := SQLDB.Close(); err != nil {
		log.Println("Error al cerrar la conexión con la base de datos:", err)
		return err
	}

	log.Println("SQL DB DISCONNECTED")
	return nil
}
