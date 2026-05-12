package service

import (
	"context"
	"transporteRuta/src/internal/coordenada/model"
	"transporteRuta/src/internal/coordenada/repository"

	"transporteRuta/src/internal/coordenada/dto"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Coordenada struct {
	coordenadaRepository repository.Coordenada
	cliente              *mongo.Client
}

func NewCoordenadaService(coordenadaRepository repository.Coordenada, cliente *mongo.Client) *Coordenada {
	return &Coordenada{
		coordenadaRepository: coordenadaRepository,
		cliente:              cliente,
	}
}

func (s *Coordenada) CrearCoordenada(coordenada *dto.CoordenadaDto, ctx context.Context) error {
	data := model.Coordenada{
		Latitud:  coordenada.Latitud,
		Longitud: coordenada.Longitud,
	}
	err := s.coordenadaRepository.CrearCoordenada(&data, ctx)
	if err != nil {
		return err
	}
	return nil
}

func (s *Coordenada) ListarCoordenada(id *bson.ObjectID, ctx context.Context) (interface{}, error) {
	return nil, nil
}

func (s *Coordenada) ActualizarCoordenada(id *bson.ObjectID, coordenada *dto.CoordenadaDto, ctx context.Context) error {
	return nil
}

func (s *Coordenada) EliminarCoordenada(id *bson.ObjectID, ctx context.Context) error {
	return nil
}
