package router

import (
	"net/http"
	"transporteRuta/src/internal/linea/controller"
)

func NewLineaRouter(mux *http.ServeMux, controller *controller.Linea) {
	mux.HandleFunc("GET /api/linea/listar", controller.ListarLinea)
}
