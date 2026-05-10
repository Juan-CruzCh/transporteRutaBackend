package repository

import (
	"context"
	"transporteRuta/src/internal/ubicacion/model"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Ubicacion interface {
	CrearUbicacion(ubicacion *model.Ubicacion, ctx context.Context) error
	ListarUbicacion(id *bson.ObjectID, ctx context.Context) (interface{}, error)
	ActualizarUbicacion(id *bson.ObjectID, ubicacion *model.Ubicacion, ctx context.Context) error
	EliminarUbicacion(id *bson.ObjectID, ctx context.Context) error
}

type ubicacion struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewUbicacionRepository(db *mongo.Database) *ubicacion {
	collection := db.Collection("Ubicacion")
	return &ubicacion{db: db, collection: collection}
}

func (r *ubicacion) CrearUbicacion(ubicacion *model.Ubicacion, ctx context.Context) error {
	return nil
}

func (r *ubicacion) ListarUbicacion(id *bson.ObjectID, ctx context.Context) (interface{}, error) {
	return nil, nil
}

func (r *ubicacion) ActualizarUbicacion(id *bson.ObjectID, ubicacion *model.Ubicacion, ctx context.Context) error {
	return nil
}

func (r *ubicacion) EliminarUbicacion(id *bson.ObjectID, ctx context.Context) error {
	return nil
}
