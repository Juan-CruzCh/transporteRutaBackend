package repository

import (
	"context"
	"fmt"
	"transporteRuta/src/app/enum"
	"transporteRuta/src/app/utils"
	"transporteRuta/src/internal/ruta/model"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Ruta interface {
	CrearRuta(ruta *model.Ruta, ctx context.Context) (*bson.ObjectID, error)
	ListarRuta(ctx context.Context) (*[]model.Ruta, error)
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

func (r *ruta) CrearRuta(ruta *model.Ruta, ctx context.Context) (*bson.ObjectID, error) {
	ruta.Fecha = utils.FechaHoraBolivia()
	ruta.Flag = enum.FlagNuevo
	resultado, err := r.collection.InsertOne(ctx, ruta)
	if err != nil {
		return nil, err
	}
	ID, ok := resultado.InsertedID.(bson.ObjectID)
	if !ok {
		return nil, fmt.Errorf("Error de parseo")
	}
	return &ID, nil
}

func (r *ruta) ListarRuta(ctx context.Context) (*[]model.Ruta, error) {
	cursor, err := r.collection.Find(ctx, bson.M{"flag": enum.FlagNuevo})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var data []model.Ruta = []model.Ruta{}
	for cursor.Next(ctx) {
		var r model.Ruta
		err = cursor.Decode(&r)
		if err != nil {
			return nil, err
		}
		data = append(data, r)

	}
	return &data, nil
}

func (r *ruta) ActualizarRuta(id *bson.ObjectID, ruta *model.Ruta, ctx context.Context) error {
	return nil
}

func (r *ruta) EliminarRuta(id *bson.ObjectID, ctx context.Context) error {
	return nil
}
