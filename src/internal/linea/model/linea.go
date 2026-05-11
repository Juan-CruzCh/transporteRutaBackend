package model

import (
	"time"
	"transporteRuta/src/app/enum"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Linea struct {
	ID          bson.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Fecha       time.Time     `bson:"fecha" json:"fecha"`
	Flag        enum.FlagE    `bson:"flag" json:"flag"`
	Nombre      string        `bson:"nombre" json:"nombre"`
	Descripcion string        `bson:"descripcion" json:"descripcion"`
	Color       string        `bson:"color" json:"color"`
}
