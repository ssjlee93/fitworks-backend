package main

import (
	"github.com/ssjlee93/fitworks-backend/pkg/db"
	"os"
)

func main() {
	// db config
	//Database
	connection, err := db.DatabaseConnection()
	if err != nil {
		os.Exit(1)
	}
	defer connection.Close()

}
