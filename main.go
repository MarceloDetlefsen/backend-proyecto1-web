package main

import (
	"backend-proyecto1-web/db"
	"backend-proyecto1-web/handlers"
	"log"
	"net/http"
)

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	db.Init()

	mux := http.NewServeMux()

	// Series CRUD
	mux.HandleFunc("GET /series", handlers.GetAllSeries)
	mux.HandleFunc("GET /series/{id}", handlers.GetSerieByID)
	mux.HandleFunc("POST /series", handlers.CreateSerie)
	mux.HandleFunc("PUT /series/{id}", handlers.UpdateSerie)
	mux.HandleFunc("DELETE /series/{id}", handlers.DeleteSerie)

	// Episodios
	mux.HandleFunc("PATCH /series/{id}/episodio/incrementar", handlers.IncrementarEpisodio)
	mux.HandleFunc("PATCH /series/{id}/episodio/decrementar", handlers.DecrementarEpisodio)

	// Ratings
	mux.HandleFunc("POST /series/{id}/ratings", handlers.CreateRating)
	mux.HandleFunc("GET /series/{id}/ratings", handlers.GetRatings)
	mux.HandleFunc("DELETE /ratings/{id}", handlers.DeleteRating)

	log.Println("Server running on http://localhost:8080")

	err := http.ListenAndServe(":8080", enableCORS(mux))
	if err != nil {
		log.Fatal(err)
	}
}
