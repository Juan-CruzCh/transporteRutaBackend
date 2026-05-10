package controller

import (
	"net/http"
	"transporteRuta/src/internal/ubicacion/service"
)

type Ubicacion struct {
	ubicacionService *service.Ubicacion
}

func NewUbicacionController(ubicacionService *service.Ubicacion) *Ubicacion {
	return &Ubicacion{
		ubicacionService: ubicacionService,
	}
}
func (c *Ubicacion) CrearUbicacion(w http.ResponseWriter, r *http.Request) {
}

func (c *Ubicacion) ListarUbicacion(w http.ResponseWriter, r *http.Request) {
}

func (c *Ubicacion) ActualizarUbicacion(w http.ResponseWriter, r *http.Request) {
}

func (c *Ubicacion) EliminarUbicacion(w http.ResponseWriter, r *http.Request) {
}
