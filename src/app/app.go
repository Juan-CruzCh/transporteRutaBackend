package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"transporteRuta/src/app/config"
	"transporteRuta/src/app/middleware"
	LineaController "transporteRuta/src/internal/linea/controller"
	LineaRepository "transporteRuta/src/internal/linea/repository"
	LineaRouter "transporteRuta/src/internal/linea/router"
	LineaService "transporteRuta/src/internal/linea/service"

	RutaRepository "transporteRuta/src/internal/ruta/repository"

	RutaController "transporteRuta/src/internal/ruta/controller"
	RutaRouter "transporteRuta/src/internal/ruta/router"
	RutaService "transporteRuta/src/internal/ruta/service"

	UbicacionController "transporteRuta/src/internal/ubicacion/controller"
	UbicacionRepository "transporteRuta/src/internal/ubicacion/repository"
	UbicacionRouter "transporteRuta/src/internal/ubicacion/router"
	UbicacionService "transporteRuta/src/internal/ubicacion/service"

	UsuarioController "transporteRuta/src/internal/usuario/controller"
	UsuarioRepository "transporteRuta/src/internal/usuario/repository"
	UsuarioRouter "transporteRuta/src/internal/usuario/router"
	UsuarioService "transporteRuta/src/internal/usuario/service"

	CoordenadaController "transporteRuta/src/internal/coordenada/controller"
	CoordenadaRepository "transporteRuta/src/internal/coordenada/repository"
	CoordenadaRouter "transporteRuta/src/internal/coordenada/router"
	CoordenadaService "transporteRuta/src/internal/coordenada/service"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Repositories struct {
	LineaRepository        LineaRepository.Linea
	DetalleLineaRepository LineaRepository.DetalleLinea
	RutaRepository         RutaRepository.Ruta
	DetalleRutaRepository  RutaRepository.DetalleRuta
	UbicacionRepository    UbicacionRepository.Ubicacion
	UsuarioRepository      UsuarioRepository.Usuario
	CoordenadaRepository   CoordenadaRepository.Coordenada
}

func NewRepositories(db *mongo.Database) *Repositories {
	return &Repositories{
		LineaRepository:        LineaRepository.NewLineaRepository(db),
		RutaRepository:         RutaRepository.NewRutaRepository(db),
		UbicacionRepository:    UbicacionRepository.NewUbicacionRepository(db),
		UsuarioRepository:      UsuarioRepository.NewUsuarioRepository(db),
		CoordenadaRepository:   CoordenadaRepository.NewCoordenadaRepository(db),
		DetalleLineaRepository: LineaRepository.NewDetalleLineaRepository(db),
		DetalleRutaRepository:  RutaRepository.NewDetalleRutaRepository(db),
	}
}

type App struct {
	ServerMux    *http.ServeMux
	Repositories *Repositories
	Validate     *validator.Validate
	Cliente      *mongo.Client
}

func StartApp() *App {
	db, cliente, err := config.ConnectMongo(config.UrlMongo, "transporte")
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
	initLinea(app)
	initRuta(app)
	initUbicacion(app)
	initUsuario(app)
	initCoordenada(app)

	return app

}
func (app *App) Run() {
	var port string = config.Port
	log.Printf("Servidor corriendo en el http://localhost:%s", port)
	var handler http.Handler = app.ServerMux
	handler = middleware.EnableCORS(handler)
	err := http.ListenAndServe(":"+port, handler)
	if err != nil {
		j, _ := json.MarshalIndent(err, " ", " ")
		fmt.Println(err, string(j))
		log.Fatal(err, string(j))
	}
}

func initLinea(app *App) {
	LineaService := LineaService.NewLineaService(app.Repositories.LineaRepository, app.Repositories.DetalleLineaRepository, app.Cliente)
	LineaController := LineaController.NewLineaController(LineaService, app.Validate)
	LineaRouter.NewLineaRouter(app.ServerMux, LineaController)
}

func initRuta(app *App) {
	RutaService := RutaService.NewRutaService(app.Repositories.RutaRepository, app.Repositories.DetalleRutaRepository, app.Cliente)
	RutaController := RutaController.NewRutaController(RutaService, app.Validate)
	RutaRouter.NewRutaRouter(app.ServerMux, RutaController)

}
func initUbicacion(app *App) {
	UbicacionService := UbicacionService.NewUbicacionService(
		app.Repositories.UbicacionRepository,
		app.Cliente,
	)

	UbicacionController := UbicacionController.NewUbicacionController(
		UbicacionService,
	)

	UbicacionRouter.NewUbicacionRouter(
		app.ServerMux,
		UbicacionController,
	)

}

func initUsuario(app *App) {
	UsuarioService := UsuarioService.NewUsuarioService(
		app.Repositories.UsuarioRepository,
		app.Cliente,
	)

	UsuarioController := UsuarioController.NewUsuarioController(
		UsuarioService,
	)

	UsuarioRouter.NewUsuarioRouter(
		app.ServerMux,
		UsuarioController,
	)

}
func initCoordenada(app *App) {
	CoordenadaService := CoordenadaService.NewCoordenadaService(app.Repositories.CoordenadaRepository, app.Cliente)
	CoordenadaController := CoordenadaController.NewCoordenadaController(CoordenadaService, app.Validate)
	CoordenadaRouter.NewCoordenadaRouter(app.ServerMux, CoordenadaController)

}
