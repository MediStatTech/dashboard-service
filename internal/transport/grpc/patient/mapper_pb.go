package patient

import (
	pb_models "github.com/MediStatTech/dashboard-client/pb/go/models/v1"
	"github.com/MediStatTech/dashboard-service/internal/app/dashboard/domain"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func patientToPb(p domain.Patient) *pb_models.Patient_Read {
	result := &pb_models.Patient_Read{
		PatientId: p.PatientID,
		FirstName: p.FirstName,
		LastName:  p.LastName,
		Gender:    p.Gender,
		Dob:       p.Dob,
		Status:    p.Status,
	}

	if p.ContactInfo != nil {
		result.ContactInfo = contactInfoToPb(*p.ContactInfo)
	}

	if p.Address != nil {
		result.Address = addressToPb(*p.Address)
	}

	diseases := make([]*pb_models.Diseas, 0, len(p.Diseases))
	for _, d := range p.Diseases {
		diseases = append(diseases, diseasToPb(d))
	}
	result.Diseases = diseases

	sensors := make([]*pb_models.Sensor, 0, len(p.Sensors))
	for _, s := range p.Sensors {
		sensors = append(sensors, sensorToPb(s))
	}
	result.Sensors = sensors

	return result
}

func patientListItemToPb(p domain.PatientListItem) *pb_models.Patient_ListItem {
	return &pb_models.Patient_ListItem{
		PatientId: p.PatientID,
		FirstName: p.FirstName,
		LastName:  p.LastName,
		Gender:    p.Gender,
		Dob:       p.Dob,
		Status:    p.Status,
	}
}

func contactInfoToPb(c domain.ContactInfo) *pb_models.ContactInfo {
	return &pb_models.ContactInfo{
		ContactId: c.ContactID,
		Phone:     c.Phone,
		Email:     c.Email,
		Primary:   c.Primary,
	}
}

func addressToPb(a domain.Address) *pb_models.Address {
	return &pb_models.Address{
		PlaceId: a.PlaceID,
		Line_1:  a.Line1,
		City:    a.City,
		State:   a.State,
	}
}

func diseasToPb(d domain.Diseas) *pb_models.Diseas {
	return &pb_models.Diseas{
		DiseasId: d.DiseasID,
		Name:     d.Name,
		Code:     d.Code,
	}
}

func sensorToPb(s domain.Sensor) *pb_models.Sensor {
	metricTypes := make([]*pb_models.MetricType, 0, len(s.MetricTypes))
	for _, mt := range s.MetricTypes {
		metricTypes = append(metricTypes, &pb_models.MetricType{
			MetricTypeId: mt.MetricTypeID,
			SensorId:     mt.SensorID,
			Code:         mt.Code,
			Name:         mt.Name,
			Symbol:       mt.Symbol,
			MinValue:     mt.MinValue,
			MaxValue:     mt.MaxValue,
		})
	}

	measurements := make([]*pb_models.Measurement, 0, len(s.Measurements))
	for _, m := range s.Measurements {
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
		measurements = append(measurements, &pb_models.Measurement{
			SensorId:   m.SensorID,
			PatientId:  m.PatientID,
			CreatedAt:  timestamppb.New(m.CreatedAt),
			Components: components,
		})
	}

	return &pb_models.Sensor{
		SensorId:     s.SensorID,
		Name:         s.Name,
		Code:         s.Code,
		Symbol:       s.Symbol,
		Status:       s.Status,
		MetricTypes:  metricTypes,
		Measurements: measurements,
	}
}
