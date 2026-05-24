package patient_panic_trigger

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	errInvalidRequest = status.Error(codes.InvalidArgument, "invalid request: patient_id is required")
)
