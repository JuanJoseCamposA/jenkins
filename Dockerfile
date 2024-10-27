FROM golang:latest

RUN mkdir /build
WORKDIR /build

# Clonar el repositorio directamente, no es necesario usar 'go get'
RUN git clone https://github.com/Daniel202412/API-REST.git

# Cambiar al directorio del proyecto e inicializar los módulos si no existen
WORKDIR /build/API-REST

# Construir el proyecto
RUN go mod tidy && go build -o main

# Definir el punto de entrada
ENTRYPOINT ["/build/API-REST/main"]

