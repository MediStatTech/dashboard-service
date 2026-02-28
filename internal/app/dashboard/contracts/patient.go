package contracts

import (
	"context"

	patient_v1 "github.com/MediStatTech/patient-client/pb/go/services/v1"
	"google.golang.org/grpc"
)

type PatientService interface {
	PatientCreate(ctx context.Context, in *patient_v1.PatientCreateRequest, opts ...grpc.CallOption) (*patient_v1.PatientCreateReply, error)
	PatientGet(ctx context.Context, in *patient_v1.PatientGetRequest, opts ...grpc.CallOption) (*patient_v1.PatientGetReply, error)
	PatientGetByStaffID(ctx context.Context, in *patient_v1.PatientGetByStaffIDRequest, opts ...grpc.CallOption) (*patient_v1.PatientPatientGetByStaffIDReply, error)
	PatientRetrieve(ctx context.Context, in *patient_v1.PatientRetrieveRequest, opts ...grpc.CallOption) (*patient_v1.PatientRetrieveReply, error)
}

type PatientDiseasService interface {
	PatientDiseasCreate(ctx context.Context, in *patient_v1.PatientDiseasCreateRequest, opts ...grpc.CallOption) (*patient_v1.PatientDiseasCreateReply, error)
	PatientDiseasGet(ctx context.Context, in *patient_v1.PatientDiseasGetRequest, opts ...grpc.CallOption) (*patient_v1.PatientDiseasGetReply, error)
	PatientDiseasRetrieve(ctx context.Context, in *patient_v1.PatientDiseasRetrieveRequest, opts ...grpc.CallOption) (*patient_v1.PatientDiseasRetrieveReply, error)
}

type PatientContactInfoService interface {
	PatientContactInfoCreate(ctx context.Context, in *patient_v1.PatientContactInfoCreateRequest, opts ...grpc.CallOption) (*patient_v1.PatientContactInfoCreateReply, error)
	PatientContactInfoGet(ctx context.Context, in *patient_v1.PatientContactInfoGetRequest, opts ...grpc.CallOption) (*patient_v1.PatientContactInfoGetReply, error)
	PatientContactInfoRetrieve(ctx context.Context, in *patient_v1.PatientContactInfoRetrieveRequest, opts ...grpc.CallOption) (*patient_v1.PatientContactInfoRetrieveReply, error)
}

type PatientAddressService interface {
	PatientAddressCreate(ctx context.Context, in *patient_v1.PatientAddressCreateRequest, opts ...grpc.CallOption) (*patient_v1.PatientAddressCreateReply, error)
	PatientAddressGet(ctx context.Context, in *patient_v1.PatientAddressGetRequest, opts ...grpc.CallOption) (*patient_v1.PatientAddressGetReply, error)
}
