package main

import (
    "net/http"
)

func main() {
    defer db.Close()
    http.Handle("/", routes)
    http.ListenAndServe(":8080", nil)
}

