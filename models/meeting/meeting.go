package models

type Meeting struct {
	Title     string
	Attendees string
}

func (m *Meeting) GetCollectionName() string {
	return "meetings"
}
