package service

import (
	"context"
	"transporteRuta/src/internal/linea/dto"
	"transporteRuta/src/internal/linea/model"
	"transporteRuta/src/internal/linea/repository"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Linea struct {
	lineaRepository       repository.Linea
	DetalleRutaRepository repository.DetalleLinea
	cliente               *mongo.Client
}

func NewLineaService(lineaRepository repository.Linea, DetalleRutaRepository repository.DetalleLinea, cliente *mongo.Client) *Linea {
	return &Linea{
		lineaRepository:       lineaRepository,
		DetalleRutaRepository: DetalleRutaRepository,
		cliente:               cliente,
	}
}

func (s *Linea) CrearLinea(linea *dto.LineaDto, ctx context.Context) error {
	data := model.Linea{
		Nombre:      linea.Nombre,
		Descripcion: linea.Descripcion,
		Color:       linea.Color,
	}
	err := s.lineaRepository.CrearLinea(&data, ctx)
	if err != nil {
		return err
	}
	return nil
}

func (s *Linea) ListarLinea(ctx context.Context) (*[]model.Linea, error) {
	data, err := s.lineaRepository.ListarLinea(ctx)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *Linea) ActualizarLinea(id *bson.ObjectID, linea *dto.LineaDto, ctx context.Context) error {
	return nil
}

func (s *Linea) EliminarLinea(id *bson.ObjectID, ctx context.Context) error {
	return nil
}
