package model

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Ruta struct {
	ID       bson.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Fecha    time.Time     `bson:"fecha" json:"fecha"`
	Latitud  float64       `bson:"latitud" json:"latitud"`
	Longitud float64       `bson:"longitud" json:"longitud"`
}
