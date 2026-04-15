package repository

import (
	"backend-proyecto1-web/db"
	"backend-proyecto1-web/models"
)

func GetAllSeries() ([]models.Serie, error) {
	rows, err := db.DB.Query(`SELECT id, titulo, episodio_actual, total_episodios, estado, calificacion, imagen FROM series`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var series []models.Serie
	for rows.Next() {
		var s models.Serie
		err := rows.Scan(&s.ID, &s.Titulo, &s.EpisodioActual, &s.TotalEpisodios, &s.Estado, &s.Calificacion, &s.Imagen)
		if err != nil {
			return nil, err
		}
		series = append(series, s)
	}

	return series, nil
}

func GetSerieByID(id int) (*models.Serie, error) {
	var s models.Serie
	err := db.DB.QueryRow(`SELECT id, titulo, episodio_actual, total_episodios, estado, calificacion, imagen FROM series WHERE id = ?`, id).
		Scan(&s.ID, &s.Titulo, &s.EpisodioActual, &s.TotalEpisodios, &s.Estado, &s.Calificacion, &s.Imagen)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func CreateSerie(s models.Serie) (*models.Serie, error) {
	result, err := db.DB.Exec(
		`INSERT INTO series (titulo, episodio_actual, total_episodios, estado, calificacion, imagen) VALUES (?, ?, ?, ?, ?, ?)`,
		s.Titulo, s.EpisodioActual, s.TotalEpisodios, s.Estado, s.Calificacion, s.Imagen,
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
		`UPDATE series SET titulo = ?, episodio_actual = ?, total_episodios = ?, estado = ?, calificacion = ?, imagen = ? WHERE id = ?`,
		s.Titulo, s.EpisodioActual, s.TotalEpisodios, s.Estado, s.Calificacion, s.Imagen, id,
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
