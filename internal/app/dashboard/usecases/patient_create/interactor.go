package patient_create

import (
	"context"
	"sync"

	biometric_v1 "github.com/MediStatTech/biometric-client/pb/go/services/v1"
	patient_models "github.com/MediStatTech/patient-client/pb/go/models/v1"
	patient_v1 "github.com/MediStatTech/patient-client/pb/go/services/v1"
	"github.com/MediStatTech/dashboard-service/internal/app/dashboard/contracts"
	"golang.org/x/sync/errgroup"
)

type Interactor struct {
	patientService            contracts.PatientService
	patientAddressService     contracts.PatientAddressService
	patientContactInfoService contracts.PatientContactInfoService
	patientDiseasService      contracts.PatientDiseasService
	diseasSensorService       contracts.DiseasSensorService
	sensorPatientService      contracts.SensorPatientService
	logger                    contracts.Logger
}

func New(
	patientService contracts.PatientService,
	patientAddressService contracts.PatientAddressService,
	patientContactInfoService contracts.PatientContactInfoService,
	patientDiseasService contracts.PatientDiseasService,
	diseasSensorService contracts.DiseasSensorService,
	sensorPatientService contracts.SensorPatientService,
	logger contracts.Logger,
) *Interactor {
	return &Interactor{
		patientService:            patientService,
		patientAddressService:     patientAddressService,
		patientContactInfoService: patientContactInfoService,
		patientDiseasService:      patientDiseasService,
		diseasSensorService:       diseasSensorService,
		sensorPatientService:      sensorPatientService,
		logger:                    logger,
	}
}

func (it *Interactor) Execute(ctx context.Context, req Request) (*Response, error) {
	if req.FirstName == "" || req.LastName == "" || req.Gender == "" || req.Dob == "" {
		return nil, errInvalidRequest
	}

	// 1. Create patient (sequential — need patient_id for all subsequent steps)
	patientID, err := it.createPatient(ctx, req)
	if err != nil {
		return nil, err
	}

	// 2. Create address, contact info, diseases in parallel
	eg, egCtx := errgroup.WithContext(ctx)

	if req.Address != nil {
		eg.Go(func() error {
			return it.createAddress(egCtx, patientID, req.Address)
		})
	}

	if req.ContactInfo != nil {
		eg.Go(func() error {
			return it.createContactInfo(egCtx, patientID, req.ContactInfo)
		})
	}

	if len(req.DiseasIDs) > 0 {
		eg.Go(func() error {
			return it.createDiseases(egCtx, patientID, req.DiseasIDs)
		})
	}

	if err := eg.Wait(); err != nil {
		return nil, err
	}

	// 3. Auto-assign sensors via disease-sensor mapping
	if len(req.DiseasIDs) > 0 {
		it.assignSensors(ctx, patientID, req.DiseasIDs)
	}

	return &Response{
		PatientID: patientID,
	}, nil
}

func (it *Interactor) createPatient(ctx context.Context, req Request) (string, error) {
	createReq := &patient_v1.PatientCreateRequest{
		Patient: &patient_models.Patient_Create{
			FirstName: req.FirstName,
			LastName:  req.LastName,
			Gender:    req.Gender,
			Dob:       req.Dob,
		},
	}
	if req.StaffID != "" {
		createReq.Patient.StaffId = &req.StaffID
	}

	reply, err := it.patientService.PatientCreate(ctx, createReq)
	if err != nil {
		return "", err
	}

	patientID := reply.GetPatient().GetPatientId()
	if patientID == "" {
		return "", errFailedCreatePatient
	}

	return patientID, nil
}

func (it *Interactor) createAddress(ctx context.Context, patientID string, addr *CreateAddress) error {
	_, err := it.patientAddressService.PatientAddressCreate(ctx, &patient_v1.PatientAddressCreateRequest{
		PatientAddress: &patient_models.PatientAddress_Create{
			PatientId: patientID,
			Line_1:    addr.Line1,
			City:      addr.City,
			State:     addr.State,
		},
	})
	if err != nil {
		it.logger.Error("failed to create patient address", map[string]any{
			"patient_id": patientID,
			"error":      err.Error(),
		})
		return errFailedCreateAddress
	}
	return nil
}

func (it *Interactor) createContactInfo(ctx context.Context, patientID string, contact *CreateContactInfo) error {
	_, err := it.patientContactInfoService.PatientContactInfoCreate(ctx, &patient_v1.PatientContactInfoCreateRequest{
		PatientContactInfo: &patient_models.PatientContactInfo_Create{
			PatientId: patientID,
			Phone:     contact.Phone,
			Email:     contact.Email,
			Primary:   contact.Primary,
		},
	})
	if err != nil {
		it.logger.Error("failed to create patient contact info", map[string]any{
			"patient_id": patientID,
			"error":      err.Error(),
		})
		return errFailedCreateContact
	}
	return nil
}

func (it *Interactor) createDiseases(ctx context.Context, patientID string, diseasIDs []string) error {
	eg, egCtx := errgroup.WithContext(ctx)

	for _, diseasID := range diseasIDs {
		eg.Go(func() error {
			_, err := it.patientDiseasService.PatientDiseasCreate(egCtx, &patient_v1.PatientDiseasCreateRequest{
				PatientDiseas: &patient_models.PatientDiseas_Create{
					PatientId: patientID,
					DiseasId:  diseasID,
				},
			})
			if err != nil {
				it.logger.Error("failed to create patient disease", map[string]any{
					"patient_id": patientID,
					"diseas_id":  diseasID,
					"error":      err.Error(),
				})
				return errFailedCreateDiseas
			}
			return nil
		})
	}

	return eg.Wait()
}

func (it *Interactor) assignSensors(ctx context.Context, patientID string, diseasIDs []string) {
	// Collect sensor IDs from all diseases in parallel
	var mu sync.Mutex
	sensorIDSet := make(map[string]struct{})

	eg, egCtx := errgroup.WithContext(ctx)

	for _, diseasID := range diseasIDs {
		eg.Go(func() error {
			dsReply, err := it.diseasSensorService.DiseasSensorGet(egCtx, &biometric_v1.DiseasSensorGetRequest{
				DiseasId: diseasID,
			})
			if err != nil {
				it.logger.Error("failed to get diseas sensors", map[string]any{
					"diseas_id": diseasID,
					"error":     err.Error(),
				})
				return nil // don't fail the whole flow
			}

			mu.Lock()
			for _, ds := range dsReply.GetDiseasSensors() {
				sensorIDSet[ds.GetSensorId()] = struct{}{}
			}
			mu.Unlock()

			return nil
		})
	}

	_ = eg.Wait()

	// Create sensor-patient assignments in parallel
	eg2, egCtx2 := errgroup.WithContext(ctx)

	for sensorID := range sensorIDSet {
		eg2.Go(func() error {
			_, err := it.sensorPatientService.SensorPatientCreate(egCtx2, &biometric_v1.SensorPatientCreateRequest{
				PatientId: patientID,
				SensorId:  sensorID,
			})
			if err != nil {
				it.logger.Error("failed to create sensor-patient assignment", map[string]any{
					"patient_id": patientID,
					"sensor_id":  sensorID,
					"error":      err.Error(),
				})
			}
			return nil // don't fail the whole flow
		})
	}

	_ = eg2.Wait()
}
