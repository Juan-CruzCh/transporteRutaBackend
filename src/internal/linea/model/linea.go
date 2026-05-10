package model

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Linea struct {
	ID          bson.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Fecha       time.Time     `bson:"fecha" json:"fecha"`
	Nombre      string        `bson:"nombre" json:"nombre"`
	Descripcion string        `bson:"descripcion" json:"descripcion"`
	Color       string        `bson:"color" json:"color"`
}
