package repository

import (
	"backend-proyecto1-web/db"
	"backend-proyecto1-web/models"
	"strconv"
)

func GetAllSeries(params map[string]string) ([]models.Serie, int, error) {
	query := `SELECT id, titulo, episodio_actual, total_episodios, estado, calificacion, imagen FROM series WHERE 1=1`
	countQuery := `SELECT COUNT(*) FROM series WHERE 1=1`
	args := []any{}

	// Búsqueda por nombre
	if q, ok := params["q"]; ok && q != "" {
		filter := " AND titulo LIKE ?"
		query += filter
		countQuery += filter
		args = append(args, "%"+q+"%")
	}

	// Ordenamiento
	validColumns := map[string]bool{
		"id": true, "titulo": true, "episodio_actual": true,
		"total_episodios": true, "estado": true, "calificacion": true,
	}
	sort := "id"
	if s, ok := params["sort"]; ok {
		if s == "progreso" {
			sort = "CAST(episodio_actual AS REAL) / NULLIF(total_episodios, 0)"
		} else if validColumns[s] {
			sort = s
		}
	}
	order := "asc"
	if o, ok := params["order"]; ok && (o == "asc" || o == "desc") {
		order = o
	}
	query += " ORDER BY " + sort + " " + order

	// Total de registros para paginación
	var total int
	err := db.DB.QueryRow(countQuery, args...).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	// Paginación
	limit := 10
	page := 1
	if l, ok := params["limit"]; ok && l != "" {
		if parsed, err := strconv.Atoi(l); err == nil && parsed > 0 {
			limit = parsed
		}
	}
	if p, ok := params["page"]; ok && p != "" {
		if parsed, err := strconv.Atoi(p); err == nil && parsed > 0 {
			page = parsed
		}
	}
	offset := (page - 1) * limit
	query += " LIMIT ? OFFSET ?"
	args = append(args, limit, offset)

	rows, err := db.DB.Query(query, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var series []models.Serie
	for rows.Next() {
		var s models.Serie
		err := rows.Scan(&s.ID, &s.Titulo, &s.EpisodioActual, &s.TotalEpisodios, &s.Estado, &s.Calificacion, &s.Imagen)
		if err != nil {
			return nil, 0, err
		}
		series = append(series, s)
	}

	return series, total, nil
}

func GetSerieByID(id int) (*models.Serie, error) {
	var s models.Serie
	err := db.DB.QueryRow(`SELECT id, titulo, episodio_actual, total_episodios, estado, calificacion, imagen, descripcion FROM series WHERE id = ?`, id).
		Scan(&s.ID, &s.Titulo, &s.EpisodioActual, &s.TotalEpisodios, &s.Estado, &s.Calificacion, &s.Imagen, &s.Descripcion)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func CreateSerie(s models.Serie) (*models.Serie, error) {
	result, err := db.DB.Exec(
		`INSERT INTO series (titulo, episodio_actual, total_episodios, estado, calificacion, imagen, descripcion) VALUES (?, ?, ?, ?, ?, ?, ?)`,
		s.Titulo, s.EpisodioActual, s.TotalEpisodios, s.Estado, s.Calificacion, s.Imagen, s.Descripcion,
	)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return GetSerieByID(int(id))
}

func UpdateSerie(id int, s models.Serie) (*models.Serie, error) {
	_, err := db.DB.Exec(
		`UPDATE series SET titulo = ?, episodio_actual = ?, total_episodios = ?, estado = ?, calificacion = ?, imagen = ?, descripcion = ? WHERE id = ?`,
		s.Titulo, s.EpisodioActual, s.TotalEpisodios, s.Estado, s.Calificacion, s.Imagen, s.Descripcion, id,
	)
	if err != nil {
		return nil, err
	}

	return GetSerieByID(id)
}

func DeleteSerie(id int) error {
	_, err := db.DB.Exec(`DELETE FROM series WHERE id = ?`, id)
	return err
}

func IncrementarEpisodio(id int) (*models.Serie, error) {
	_, err := db.DB.Exec(
		`UPDATE series SET episodio_actual = MIN(episodio_actual + 1, total_episodios) WHERE id = ?`,
		id,
	)
	if err != nil {
		return nil, err
	}
	return GetSerieByID(id)
}

func DecrementarEpisodio(id int) (*models.Serie, error) {
	_, err := db.DB.Exec(
		`UPDATE series SET episodio_actual = MAX(episodio_actual - 1, 0) WHERE id = ?`,
		id,
	)
	if err != nil {
		return nil, err
	}
	return GetSerieByID(id)
}
