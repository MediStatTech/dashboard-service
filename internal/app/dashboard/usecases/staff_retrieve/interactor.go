package staff_retrieve

import (
	"context"

	auth_v1 "github.com/MediStatTech/auth-client/pb/go/services/v1"
	"github.com/MediStatTech/dashboard-service/internal/app/dashboard/contracts"
	"github.com/MediStatTech/dashboard-service/internal/app/dashboard/domain"
	"golang.org/x/sync/errgroup"
)

type Interactor struct {
	staffsService   contracts.StaffsService
	positionService contracts.PositionService
	logger          contracts.Logger
}

func New(
	staffsService contracts.StaffsService,
	positionService contracts.PositionService,
	logger contracts.Logger,
) *Interactor {
	return &Interactor{
		staffsService:   staffsService,
		positionService: positionService,
		logger:          logger,
	}
}

func (it *Interactor) Execute(ctx context.Context, req Request) (*Response, error) {
	if req.StaffID == "" {
		return nil, errInvalidRequest
	}

	var staffReply *auth_v1.StaffRetrieveReply
	var positionsReply *auth_v1.PositionGetReply

	eg, egCtx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		var err error
		staffReply, err = it.fetchStaff(egCtx, req.StaffID)
		return err
	})

	eg.Go(func() error {
		var err error
		positionsReply, err = it.fetchPositions(egCtx)
		return err
	})

	if err := eg.Wait(); err != nil {
		return nil, err
	}

	staff := staffReply.GetStaff()
	if staff == nil {
		return nil, errStaffNotFound
	}

	positionName := it.findPositionName(positionsReply, staff.GetPositionId())

	return &Response{
		Staff: domain.Staff{
			StaffID:        staff.GetStaffId(),
			FirstName:      staff.GetFirstName(),
			LastName:       staff.GetLastName(),
			SelfieURL:      staff.SelfieUrl,
			SelfieThumbURL: staff.SelfieThumbUrl,
			Status:         staff.GetStatus(),
			Email:          staff.GetEmail(),
			Position: domain.Position{
				PositionID: staff.GetPositionId(),
				Name:       positionName,
			},
		},
	}, nil
}

func (it *Interactor) fetchStaff(ctx context.Context, staffID string) (*auth_v1.StaffRetrieveReply, error) {
	return it.staffsService.StaffRetrieve(ctx, &auth_v1.StaffRetrieveRequest{
		StaffId: staffID,
	})
}

func (it *Interactor) fetchPositions(ctx context.Context) (*auth_v1.PositionGetReply, error) {
	return it.positionService.PositionGet(ctx, &auth_v1.PositionGetRequest{})
}

func (it *Interactor) findPositionName(reply *auth_v1.PositionGetReply, positionID string) string {
	for _, p := range reply.GetPositions() {
		if p.GetPositionId() == positionID {
			return p.GetName()
		}
	}
	return ""
}
