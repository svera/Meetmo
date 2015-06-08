package main

import (
    "net/http"
    "github.com/svera/meetmo/routes"
)

func main() {
    http.Handle("/", routes.AddRoutes())
    http.ListenAndServe(":8080", nil)
}

