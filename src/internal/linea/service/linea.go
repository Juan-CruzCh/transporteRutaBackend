package service

import (
	"context"
	"transporteRuta/src/internal/linea/dto"
	"transporteRuta/src/internal/linea/repository"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Linea struct {
	lineaRepository repository.Linea
	cliente         *mongo.Client
}

func NewLineaService(lineaRepository repository.Linea, cliente *mongo.Client) *Linea {
	return &Linea{
		lineaRepository: lineaRepository,
		cliente:         cliente,
	}
}

func (s *Linea) CrearLinea(linea *dto.LineaDto, ctx context.Context) error {
	return nil
}

func (s *Linea) ListarLinea(id *bson.ObjectID, ctx context.Context) (interface{}, error) {
	return nil, nil
}

func (s *Linea) ActualizarLinea(id *bson.ObjectID, linea *dto.LineaDto, ctx context.Context) error {
	return nil
}

func (s *Linea) EliminarLinea(id *bson.ObjectID, ctx context.Context) error {
	return nil
}
