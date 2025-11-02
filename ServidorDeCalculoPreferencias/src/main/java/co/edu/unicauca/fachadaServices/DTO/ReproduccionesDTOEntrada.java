package co.edu.unicauca.fachadaServices.DTO;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@AllArgsConstructor
@NoArgsConstructor
public class ReproduccionesDTOEntrada {
    private Integer id;
    private Integer idUsuario;
    private Integer idCancion;
    private String titulo;
    private String artista;
    private String genero;
    private String idioma;
}