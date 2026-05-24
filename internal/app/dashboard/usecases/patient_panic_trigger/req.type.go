package patient_panic_trigger

import "time"

type Request struct {
	PatientID       string
	DurationSeconds int32
}

type Response struct {
	PanicUntil time.Time
}
