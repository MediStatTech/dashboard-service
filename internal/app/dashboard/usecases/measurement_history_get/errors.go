package measurement_history_get

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	errInvalidRequest   = status.Error(codes.InvalidArgument, "invalid request: sensor_id and patient_id are required")
	errInvalidTimeRange = status.Error(codes.InvalidArgument, "invalid time range: start_time must be before end_time")
)
