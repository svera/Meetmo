package main

import (
	"github.com/gorilla/mux"
	"github.com/maxwellhealth/bongo"
	"log"
	"net/http"
)

var (
	routes       *mux.Router
	dbConnection *bongo.Connection
)

func init() {
	config := &bongo.Config{
		ConnectionString: "localhost",
		Database:         "meetmo",
	}
	dbConnection, err := bongo.Connect(config)
	if err != nil {
		log.Println("Error connection DB")
	}
	routes = mux.NewRouter()
	setupRoutes(dbConnection)
}

func main() {
	http.Handle("/", routes)
	log.Println("HTTP server listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
