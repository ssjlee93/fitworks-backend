package main

import (
	"fmt"
	"github.com/ssjlee93/fitworks-backend/internal/controllers"
	"github.com/ssjlee93/fitworks-backend/pkg/db"
	"log"
	"net/http"
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

	// controller
	ctr := controllers.NewUserController()

	// Create a new ServeMux
	mux := http.NewServeMux()
	// Register handlers for different routes
	mux.HandleFunc("GET /users", ctr.Get)
	////mux.HandleFunc("GET /users/{id}", )
	//mux.HandleFunc("POST /users", ctr.Create)

	fmt.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
