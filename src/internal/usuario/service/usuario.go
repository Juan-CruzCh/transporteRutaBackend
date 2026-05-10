package service

import (
	"context"
	"transporteRuta/src/internal/usuario/dto"
	"transporteRuta/src/internal/usuario/repository"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Usuario struct {
	usuarioRepository repository.Usuario
	cliente           *mongo.Client
}

func NewUsuarioService(usuarioRepository repository.Usuario, cliente *mongo.Client) *Usuario {
	return &Usuario{
		usuarioRepository: usuarioRepository,
		cliente:           cliente,
	}
}

func (s *Usuario) CrearUsuario(usuario *dto.UsuarioDto, ctx context.Context) error {
	return nil
}

func (s *Usuario) ListarUsuario(id *bson.ObjectID, ctx context.Context) (interface{}, error) {
	return nil, nil
}

func (s *Usuario) ActualizarUsuario(id *bson.ObjectID, usuario *dto.UsuarioDto, ctx context.Context) error {
	return nil
}

func (s *Usuario) EliminarUsuario(id *bson.ObjectID, ctx context.Context) error {
	return nil
}
