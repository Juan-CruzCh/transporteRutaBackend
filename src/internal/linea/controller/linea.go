package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
	"transporteRuta/src/app/utils"
	"transporteRuta/src/internal/linea/dto"
	"transporteRuta/src/internal/linea/service"

	"github.com/go-playground/validator/v10"
)

type Linea struct {
	lineaService *service.Linea
	Validate     *validator.Validate
}

func NewLineaController(lineaService *service.Linea, Validate *validator.Validate) *Linea {

	return &Linea{
		lineaService: lineaService,
		Validate:     Validate,
	}
}
func (c *Linea) CrearLinea(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	var body dto.LineaDto
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
	err = c.lineaService.CrearLinea(&body, ctx)
	if err != nil {
		utils.ResponseJSON(w, http.StatusCreated, map[string]string{"mensaje": err.Error()})
		return
	}
	utils.ResponseJSON(w, http.StatusCreated, map[string]string{"mensaje": "registrado"})

}

func (c *Linea) CrearDetalleLinea(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	var body dto.DetalleLineaDto
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
	err = c.lineaService.CrearDetalleLinea(&body, ctx)
	if err != nil {
		utils.ResponseJSON(w, http.StatusCreated, map[string]string{"mensaje": err.Error()})
		return
	}
	utils.ResponseJSON(w, http.StatusCreated, map[string]string{"mensaje": "registrado"})

}
func (c *Linea) ListarLinea(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	data, err := c.lineaService.ListarLinea(ctx)

	if err != nil {
		utils.ResponseJSON(w, http.StatusCreated, map[string]string{"mensaje": err.Error()})
		return
	}
	utils.ResponseJSON(w, http.StatusOK, data)
}

func (c *Linea) ActualizarLinea(w http.ResponseWriter, r *http.Request) {
}

func (c *Linea) EliminarLinea(w http.ResponseWriter, r *http.Request) {
}
