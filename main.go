package main

import (
    "net/http"
    "html/template"
    "github.com/gorilla/mux"
)

func handlerNew(w http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles("views/new.html")
    t.Execute(w, nil)
}

func handlerCreate() {

}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/new", handlerNew)
    r.HandleFunc("/create", handlerCreate).Methods("POST")
    http.Handle("/", r)
    http.ListenAndServe(":8080", nil)
}

