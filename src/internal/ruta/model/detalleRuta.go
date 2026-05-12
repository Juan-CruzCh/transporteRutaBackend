package model

import (
	"time"
	"transporteRuta/src/app/enum"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type DetalleRuta struct {
	ID         bson.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Fecha      time.Time     `bson:"fecha" json:"fecha"`
	Flag       enum.FlagE    `bson:"flag" json:"flag"`
	Coordenada bson.ObjectID `bson:"coordenada" json:"coordenada"`
	Ruta       bson.ObjectID `bson:"ruta" json:"ruta"`
	EsParada   bool          `bson:"es_parada" json:"es_parada"`
}
