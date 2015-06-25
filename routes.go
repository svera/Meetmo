package main

import (
	"github.com/maxwellhealth/bongo"
	"github.com/svera/meetmo/controllers/meetings"
	"net/http"
)

func setupRoutes(dbConnection *bongo.Connection) {
	routes.HandleFunc("/meetings", func(w http.ResponseWriter, r *http.Request) {
		meetings.Index(w, r, dbConnection)
	})
	routes.HandleFunc("/meetings/new", meetings.New)
	routes.HandleFunc("/meetings/create", func(w http.ResponseWriter, r *http.Request) {
		meetings.Create(w, r, dbConnection)
	}).Methods("POST")
	routes.HandleFunc("/meetings/{id}/delete", func(w http.ResponseWriter, r *http.Request) {
		meetings.Delete(w, r, dbConnection)
	}).Methods("GET")
	routes.HandleFunc("/meetings/{id}/edit", func(w http.ResponseWriter, r *http.Request) {
		meetings.Edit(w, r, dbConnection)
	}).Methods("GET")
	routes.HandleFunc("/meetings/{id}/update", func(w http.ResponseWriter, r *http.Request) {
		meetings.Edit(w, r, dbConnection)
	}).Methods("POST")
}
