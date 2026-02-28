package patient_create

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	errInvalidRequest      = status.Error(codes.InvalidArgument, "first_name, last_name, gender, and dob are required")
	errFailedCreatePatient = status.Error(codes.Internal, "failed to create patient")
	errFailedCreateAddress = status.Error(codes.Internal, "failed to create patient address")
	errFailedCreateContact = status.Error(codes.Internal, "failed to create patient contact info")
	errFailedCreateDiseas  = status.Error(codes.Internal, "failed to create patient disease")
)
