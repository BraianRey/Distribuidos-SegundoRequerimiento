package vistas

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	comp "cliente.local/grpc-cliente/componentes"
	pbCancion "servidor.local/grpc-servidorcanciones/serviciosCanciones"
	pbStream "servidor.local/grpc-servidorstream/serviciosStreaming"
)

func MostrarMenuPrincipal(clientStream pbStream.AudioServiceClient, clientCancion pbCancion.CancionesServiceClient, ctx context.Context) {
	// lector de entrada estándar
	readerInput := bufio.NewReader(os.Stdin)
	// menú principal
	for {
		fmt.Print("\n1) Ver géneros\n0) Salir\nSeleccione una opción: ")
		opcion, _ := readerInput.ReadString('\n')
		opcion = strings.TrimSpace(opcion)
		if opcion == "0" {
			fmt.Println("Saliendo...")
			return
		}
		if opcion != "1" {
			fmt.Println("Opción inválida.")
			continue
		}
		// Paso 1: Seleccionar género usando componentes
		idGenero, ok, err := comp.SeleccionarGenero(clientCancion, ctx, readerInput)
		if err != nil {
			fmt.Println("Error al obtener géneros:", err)
			continue
		}
		if !ok {
			// usuario eligió volver o no hay géneros
			continue
		}

		// Paso 2 + 3: Seleccionar canción y obtener detalles
		_, detalle, ok, err := comp.SeleccionarCancion(clientCancion, ctx, idGenero, readerInput)
		if err != nil {
			fmt.Println("Error al obtener canción o detalles:", err)
			continue
		}
		if !ok {
			// volver al menú de géneros
			continue
		}

		// Mostrar información de la canción seleccionada
		fmt.Printf("\nDetalles de la canción:\nTítulo: %s\nArtista: %s\nAlbum: %s\nAño: %d\nDuración: %s\nGénero: %s\n",
			detalle.Titulo, detalle.Artista, detalle.Album, detalle.Anio, detalle.Duracion, detalle.Genero.Nombre)

		// Paso 4: Opción de reproducir o volver
		for {
			fmt.Print("\n1) Reproducir\n0) Volver\nSeleccione una opción: ")
			opc, _ := readerInput.ReadString('\n')
			opc = strings.TrimSpace(opc)
			if opc == "0" {
				break
			}
			if opc == "1" {
				if err := comp.ReproducirCancion(clientStream, detalle, ctx, readerInput); err != nil {
					log.Println("Error durante la reproducción:", err)
				}
				break
			}
			fmt.Println("Opción inválida.")
		}
	}
}
