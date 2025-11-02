# 🎯 Cliente RMI - Sistema de Streaming Musical

## 📦 Proyecto Completo y Funcional

Este es el **Cliente RMI completo** con todos los archivos necesarios, corregidos y listos para compilar.

---

## ✅ Archivos Incluidos

### Estructura Completa:
```
Cliente_RMI/
├── pom.xml                                      ✅ Con Lombok
├── src/
│   ├── main/
│   │   ├── java/co/edu/unicauca/
│   │   │   ├── capaDeControladores/
│   │   │   │   └── ControladorPreferenciasUsuariosInt.java
│   │   │   ├── configuracion/
│   │   │   │   ├── lector/
│   │   │   │   │   └── LectorPropiedadesConfig.java
│   │   │   │   └── servicios/
│   │   │   │       └── ClienteDeObjetos.java
│   │   │   ├── fachadaServices/
│   │   │   │   ├── DTO/
│   │   │   │   │   ├── PreferenciaGeneroDTORespuesta.java      ✅ Corregido
│   │   │   │   │   ├── PreferenciaArtistaDTORespuesta.java     ✅ Corregido
│   │   │   │   │   ├── PreferenciaIdiomaDTORespuesta.java      ✅ NUEVO
│   │   │   │   │   └── PreferenciasDTORespuesta.java           ✅ Corregido
│   │   │   │   └── services/
│   │   │   │       ├── DTO/                     (DTOs duplicados - ignorar)
│   │   │   │       └── FachadaGestorUsuariosIml.java
│   │   │   ├── main/
│   │   │   │   └── Main.java                                   ✅ Corregido
│   │   │   ├── utilidades/
│   │   │   │   └── UtilidadesConsola.java
│   │   │   └── vista/
│   │   │       └── Menu.java                                   ✅ Corregido
│   │   └── resources/
│   │       └── application.properties
│   └── test/
│       └── java/
└── README.md                                    ✅ Este archivo
```

**Total:** 14 archivos Java + configuración

---

## 🚀 Inicio Rápido

### Paso 1: Requisitos
- Java 11 o superior
- Maven 3.6+
- ServidorDeCalculoPreferencias ejecutándose en localhost:1099

### Paso 2: Compilar
```bash
cd Cliente_RMI

# Limpiar y compilar
mvn clean compile
```

**Salida esperada:**
```
[INFO] BUILD SUCCESS
[INFO] Total time: X.XXX s
```

### Paso 3: Ejecutar
```bash
mvn exec:java -Dexec.mainClass="co.edu.unicauca.main.Main"
```

---

## 🔧 Configuración

### application.properties
```properties
servidor.ip=localhost
servidor.puerto=1099
```

Puedes modificar estos valores si tu servidor RMI está en otra IP o puerto.

---

## 📝 Funcionalidades

### Menú del Cliente:
```
========================================
       MENÚ CLIENTE RMI
========================================
1. Consultar preferencias de usuario
2. Salir
========================================
```

### ¿Qué hace?

1. **Consultar preferencias:**
   - Solicita ID de usuario
   - Se conecta al ServidorDeCalculoPreferencias vía RMI
   - El servidor consulta:
     - Canciones disponibles (desde ServidorDeCanciones)
     - Reproducciones del usuario (desde ServidorDeReproducciones)
   - Calcula y retorna preferencias por:
     - 🎵 Género
     - 👤 Artista
     - 🌍 Idioma

---

## 🛠️ Tecnologías Utilizadas

- **Java 11+**
- **Maven** - Gestión de dependencias
- **RMI** (Remote Method Invocation) - Comunicación remota
- **Lombok** - Generación automática de código (getters, setters, etc.)

---

## 📚 Dependencias (pom.xml)

```xml
<dependency>
    <groupId>org.projectlombok</groupId>
    <artifactId>lombok</artifactId>
    <version>1.18.30</version>
    <scope>provided</scope>
</dependency>
```

---

## 🔍 Componentes Principales

### 1. Main.java
Punto de entrada de la aplicación. Inicia el menú principal.

### 2. Menu.java
Maneja la interfaz de usuario:
- Conexión al servidor RMI
- Consulta de preferencias
- Presentación de resultados

### 3. DTOs (Data Transfer Objects)
Clases para transferir datos entre cliente y servidor:
- `PreferenciasDTORespuesta` - Contiene todas las preferencias
- `PreferenciaGeneroDTORespuesta` - Preferencia por género
- `PreferenciaArtistaDTORespuesta` - Preferencia por artista
- `PreferenciaIdiomaDTORespuesta` - Preferencia por idioma

### 4. ClienteDeObjetos.java
Maneja la conexión RMI con el servidor.

### 5. UtilidadesConsola.java
Funciones auxiliares para leer entrada del usuario.

---

## 🧪 Pruebas

### Prueba 1: Compilación
```bash
mvn clean compile
```
✅ Debe compilar sin errores

### Prueba 2: Verificar DTOs
```bash
ls target/classes/co/edu/unicauca/fachadaServices/DTO/
```
✅ Deben existir 4 archivos .class

### Prueba 3: Ejecución
```bash
mvn exec:java -Dexec.mainClass="co.edu.unicauca.main.Main"
```
✅ Debe iniciar el menú

### Prueba 4: Consultar Preferencias
1. Iniciar ServidorDeCalculoPreferencias
2. Ejecutar cliente
3. Opción 1 → Ingresar ID de usuario
4. Ver preferencias calculadas

---

## 🔗 Integración con Servidores

### Servidor Requerido:
**ServidorDeCalculoPreferencias** (RMI - Puerto 1099)

### Servidores que consulta el Servidor de Preferencias:
- ServidorDeCanciones (gRPC - Puerto 50051)
- ServidorDeReproducciones (REST - Puerto 3000)

---

## ❓ Solución de Problemas

### Error: "Cannot find symbol @Data"
**Causa:** Lombok no está configurado correctamente  
**Solución:** Verifica que el pom.xml tiene la dependencia de Lombok

### Error: "Connection refused"
**Causa:** ServidorDeCalculoPreferencias no está ejecutándose  
**Solución:** 
```bash
cd ServidorDeCalculoPreferencias
mvn exec:java -Dexec.mainClass="co.edu.unicauca.main.Main"
```

### Error: "Cannot find symbol getNombreGenero()"
**Causa:** Los DTOs no tienen Lombok o están mal configurados  
**Solución:** Este proyecto ya tiene los DTOs corregidos

### Error: "method leerEntero(String) not found"
**Causa:** Menu.java está llamando mal a UtilidadesConsola  
**Solución:** Este proyecto ya tiene Menu.java corregido

---

## 📊 Ejemplo de Salida

```
===============================================
   CLIENTE RMI - SISTEMA DE STREAMING
===============================================

🔌 Conectando al servidor RMI...
   - IP: localhost
   - Puerto: 1099
   - Objeto: ObjetoRemotoPreferencias
✅ Conexión exitosa al servidor RMI

========================================
       MENÚ CLIENTE RMI
========================================
1. Consultar preferencias de usuario
2. Salir
========================================
Seleccione una opción: 1

========================================
   CONSULTAR PREFERENCIAS MUSICALES
========================================
Ingrese el ID del usuario: 1

⏳ Consultando preferencias...
   (El servidor está consultando canciones y reproducciones)

╔════════════════════════════════════════╗
║   PREFERENCIAS DEL USUARIO #1          ║
╚════════════════════════════════════════╝

🎵 PREFERENCIAS POR GÉNERO:
   ────────────────────────────
   • Rock: 5 canciones
   • Pop: 3 canciones
   • EDM: 2 canciones

👤 PREFERENCIAS POR ARTISTA:
   ────────────────────────────
   • Imagine Dragons: 5 canciones
   • Taylor Swift: 3 canciones

🌍 PREFERENCIAS POR IDIOMA:
   ────────────────────────────
   • Inglés: 8 canciones
   • Español: 2 canciones

════════════════════════════════════════
```

---

## ✅ Verificación de Archivos Corregidos

Los siguientes archivos fueron **corregidos** de las versiones corruptas del RAR:

- [x] **pom.xml** - Agregada dependencia de Lombok
- [x] **Main.java** - Creado desde cero (estaba vacío)
- [x] **Menu.java** - Creado desde cero (estaba vacío)
- [x] **PreferenciaGeneroDTORespuesta.java** - Creado (estaba vacío)
- [x] **PreferenciaArtistaDTORespuesta.java** - Creado (estaba vacío)
- [x] **PreferenciaIdiomaDTORespuesta.java** - Creado (no existía)
- [x] **PreferenciasDTORespuesta.java** - Corregido

Todos los demás archivos originales se mantienen intactos.

---

## 🎓 Cumplimiento del Requerimiento

✅ **Cliente mediante un menú puede:**
- Iniciar sesión (simulado con ID de usuario)
- Ver las preferencias (✅ **implementado**)

✅ **Comunicación RMI:**
- Cliente se conecta a ServidorDeCalculoPreferencias
- Transferencia de DTOs con serialización

✅ **Patrones de Diseño:**
- MVC: Main (controlador), Menu (vista), DTOs (modelo)
- DTO: Clases de transferencia de datos
- Capas: Separación en paquetes

---

## 📞 Comandos Útiles

```bash
# Compilar
mvn clean compile

# Ejecutar
mvn exec:java -Dexec.mainClass="co.edu.unicauca.main.Main"

# Crear JAR
mvn clean package

# Ejecutar JAR
java -jar target/Cliente_RMI-1.0-SNAPSHOT.jar

# Limpiar
mvn clean
```

---

## 🎯 Arquitectura

```
Cliente_RMI
    │
    └──[RMI]──► ServidorDeCalculoPreferencias
                    │
                    ├──[gRPC]──► ServidorDeCanciones
                    │
                    └──[REST]──► ServidorDeReproducciones
```

---

## 📄 Licencia

Proyecto académico - Universidad del Cauca  
Laboratorio de Sistemas Distribuidos

---

**¡Proyecto completo, corregido y listo para usar!** 🚀

---

## 🔄 Changelog

### v1.0 (Corregido)
- ✅ Agregada dependencia de Lombok al pom.xml
- ✅ Creado Main.java completo
- ✅ Creado Menu.java con lógica RMI completa
- ✅ Creados todos los DTOs con anotaciones de Lombok
- ✅ Agregado PreferenciaIdiomaDTORespuesta para cumplir requerimiento
- ✅ Corregidos getters en Menu.java (getNombreGenero, getNumeroPreferencias)
- ✅ Eliminadas llamadas incorrectas a UtilidadesConsola con parámetros

---

Fecha: 2 de noviembre de 2025
