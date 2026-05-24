package patient_get

import (
	"context"

	biometric_v1 "github.com/MediStatTech/biometric-client/pb/go/services/v1"
	patient_v1 "github.com/MediStatTech/patient-client/pb/go/services/v1"
	"github.com/MediStatTech/dashboard-service/internal/app/dashboard/contracts"
	"github.com/MediStatTech/dashboard-service/internal/app/dashboard/domain"
)

const defaultStatus = "ok"

type Interactor struct {
	patientService       contracts.PatientService
	patientStatusService contracts.PatientStatusService
	logger               contracts.Logger
}

func New(
	patientService contracts.PatientService,
	patientStatusService contracts.PatientStatusService,
	logger contracts.Logger,
) *Interactor {
	return &Interactor{
		patientService:       patientService,
		patientStatusService: patientStatusService,
		logger:               logger,
	}
}

func (it *Interactor) Execute(ctx context.Context, _ Request) (*Response, error) {
	reply, err := it.patientService.PatientGet(ctx, &patient_v1.PatientGetRequest{})
	if err != nil {
		return nil, err
	}

	rawPatients := reply.GetPatients()
	if len(rawPatients) == 0 {
		return &Response{Patients: nil}, nil
	}

	patientIDs := make([]string, 0, len(rawPatients))
	for _, p := range rawPatients {
		patientIDs = append(patientIDs, p.GetPatientId())
	}

	statusByID := make(map[string]string, len(patientIDs))
	statusReply, err := it.patientStatusService.PatientStatusGetBatch(ctx, &biometric_v1.PatientStatusGetBatchRequest{
		PatientIds: patientIDs,
	})
	if err != nil {
		it.logger.Error("failed to fetch patient statuses", map[string]any{
			"error": err.Error(),
		})
	} else {
		for _, s := range statusReply.GetPatientStatuses() {
			statusByID[s.GetPatientId()] = s.GetStatus()
		}
	}

	patients := make([]domain.PatientListItem, 0, len(rawPatients))
	for _, p := range rawPatients {
		status, ok := statusByID[p.GetPatientId()]
		if !ok {
			status = defaultStatus
		}
		patients = append(patients, domain.PatientListItem{
			PatientID: p.GetPatientId(),
			FirstName: p.GetFirstName(),
			LastName:  p.GetLastName(),
			Gender:    p.GetGender(),
			Dob:       p.GetDob(),
			Status:    status,
		})
	}

	return &Response{
		Patients: patients,
	}, nil
}
