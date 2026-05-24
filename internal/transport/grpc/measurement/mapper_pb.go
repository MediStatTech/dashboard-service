package measurement

import (
	pb_models "github.com/MediStatTech/dashboard-client/pb/go/models/v1"
	"github.com/MediStatTech/dashboard-service/internal/app/dashboard/domain"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func measurementToPb(m domain.Measurement) *pb_models.Measurement {
	components := make([]*pb_models.Component, 0, len(m.Components))
	for _, c := range m.Components {
		components = append(components, &pb_models.Component{
			MetricTypeId: c.MetricTypeID,
			Code:         c.Code,
			Name:         c.Name,
			Value:        c.Value,
			Symbol:       c.Symbol,
		})
	}
	return &pb_models.Measurement{
		SensorId:   m.SensorID,
		PatientId:  m.PatientID,
		CreatedAt:  timestamppb.New(m.CreatedAt),
		Components: components,
	}
}
