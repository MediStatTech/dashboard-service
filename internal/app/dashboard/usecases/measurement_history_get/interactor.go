package measurement_history_get

import (
	"context"
	"time"

	biometric_v1 "github.com/MediStatTech/biometric-client/pb/go/services/v1"
	"github.com/MediStatTech/dashboard-service/internal/app/dashboard/contracts"
	"github.com/MediStatTech/dashboard-service/internal/app/dashboard/domain"
	"google.golang.org/protobuf/types/known/timestamppb"
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
	if req.EndTime.IsZero() {
		req.EndTime = time.Now().UTC()
	}
	if req.StartTime.IsZero() || !req.StartTime.Before(req.EndTime) {
		return nil, errInvalidTimeRange
	}

	reply, err := it.sensorPatientMetricService.SensorPatientMetricHistoryGet(ctx, &biometric_v1.SensorPatientMetricHistoryGetRequest{
		SensorId:  req.SensorID,
		PatientId: req.PatientID,
		StartTime: timestamppb.New(req.StartTime),
		EndTime:   timestamppb.New(req.EndTime),
		Limit:     req.Limit,
		Offset:    req.Offset,
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

	return &Response{
		Measurements: measurements,
		Total:        reply.GetTotal(),
	}, nil
}
