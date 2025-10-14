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

	// Create a new request multiplexer
	// Take incoming requests and dispatch them to the matching handlers
	mux := http.NewServeMux()

	// Register the routes and handlers
	mux.Handle("/", &controllers.HomeHandler{})
	mux.HandleFunc("GET /users", ctr.Get)

	// Run the server
	fmt.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
