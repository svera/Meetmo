package routes

import (
    "github.com/gorilla/mux"
    "github.com/svera/meetmo/controllers/meetings"
)

func AddRoutes() *mux.Router{
    r := mux.NewRouter()
    r.HandleFunc("/meetings/new", meetings.HandlerNew)
    r.HandleFunc("/meetings/create", meetings.HandlerCreate).Methods("POST")
    return r
}