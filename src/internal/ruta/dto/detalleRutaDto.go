package dto

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type DetalleRutaDto struct {
	Coordenada bson.ObjectID `json:"coordenada" validate:"required"`
	Ruta       bson.ObjectID `json:"ruta" validate:"required"`
	EsParada   *bool         `json:"esParada" validate:"required"`
}
