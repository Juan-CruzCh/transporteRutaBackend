package dto

type LineaDto struct {
	Nombre      string `json:"nombre" bson:"nombre" validate:"required"`
	Descripcion string `json:"descripcion" bson:"descripcion" validate:"required"`
	Color       string `json:"color" bson:"color" validate:"required"`
}
