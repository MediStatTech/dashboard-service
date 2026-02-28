package sign_in

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	errInvalidRequest = status.Error(codes.InvalidArgument, "email and password are required")
)
