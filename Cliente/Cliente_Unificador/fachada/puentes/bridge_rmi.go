package puentes

import (
	"context"
)

type ModuloRMI struct{}

func (m *ModuloRMI) Nombre() string {
	return "Cliente RMI"
}

func (m *ModuloRMI) Iniciar(ctx context.Context) error {
	// IMPLEMENTAR LÓGICA DEL CLIENTE RMI AQUÍ
	return nil
}

func (m *ModuloRMI) Cerrar() error {
	return nil
}
