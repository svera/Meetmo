package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/maxwellhealth/bongo"
	"github.com/syb-devs/dockerlink"
)

var (
	routes       *mux.Router
	dbConnection *bongo.Connection
)

func init() {
	config := &bongo.Config{
		ConnectionString: getMongoURI(),
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
	port := getEnv("PORT", "8080")
	http.Handle("/", routes)
	log.Printf("HTTP server listening on port %s\n", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}

func getEnv(key, defVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defVal
}

func getMongoURI() string {
	if uri := os.Getenv("MONGO_URL"); uri != "" {
		return uri
	}
	if link, err := dockerlink.GetLink("mongodb", 27017, "tcp"); err == nil {
		return fmt.Sprintf("%s:%d", link.Address, link.Port)
	}
	panic("mongodb connection not found, use MONGO_URL env var or a docker link with mongodb name")
}
