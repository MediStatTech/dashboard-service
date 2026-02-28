package staff

import (
	"context"

	pb "github.com/MediStatTech/dashboard-client/pb/go/services/v1"
	"github.com/MediStatTech/dashboard-service/internal/app/dashboard/usecases/staff_retrieve"
	"github.com/MediStatTech/dashboard-service/pkg/auth"
)

func (h *Handler) StaffRetrieve(
	ctx context.Context,
	_ *pb.StaffRetrieveRequest,
) (*pb.StaffRetrieveReply, error) {
	authInfo := auth.GetAuth(ctx)
	if authInfo.StaffID == "" {
		return nil, errMissingStaffID
	}

	resp, err := h.queries.StaffRetrieve.Execute(ctx, staff_retrieve.Request{
		StaffID: authInfo.StaffID,
	})
	if err != nil {
		return nil, err
	}

	return &pb.StaffRetrieveReply{
		Staff: staffToPb(resp.Staff),
	}, nil
}
