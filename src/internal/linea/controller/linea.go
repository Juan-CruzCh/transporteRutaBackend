package controller

import (
	"net/http"
	"transporteRuta/src/app/utils"
	"transporteRuta/src/internal/linea/service"
)

type Linea struct {
	lineaService *service.Linea
}

func NewLineaController(lineaService *service.Linea) *Linea {

	return &Linea{
		lineaService: lineaService,
	}
}
func (c *Linea) CrearLinea(w http.ResponseWriter, r *http.Request) {

	utils.ResponseJSON(w, http.StatusOK, map[string]string{"me": "ok"})
}

func (c *Linea) ListarLinea(w http.ResponseWriter, r *http.Request) {
	utils.ResponseJSON(w, http.StatusOK, map[string]string{"me": "ok"})
}

func (c *Linea) ActualizarLinea(w http.ResponseWriter, r *http.Request) {
}

func (c *Linea) EliminarLinea(w http.ResponseWriter, r *http.Request) {
}
