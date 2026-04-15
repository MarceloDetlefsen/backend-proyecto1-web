package handlers

import (
	"backend-proyecto1-web/models"
	"backend-proyecto1-web/repository"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"error": message})
}

func parseID(r *http.Request) (int, error) {
	parts := strings.Split(r.URL.Path, "/")
	return strconv.Atoi(parts[len(parts)-1])
}

// GET /series
func GetAllSeries(w http.ResponseWriter, r *http.Request) {
	series, err := repository.GetAllSeries()
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Error obteniendo series")
		return
	}

	if series == nil {
		series = []models.Serie{}
	}

	writeJSON(w, http.StatusOK, series)
}

// GET /series/{id}
func GetSerieByID(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, "ID inválido")
		return
	}

	serie, err := repository.GetSerieByID(id)
	if err == sql.ErrNoRows {
		writeError(w, http.StatusNotFound, "Serie no encontrada")
		return
	}
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Error obteniendo serie")
		return
	}

	writeJSON(w, http.StatusOK, serie)
}

// POST /series
func CreateSerie(w http.ResponseWriter, r *http.Request) {
	var s models.Serie
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		writeError(w, http.StatusBadRequest, "Body inválido")
		return
	}

	if strings.TrimSpace(s.Titulo) == "" {
		writeError(w, http.StatusBadRequest, "El título es requerido")
		return
	}

	estados := map[string]bool{"pendiente": true, "viendo": true, "completada": true}
	if !estados[s.Estado] {
		writeError(w, http.StatusBadRequest, "Estado inválido. Use: pendiente, viendo, completada")
		return
	}

	if s.TotalEpisodios < 0 || s.EpisodioActual < 0 {
		writeError(w, http.StatusBadRequest, "Los episodios no pueden ser negativos")
		return
	}

	if s.EpisodioActual > s.TotalEpisodios {
		writeError(w, http.StatusBadRequest, "El episodio actual no puede superar el total")
		return
	}

	nueva, err := repository.CreateSerie(s)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Error creando serie")
		return
	}

	writeJSON(w, http.StatusCreated, nueva)
}

// PUT /series/{id}
func UpdateSerie(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, "ID inválido")
		return
	}

	_, err = repository.GetSerieByID(id)
	if err == sql.ErrNoRows {
		writeError(w, http.StatusNotFound, "Serie no encontrada")
		return
	}

	var s models.Serie
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		writeError(w, http.StatusBadRequest, "Body inválido")
		return
	}

	if strings.TrimSpace(s.Titulo) == "" {
		writeError(w, http.StatusBadRequest, "El título es requerido")
		return
	}

	estados := map[string]bool{"pendiente": true, "viendo": true, "completada": true}
	if !estados[s.Estado] {
		writeError(w, http.StatusBadRequest, "Estado inválido. Use: pendiente, viendo, completada")
		return
	}

	if s.TotalEpisodios < 0 || s.EpisodioActual < 0 {
		writeError(w, http.StatusBadRequest, "Los episodios no pueden ser negativos")
		return
	}

	if s.EpisodioActual > s.TotalEpisodios {
		writeError(w, http.StatusBadRequest, "El episodio actual no puede superar el total")
		return
	}

	actualizada, err := repository.UpdateSerie(id, s)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Error actualizando serie")
		return
	}

	writeJSON(w, http.StatusOK, actualizada)
}

// DELETE /series/{id}
func DeleteSerie(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, "ID inválido")
		return
	}

	_, err = repository.GetSerieByID(id)
	if err == sql.ErrNoRows {
		writeError(w, http.StatusNotFound, "Serie no encontrada")
		return
	}

	if err := repository.DeleteSerie(id); err != nil {
		writeError(w, http.StatusInternalServerError, "Error eliminando serie")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// PATCH /series/{id}/episodio/incrementar
func IncrementarEpisodio(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	id, err := strconv.Atoi(parts[2])
	if err != nil {
		writeError(w, http.StatusBadRequest, "ID inválido")
		return
	}

	_, err = repository.GetSerieByID(id)
	if err == sql.ErrNoRows {
		writeError(w, http.StatusNotFound, "Serie no encontrada")
		return
	}

	serie, err := repository.IncrementarEpisodio(id)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Error incrementando episodio")
		return
	}

	writeJSON(w, http.StatusOK, serie)
}

// PATCH /series/{id}/episodio/decrementar
func DecrementarEpisodio(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	id, err := strconv.Atoi(parts[2])
	if err != nil {
		writeError(w, http.StatusBadRequest, "ID inválido")
		return
	}

	_, err = repository.GetSerieByID(id)
	if err == sql.ErrNoRows {
		writeError(w, http.StatusNotFound, "Serie no encontrada")
		return
	}

	serie, err := repository.DecrementarEpisodio(id)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Error decrementando episodio")
		return
	}

	writeJSON(w, http.StatusOK, serie)
}
