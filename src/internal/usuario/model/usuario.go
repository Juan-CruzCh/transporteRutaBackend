package model

import (
	"time"
	"transporteRuta/src/app/enum"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Usuario struct {
	ID    bson.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Fecha time.Time     `bson:"fecha" json:"fecha"`
	Flag  enum.FlagE    `bson:"flag" json:"flag"`
}
