package measurement_get

import "github.com/MediStatTech/dashboard-service/internal/app/dashboard/domain"

type Request struct {
	SensorID  string
	PatientID string
}

type Response struct {
	Measurements []domain.Measurement
}
