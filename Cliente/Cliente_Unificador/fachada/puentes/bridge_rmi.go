package bridges

import (
	"context"
	"fmt"
)

type ModuloRMI struct{}

func (m *ModuloRMI) Nombre() string {
	return "Cliente RMI"
}

func (m *ModuloRMI) Iniciar(ctx context.Context) error {
	fmt.Println("Módulo RMI aún no implementado.")
	return nil
}

func (m *ModuloRMI) Cerrar() error {
	return nil
}
