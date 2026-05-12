package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
	"transporteRuta/src/app/utils"
	"transporteRuta/src/internal/coordenada/dto"
	"transporteRuta/src/internal/coordenada/service"

	"github.com/go-playground/validator/v10"
)

type Coordenada struct {
	coordenadaService *service.Coordenada
	Validate          *validator.Validate
}

func NewCoordenadaController(coordenadaService *service.Coordenada, Validate *validator.Validate) *Coordenada {
	return &Coordenada{
		coordenadaService: coordenadaService,
		Validate:          Validate,
	}
}
func (c *Coordenada) CrearCoordenada(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	var body dto.CoordenadaDto
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]string{"mensaje": err.Error()})
		return
	}
	err = c.Validate.Struct(&body)
	if err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]string{"mensaje": err.Error()})
		return
	}
	err = c.coordenadaService.CrearCoordenada(&body, ctx)
	if err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]string{"mensaje": err.Error()})
		return
	}
	utils.ResponseJSON(w, http.StatusCreated, map[string]string{"mensaje": "registrado"})
}

func (c *Coordenada) ListarCoordenada(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	data, err := c.coordenadaService.ListarCoordenada(ctx)
	if err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]string{"mensaje": err.Error()})
		return
	}
	utils.ResponseJSON(w, http.StatusOK, data)
}

func (c *Coordenada) ActualizarCoordenada(w http.ResponseWriter, r *http.Request) {
}

func (c *Coordenada) EliminarCoordenada(w http.ResponseWriter, r *http.Request) {
}
