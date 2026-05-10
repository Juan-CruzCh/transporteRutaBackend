package repository

import (
	"context"
	"transporteRuta/src/internal/ruta/model"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Ruta interface {
	CrearRuta(ruta *model.Ruta, ctx context.Context) error
	ListarRuta(id *bson.ObjectID, ctx context.Context) (interface{}, error)
	ActualizarRuta(id *bson.ObjectID, ruta *model.Ruta, ctx context.Context) error
	EliminarRuta(id *bson.ObjectID, ctx context.Context) error
}

type ruta struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewRutaRepository(db *mongo.Database) *ruta {
	collection := db.Collection("Ruta")
	return &ruta{db: db, collection: collection}
}

func (r *ruta) CrearRuta(ruta *model.Ruta, ctx context.Context) error {
	return nil
}

func (r *ruta) ListarRuta(id *bson.ObjectID, ctx context.Context) (interface{}, error) {
	return nil, nil
}

func (r *ruta) ActualizarRuta(id *bson.ObjectID, ruta *model.Ruta, ctx context.Context) error {
	return nil
}

func (r *ruta) EliminarRuta(id *bson.ObjectID, ctx context.Context) error {
	return nil
}
