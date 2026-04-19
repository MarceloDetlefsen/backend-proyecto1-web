# Series Tracker - Backend 🎬

Backend REST API hecho en Go desde cero (sin frameworks) con SQLite como base de datos.

## Cómo correr el proyecto

### Requisitos
- Go 1.22+

### Instalación
```bash
git clone https://github.com/MarceloDetlefsen/backend-proyecto1-web.git
cd backend-proyecto1-web
go run main.go
```

El servidor corre en `http://localhost:8080`

### Estructura del proyecto
```
.
├── main.go              # Servidor HTTP, CORS y registro de rutas
├── db/
│   └── db.go            # Conexión a SQLite y creación de tablas
├── models/
│   ├── series.go        # Struct Serie
│   └── rating.go        # Structs Rating y RatingSummary
├── repository/
│   ├── series.go        # Queries de series a la DB
│   └── ratings.go       # Queries de ratings a la DB
├── handlers/
│   ├── series.go        # Handlers HTTP de series
│   └── ratings.go       # Handlers HTTP de ratings
├── series.db            # Base de datos SQLite (se crea automáticamente)
└── README.md
```

## Endpoints

### Series
| Método | Ruta                              | Descripción                        |
|--------|-----------------------------------|------------------------------------|
| GET    | `/series`                         | Listar todas las series            |
| GET    | `/series/{id}`                    | Obtener una serie por ID           |
| POST   | `/series`                         | Crear una serie nueva              |
| PUT    | `/series/{id}`                    | Editar una serie existente         |
| DELETE | `/series/{id}`                    | Eliminar una serie                 |
| PATCH  | `/series/{id}/episodio/incrementar` | Sumar +1 al episodio actual      |
| PATCH  | `/series/{id}/episodio/decrementar` | Restar -1 al episodio actual     |

### Ratings
| Método | Ruta                    | Descripción                                      |
|--------|-------------------------|--------------------------------------------------|
| POST   | `/series/{id}/ratings`  | Agregar un rating a una serie (puntuación 1–10)  |
| GET    | `/series/{id}/ratings`  | Obtener ratings y promedio de una serie          |
| DELETE | `/ratings/{id}`         | Eliminar un rating por ID                        |

### Query params disponibles en GET /series
| Parámetro | Ejemplo              | Descripción                              |
|-----------|----------------------|------------------------------------------|
| `q`       | `?q=breaking`        | Buscar series por nombre                 |
| `sort`    | `?sort=calificacion` | Ordenar por columna                      |
| `order`   | `?order=desc`        | Dirección del orden (`asc` o `desc`)     |
| `page`    | `?page=2`            | Página actual (default: 1)               |
| `limit`   | `?limit=5`           | Resultados por página (default: 10)      |

## Challenges implementados

| Challenge | Puntos |
|-----------|--------|
| Códigos HTTP correctos (201, 204, 404, 400) | 20 |
| Validación server-side con errores en JSON | 20 |
| Paginación con `?page=` y `?limit=` | 30 |
| Búsqueda por nombre con `?q=` | 15 |
| Ordenamiento con `?sort=` y `?order=` | 15 |
| Sistema de ratings con tabla propia en DB | 30 |

**Total: 130 puntos**


## Detalles técnicos

- El servidor usa únicamente **`net/http`** de la librería estándar de Go, sin frameworks externos
- La base de datos es **SQLite** usando `modernc.org/sqlite`, un driver pure Go que no requiere CGo ni GCC
- La tabla `ratings` tiene una **foreign key** hacia `series` con `ON DELETE CASCADE`, lo que elimina automáticamente los ratings al borrar una serie
- El campo `imagen` en series almacena una **URL o path** a la imagen, no el binario
- La función `GetAllSeries` construye la query dinámicamente con **protección contra SQL injection** usando `?` como placeholders
- Los endpoints de incrementar/decrementar episodios usan `MIN` y `MAX` directamente en SQL para evitar que el episodio se salga del rango válido sin validación extra en Go


## 👨‍💻 Autor
Marcelo Detlefsen - 24554

## 🔗 Links
- [Repositorio Frontend](https://github.com/MarceloDetlefsen/frontend-proyecto1-web)