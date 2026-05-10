package main

import (
	"log"
	"transporteRuta/src/app"
	"transporteRuta/src/app/config"

	"github.com/joho/godotenv"
)

func main() {
	config.ConfiguracionLog()
	defer config.CerrarLog()
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	config.VariablesEntorno()
	app := app.StartApp()
	app.Run()
}
