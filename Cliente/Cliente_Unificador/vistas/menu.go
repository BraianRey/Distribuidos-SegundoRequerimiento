package ui

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	interfaces "cliente.local/unificador/fachada"
)

func MostrarMenuUnificado(modulos []interfaces.ClienteModulo, ctx context.Context) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n===== CLIENTE UNIFICADOR =====")
		for i, m := range modulos {
			fmt.Printf("%d) %s\n", i+1, m.Nombre())
		}
		fmt.Println("0) Salir")
		fmt.Print("Seleccione una opción: ")

		opcion, _ := reader.ReadString('\n')
		opcion = strings.TrimSpace(opcion)

		switch opcion {
		case "0":
			fmt.Println("Saliendo del cliente unificador...")
			return

		default:
			idx := int(opcion[0] - '1')
			if idx >= 0 && idx < len(modulos) {
				fmt.Printf("Ejecutando módulo: %s...\n", modulos[idx].Nombre())
				if err := modulos[idx].Iniciar(ctx); err != nil {
					fmt.Println("Error al ejecutar módulo:", err)
				}
			} else {
				fmt.Println("Opción inválida.")
			}
		}
	}
}
