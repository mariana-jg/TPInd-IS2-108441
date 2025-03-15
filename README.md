# TPInd-IS2-108441

## Alumna: Mariana Noelia Juarez Goldemberg
## Padrón: 108441

### Contenidos
### Introducción
### Desafíos
### Pre-requisitos
### User-guides
### Cómo construir la imagen de Docker.

  ```docker build -t [NOMBRE] .```
  
  ```docker run --rm   -p 8080:8080   -e ENVIRONMENT=production   -e HOST=0.0.0.0   -e PORT=8080   -v $(pwd)/logs:/root/logs   [NOMBRE]```
  Con docker compose

  ```docker-compose up --build -d```
  ```docker-compose down -v```
  ```docker logs -f api-container```

### Cómo correr la base de datos.
### Como correr la imagen del servicio.

### Fecha máxima de entrega: 20/03/2025
