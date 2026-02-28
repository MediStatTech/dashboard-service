package patient_retrieve

import "github.com/MediStatTech/dashboard-service/internal/app/dashboard/domain"

type Request struct {
	PatientID string
}

type Response struct {
	Patient domain.Patient
}
