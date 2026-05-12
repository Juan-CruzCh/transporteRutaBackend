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
	ListarLinea(ctx context.Context) (*[]model.Linea, error)
	ActualizarLinea(id *bson.ObjectID, linea *model.Linea, ctx context.Context) error
	EliminarLinea(id *bson.ObjectID, ctx context.Context) error
}

type linea struct {
	collection *mongo.Collection
}

func NewLineaRepository(db *mongo.Database) *linea {
	collection := db.Collection("Linea")
	return &linea{collection: collection}
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

func (r *linea) ListarLinea(ctx context.Context) (*[]model.Linea, error) {

	var data []model.Linea = []model.Linea{}

	cursor, err := r.collection.Find(ctx, bson.M{"flag": enum.FlagNuevo})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	err = cursor.All(ctx, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *linea) ActualizarLinea(id *bson.ObjectID, linea *model.Linea, ctx context.Context) error {
	return nil
}

func (r *linea) EliminarLinea(id *bson.ObjectID, ctx context.Context) error {
	return nil
}
