# TPInd-IS2-108441

## Alumna: Mariana Noelia Juarez Goldemberg
## Padrón: 108441

### Contenidos

1. [Introducción](#introducción)
2. [Desafíos](#desafíos)
3. [Pre-requisitos](#pre-requisitos)
4. [Cómo levantar el servicio y la base de datos utilizando Docker Compose](#cómo-levantar-el-servicio-y-la-base-de-datos-utilizando-docker-compose)
5. [User-guides](#user-guides)
   1. [Cómo probar la API](#cómo-probar-la-api-con-el-servicio-activo)
   2. [Cómo correr los tests End to End](#cómo-correr-los-tests-end-to-end)

### Introducción

Este trabajo práctico implementa una API RESTful para la gestión de cursos utilizando endpoints HTTP que permiten crear, obtener, listar y eliminar cursos en particular.

El desarrollo backend se realizó utilizando el lenguaje Go con el framework Gin para la gestión del router y la administración de las requests y responses HTTP. Además, para la persistencia de la información de los cursos, se utilizó PostgreSQL como base de datos. Además, cuenta con el feature de loggear registrando en un archivo.

El servicio completo (API + DB) se levanta utilizando el orquestador Docker Compose.

La arquitectura de la API sigue el patrón Repository-Service-Controller:
- Repository: capa de acceso a los datos, interactúa directamente con la DB.
- Service: capa de business logic, valida y aplica reglas antes de llamar al repository. Aún no muy desarrollada en esta etapa temprana de la API, pero la división de este módulo puede ser muy útil si la aplicación escala.
- Controller: capa de gestión de router y conexión con HTTP.

### Desafíos

Definitivamente el despliegue de la base de datos y el uso de Docker Compose fue lo más desafiante, por mi parte fue la primera vez donde me encargué de añadir esas features a un proyecto y considero que, a pesar de que me resultó un poco dificil al inicio, logré el resultado esperado. Pasar de un manejo en memoria de la información al servicio de base de datos resultó un desafío de refactorización casi completa del código, más que nada del repositorio. 

### Pre-requisitos

*Importante*: no hay archivo .env en el repositorio, el usuario debe tener configuración de variables de entorno en un archivo .env. Un ejemplo del archivo es:

```
DATABASE_USER=myuser
DATABASE_PASSWORD=mypassword
DATABASE_HOST=db
DATABASE_PORT=5432
DATABASE_NAME=apirest_is2
ENVIRONMENT=development
HOST=0.0.0.0
PORT=8080
```

### Cómo levantar el servicio y la base de datos utilizando Docker Compose.
  ```docker-compose up --build -d```

Para detener los containers y borrar el volumen
 
  ```docker-compose down -v```

Para ver los logs en ejecución
  
  ```docker logs -f api-container```

### User-guides

#### Cómo probar la API (con el servicio activo)

Reemplazar ```PORT``` por el puerto de despliegue asignado en el archivo de variables de entorno.

Para crear un curso:

```
curl -X POST http://localhost:PORT/courses \
  -H "Content-Type: application/json" \
  -d '{"title": "Fundamentals of Software Engineer", "description": "Learn the fundamentals of software engineer and build your own API using go"}'
```

Para obtener todos los cursos:

```
curl -X GET http://localhost:PORT/courses
```

Para obtener un curso en particular:

```
curl -X GET http://localhost:PORT/courses/id
```

Para eliminar un curso:

```
curl -X DELETE http://localhost:PORT/courses/id
```

#### Cómo correr los tests End to End

1. Hay que levantar el servicio primero con  ```docker-compose up --build -d```
2. Ejecutar en la terminal ``` go test -v ```

#### Fecha máxima de entrega: 20/03/2025
