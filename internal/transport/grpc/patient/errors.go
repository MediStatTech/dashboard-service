package patient

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	errRequestNil        = status.Error(codes.InvalidArgument, "request is nil")
	errInvalidPatientData = status.Error(codes.InvalidArgument, "invalid patient data")
	errMissingStaffID     = status.Error(codes.Unauthenticated, "missing staff_id in context")
)
