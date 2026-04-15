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

// POST /series/{id}/ratings
func CreateRating(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	serieID, err := strconv.Atoi(parts[2])
	if err != nil {
		writeError(w, http.StatusBadRequest, "ID inválido")
		return
	}

	_, err = repository.GetSerieByID(serieID)
	if err == sql.ErrNoRows {
		writeError(w, http.StatusNotFound, "Serie no encontrada")
		return
	}
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Error verificando serie")
		return
	}

	var rating models.Rating
	if err := json.NewDecoder(r.Body).Decode(&rating); err != nil {
		writeError(w, http.StatusBadRequest, "Body inválido")
		return
	}

	if rating.Puntuacion < 1 || rating.Puntuacion > 10 {
		writeError(w, http.StatusBadRequest, "La puntuación debe estar entre 1 y 10")
		return
	}

	created, err := repository.CreateRating(serieID, rating)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Error creando rating")
		return
	}

	writeJSON(w, http.StatusCreated, created)
}

// GET /series/{id}/ratings
func GetRatings(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	serieID, err := strconv.Atoi(parts[2])
	if err != nil {
		writeError(w, http.StatusBadRequest, "ID inválido")
		return
	}

	_, err = repository.GetSerieByID(serieID)
	if err == sql.ErrNoRows {
		writeError(w, http.StatusNotFound, "Serie no encontrada")
		return
	}
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Error verificando serie")
		return
	}

	summary, err := repository.GetRatingsBySerie(serieID)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Error obteniendo ratings")
		return
	}

	writeJSON(w, http.StatusOK, summary)
}

// DELETE /ratings/{id}
func DeleteRating(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	id, err := strconv.Atoi(parts[len(parts)-1])
	if err != nil {
		writeError(w, http.StatusBadRequest, "ID inválido")
		return
	}

	if err := repository.DeleteRating(id); err != nil {
		writeError(w, http.StatusInternalServerError, "Error eliminando rating")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
