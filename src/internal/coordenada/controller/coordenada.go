package controller

import (
	"net/http"
	"transporteRuta/src/internal/coordenada/service"
)

type Coordenada struct {
	coordenadaService *service.Coordenada
}

func NewCoordenadaController(coordenadaService *service.Coordenada) *Coordenada {
	return &Coordenada{
		coordenadaService: coordenadaService,
	}
}
func (c *Coordenada) CrearCoordenada(w http.ResponseWriter, r *http.Request) {

}

func (c *Coordenada) ListarCoordenada(w http.ResponseWriter, r *http.Request) {
}

func (c *Coordenada) ActualizarCoordenada(w http.ResponseWriter, r *http.Request) {
}

func (c *Coordenada) EliminarCoordenada(w http.ResponseWriter, r *http.Request) {
}
