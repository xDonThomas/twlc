# TWLC - T Writer Logger Console (Go)

`twlc` es un paquete de logging flexible y colorido para aplicaciones en Go, que permite mostrar mensajes en consola y/o guardarlos en archivos, con soporte para colores, etiquetas por tipo de mensaje, marcas de tiempo, y conversión de estructuras a texto o JSON.

## 🚀 Características

- ✅ Soporte para múltiples niveles de log: `INFO`, `SUCCESS`, `WARNING`, `ERROR`, `DEBUG`, `TRACE`.
- 🎨 Colores configurables para mensajes y etiquetas (`FGColor`, `BGColor`).
- 🕒 Opción para incluir marcas de tiempo.
- 🖨 Mostrar mensajes en consola (`ShowInConsole`).
- 📂 Guardar logs en archivos separados por día (`SaveInLogFile`).
- 🧾 Conversión de structs a string o JSON con identación.
- 📁 Crea automáticamente el directorio y archivos de log.

## 📦 Instalación

```bash
go get github.com/xxxAlvaDevxxx/twlc
```

## 🧑‍💻 Uso

### Importar

``` go
import "github.com/xxxAlvaDevxxx/twlc"
```

### Inicialización básica

``` go
logger := twlc.DefaultTwlc()
```

#### 📝 Descripción
La función `DefaultTwlc()` crea una instancia preconfigurada del logger Twlc con los valores más comunes ya activados. Está pensada para que puedas empezar rápidamente sin preocuparte por la configuración inicial.

#### ⚙️ Configuración por defecto

| **Parámetro**     | **Valor** | **Descripción**                                                                 |
|-------------------|-----------|----------------------------------------------------------------------------------|
| `SaveInLogFile`   | `true`    | Guarda los logs en archivos diarios (`./logs/twlc_YYYYMMDD.log`).               |
| `ShowInConsole`   | `true`    | Muestra los logs en la consola (terminal).                                      |
| `ColorMessages`   | `true`    | Activa los colores para los mensajes.                                           |
| `BGColor`         | `true`    | Colorea el fondo del tipo de mensaje (`[INFO]`, `[ERROR]`, etc.).               |
| `FGColor`         | `true`    | Colorea el contenido del mensaje.                                               |
| `WithTime`        | `true`    | Incluye la hora y fecha en los mensajes de log.                                 |
| `LogDir`          | `./logs/` | Crea un directorio llamado `logs` al lado del ejecutable.                       |

#### 📁 Ejemplo de ubicación del log

Si tu ejecutable está en:

``` bash
/home/usuario/proyecto/
```

Se generará el archivo:

```bash
/home/usuario/proyecto/logs/twlc_20250724.log
```

#### ✅ ¿Cuándo usar DefaultTwlc()?
* Para pruebas rápidas o proyectos pequeños.
* Cuando quieres una configuración lista para usar.
* Si no necesitas personalizar la ruta de logs o el formato.

### Inicialización personalizada

``` go
logger := twlc.NewTwlc(
    true,   // SaveInLogFile
    true,   // ShowInConsole
    true,   // ColorMessages
    true,   // BGColor
    true,   // FGColor
    true,   // WithTime
    "./logs",
)
```

### Escribir mensajes

``` go
logger.WriteInfo("Información general")
logger.WriteSuccess("Operación exitosa")
logger.WriteWarning("Advertencia importante")
logger.WriteError("Ha ocurrido un error")
logger.WriteDebug("Mensaje de depuración")
logger.WriteTrace("Mensaje de traza")
```

### Convertir structs

``` go
type Animal struct {
    Name string
    Age  int
}

a := Animal{"Dog", 5}

fmt.Println(logger.StructToString(a, true))  // {Name:Dog Age:5}
fmt.Println(logger.StructToString(a, false)) // main.Animal{Name:"Dog", Age:5}

jsonStr, _ := logger.StructToJson(a)
fmt.Println(jsonStr)
```

## 📁 Estructura de logs

* Los logs se guardan automáticamente en el directorio indicado (logDir) con el nombre twlc_YYYYMMDD.log.

Ejemplo:

``` lua
logs/
 └── twlc_20250724.log
```

## 🛠 Configuración de colores

Puedes personalizar si los colores se muestran para:

* Mensaje: `FGColor`
* Tipo de mensaje: `BGColor`
* Ambos: `ColorMessages`

Los colores se muestran en consola si tu terminal soporta ANSI.

## 🧪 Ejemplo completo

``` go
package main

import (
    "github.com/xxxAlvaDevxxx/twlc"
)

func main() {
    logger := twlc.DefaultTwlc()
    logger.WriteInfo("Inicio del programa")
    logger.WriteSuccess("Operación completada")
}
```

## 🧾 Licencia

Este proyecto está licenciado bajo la licencia MIT. Puedes modificarlo y usarlo libremente.
