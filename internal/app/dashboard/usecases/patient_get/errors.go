package patient_get

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	errFailedToGetPatients = status.Error(codes.Internal, "failed to get patients")
)
