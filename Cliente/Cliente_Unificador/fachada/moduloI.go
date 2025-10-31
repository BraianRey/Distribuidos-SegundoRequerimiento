package interfaces

import "context"

type ClienteModulo interface {
	Nombre() string
	Iniciar(ctx context.Context) error
}
