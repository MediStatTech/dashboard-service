package patient_panic_trigger

import (
	"context"

	biometric_v1 "github.com/MediStatTech/biometric-client/pb/go/services/v1"
	"github.com/MediStatTech/dashboard-service/internal/app/dashboard/contracts"
)

type Interactor struct {
	patientStatusService contracts.PatientStatusService
	logger               contracts.Logger
}

func New(
	patientStatusService contracts.PatientStatusService,
	logger contracts.Logger,
) *Interactor {
	return &Interactor{
		patientStatusService: patientStatusService,
		logger:               logger,
	}
}

func (it *Interactor) Execute(ctx context.Context, req Request) (*Response, error) {
	if req.PatientID == "" {
		return nil, errInvalidRequest
	}

	reply, err := it.patientStatusService.PatientPanicTrigger(ctx, &biometric_v1.PatientPanicTriggerRequest{
		PatientId:       req.PatientID,
		DurationSeconds: req.DurationSeconds,
	})
	if err != nil {
		return nil, err
	}

	return &Response{
		PanicUntil: reply.GetPanicUntil().AsTime(),
	}, nil
}
