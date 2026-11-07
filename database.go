package main

import (
	"database/sql"

	"context"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func openDatabase() (*sql.DB, error) {
	db, err := sql.Open("pgx", "host=localhost port=5432 user=agnieszkagrzybowska dbname=record_collection sslmode=disable")
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
