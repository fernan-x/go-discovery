package main

import (
	"fmt"
	"log"
	"microservice/pkg/api"
	"microservice/pkg/db"
	"net/http"
	"os"
)

func main() {
	log.Print("Server has started")

	// Start the db
	pgdb, err := db.StartDB()
	if err != nil {
		log.Fatalf("Error starting the database %v", err)
	}

	router := api.StartAPI(pgdb)
	port := os.Getenv("PORT")

	err = http.ListenAndServe(fmt.Sprintf(":%s", port), router)
	if err != nil {
		log.Fatalf("Error from router %v\n", err)
	}
}
