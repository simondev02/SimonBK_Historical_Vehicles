# Usar la imagen oficial de Go 1.20 como base
FROM golang:1.20

# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar el go.mod y go.sum al directorio de trabajo actual
COPY go.mod go.sum ./

# Descargar todas las dependencias
RUN go mod download

# Copiar el código fuente al directorio de trabajo
COPY . .

# Compilar la aplicación
RUN go build -o SimonBK_vehiculos .

# Exponer el puerto 8003
EXPOSE 60030

# Comando para ejecutar la aplicación
CMD ["./SimonBK_vehiculos"]
