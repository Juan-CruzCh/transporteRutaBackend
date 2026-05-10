package service

import (
	"context"
	"transporteRuta/src/internal/ubicacion/dto"
	"transporteRuta/src/internal/ubicacion/repository"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Ubicacion struct {
	ubicacionRepository repository.Ubicacion
	cliente             *mongo.Client
}

func NewUbicacionService(ubicacionRepository repository.Ubicacion, cliente *mongo.Client) *Ubicacion {
	return &Ubicacion{
		ubicacionRepository: ubicacionRepository,
		cliente:             cliente,
	}
}

func (s *Ubicacion) CrearUbicacion(ubicacion *dto.UbicacionDto, ctx context.Context) error {
	return nil
}

func (s *Ubicacion) ListarUbicacion(id *bson.ObjectID, ctx context.Context) (interface{}, error) {
	return nil, nil
}

func (s *Ubicacion) ActualizarUbicacion(id *bson.ObjectID, ubicacion *dto.UbicacionDto, ctx context.Context) error {
	return nil
}

func (s *Ubicacion) EliminarUbicacion(id *bson.ObjectID, ctx context.Context) error {
	return nil
}
