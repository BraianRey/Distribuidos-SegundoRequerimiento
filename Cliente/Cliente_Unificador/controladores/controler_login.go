package controladores

import "cliente.local/unificador/servicios"

type LoginController struct {
	servicio *servicios.ServicioLogin
}

func NuevoLoginController(servicio *servicios.ServicioLogin) *LoginController {
	return &LoginController{servicio: servicio}
}

func (c *LoginController) ValidarUsuario(usuario, password string) (bool, error) {
	return c.servicio.VerificarCredenciales(usuario, password)
}
