package service

import (
	"context"
	"transporteRuta/src/internal/ruta/dto"
	"transporteRuta/src/internal/ruta/model"
	"transporteRuta/src/internal/ruta/repository"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Ruta struct {
	rutaRepository repository.Ruta
	cliente        *mongo.Client
}

func NewRutaService(rutaRepository repository.Ruta, cliente *mongo.Client) *Ruta {
	return &Ruta{
		rutaRepository: rutaRepository,
		cliente:        cliente,
	}
}

func (s *Ruta) CrearRuta(ruta *dto.RutaDto, ctx context.Context) error {

	data := model.Ruta{
		Latitud:  ruta.Latitud,
		Longitud: ruta.Longitud,
	}
	err := s.rutaRepository.CrearRuta(&data, ctx)
	if err != nil {
		return err
	}
	return nil
}

func (s *Ruta) ListarRuta(id *bson.ObjectID, ctx context.Context) (interface{}, error) {
	return nil, nil
}

func (s *Ruta) ActualizarRuta(id *bson.ObjectID, ruta *dto.RutaDto, ctx context.Context) error {
	return nil
}

func (s *Ruta) EliminarRuta(id *bson.ObjectID, ctx context.Context) error {
	return nil
}
