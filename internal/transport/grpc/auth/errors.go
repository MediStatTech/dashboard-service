package auth

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	errRequestNil = status.Error(codes.InvalidArgument, "request is nil")
)
