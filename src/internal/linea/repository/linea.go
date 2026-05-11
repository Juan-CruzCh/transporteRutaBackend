package repository

import (
	"context"
	"transporteRuta/src/app/enum"
	"transporteRuta/src/app/utils"
	"transporteRuta/src/internal/linea/model"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Linea interface {
	CrearLinea(linea *model.Linea, ctx context.Context) error
	ListarLinea(id *bson.ObjectID, ctx context.Context) (interface{}, error)
	ActualizarLinea(id *bson.ObjectID, linea *model.Linea, ctx context.Context) error
	EliminarLinea(id *bson.ObjectID, ctx context.Context) error
}

type linea struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewLineaRepository(db *mongo.Database) *linea {
	collection := db.Collection("Linea")
	return &linea{db: db, collection: collection}
}

func (r *linea) CrearLinea(linea *model.Linea, ctx context.Context) error {
	linea.Flag = enum.FlagNuevo
	linea.Fecha = utils.FechaHoraBolivia()
	_, err := r.collection.InsertOne(ctx, linea)
	if err != nil {
		return err
	}
	return nil
}

func (r *linea) ListarLinea(id *bson.ObjectID, ctx context.Context) (interface{}, error) {
	return nil, nil
}

func (r *linea) ActualizarLinea(id *bson.ObjectID, linea *model.Linea, ctx context.Context) error {
	return nil
}

func (r *linea) EliminarLinea(id *bson.ObjectID, ctx context.Context) error {
	return nil
}
