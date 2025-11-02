package co.edu.unicauca.fachadaServices.DTO;

import java.io.Serializable;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@AllArgsConstructor
@NoArgsConstructor
public class CancionDTOEntrada implements Serializable {
    private Integer id;
    private String titulo;
    private String artista;
    private String genero;
    private String idioma;  // ← AGREGADO
}

