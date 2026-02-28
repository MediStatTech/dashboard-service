package staff

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	errMissingStaffID = status.Error(codes.Unauthenticated, "missing staff_id in context")
)
