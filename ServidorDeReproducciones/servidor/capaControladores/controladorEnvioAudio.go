package capacontroladores

import (
	"fmt"

	"golang.org/x/net/context"
	"google.golang.org/grpc/peer"
	comunicacioneservidopreferencias "servidor.local/grpc-servidor/capaComunicacionExterna/comunicacionservidorPreferencias"
	capafachadaservices "servidor.local/grpc-servidor/capaFachadaServices"
	"servidor.local/grpc-servidor/capalogger"
	pb "servidor.local/grpc-servidor/serviciosCancion"
)

type ControladorServidor struct {
	pb.UnimplementedAudioServiceServer
	logger *capalogger.Logger
}

func NewControladorServidor(logger *capalogger.Logger) *ControladorServidor {
	return &ControladorServidor{
		logger: logger,
	}
}

func ObtenerDireccionCliente(ctx context.Context) string {
	if p, ok := peer.FromContext(ctx); ok {
		return p.Addr.String()
	}
	return "descononcido"
}

// Implementación del procedimiento remoto
func (thisC *ControladorServidor) EnviarCancionMedianteStream(
	req *pb.PeticionDTO, stream pb.AudioService_EnviarCancionMedianteStreamServer) error {

	//invocación a operación síncrona
	direccionCliente := ObtenerDireccionCliente(stream.Context())

	//invocación a operación asincrona que envia datos a otro proceso
	go func() {
		err := comunicacioneservidopreferencias.RegistrarReproduccionEnTendencias(req.Titulo, direccionCliente)
		if err != nil {
			fmt.Println("Error registrando tendencia:", err)
		}
	}()

	//invocación a operación asincrona que envia datos dentro del proceso
	go thisC.logger.AlmacenarSolicitud(req.Titulo, direccionCliente)

	return capafachadaservices.StreamAudioFile(
		req.Titulo,
		func(data []byte) error {
			return stream.Send(&pb.FragmentoCancion{Data: data})
		})
}
