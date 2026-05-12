package router

import (
	"net/http"
	"transporteRuta/src/internal/ruta/controller"
)

func NewRutaRouter(mux *http.ServeMux, controller *controller.Ruta) {
	mux.HandleFunc("POST /api/ruta", controller.CrearRuta)
	mux.HandleFunc("GET /api/ruta", controller.ListarRuta)
}
