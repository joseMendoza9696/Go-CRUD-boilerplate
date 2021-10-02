package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/josemendoza/restapi/app"
)

func main() {
	// cargamos el archivo de las variables de entorno
	er := godotenv.Load()
	if er != nil {
		log.Fatal("Error loading .env file")
	}

	app.Initilize()
}
