package diseas_get

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	errFailedToGetDiseases = status.Error(codes.Internal, "failed to get diseases")
)
