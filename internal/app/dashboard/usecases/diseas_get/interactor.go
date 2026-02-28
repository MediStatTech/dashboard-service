package diseas_get

import (
	"context"

	biometric_v1 "github.com/MediStatTech/biometric-client/pb/go/services/v1"
	"github.com/MediStatTech/dashboard-service/internal/app/dashboard/contracts"
	"github.com/MediStatTech/dashboard-service/internal/app/dashboard/domain"
)

type Interactor struct {
	diseasService contracts.DiseasService
	logger        contracts.Logger
}

func New(
	diseasService contracts.DiseasService,
	logger contracts.Logger,
) *Interactor {
	return &Interactor{
		diseasService: diseasService,
		logger:        logger,
	}
}

func (it *Interactor) Execute(ctx context.Context, _ Request) (*Response, error) {
	reply, err := it.diseasService.DiseasGet(ctx, &biometric_v1.DiseasGetRequest{})
	if err != nil {
		return nil, err
	}

	diseases := make([]domain.Diseas, 0, len(reply.GetDiseases()))
	for _, d := range reply.GetDiseases() {
		diseases = append(diseases, domain.Diseas{
			DiseasID: d.GetDiseasId(),
			Name:     d.GetName(),
			Code:     d.GetCode(),
		})
	}

	return &Response{
		Diseases: diseases,
	}, nil
}
