package models

import (
	"errors"
	"fmt"
	"github.com/maxwellhealth/bongo"
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
	err := make([]error, 1)

	if len(m.Title) == 0 {
		err = append(err, errors.New("Title cannot be empty"))
		fmt.Println("error en el title")
	}
	return err
}
