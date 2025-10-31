package main

import (
	"context"
	"fmt"

	interfaces "cliente.local/unificador/fachada"
	bridges "cliente.local/unificador/fachada/puentes"
	ui "cliente.local/unificador/vistas"
)

func main() {
	fmt.Println("=== Cliente Unificador ===")

	ctx := context.Background()

	// Módulos disponibles: gRPC y RMI
	modulos := []interfaces.ClienteModulo{
		&bridges.ModuloGRPC{},
		&bridges.ModuloRMI{},
	}

	// Ejecutar menú principal
	ui.MostrarMenuUnificado(modulos, ctx)
}
