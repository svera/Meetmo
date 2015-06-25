package meetings

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/maxwellhealth/bongo"
	"github.com/svera/meetmo/core/form"
	"github.com/svera/meetmo/models/meeting"
	"gopkg.in/mgo.v2/bson"
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request, dbConnection *bongo.Connection) {
	results := dbConnection.Collection("meetings").Find(nil)
	meeting := &models.Meeting{}
	meetings := make([]models.Meeting, 3)

	for results.Next(meeting) {
		meetings = append(meetings, *meeting)
		//fmt.Println(meeting.Title)
	}

	t := template.Must(template.ParseFiles("views/meetings/index.html", "views/shared/header.html"))
	t.Execute(w, meetings)
}

func New(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("views/meetings/new.html")
	t.Execute(w, nil)
}

func Create(w http.ResponseWriter, r *http.Request, dbConnection *bongo.Connection) {
	meeting := &models.Meeting{
		Title:     r.FormValue("title"),
		Attendees: r.FormValue("attendees"),
		Agenda:    r.FormValue("agenda"),
		Outcome:   r.FormValue("outcome"),
	}
	err := dbConnection.Collection("meetings").Save(meeting)
	if vErr, ok := err.(*bongo.ValidationError); ok {
		formErrors := getFormErrors(vErr)
		data := struct {
			Meeting    *models.Meeting
			FormErrors map[string]string
		}{
			meeting,
			formErrors,
		}
		t, _ := template.ParseFiles("views/meetings/new.html")
		t.Execute(w, data)
	}
}

func Edit(w http.ResponseWriter, r *http.Request, dbConnection *bongo.Connection) {
	meeting := &models.Meeting{}
	vars := mux.Vars(r)
	id := vars["id"]
	err := dbConnection.Collection("meetings").FindById(bson.ObjectIdHex(id), meeting)
	if _, ok := err.(*bongo.DocumentNotFoundError); ok {
		fmt.Println("document not found")
	} else {
		data := struct {
			Meeting    *models.Meeting
			FormErrors map[string]string
		}{
			meeting,
			nil,
		}
		t := template.Must(template.ParseFiles("views/layouts/base.html", "views/meetings/edit.html"))
		t.Execute(w, data)
	}
}

func Update(w http.ResponseWriter, r *http.Request, dbConnection *bongo.Connection) {
	meeting := &models.Meeting{}
	vars := mux.Vars(r)
	id := vars["id"]
	err := dbConnection.Collection("meetings").FindById(bson.ObjectIdHex(id), meeting)
	if vErr, ok := err.(*bongo.ValidationError); ok {
		formErrors := getFormErrors(vErr)
		data := struct {
			Meeting    *models.Meeting
			FormErrors map[string]string
		}{
			meeting,
			formErrors,
		}
		t, _ := template.ParseFiles("views/meetings/edit.html")
		t.Execute(w, data)
	}
	meeting.Title = r.FormValue("title")
	meeting.Attendees = r.FormValue("attendees")
	meeting.Agenda = r.FormValue("agenda")
	meeting.Outcome = r.FormValue("outcome")

	dbConnection.Collection("meetings").Save(meeting)
	http.Redirect(w, r, "/meetings", 301)
}

func Delete(w http.ResponseWriter, r *http.Request, dbConnection *bongo.Connection) {
	meeting := &models.Meeting{}
	vars := mux.Vars(r)
	id := vars["id"]
	err := dbConnection.Collection("meetings").FindById(bson.ObjectIdHex(id), meeting)
	if _, ok := err.(*bongo.DocumentNotFoundError); ok {
		fmt.Println("document not found")
	} else {
		dbConnection.Collection("meetings").DeleteDocument(meeting)
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
