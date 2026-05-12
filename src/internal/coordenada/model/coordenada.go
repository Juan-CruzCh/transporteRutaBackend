package model

import (
	"time"
	"transporteRuta/src/app/enum"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Coordenada struct {
	ID       bson.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Fecha    time.Time     `bson:"fecha" json:"fecha"`
	Flag     enum.FlagE    `bson:"flag" json:"flag"`
	Latitud  float64       `bson:"latitud" json:"latitud"`
	Longitud float64       `bson:"longitud" json:"longitud"`
	Ruta     bson.ObjectID `bson:"ruta" json:"ruta"`
}
