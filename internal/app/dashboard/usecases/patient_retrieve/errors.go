package patient_retrieve

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	errInvalidRequest = status.Error(codes.InvalidArgument, "patient_id is required")
	errPatientNotFound = status.Error(codes.NotFound, "patient not found")
)
