package router

import (
	"net/http"
	"transporteRuta/src/internal/coordenada/controller"
)

func NewCoordenadaRouter(mux *http.ServeMux, controller *controller.Coordenada) {
	mux.HandleFunc("POST /api/coordenada", controller.CrearCoordenada)
	mux.HandleFunc("GET /api/coordenada", controller.ListarCoordenada)
}
