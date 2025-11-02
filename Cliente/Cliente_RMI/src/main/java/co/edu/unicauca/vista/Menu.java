package co.edu.unicauca.vista;

import java.rmi.RemoteException;

import co.edu.unicauca.capaDeControladores.ControladorPreferenciasUsuariosInt;
import co.edu.unicauca.configuracion.lector.LectorPropiedadesConfig;
import co.edu.unicauca.configuracion.servicios.ClienteDeObjetos;
import co.edu.unicauca.fachadaServices.DTO.PreferenciasDTORespuesta;
import co.edu.unicauca.fachadaServices.services.FachadaGestorUsuariosIml;
import co.edu.unicauca.utilidades.UtilidadesConsola;

/**
 * Menú principal del Cliente RMI
 * Permite consultar las preferencias musicales de los usuarios
 */
public class Menu {
    
    private FachadaGestorUsuariosIml fachada;
    private LectorPropiedadesConfig config;
    
    public Menu() {
        this.config = new LectorPropiedadesConfig();
    }
    
    /**
     * Menú principal del cliente
     */
    public void ejecutarMenuPrincipal() {
        int opcion = 0;
        
        // Conectar al servidor RMI
        if (!conectarServidorRMI()) {
            System.out.println("❌ No se pudo conectar al servidor RMI");
            System.out.println("Verifique que el ServidorDeCalculoPreferencias esté ejecutándose");
            return;
        }
        
        do {
            System.out.println("\n========================================");
            System.out.println("       MENÚ CLIENTE RMI");
            System.out.println("========================================");
            System.out.println("1. Consultar preferencias de usuario");
            System.out.println("2. Salir");
            System.out.println("========================================");
            System.out.print("Seleccione una opción: ");
            
            opcion = UtilidadesConsola.leerEntero();
            
            switch (opcion) {
                case 1:
                    consultarPreferencias();
                    break;
                case 2:
                    System.out.println("\n¡Hasta pronto!");
                    break;
                default:
                    System.out.println("❌ Opción inválida");
            }
            
        } while (opcion != 2);
    }
    
    /**
     * Conectar al servidor RMI de cálculo de preferencias
     */
    private boolean conectarServidorRMI() {
        try {
            System.out.println("\n🔌 Conectando al servidor RMI...");
            
            // Leer configuración
            String ipServidor = config.obtenerIPServidor();
            int puertoServidor = config.obtenerPuertoServidor();
            String nombreObjeto = "ObjetoRemotoPreferencias";
            
            System.out.println("   - IP: " + ipServidor);
            System.out.println("   - Puerto: " + puertoServidor);
            System.out.println("   - Objeto: " + nombreObjeto);
            
            // Obtener objeto remoto
            ControladorPreferenciasUsuariosInt objRemoto = 
                ClienteDeObjetos.obtenerObjetoRemoto(ipServidor, puertoServidor, nombreObjeto);
            
            if (objRemoto == null) {
                return false;
            }
            
            // Crear fachada
            this.fachada = new FachadaGestorUsuariosIml(objRemoto);
            
            System.out.println("✅ Conexión exitosa al servidor RMI\n");
            return true;
            
        } catch (Exception e) {
            System.out.println("❌ Error al conectar: " + e.getMessage());
            return false;
        }
    }
    
    /**
     * Consultar las preferencias musicales de un usuario
     */
    private void consultarPreferencias() {
        try {
            System.out.println("\n========================================");
            System.out.println("   CONSULTAR PREFERENCIAS MUSICALES");
            System.out.println("========================================");
            
            System.out.print("Ingrese el ID del usuario: ");
            int idUsuario = UtilidadesConsola.leerEntero();
            
            System.out.println("\n⏳ Consultando preferencias...");
            System.out.println("   (El servidor está consultando canciones y reproducciones)");
            
            // Llamar al servidor RMI
            PreferenciasDTORespuesta preferencias = fachada.getReferencias(idUsuario);
            
            if (preferencias != null) {
                mostrarPreferencias(idUsuario, preferencias);
            } else {
                System.out.println("\n❌ No se pudieron obtener las preferencias");
            }
            
        } catch (RemoteException e) {
            System.out.println("\n❌ Error en la comunicación RMI: " + e.getMessage());
        } catch (Exception e) {
            System.out.println("\n❌ Error: " + e.getMessage());
        }
    }
    
    /**
     * Mostrar las preferencias del usuario de forma organizada
     */
    private void mostrarPreferencias(int idUsuario, PreferenciasDTORespuesta preferencias) {
        System.out.println("\n╔════════════════════════════════════════╗");
        System.out.println("║   PREFERENCIAS DEL USUARIO #" + idUsuario + "         ║");
        System.out.println("╚════════════════════════════════════════╝");
        
        // Preferencias por Género
        System.out.println("\n🎵 PREFERENCIAS POR GÉNERO:");
        System.out.println("   ────────────────────────────");
        if (preferencias.getPreferenciasGenero() != null && !preferencias.getPreferenciasGenero().isEmpty()) {
            preferencias.getPreferenciasGenero().forEach(pref -> {
                System.out.println("   • " + pref.getGenero() + ": " + pref.getCantidad() + " canciones");
            });
        } else {
            System.out.println("   (No hay reproducciones registradas)");
        }
        
        // Preferencias por Artista
        System.out.println("\n👤 PREFERENCIAS POR ARTISTA:");
        System.out.println("   ────────────────────────────");
        if (preferencias.getPreferenciasArtista() != null && !preferencias.getPreferenciasArtista().isEmpty()) {
            preferencias.getPreferenciasArtista().forEach(pref -> {
                System.out.println("   • " + pref.getArtista() + ": " + pref.getCantidad() + " canciones");
            });
        } else {
            System.out.println("   (No hay reproducciones registradas)");
        }
        
        // Preferencias por Idioma
        System.out.println("\n🌍 PREFERENCIAS POR IDIOMA:");
        System.out.println("   ────────────────────────────");
        if (preferencias.getPreferenciasIdioma() != null && !preferencias.getPreferenciasIdioma().isEmpty()) {
            preferencias.getPreferenciasIdioma().forEach(pref -> {
                System.out.println("   • " + pref.getIdioma() + ": " + pref.getCantidad() + " canciones");
            });
        } else {
            System.out.println("   (No hay reproducciones registradas)");
        }
        
        System.out.println("\n════════════════════════════════════════");
    }
}