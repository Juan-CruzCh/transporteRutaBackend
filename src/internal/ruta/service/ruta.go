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
	rutaRepository        repository.Ruta
	DetalleRutaRepository repository.DetalleRuta
	cliente               *mongo.Client
}

func NewRutaService(rutaRepository repository.Ruta, DetalleRutaRepository repository.DetalleRuta, cliente *mongo.Client) *Ruta {
	return &Ruta{
		rutaRepository:        rutaRepository,
		DetalleRutaRepository: DetalleRutaRepository,
		cliente:               cliente,
	}
}

func (s *Ruta) CrearRuta(ruta *dto.RutaDto, ctx context.Context) (*map[string]bson.ObjectID, error) {
	data := model.Ruta{
		Nombre: ruta.Nombre,
		Color:  ruta.Color,
	}
	id, err := s.rutaRepository.CrearRuta(&data, ctx)
	if err != nil {
		return nil, err
	}
	return &map[string]bson.ObjectID{"ruta": *id}, nil
}

func (s *Ruta) CrearDetalleRuta(ruta *dto.DetalleRutaDto, ctx context.Context) error {
	data := model.DetalleRuta{
		Coordenada: ruta.Coordenada,
		EsParada:   *ruta.EsParada,
		Ruta:       ruta.Ruta,
	}
	err := s.DetalleRutaRepository.CrearDetalleRuta(&data, ctx)
	if err != nil {
		return err
	}
	return nil
}

func (s *Ruta) ListarRuta(ctx context.Context) (*[]model.Ruta, error) {
	data, err := s.rutaRepository.ListarRuta(ctx)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *Ruta) ActualizarRuta(id *bson.ObjectID, ruta *dto.RutaDto, ctx context.Context) error {
	return nil
}

func (s *Ruta) EliminarRuta(id *bson.ObjectID, ctx context.Context) error {
	return nil
}
