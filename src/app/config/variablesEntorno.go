package config

import "os"

var Port string = ""

var UrlMongo string = ""

func VariablesEntorno() {
	UrlMongo = os.Getenv("URL_MONGO")
	Port = os.Getenv("PORT")

}
