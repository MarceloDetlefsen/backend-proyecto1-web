package models

type Rating struct {
	ID         int     `json:"id"`
	SerieID    int     `json:"serie_id"`
	Puntuacion float64 `json:"puntuacion"`
	Comentario *string `json:"comentario"`
	CreatedAt  string  `json:"created_at"`
}

type RatingSummary struct {
	SerieID  int      `json:"serie_id"`
	Promedio float64  `json:"promedio"`
	Total    int      `json:"total"`
	Ratings  []Rating `json:"ratings"`
}
