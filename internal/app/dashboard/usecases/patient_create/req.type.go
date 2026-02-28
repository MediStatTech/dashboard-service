package patient_create

type CreateContactInfo struct {
	Phone   string
	Email   string
	Primary bool
}

type CreateAddress struct {
	Line1 string
	City  string
	State string
}

type Request struct {
	StaffID     string
	FirstName   string
	LastName    string
	Gender      string
	Dob         string
	ContactInfo *CreateContactInfo
	Address     *CreateAddress
	DiseasIDs   []string
}

type Response struct {
	PatientID string
}
