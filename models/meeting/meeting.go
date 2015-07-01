package meeting

import (
	//"errors"
	"fmt"
	"github.com/maxwellhealth/bongo"
	"net/http"
	"time"
)

type Meeting struct {
	bongo.DocumentBase `bson:",inline"`
	Title              string
	Date               time.Time
	Attendees          string
	Agenda             string
	Outcome            string
}

func (m *Meeting) Validate(r *http.Request) map[string]string {
	errs := make(map[string]string, 0)
	var tErr error

	m.Title = r.FormValue("title")
	m.Date, tErr = time.Parse(time.RFC3339, fmt.Sprintf("%sT00:00:00+00:00", r.FormValue("date")))
	m.Attendees = r.FormValue("attendees")
	m.Agenda = r.FormValue("agenda")
	m.Outcome = r.FormValue("outcome")

	if tErr != nil {
		errs["Date"] = "Wrong date format"
	}

	if len(m.Title) == 0 {
		errs["Title"] = "Title cannot be empty"
	}

	if len(m.Title) > 255 {
		errs["Title"] = "Title too large"
	}

	if len(m.Attendees) == 0 {
		errs["Attendees"] = "Attendees cannot be empty"
	}

	if len(m.Attendees) > 4000 {
		errs["Attendees"] = "Attendees too large"
	}

	if len(m.Agenda) == 0 {
		errs["Agenda"] = "Agenda cannot be empty"
	}

	if len(m.Agenda) > 4000 {
		errs["Agenda"] = "Agenda too large"
	}

	if len(m.Outcome) == 0 {
		errs["Outcome"] = "Outcome cannot be empty"
	}

	if len(m.Outcome) > 4000 {
		errs["Outcome"] = "Outcome too large"
	}

	return errs
}
