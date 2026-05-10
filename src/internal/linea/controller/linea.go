package controller

import (
	"net/http"
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
}

func (c *Linea) ListarLinea(w http.ResponseWriter, r *http.Request) {
}

func (c *Linea) ActualizarLinea(w http.ResponseWriter, r *http.Request) {
}

func (c *Linea) EliminarLinea(w http.ResponseWriter, r *http.Request) {
}
