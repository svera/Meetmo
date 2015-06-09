package meetings

import (
	"github.com/svera/meetmo/core/database"
	"github.com/svera/meetmo/models/meeting"
	"html/template"
	"net/http"
)

func New(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("views/new.html")
	t.Execute(w, nil)
}

func Create(w http.ResponseWriter, r *http.Request, db *database.Database) {
	meeting := models.Meeting{Title: "Test"}
	db.Insert(&meeting)
}
