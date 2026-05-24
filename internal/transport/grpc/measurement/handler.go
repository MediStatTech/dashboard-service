package measurement

import (
	"context"

	pb_models "github.com/MediStatTech/dashboard-client/pb/go/models/v1"
	pb "github.com/MediStatTech/dashboard-client/pb/go/services/v1"
	"github.com/MediStatTech/dashboard-service/internal/app/dashboard/usecases/measurement_get"
	"github.com/MediStatTech/dashboard-service/internal/app/dashboard/usecases/measurement_history_get"
)

func (h *Handler) MeasurementGet(
	ctx context.Context,
	req *pb.MeasurementGetRequest,
) (*pb.MeasurementGetReply, error) {
	resp, err := h.queries.MeasurementGet.Execute(ctx, measurement_get.Request{
		SensorID:  req.GetSensorId(),
		PatientID: req.GetPatientId(),
	})
	if err != nil {
		return nil, err
	}

	measurements := make([]*pb_models.Measurement, 0, len(resp.Measurements))
	for _, m := range resp.Measurements {
		measurements = append(measurements, measurementToPb(m))
	}

	return &pb.MeasurementGetReply{
		Measurements: measurements,
	}, nil
}

func (h *Handler) MeasurementHistoryGet(
	ctx context.Context,
	req *pb.MeasurementHistoryGetRequest,
) (*pb.MeasurementHistoryGetReply, error) {
	resp, err := h.queries.MeasurementHistoryGet.Execute(ctx, measurement_history_get.Request{
		SensorID:  req.GetSensorId(),
		PatientID: req.GetPatientId(),
		StartTime: req.GetStartTime().AsTime(),
		EndTime:   req.GetEndTime().AsTime(),
		Limit:     req.GetLimit(),
		Offset:    req.GetOffset(),
	})
	if err != nil {
		return nil, err
	}

	measurements := make([]*pb_models.Measurement, 0, len(resp.Measurements))
	for _, m := range resp.Measurements {
		measurements = append(measurements, measurementToPb(m))
	}

	return &pb.MeasurementHistoryGetReply{
		Measurements: measurements,
		Total:        resp.Total,
	}, nil
}
