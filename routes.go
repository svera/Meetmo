package main

import (
    "github.com/svera/meetmo/controllers/meetings"
    "github.com/svera/meetmo/core/database"
    "net/http"
)

func setupRoutes(db *database.Database) {
    routes.HandleFunc("/meetings/new", meetings.New)
    routes.HandleFunc("/meetings/create", func(w http.ResponseWriter, r *http.Request) {
        meetings.Create(w, r, db)
    }).Methods("POST")
}