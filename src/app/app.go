package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"transporteRuta/src/app/config"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Repositories struct {
}

func NewRepositories(db *mongo.Database) *Repositories {
	return &Repositories{}
}

type App struct {
	ServerMux    *http.ServeMux
	Repositories *Repositories
	Validate     *validator.Validate
	Cliente      *mongo.Client
}

func StartApp() *App {
	db, cliente, err := config.ConnectMongo("", "")
	if err != nil {
		log.Fatal(err)
	}
	validate := validator.New()
	serverMux := http.NewServeMux()

	app := &App{
		ServerMux:    serverMux,
		Repositories: NewRepositories(db),
		Validate:     validate,
		Cliente:      cliente,
	}
	return app

}
func (app *App) Run() {
	var port string = "3000"
	log.Printf("Servidor corriendo en el http://localhost:%s", port)
	var handler http.Handler = app.ServerMux
	err := http.ListenAndServe(":"+port, handler)
	if err != nil {
		j, _ := json.MarshalIndent(err, " ", " ")
		fmt.Println(err, string(j))
		log.Fatal(err, string(j))
	}
}
