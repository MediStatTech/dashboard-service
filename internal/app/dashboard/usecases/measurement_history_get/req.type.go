package measurement_history_get

import (
	"time"

	"github.com/MediStatTech/dashboard-service/internal/app/dashboard/domain"
)

type Request struct {
	SensorID  string
	PatientID string
	StartTime time.Time
	EndTime   time.Time
	Limit     int32
	Offset    int32
}

type Response struct {
	Measurements []domain.Measurement
	Total        int32
}
