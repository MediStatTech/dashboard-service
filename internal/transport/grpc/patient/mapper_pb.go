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
	metrics := make([]*pb_models.Metric, 0, len(s.Metrics))
	for _, m := range s.Metrics {
		metrics = append(metrics, &pb_models.Metric{
			Value:     m.Value,
			Symbol:    m.Symbol,
			CreatedAt: timestamppb.New(m.CreatedAt),
		})
	}

	return &pb_models.Sensor{
		SensorId: s.SensorID,
		Name:     s.Name,
		Code:     s.Code,
		Symbol:   s.Symbol,
		Metrics:  metrics,
	}
}
