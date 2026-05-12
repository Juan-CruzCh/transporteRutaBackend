package repository

import (
	"context"
	"transporteRuta/src/internal/usuario/model"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Usuario interface {
	CrearUsuario(usuario *model.Usuario, ctx context.Context) error
	ListarUsuario(id *bson.ObjectID, ctx context.Context) (interface{}, error)
	ActualizarUsuario(id *bson.ObjectID, usuario *model.Usuario, ctx context.Context) error
	EliminarUsuario(id *bson.ObjectID, ctx context.Context) error
}

type usuario struct {
	collection *mongo.Collection
}

func NewUsuarioRepository(db *mongo.Database) *usuario {
	collection := db.Collection("Usuario")
	return &usuario{collection: collection}
}

func (r *usuario) CrearUsuario(usuario *model.Usuario, ctx context.Context) error {
	return nil
}

func (r *usuario) ListarUsuario(id *bson.ObjectID, ctx context.Context) (interface{}, error) {
	return nil, nil
}

func (r *usuario) ActualizarUsuario(id *bson.ObjectID, usuario *model.Usuario, ctx context.Context) error {
	return nil
}

func (r *usuario) EliminarUsuario(id *bson.ObjectID, ctx context.Context) error {
	return nil
}
