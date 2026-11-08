package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
		log.Fatal("Could not load .env:", err)
	}

	db, err := openDatabase()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	log.Printf("Connected to PostgreSQL")

	app := application{db: db}

	http.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	http.HandleFunc("POST /records", app.createRecord)
	http.HandleFunc("GET /records", app.getRecords)

	log.Println("Server listening on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Println("Server stopped:", err)
	}
}
