package meetings

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/maxwellhealth/bongo"
	"github.com/svera/meetmo/core/form"
	"github.com/svera/meetmo/models/meeting"
	"html/template"
	"log"
	"net/http"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request, dbConnection *bongo.Connection) {
	t := template.Must(template.ParseFiles("views/layouts/base.html", "views/meetings/index.html"))
	t.Execute(w, meeting.GetAll(dbConnection))
}

func New(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("views/layouts/base.html", "views/meetings/new.html")
	t.Execute(w, nil)
}

func Create(w http.ResponseWriter, r *http.Request, dbConnection *bongo.Connection) {
	date, err := time.Parse(time.RFC3339, fmt.Sprintf("%sT00:00:00+00:00", r.FormValue("date")))
	if err != nil {
		log.Println(err)
	}
	m := &meeting.Meeting{
		Title:     r.FormValue("title"),
		Date:      date,
		Attendees: r.FormValue("attendees"),
		Agenda:    r.FormValue("agenda"),
		Outcome:   r.FormValue("outcome"),
	}
	err = dbConnection.Collection(meeting.CollectionName).Save(m)
	if vErr, ok := err.(*bongo.ValidationError); ok {
		formErrors := getFormErrors(vErr)
		data := struct {
			Meeting    *meeting.Meeting
			FormErrors map[string]string
		}{
			m,
			formErrors,
		}
		t, _ := template.ParseFiles("views/layouts/base.html", "views/meetings/new.html")
		t.Execute(w, data)
	} else {
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
	m.Title = r.FormValue("title")
	m.Attendees = r.FormValue("attendees")
	m.Agenda = r.FormValue("agenda")
	m.Outcome = r.FormValue("outcome")
	err = dbConnection.Collection(meeting.CollectionName).Save(m)
	if vErr, ok := err.(*bongo.ValidationError); ok {
		formErrors := getFormErrors(vErr)
		data := struct {
			Meeting    *meeting.Meeting
			FormErrors map[string]string
		}{
			m,
			formErrors,
		}
		t, _ := template.ParseFiles("views/layouts/base.html", "views/meetings/edit.html")
		t.Execute(w, data)
	} else {
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

func getFormErrors(err *bongo.ValidationError) map[string]string {
	var formErrors map[string]string
	formErrors = make(map[string]string)
	for _, v := range err.Errors {
		fErr := v.(*form.Error)
		formErrors[fErr.Field] = fErr.Message
	}
	return formErrors
}
