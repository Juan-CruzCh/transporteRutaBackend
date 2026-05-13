package repository

import (
	"context"
	"transporteRuta/src/app/enum"
	"transporteRuta/src/app/utils"
	"transporteRuta/src/internal/linea/model"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type DetalleLinea interface {
	CrearDetalleLinea(detalleLinea *model.DetalleLinea, ctx context.Context) error
	ListarDetalleLinea(id *bson.ObjectID, ctx context.Context) (interface{}, error)
	ActualizarDetalleLinea(id *bson.ObjectID, detalleLinea *model.DetalleLinea, ctx context.Context) error
	EliminarDetalleLinea(id *bson.ObjectID, ctx context.Context) error
}

type detalleLinea struct {
	collection *mongo.Collection
}

func NewDetalleLineaRepository(db *mongo.Database) *detalleLinea {
	collection := db.Collection("DetalleLinea")
	return &detalleLinea{collection: collection}
}

func (r *detalleLinea) CrearDetalleLinea(detalleLinea *model.DetalleLinea, ctx context.Context) error {
	detalleLinea.Fecha = utils.FechaHoraBolivia()
	detalleLinea.Flag = enum.FlagNuevo
	_, err := r.collection.InsertOne(ctx, detalleLinea)
	if err != nil {
		return err
	}
	return nil
}

func (r *detalleLinea) ListarDetalleLinea(id *bson.ObjectID, ctx context.Context) (interface{}, error) {
	return nil, nil
}

func (r *detalleLinea) ActualizarDetalleLinea(id *bson.ObjectID, detalleLinea *model.DetalleLinea, ctx context.Context) error {
	return nil
}

func (r *detalleLinea) EliminarDetalleLinea(id *bson.ObjectID, ctx context.Context) error {
	return nil
}
