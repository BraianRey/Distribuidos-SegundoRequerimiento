package bridges

import (
	"context"

	clientgrpc "cliente.local/grpc-cliente/conexion"
)

type ModuloGRPC struct{}

func (m *ModuloGRPC) Nombre() string { return "Cliente gRPC" }

func (m *ModuloGRPC) Iniciar(ctx context.Context) error {
	clientgrpc.RunClienteGRPC()
	return nil
}
