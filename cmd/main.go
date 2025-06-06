package main

import (
	"fmt"
	"github.com/ssjlee93/fitworks-backend/pkg/db"
	"os"
)

func main() {
	// db config
	//Database
	db, err := db.DatabaseConnection()
	if err != nil {
		os.Exit(1)
	}

	fmt.Println(db)

}
