package auth

import (
	"context"

	pb "github.com/MediStatTech/dashboard-client/pb/go/services/v1"
	"github.com/MediStatTech/dashboard-service/internal/app/dashboard/usecases/sign_in"
)

func (h *Handler) SignIn(
	ctx context.Context,
	req *pb.SignInRequest,
) (*pb.SignInReply, error) {
	if req == nil {
		return nil, errRequestNil
	}

	resp, err := h.commands.SignIn.Execute(ctx, sign_in.Request{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	})
	if err != nil {
		return nil, err
	}

	return &pb.SignInReply{
		Token: resp.Token,
	}, nil
}
