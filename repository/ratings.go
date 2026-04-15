package repository

import (
	"backend-proyecto1-web/db"
	"backend-proyecto1-web/models"
)

func CreateRating(serieID int, r models.Rating) (*models.Rating, error) {
	result, err := db.DB.Exec(
		`INSERT INTO ratings (serie_id, puntuacion, comentario) VALUES (?, ?, ?)`,
		serieID, r.Puntuacion, r.Comentario,
	)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	var created models.Rating
	err = db.DB.QueryRow(
		`SELECT id, serie_id, puntuacion, comentario, created_at FROM ratings WHERE id = ?`, id,
	).Scan(&created.ID, &created.SerieID, &created.Puntuacion, &created.Comentario, &created.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &created, nil
}

func GetRatingsBySerie(serieID int) (*models.RatingSummary, error) {
	rows, err := db.DB.Query(
		`SELECT id, serie_id, puntuacion, comentario, created_at FROM ratings WHERE serie_id = ? ORDER BY created_at DESC`,
		serieID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ratings []models.Rating
	for rows.Next() {
		var r models.Rating
		err := rows.Scan(&r.ID, &r.SerieID, &r.Puntuacion, &r.Comentario, &r.CreatedAt)
		if err != nil {
			return nil, err
		}
		ratings = append(ratings, r)
	}

	if ratings == nil {
		ratings = []models.Rating{}
	}

	// Calcular promedio
	var promedio float64
	if len(ratings) > 0 {
		var suma float64
		for _, r := range ratings {
			suma += r.Puntuacion
		}
		promedio = suma / float64(len(ratings))
	}

	return &models.RatingSummary{
		SerieID:  serieID,
		Promedio: promedio,
		Total:    len(ratings),
		Ratings:  ratings,
	}, nil
}

func DeleteRating(id int) error {
	_, err := db.DB.Exec(`DELETE FROM ratings WHERE id = ?`, id)
	return err
}
