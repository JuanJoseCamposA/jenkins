# Usa una imagen base de Go
FROM golang:1.20

# Instala sqlite3 y limpia archivos temporales
RUN apt-get update && \
    apt-get install -y sqlite3 && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

# Establece el directorio de trabajo en el contenedor
WORKDIR /home/juanjose/Api

# Copia el archivo principal de tu aplicación en el contenedor
COPY mainprueba.go .

# Crea un directorio para la base de datos
RUN mkdir -p data

# Copia la base de datos en el contenedor
COPY data/actividades.db ./data/

# Inicializa el módulo de Go
RUN go mod init miapi && go mod tidy

# Compila la aplicación
RUN go build -o miapi mainprueba.go

# Expone el puerto en el que corre la API
EXPOSE 8080

# Comando para ejecutar la API
CMD ["./miapi"]



