package meetings

import (
	"fmt"
	"github.com/maxwellhealth/bongo"
	"github.com/svera/meetmo/core/form"
	"github.com/svera/meetmo/models/meeting"
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
		for _, a := range vErr.Errors {
			b := a.(*form.Error)
			fmt.Println(b.Message)
		}
		data := struct {
			Meeting    *models.Meeting
			FormErrors *bongo.ValidationError
		}{
			meeting,
			vErr,
		}
		t, _ := template.ParseFiles("views/meetings/new.html")
		t.Execute(w, data)
	}
}
