package middleware

import (
	"context"
	"strings"

	auth_v1 "github.com/MediStatTech/auth-client/pb/go/services/v1"
	"github.com/MediStatTech/dashboard-service/internal/app/dashboard/contracts"
	"github.com/MediStatTech/dashboard-service/pkg/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var skipAuthMethods = map[string]bool{
	"/services.v1.AuthService/SignIn": true,
}

func AuthInterceptor(jwtService contracts.JwtService) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (any, error) {
		if skipAuthMethods[info.FullMethod] {
			return handler(ctx, req)
		}

		token, err := extractToken(ctx)
		if err != nil {
			return nil, err
		}

		reply, err := jwtService.VerifyToken(ctx, &auth_v1.VerifyTokenRequest{
			Token: token,
		})
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, "invalid token")
		}

		ctx = auth.WithAuth(ctx, auth.Auth{
			StaffID:    reply.GetStaffId(),
			PositionID: reply.GetPositionId(),
		})

		return handler(ctx, req)
	}
}

func extractToken(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", status.Error(codes.Unauthenticated, "missing metadata")
	}

	values := md.Get("authorization")
	if len(values) == 0 {
		return "", status.Error(codes.Unauthenticated, "missing authorization token")
	}

	token, _ := strings.CutPrefix(values[0], "Bearer ")

	return token, nil
}
