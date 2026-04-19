package models

type Serie struct {
	ID             int      `json:"id"`
	Titulo         string   `json:"titulo"`
	EpisodioActual int      `json:"episodio_actual"`
	TotalEpisodios int      `json:"total_episodios"`
	Estado         string   `json:"estado"`
	Calificacion   *float64 `json:"calificacion"`
	Imagen         *string  `json:"imagen"`
	Descripcion    *string  `json:"descripcion"`
}
