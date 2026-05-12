package dto

type CoordenadaDto struct {
	Latitud  float64 `json:"latitud" validate:"required"`
	Longitud float64 `json:"longitud" validate:"required"`
}
