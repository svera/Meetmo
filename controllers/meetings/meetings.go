package meetings

import (
	"github.com/gorilla/mux"
	"github.com/maxwellhealth/bongo"
	"github.com/svera/meetmo/models/meeting"
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request, dbConnection *bongo.Connection) {
	t := template.Must(template.ParseFiles("views/layouts/base.html", "views/meetings/index.html"))
	t.Execute(w, meeting.GetAll(dbConnection))
}

func New(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("views/layouts/base.html", "views/meetings/new.html"))
	t.Execute(w, nil)
}

func Create(w http.ResponseWriter, r *http.Request, dbConnection *bongo.Connection) {
	m := &meeting.Meeting{}
	formErrors := m.Validate(r)
	if len(formErrors) > 0 {
		data := struct {
			Meeting    *meeting.Meeting
			FormErrors map[string]string
		}{
			m,
			formErrors,
		}
		t := template.Must(template.ParseFiles("views/layouts/base.html", "views/meetings/new.html"))
		t.Execute(w, data)
	} else {
		dbConnection.Collection(meeting.CollectionName).Save(m)
		http.Redirect(w, r, "/meetings", 301)
	}
}

func Edit(w http.ResponseWriter, r *http.Request, dbConnection *bongo.Connection) {
	vars := mux.Vars(r)
	m, err := meeting.GetOne(vars["id"], dbConnection)
	if _, ok := err.(*bongo.DocumentNotFoundError); ok {
		http.NotFound(w, r)
	} else {
		data := struct {
			Meeting    *meeting.Meeting
			FormErrors map[string]string
		}{
			m,
			nil,
		}
		t := template.Must(template.ParseFiles("views/layouts/base.html", "views/meetings/edit.html"))
		t.Execute(w, data)
	}
}

func Update(w http.ResponseWriter, r *http.Request, dbConnection *bongo.Connection) {
	m, err := meeting.GetOne(r.FormValue("id"), dbConnection)
	if _, ok := err.(*bongo.DocumentNotFoundError); ok {
		http.NotFound(w, r)
	}

	formErrors := m.Validate(r)
	if len(formErrors) > 0 {
		data := struct {
			Meeting    *meeting.Meeting
			FormErrors map[string]string
		}{
			m,
			formErrors,
		}
		t := template.Must(template.ParseFiles("views/layouts/base.html", "views/meetings/edit.html"))
		t.Execute(w, data)
	} else {
		err = dbConnection.Collection(meeting.CollectionName).Save(m)
		http.Redirect(w, r, "/meetings", 301)
	}
}

func Delete(w http.ResponseWriter, r *http.Request, dbConnection *bongo.Connection) {
	vars := mux.Vars(r)
	err := meeting.Delete(vars["id"], dbConnection)
	if _, ok := err.(*bongo.DocumentNotFoundError); ok {
		http.NotFound(w, r)
	}

	http.Redirect(w, r, "/meetings", 301)
}
