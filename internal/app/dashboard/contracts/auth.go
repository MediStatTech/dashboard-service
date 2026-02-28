package contracts

import (
	"context"

	auth_v1 "github.com/MediStatTech/auth-client/pb/go/services/v1"
	"google.golang.org/grpc"
)

type JwtService interface {
	VerifyToken(ctx context.Context, in *auth_v1.VerifyTokenRequest, opts ...grpc.CallOption) (*auth_v1.VerifyTokenReply, error)
}

type StaffsService interface {
	StaffCreate(ctx context.Context, in *auth_v1.StaffCreateRequest, opts ...grpc.CallOption) (*auth_v1.StaffCreateReply, error)
	StaffDeactivate(ctx context.Context, in *auth_v1.StaffDeactivateRequest, opts ...grpc.CallOption) (*auth_v1.StaffDeactivateReply, error)
	SignIn(ctx context.Context, in *auth_v1.SignInRequest, opts ...grpc.CallOption) (*auth_v1.SignInReply, error)
	StaffRetrieve(ctx context.Context, in *auth_v1.StaffRetrieveRequest, opts ...grpc.CallOption) (*auth_v1.StaffRetrieveReply, error)
	StaffGet(ctx context.Context, in *auth_v1.StaffGetRequest, opts ...grpc.CallOption) (*auth_v1.StaffGetReply, error)
}

type PositionService interface {
	PositionCreate(ctx context.Context, in *auth_v1.PositionCreateRequest, opts ...grpc.CallOption) (*auth_v1.PositionCreateReply, error)
	PositionGet(ctx context.Context, in *auth_v1.PositionGetRequest, opts ...grpc.CallOption) (*auth_v1.PositionGetReply, error)
}
