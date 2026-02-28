package sign_in

import (
	"context"

	auth_v1 "github.com/MediStatTech/auth-client/pb/go/services/v1"
	"github.com/MediStatTech/dashboard-service/internal/app/dashboard/contracts"
)

type Interactor struct {
	staffsService contracts.StaffsService
	logger        contracts.Logger
}

func New(
	staffsService contracts.StaffsService,
	logger contracts.Logger,
) *Interactor {
	return &Interactor{
		staffsService: staffsService,
		logger:        logger,
	}
}

func (it *Interactor) Execute(ctx context.Context, req Request) (*Response, error) {
	if req.Email == "" || req.Password == "" {
		return nil, errInvalidRequest
	}

	reply, err := it.staffsService.SignIn(ctx, &auth_v1.SignInRequest{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	return &Response{
		Token: reply.GetToken(),
	}, nil
}
