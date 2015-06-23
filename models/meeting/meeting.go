package models

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
	err := make([]error, 0)

	if len(m.Title) == 0 {
		err = append(err, &form.Error{Field: "Title", Message: "Title cannot be empty"})
	}
	return err
}
