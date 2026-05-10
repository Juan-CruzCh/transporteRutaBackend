package controller

import (
	"net/http"
	"transporteRuta/src/internal/usuario/service"
)

type Usuario struct {
	usuarioService *service.Usuario
}

func NewUsuarioController(usuarioService *service.Usuario) *Usuario {
	return &Usuario{
		usuarioService: usuarioService,
	}
}
func (c *Usuario) CrearUsuario(w http.ResponseWriter, r *http.Request) {
}

func (c *Usuario) ListarUsuario(w http.ResponseWriter, r *http.Request) {
}

func (c *Usuario) ActualizarUsuario(w http.ResponseWriter, r *http.Request) {
}

func (c *Usuario) EliminarUsuario(w http.ResponseWriter, r *http.Request) {
}
