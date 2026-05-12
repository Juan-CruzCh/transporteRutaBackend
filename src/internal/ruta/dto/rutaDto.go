package dto

type RutaDto struct {
	Nombre string `bson:"nombre" json:"nombre" validate:"required"`
	Color  string `bson:"color" json:"color" validate:"required"`
}
