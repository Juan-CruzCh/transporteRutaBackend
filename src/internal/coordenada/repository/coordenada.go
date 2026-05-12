package repository

import (
	"context"
	"transporteRuta/src/app/enum"
	"transporteRuta/src/app/utils"
	"transporteRuta/src/internal/coordenada/model"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Coordenada interface {
	CrearCoordenada(coordenada *model.Coordenada, ctx context.Context) error
	ListarCoordenada(ctx context.Context) (*[]model.Coordenada, error)
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
	coordenada.Fecha = utils.FechaHoraBolivia()
	coordenada.Flag = enum.FlagNuevo
	_, err := r.collection.InsertOne(ctx, coordenada)
	if err != nil {
		return err
	}
	return nil
}

func (r *coordenada) ListarCoordenada(ctx context.Context) (*[]model.Coordenada, error) {
	cursor, err := r.collection.Find(ctx, bson.M{"flag": enum.FlagNuevo})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var data []model.Coordenada = []model.Coordenada{}
	for cursor.Next(ctx) {
		var c model.Coordenada
		err = cursor.Decode(&c)
		if err != nil {
			return nil, err
		}
		data = append(data, c)
	}
	return &data, nil
}

func (r *coordenada) ActualizarCoordenada(id *bson.ObjectID, coordenada *model.Coordenada, ctx context.Context) error {
	return nil
}

func (r *coordenada) EliminarCoordenada(id *bson.ObjectID, ctx context.Context) error {
	return nil
}
