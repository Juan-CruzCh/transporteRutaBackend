package controller

import (
	"net/http"
	"transporteRuta/src/internal/ruta/service"
)

type Ruta struct {
	rutaService *service.Ruta
}

func NewRutaController(rutaService *service.Ruta) *Ruta {
	return &Ruta{
		rutaService: rutaService,
	}
}
func (c *Ruta) CrearRuta(w http.ResponseWriter, r *http.Request) {
}

func (c *Ruta) ListarRuta(w http.ResponseWriter, r *http.Request) {
}

func (c *Ruta) ActualizarRuta(w http.ResponseWriter, r *http.Request) {
}

func (c *Ruta) EliminarRuta(w http.ResponseWriter, r *http.Request) {
}
