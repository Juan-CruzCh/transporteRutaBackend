package dto

import "go.mongodb.org/mongo-driver/v2/bson"

type DetalleLineaDto struct {
	Ruta  bson.ObjectID `json:"ruta" validate:"required"`
	Linea bson.ObjectID `json:"linea" validate:"required"`
}
