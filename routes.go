package main

import (
    "github.com/gorilla/mux"
    "github.com/svera/meetmo/controllers/meetings"
    "github.com/svera/meetmo/core/database"
    "net/http"
)

var (
    routes *mux.Router
    db *database.Database
)

func init() {
    routes = mux.NewRouter()
    db = database.Connect("localhost", "meetmo")
    routes.HandleFunc("/meetings/new", meetings.New)
    routes.HandleFunc("/meetings/create", func(w http.ResponseWriter, r *http.Request) {
        meetings.Create(w, r, db)
    }).Methods("POST")
}