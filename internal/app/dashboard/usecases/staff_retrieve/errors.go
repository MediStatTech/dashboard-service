package staff_retrieve

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	errInvalidRequest = status.Error(codes.InvalidArgument, "staff_id is required")
	errStaffNotFound  = status.Error(codes.NotFound, "staff not found")
)
