package repository

import (
	"context"
	"transporteRuta/src/internal/coordenada/model"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Coordenada interface {
	CrearCoordenada(coordenada *model.Coordenada, ctx context.Context) error
	ListarCoordenada(id *bson.ObjectID, ctx context.Context) (interface{}, error)
	ActualizarCoordenada(id *bson.ObjectID, coordenada *model.Coordenada, ctx context.Context) error
	EliminarCoordenada(id *bson.ObjectID, ctx context.Context) error
}

type coordenada struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewCoordenadaRepository(db *mongo.Database) *coordenada {
	collection := db.Collection("Coordenada")
	return &coordenada{db: db, collection: collection}
}

func (r *coordenada) CrearCoordenada(coordenada *model.Coordenada, ctx context.Context) error {

	return nil
}

func (r *coordenada) ListarCoordenada(id *bson.ObjectID, ctx context.Context) (interface{}, error) {
	return nil, nil
}

func (r *coordenada) ActualizarCoordenada(id *bson.ObjectID, coordenada *model.Coordenada, ctx context.Context) error {
	return nil
}

func (r *coordenada) EliminarCoordenada(id *bson.ObjectID, ctx context.Context) error {
	return nil
}
