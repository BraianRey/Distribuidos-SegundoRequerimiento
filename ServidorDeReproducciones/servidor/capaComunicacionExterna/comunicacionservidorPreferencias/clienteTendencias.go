package comunicacioneservidopreferencias

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Estructura que coincide con el DTO del servicio de tendencias
type ReproduccionDTOInput struct {
	Titulo  string `json:"titulo"`
	Cliente string `json:"cliente"`
}

// Función que envía una reproducción al microservicio de tendencias
func RegistrarReproduccionEnTendencias(titulo string, cliente string) error {
	url := "http://localhost:5000/tendencias/reproduccion" // endpoint del microservicio tendencias

	// Crear el cuerpo JSON
	body := ReproduccionDTOInput{
		Titulo:  titulo,
		Cliente: cliente,
	}

	jsonData, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("error convirtiendo a JSON: %v", err)
	}

	// Enviar solicitud POST al microservicio de tendencias
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("error al enviar POST a tendencias: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("el servicio de tendencias respondió con código %d", resp.StatusCode)
	}

	fmt.Println("Reproducción registrada en el microservicio de tendencias:", titulo)
	return nil
}
