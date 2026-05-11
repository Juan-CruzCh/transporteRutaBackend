package dto

type RutaDto struct {
	Latitud  float64 `json:"latitud" bson:"latitud" validate:"required"`
	Longitud float64 `json:"longitud" bson:"longitud" validate:"required"`
}
