package model

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type DetalleRuta struct {
	ID    bson.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Fecha time.Time     `bson:"fecha" json:"fecha"`
	Linea bson.ObjectID `bson:"linea" json:"linea"`
	Ruta  bson.ObjectID `bson:"ruta" json:"ruta"`
}
