package co.edu.unicauca.fachadaServices.DTO;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@AllArgsConstructor
@NoArgsConstructor
@JsonIgnoreProperties(ignoreUnknown = true)  // ← Esto ignora campos como fechaHora
public class CancionDTOEntrada {
    private Integer id;
    private Integer idUsuario;
    private Integer idCancion;
    private String titulo;
    private String artista;
    private String genero;
    private String idioma;

    @Override
    public String toString() {
        return titulo + " - " + artista + " [" + genero + ", " + idioma + "]";
    }
}