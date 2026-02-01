package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

func enableCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "GET,POST,DELETE,OPTIONS")
}


type Note struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Heading   string `json:"heading"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}

func main() {
	db := connectDB()
	defer db.Close()

	http.HandleFunc("/api/health", healthHandler)
	http.HandleFunc("/api/notes", notesHandler(db))
	http.HandleFunc("/api/notes/", deleteHandler(db))

	log.Println("Go API running on :5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}

func connectDB() *sql.DB {
	connStr := "host=postgres user=" + os.Getenv("POSTGRES_USER") +
		" password=" + os.Getenv("POSTGRES_PASSWORD") +
		" dbname=" + os.Getenv("POSTGRES_DB") +
		" sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)
	if r.Method == "OPTIONS" {
		return
	}
	w.Write([]byte(`{"status":"ok"}`))
}


func notesHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		enableCORS(w)
		if r.Method == "OPTIONS" {
			return
		}

		if r.Method == "GET" {
			rows, _ := db.Query("SELECT id,name,heading,content,created_at FROM notes ORDER BY id DESC")
			var notes []Note

			for rows.Next() {
				var n Note
				rows.Scan(&n.ID, &n.Name, &n.Heading, &n.Content, &n.CreatedAt)
				notes = append(notes, n)
			}

			json.NewEncoder(w).Encode(notes)
		}

		if r.Method == "POST" {
			var n Note
			json.NewDecoder(r.Body).Decode(&n)

			db.Exec(
				"INSERT INTO notes(name,heading,content) VALUES($1,$2,$3)",
				n.Name, n.Heading, n.Content,
			)

			w.WriteHeader(http.StatusCreated)
		}
	}
}

func deleteHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		enableCORS(w)
		if r.Method == "OPTIONS" {
			return
		}

		idStr := strings.TrimPrefix(r.URL.Path, "/api/notes/")
		id, _ := strconv.Atoi(idStr)

		if r.Method == "DELETE" {
			db.Exec("DELETE FROM notes WHERE id=$1", id)
			w.WriteHeader(http.StatusOK)
		}
	}
}
