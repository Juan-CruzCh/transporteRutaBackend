package repository

import (
	"context"
	"transporteRuta/src/internal/ruta/model"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type DetalleRuta interface {
	CrearDetalleRuta(detalleRuta *model.DetalleRuta, ctx context.Context) error
	ListarDetalleRuta(id *bson.ObjectID, ctx context.Context) (interface{}, error)
	ActualizarDetalleRuta(id *bson.ObjectID, detalleRuta *model.DetalleRuta, ctx context.Context) error
	EliminarDetalleRuta(id *bson.ObjectID, ctx context.Context) error
}
type detalleRuta struct {
	collection *mongo.Collection
}

func NewDetalleRutaRepository(db *mongo.Database) *detalleRuta {
	collection := db.Collection("DetalleRuta")
	return &detalleRuta{collection: collection}
}

func (r *detalleRuta) CrearDetalleRuta(detalleRuta *model.DetalleRuta, ctx context.Context) error {
	return nil
}

func (r *detalleRuta) ListarDetalleRuta(id *bson.ObjectID, ctx context.Context) (interface{}, error) {
	return nil, nil
}

func (r *detalleRuta) ActualizarDetalleRuta(id *bson.ObjectID, detalleRuta *model.DetalleRuta, ctx context.Context) error {
	return nil
}

func (r *detalleRuta) EliminarDetalleRuta(id *bson.ObjectID, ctx context.Context) error {
	return nil
}
