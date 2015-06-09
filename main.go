package main

import (
    "net/http"
    "github.com/svera/meetmo/core/database"
    "github.com/gorilla/mux"        
)

var (
    routes *mux.Router
    db *database.Database
)

func init() {
    db = database.Connect("localhost", "meetmo")
    routes = mux.NewRouter()
    setupRoutes(db)
}

func main() {
    defer db.Close()
    http.Handle("/", routes)
    http.ListenAndServe(":8080", nil)
}

