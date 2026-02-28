package patient_get

import (
	"context"

	patient_v1 "github.com/MediStatTech/patient-client/pb/go/services/v1"
	"github.com/MediStatTech/dashboard-service/internal/app/dashboard/contracts"
	"github.com/MediStatTech/dashboard-service/internal/app/dashboard/domain"
)

type Interactor struct {
	patientService contracts.PatientService
	logger         contracts.Logger
}

func New(
	patientService contracts.PatientService,
	logger contracts.Logger,
) *Interactor {
	return &Interactor{
		patientService: patientService,
		logger:         logger,
	}
}

func (it *Interactor) Execute(ctx context.Context, _ Request) (*Response, error) {
	reply, err := it.patientService.PatientGet(ctx, &patient_v1.PatientGetRequest{})
	if err != nil {
		return nil, err
	}

	patients := make([]domain.PatientListItem, 0, len(reply.GetPatients()))
	for _, p := range reply.GetPatients() {
		patients = append(patients, domain.PatientListItem{
			PatientID: p.GetPatientId(),
			FirstName: p.GetFirstName(),
			LastName:  p.GetLastName(),
			Gender:    p.GetGender(),
			Dob:       p.GetDob(),
			Status:    "active",
		})
	}

	return &Response{
		Patients: patients,
	}, nil
}
