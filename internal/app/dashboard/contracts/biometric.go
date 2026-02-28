package contracts

import (
	"context"

	biometric_v1 "github.com/MediStatTech/biometric-client/pb/go/services/v1"
	"google.golang.org/grpc"
)

type DiseasSensorService interface {
	DiseasSensorCreate(ctx context.Context, in *biometric_v1.DiseasSensorCreateRequest, opts ...grpc.CallOption) (*biometric_v1.DiseasSensorCreateReply, error)
	DiseasSensorGet(ctx context.Context, in *biometric_v1.DiseasSensorGetRequest, opts ...grpc.CallOption) (*biometric_v1.DiseasSensorGetReply, error)
}

type DiseasService interface {
	DiseasCreate(ctx context.Context, in *biometric_v1.DiseasCreateRequest, opts ...grpc.CallOption) (*biometric_v1.DiseasCreateReply, error)
	DiseasGet(ctx context.Context, in *biometric_v1.DiseasGetRequest, opts ...grpc.CallOption) (*biometric_v1.DiseasGetReply, error)
}

type SensorService interface {
	SensorCreate(ctx context.Context, in *biometric_v1.SensorCreateRequest, opts ...grpc.CallOption) (*biometric_v1.SensorCreateReply, error)
	SensorGet(ctx context.Context, in *biometric_v1.SensorGetRequest, opts ...grpc.CallOption) (*biometric_v1.SensorGetReply, error)
}

type SensorPatientService interface {
	SensorPatientCreate(ctx context.Context, in *biometric_v1.SensorPatientCreateRequest, opts ...grpc.CallOption) (*biometric_v1.SensorPatientCreateReply, error)
	SensorPatientGet(ctx context.Context, in *biometric_v1.SensorPatientGetRequest, opts ...grpc.CallOption) (*biometric_v1.SensorPatientGetReply, error)
	SensorPatientRetrieve(ctx context.Context, in *biometric_v1.SensorPatientRetrieveRequest, opts ...grpc.CallOption) (*biometric_v1.SensorPatientRetrieveReply, error)
}
