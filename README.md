# 🚀 Read File S3 - Procesador de Archivos Concurrente

[![Go Version](https://img.shields.io/badge/Go-1.24+-00ADD8?style=for-the-badge&logo=go)](https://golang.org/)
[![AWS S3](https://img.shields.io/badge/AWS-S3-FF9900?style=for-the-badge&logo=amazon-aws)](https://aws.amazon.com/s3/)
[![License](https://img.shields.io/badge/License-MIT-green?style=for-the-badge)](LICENSE)

> **Procesador de archivos ultra-rápido que lee archivos desde AWS S3 y los procesa usando concurrencia en Go**

## 📋 Tabla de Contenidos

- [🎯 Descripción](#-descripción)
- [✨ Características](#-características)
- [🛠️ Prerrequisitos](#️-prerrequisitos)
- [⚡ Instalación](#-instalación)
- [🔧 Configuración](#-configuración)
- [🚀 Uso](#-uso)
- [📊 Rendimiento](#-rendimiento)
- [🏗️ Arquitectura](#️-arquitectura)
- [🧪 Testing](#-testing)
- [🤝 Contribuir](#-contribuir)
- [📄 Licencia](#-licencia)

## 🎯 Descripción

Este proyecto demuestra cómo procesar archivos masivos de forma eficiente usando:

- **Go (Golang)** para concurrencia nativa
- **AWS S3 SDK v2** para acceso directo a archivos en la nube
- **Goroutines** para procesamiento paralelo
- **Streaming** para optimización de memoria

El caso de uso específico es **contar líneas que contienen la palabra "error"** en archivos CSV grandes, pero la arquitectura es extensible para cualquier tipo de procesamiento de texto.

## ✨ Características

- 🚀 **Ultra-rápido**: Procesa millones de líneas en segundos
- 🔄 **Concurrente**: Usa goroutines para procesamiento paralelo
- 💾 **Eficiente en memoria**: Streaming sin cargar todo el archivo
- ☁️ **Nativo en la nube**: Lectura directa desde AWS S3
- 🔒 **Thread-safe**: Previene race conditions con sync.Mutex
- 📊 **Métricas**: Tiempo de ejecución y contadores en tiempo real
- 🛡️ **Robusto**: Manejo de errores y cleanup automático

## 🛠️ Prerrequisitos

### Software Requerido

- **Go 1.24+** - [Descargar aquí](https://golang.org/dl/)
- **AWS CLI** (opcional, para configuración) - [Guía de instalación](https://aws.amazon.com/cli/)

### Servicios de AWS

- **Cuenta AWS activa**
- **Bucket S3** con el archivo a procesar
- **Credenciales AWS** configuradas (ver [Configuración](#-configuración))

### Permisos IAM Requeridos

```json
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
                "s3:GetObject"
            ],
            "Resource": "arn:aws:s3:::tu-bucket/*"
        }
    ]
}
```

## ⚡ Instalación

### 1. Clonar el repositorio

```bash
git clone https://github.com/WilsonSayago/read-file-s3.git
cd read-file-s3
```

### 2. Inicializar módulo Go

```bash
go mod init read-file-s3
go mod tidy
```

### 3. Instalar dependencias

```bash
go get github.com/aws/aws-sdk-go-v2/aws
go get github.com/aws/aws-sdk-go-v2/config
go get github.com/aws/aws-sdk-go-v2/service/s3
```

## 🔧 Configuración

### Opción 1: Variables de Entorno

```bash
export AWS_ACCESS_KEY_ID="tu_access_key"
export AWS_SECRET_ACCESS_KEY="tu_secret_key"
export AWS_DEFAULT_REGION="us-east-1"
```

### Opción 2: AWS CLI

```bash
aws configure
```

### Opción 3: IAM Roles (Recomendado para EC2/ECS)

Si ejecutas en AWS, usa IAM Roles en lugar de credenciales hardcodeadas.

### Configuración del Proyecto

Edita las variables en `main.go`:

```go
bucket := "tu-bucket-name"        // 🔄 Cambia por tu bucket
key := "tu-archivo.csv"           // 🔄 Cambia por tu archivo
```

## 🚀 Uso

### Ejecución Básica

```bash
go run main.go
```

### Compilar y Ejecutar

```bash
# Compilar
go build -o read-file-s3 main.go

# Ejecutar
./read-file-s3
```

### Ejemplo de Salida


## 🔧 Troubleshooting

### Errores Comunes

#### 1. Error de Credenciales AWS

```
Error: NoCredentialProviders: no valid providers in chain
```

**Solución**: Configura tus credenciales AWS (ver [Configuración](#-configuración))

#### 2. Bucket No Encontrado

```
Error: NoSuchBucket: The specified bucket does not exist
```

**Solución**: Verifica el nombre del bucket y región

#### 3. Archivo No Encontrado

```
Error: NoSuchKey: The specified key does not exist
```

**Solución**: Verifica que el archivo existe en el bucket

#### 4. Permisos Insuficientes

```
Error: AccessDenied: Access Denied
```

**Solución**: Verifica los permisos IAM (ver [Prerrequisitos](#-prerrequisitos))

### Optimizaciones

#### Para Archivos Muy Grandes (>10GB)

```go
// Limitar número de goroutines
const maxGoroutines = 1000
semaphore := make(chan struct{}, maxGoroutines)

for scanner.Scan() {
    semaphore <- struct{}{}  // Adquirir
    wg.Add(1)
    go func(line string) {
        defer func() { <-semaphore }()  // Liberar
        processLine(line)
    }(scanner.Text())
}
```

## 🤝 Contribuir

¡Las contribuciones son bienvenidas! Por favor:

1. Fork el proyecto
2. Crea una branch para tu feature (`git checkout -b feature/AmazingFeature`)
3. Commit tus cambios (`git commit -m 'Add some AmazingFeature'`)
4. Push a la branch (`git push origin feature/AmazingFeature`)
5. Abre un Pull Request

### Guías de Contribución

- Sigue las convenciones de Go (`go fmt`, `go vet`)
- Agrega tests para nuevas funcionalidades
- Actualiza la documentación
- Usa commits descriptivos

## 📄 Licencia

Este proyecto está bajo la Licencia MIT. Ver el archivo [LICENSE](LICENSE) para más detalles.

