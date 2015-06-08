package meetings

import (
    "net/http"
    "html/template"
    "github.com/svera/meetmo/models/meeting"
)

func HandlerNew(w http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles("views/new.html")
    t.Execute(w, nil)
}

func HandlerCreate(w http.ResponseWriter, r *http.Request) {
    meeting := models.Meeting{}
    meeting.Title = "Test"
}
