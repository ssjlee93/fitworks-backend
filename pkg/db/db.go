package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func DatabaseConnection() (*sql.DB, error) {
	// Make a connection
	// TODO extract db properties to config.yaml
	connStr := "user=postgres dbname=fitworks sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Check the connection
	err = db.Ping()
	if err != nil {
		log.Fatal("Could not connect to the database:", err)
		return nil, err
	}

	log.Println("Successfully connected to the database!")
	return db, nil
}
