package domain

type PatientListItem struct {
	PatientID string
	FirstName string
	LastName  string
	Gender    string
	Dob       string
	Status    string
}

type Patient struct {
	PatientID   string
	FirstName   string
	LastName    string
	Gender      string
	Dob         string
	Status      string
	ContactInfo *ContactInfo
	Address     *Address
	Diseases    []Diseas
	Sensors     []Sensor
}
