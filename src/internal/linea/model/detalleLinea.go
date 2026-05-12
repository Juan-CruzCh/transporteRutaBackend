package model

import (
	"time"
	"transporteRuta/src/app/enum"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type DetalleLinea struct {
	ID    bson.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Fecha time.Time     `bson:"fecha" json:"fecha"`
	Flag  enum.FlagE    `bson:"flag" json:"flag"`
	Ruta  bson.ObjectID `bson:"ruta" json:"ruta"`
	Linea bson.ObjectID `bson:"linea" json:"linea"`
}
