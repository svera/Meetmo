package meeting

import (
	//"errors"
	//"fmt"
	"github.com/maxwellhealth/bongo"
	"github.com/svera/meetmo/core/form"
	//"time"
)

type Meeting struct {
	bongo.DocumentBase `bson:",inline"`
	Title              string
	//Date      time.Date
	Attendees string
	Agenda    string
	Outcome   string
}

func (m *Meeting) Validate(*bongo.Collection) []error {
	errs := make([]error, 0)

	if len(m.Title) == 0 {
		errs = append(errs, &form.Error{Field: "Title", Message: "Title cannot be empty"})
	}

	if len(m.Title) > 255 {
		errs = append(errs, &form.Error{Field: "Title", Message: "Title too large"})
	}

	if len(m.Attendees) == 0 {
		errs = append(errs, &form.Error{Field: "Attendees", Message: "Attendees cannot be empty"})
	}

	if len(m.Attendees) > 4000 {
		errs = append(errs, &form.Error{Field: "Attendees", Message: "Attendees too large"})
	}

	if len(m.Agenda) == 0 {
		errs = append(errs, &form.Error{Field: "Agenda", Message: "Agenda cannot be empty"})
	}

	if len(m.Agenda) > 4000 {
		errs = append(errs, &form.Error{Field: "Agenda", Message: "Agenda too large"})
	}

	if len(m.Outcome) == 0 {
		errs = append(errs, &form.Error{Field: "Outcome", Message: "Outcome cannot be empty"})
	}

	if len(m.Outcome) > 4000 {
		errs = append(errs, &form.Error{Field: "Outcome", Message: "Outcome too large"})
	}
	return errs
}
