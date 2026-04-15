package db

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func Init() {
	var err error
	DB, err = sql.Open("sqlite", "file:series.db")
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	createTable()
	log.Println("Database connected and ready")
}

func createTable() {
	seriesQuery := `
	CREATE TABLE IF NOT EXISTS series (
		id               INTEGER PRIMARY KEY AUTOINCREMENT,
		titulo           TEXT NOT NULL,
		episodio_actual  INTEGER NOT NULL DEFAULT 0,
		total_episodios  INTEGER NOT NULL DEFAULT 0,
		estado           TEXT NOT NULL DEFAULT 'pendiente',
		calificacion     REAL DEFAULT NULL,
		imagen           TEXT DEFAULT NULL
	);`

	ratingsQuery := `
	CREATE TABLE IF NOT EXISTS ratings (
		id         INTEGER PRIMARY KEY AUTOINCREMENT,
		serie_id   INTEGER NOT NULL,
		puntuacion REAL NOT NULL,
		comentario TEXT DEFAULT NULL,
		created_at TEXT NOT NULL DEFAULT (datetime('now')),
		FOREIGN KEY (serie_id) REFERENCES series(id) ON DELETE CASCADE
	);`

	_, err := DB.Exec(seriesQuery)
	if err != nil {
		log.Fatalf("Error creating series table: %v", err)
	}

	_, err = DB.Exec(ratingsQuery)
	if err != nil {
		log.Fatalf("Error creating ratings table: %v", err)
	}
}
