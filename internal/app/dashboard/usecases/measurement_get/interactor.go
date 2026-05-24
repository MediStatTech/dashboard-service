package measurement_get

import (
	"context"

	biometric_v1 "github.com/MediStatTech/biometric-client/pb/go/services/v1"
	"github.com/MediStatTech/dashboard-service/internal/app/dashboard/contracts"
	"github.com/MediStatTech/dashboard-service/internal/app/dashboard/domain"
)

type Interactor struct {
	sensorPatientMetricService contracts.SensorPatientMetricService
	logger                     contracts.Logger
}

func New(
	sensorPatientMetricService contracts.SensorPatientMetricService,
	logger contracts.Logger,
) *Interactor {
	return &Interactor{
		sensorPatientMetricService: sensorPatientMetricService,
		logger:                     logger,
	}
}

func (it *Interactor) Execute(ctx context.Context, req Request) (*Response, error) {
	if req.SensorID == "" || req.PatientID == "" {
		return nil, errInvalidRequest
	}

	reply, err := it.sensorPatientMetricService.SensorPatientMetricGet(ctx, &biometric_v1.SensorPatientMetricGetRequest{
		SensorId:  req.SensorID,
		PatientId: req.PatientID,
	})
	if err != nil {
		return nil, err
	}

	measurements := make([]domain.Measurement, 0, len(reply.GetMeasurements()))
	for _, m := range reply.GetMeasurements() {
		components := make([]domain.MeasurementComponent, 0, len(m.GetComponents()))
		for _, c := range m.GetComponents() {
			components = append(components, domain.MeasurementComponent{
				MetricTypeID: c.GetMetricTypeId(),
				Code:         c.GetCode(),
				Name:         c.GetName(),
				Value:        c.GetValue(),
				Symbol:       c.GetSymbol(),
			})
		}
		measurements = append(measurements, domain.Measurement{
			SensorID:   m.GetSensorId(),
			PatientID:  m.GetPatientId(),
			CreatedAt:  m.GetCreatedAt().AsTime(),
			Components: components,
		})
	}

	return &Response{Measurements: measurements}, nil
}
