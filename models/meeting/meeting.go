package models

//import "time"

type Meeting struct {
	Title string
	//Date      time.Date
	Attendees string
	Agenda    string
	Outcome   string
}

func (m *Meeting) GetCollectionName() string {
	return "meetings"
}
