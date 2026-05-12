package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
	"transporteRuta/src/app/utils"
	"transporteRuta/src/internal/ruta/dto"
	"transporteRuta/src/internal/ruta/service"

	"github.com/go-playground/validator/v10"
)

type Ruta struct {
	rutaService *service.Ruta
	Validate    *validator.Validate
}

func NewRutaController(rutaService *service.Ruta, validate *validator.Validate) *Ruta {
	return &Ruta{
		rutaService: rutaService,
		Validate:    validate,
	}
}
func (c *Ruta) CrearRuta(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	var body dto.RutaDto
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
	resultado, err :=
		c.rutaService.CrearRuta(&body, ctx)
	if err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]string{"mensaje": err.Error()})
		return
	}
	utils.ResponseJSON(w, http.StatusCreated, resultado)
}

func (c *Ruta) ListarRuta(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	data, err := c.rutaService.ListarRuta(ctx)
	if err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]string{"mensaje": err.Error()})
		return
	}
	utils.ResponseJSON(w, http.StatusOK, data)
}

func (c *Ruta) ActualizarRuta(w http.ResponseWriter, r *http.Request) {
}

func (c *Ruta) EliminarRuta(w http.ResponseWriter, r *http.Request) {
}
