package patient

import (
	"context"

	pb_models "github.com/MediStatTech/dashboard-client/pb/go/models/v1"
	pb "github.com/MediStatTech/dashboard-client/pb/go/services/v1"
	"github.com/MediStatTech/dashboard-service/internal/app/dashboard/usecases/patient_create"
	"github.com/MediStatTech/dashboard-service/internal/app/dashboard/usecases/patient_get"
	"github.com/MediStatTech/dashboard-service/internal/app/dashboard/usecases/patient_retrieve"
	"github.com/MediStatTech/dashboard-service/pkg/auth"
)

func (h *Handler) PatientGet(
	ctx context.Context,
	_ *pb.PatientGetRequest,
) (*pb.PatientGetReply, error) {
	resp, err := h.queries.PatientGet.Execute(ctx, patient_get.Request{})
	if err != nil {
		return nil, err
	}

	patients := make([]*pb_models.Patient_ListItem, 0, len(resp.Patients))
	for _, p := range resp.Patients {
		patients = append(patients, patientListItemToPb(p))
	}

	return &pb.PatientGetReply{
		Patients: patients,
	}, nil
}

func (h *Handler) PatientRetrieve(
	ctx context.Context,
	req *pb.PatientRetrieveRequest,
) (*pb.PatientRetrieveReply, error) {
	if req == nil {
		return nil, errRequestNil
	}

	resp, err := h.queries.PatientRetrieve.Execute(ctx, patient_retrieve.Request{
		PatientID: req.GetPatientId(),
	})
	if err != nil {
		return nil, err
	}

	return &pb.PatientRetrieveReply{
		Patient: patientToPb(resp.Patient),
	}, nil
}

func (h *Handler) PatientCreate(
	ctx context.Context,
	req *pb.PatientCreateRequest,
) (*pb.PatientCreateReply, error) {
	if req == nil || req.GetPatient() == nil {
		return nil, errRequestNil
	}

	authInfo := auth.GetAuth(ctx)
	if authInfo.StaffID == "" {
		return nil, errMissingStaffID
	}

	patientData := req.GetPatient()

	createReq := patient_create.Request{
		StaffID:   authInfo.StaffID,
		FirstName: patientData.GetFirstName(),
		LastName:  patientData.GetLastName(),
		Gender:    patientData.GetGender(),
		Dob:       patientData.GetDob(),
		DiseasIDs: patientData.GetDiseasIds(),
	}

	if patientData.GetContactInfo() != nil {
		createReq.ContactInfo = &patient_create.CreateContactInfo{
			Phone:   patientData.GetContactInfo().GetPhone(),
			Email:   patientData.GetContactInfo().GetEmail(),
			Primary: patientData.GetContactInfo().GetPrimary(),
		}
	}

	if patientData.GetAddress() != nil {
		createReq.Address = &patient_create.CreateAddress{
			Line1: patientData.GetAddress().GetLine_1(),
			City:  patientData.GetAddress().GetCity(),
			State: patientData.GetAddress().GetState(),
		}
	}

	createResp, err := h.commands.PatientCreate.Execute(ctx, createReq)
	if err != nil {
		return nil, err
	}

	// Retrieve the created patient with full data
	retrieveResp, err := h.queries.PatientRetrieve.Execute(ctx, patient_retrieve.Request{
		PatientID: createResp.PatientID,
	})
	if err != nil {
		return nil, err
	}

	return &pb.PatientCreateReply{
		Patient: patientToPb(retrieveResp.Patient),
	}, nil
}
