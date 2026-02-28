package domain

type Staff struct {
	StaffID        string
	FirstName      string
	LastName       string
	SelfieURL      *string
	SelfieThumbURL *string
	Status         string
	Email          string
	Position       Position
}
