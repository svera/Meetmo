package models

import (
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
