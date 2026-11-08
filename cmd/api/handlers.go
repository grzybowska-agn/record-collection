package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

type application struct {
	db *sql.DB
}

func (app *application) createRecord(w http.ResponseWriter, r *http.Request) {
	var record Record

	err := json.NewDecoder(r.Body).Decode(&record)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	err = app.insertRecord(r.Context(), &record)
	if err != nil {
		log.Println("Failed to insert record:", err)
		http.Error(w, "Could not create record", http.StatusInternalServerError)
		return
	}

	w.Header().Set("ContentType", "application/json")
	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(record)
	if err != nil {
		log.Println("Failed to write response:", err)
	}

}

func (app *application) getRecords(w http.ResponseWriter, r *http.Request) {
	collection, err := app.listRecords(r.Context())
	if err != nil {
		log.Println("Failed to list records:", err)
		http.Error(w, "Could not load records", http.StatusInternalServerError)
	}

	w.Header().Set("Content/Type", "application/json")
	if err := json.NewEncoder(w).Encode(collection); err != nil {
		log.Println("Failer to write response:", err)
	}
}
