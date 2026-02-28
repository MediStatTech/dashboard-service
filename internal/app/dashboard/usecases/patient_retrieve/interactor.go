package patient_retrieve

import (
	"context"
	"sync"

	biometric_v1 "github.com/MediStatTech/biometric-client/pb/go/services/v1"
	patient_v1 "github.com/MediStatTech/patient-client/pb/go/services/v1"
	"github.com/MediStatTech/dashboard-service/internal/app/dashboard/contracts"
	"github.com/MediStatTech/dashboard-service/internal/app/dashboard/domain"
	"golang.org/x/sync/errgroup"
)

type Interactor struct {
	patientService            contracts.PatientService
	patientContactInfoService contracts.PatientContactInfoService
	patientAddressService     contracts.PatientAddressService
	patientDiseasService      contracts.PatientDiseasService
	diseasService             contracts.DiseasService
	diseasSensorService       contracts.DiseasSensorService
	sensorService             contracts.SensorService
	logger                    contracts.Logger
}

func New(
	patientService contracts.PatientService,
	patientContactInfoService contracts.PatientContactInfoService,
	patientAddressService contracts.PatientAddressService,
	patientDiseasService contracts.PatientDiseasService,
	diseasService contracts.DiseasService,
	diseasSensorService contracts.DiseasSensorService,
	sensorService contracts.SensorService,
	logger contracts.Logger,
) *Interactor {
	return &Interactor{
		patientService:            patientService,
		patientContactInfoService: patientContactInfoService,
		patientAddressService:     patientAddressService,
		patientDiseasService:      patientDiseasService,
		diseasService:             diseasService,
		diseasSensorService:       diseasSensorService,
		sensorService:             sensorService,
		logger:                    logger,
	}
}

// patientData holds the result of the initial parallel fetch phase
type patientData struct {
	patientID   string
	firstName   string
	lastName    string
	gender      string
	dob         string
	contactInfo *domain.ContactInfo
	address     *domain.Address
	diseasIDs   []string
}

func (it *Interactor) Execute(ctx context.Context, req Request) (*Response, error) {
	if req.PatientID == "" {
		return nil, errInvalidRequest
	}

	// Phase 1: fetch patient, contact info, address, diseases in parallel
	data, err := it.fetchPatientData(ctx, req.PatientID)
	if err != nil {
		return nil, err
	}

	// Phase 2: fetch disease details + sensors in parallel
	var diseases []domain.Diseas
	var sensors []domain.Sensor

	eg, egCtx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		var err error
		diseases, err = it.fetchDiseaseDetails(egCtx, data.diseasIDs)
		return err
	})

	eg.Go(func() error {
		var err error
		sensors, err = it.fetchSensors(egCtx, data.diseasIDs)
		return err
	})

	if err := eg.Wait(); err != nil {
		return nil, err
	}

	return &Response{
		Patient: domain.Patient{
			PatientID:   data.patientID,
			FirstName:   data.firstName,
			LastName:    data.lastName,
			Gender:      data.gender,
			Dob:         data.dob,
			Status:      "active",
			ContactInfo: data.contactInfo,
			Address:     data.address,
			Diseases:    diseases,
			Sensors:     sensors,
		},
	}, nil
}

func (it *Interactor) fetchPatientData(ctx context.Context, patientID string) (*patientData, error) {
	data := &patientData{}

	eg, egCtx := errgroup.WithContext(ctx)

	// Fetch patient info
	eg.Go(func() error {
		reply, err := it.patientService.PatientRetrieve(egCtx, &patient_v1.PatientRetrieveRequest{
			PatientId: patientID,
		})
		if err != nil {
			return err
		}
		p := reply.GetPatient()
		if p == nil {
			return errPatientNotFound
		}
		data.patientID = p.GetPatientId()
		data.firstName = p.GetFirstName()
		data.lastName = p.GetLastName()
		data.gender = p.GetGender()
		data.dob = p.GetDob()
		return nil
	})

	// Fetch contact info
	eg.Go(func() error {
		var err error
		data.contactInfo, err = it.fetchContactInfo(egCtx, patientID)
		return err
	})

	// Fetch address
	eg.Go(func() error {
		var err error
		data.address, err = it.fetchAddress(egCtx, patientID)
		return err
	})

	// Fetch patient diseases
	eg.Go(func() error {
		var err error
		data.diseasIDs, err = it.fetchPatientDiseases(egCtx, patientID)
		return err
	})

	if err := eg.Wait(); err != nil {
		return nil, err
	}

	return data, nil
}

func (it *Interactor) fetchContactInfo(ctx context.Context, patientID string) (*domain.ContactInfo, error) {
	reply, err := it.patientContactInfoService.PatientContactInfoGet(ctx, &patient_v1.PatientContactInfoGetRequest{})
	if err != nil {
		return nil, err
	}

	for _, c := range reply.GetPatientContactInfos() {
		if c.GetPatientId() == patientID {
			return &domain.ContactInfo{
				ContactID: c.GetContactId(),
				Phone:     c.GetPhone(),
				Email:     c.GetEmail(),
				Primary:   c.GetPrimary(),
			}, nil
		}
	}

	return nil, nil
}

func (it *Interactor) fetchAddress(ctx context.Context, patientID string) (*domain.Address, error) {
	reply, err := it.patientAddressService.PatientAddressGet(ctx, &patient_v1.PatientAddressGetRequest{})
	if err != nil {
		return nil, err
	}

	for _, a := range reply.GetPatientAddresses() {
		if a.GetPatientId() == patientID {
			return &domain.Address{
				PlaceID: a.GetPlaceId(),
				Line1:   a.GetLine_1(),
				City:    a.GetCity(),
				State:   a.GetState(),
			}, nil
		}
	}

	return nil, nil
}

func (it *Interactor) fetchPatientDiseases(ctx context.Context, patientID string) ([]string, error) {
	reply, err := it.patientDiseasService.PatientDiseasGet(ctx, &patient_v1.PatientDiseasGetRequest{})
	if err != nil {
		return nil, err
	}

	diseasIDs := make([]string, 0)
	for _, pd := range reply.GetPatientDiseases() {
		if pd.GetPatientId() == patientID {
			diseasIDs = append(diseasIDs, pd.GetDiseasId())
		}
	}

	return diseasIDs, nil
}

func (it *Interactor) fetchDiseaseDetails(ctx context.Context, diseasIDs []string) ([]domain.Diseas, error) {
	if len(diseasIDs) == 0 {
		return nil, nil
	}

	reply, err := it.diseasService.DiseasGet(ctx, &biometric_v1.DiseasGetRequest{})
	if err != nil {
		return nil, err
	}

	diseasMap := make(map[string]domain.Diseas, len(reply.GetDiseases()))
	for _, d := range reply.GetDiseases() {
		diseasMap[d.GetDiseasId()] = domain.Diseas{
			DiseasID: d.GetDiseasId(),
			Name:     d.GetName(),
			Code:     d.GetCode(),
		}
	}

	diseases := make([]domain.Diseas, 0, len(diseasIDs))
	for _, id := range diseasIDs {
		if d, ok := diseasMap[id]; ok {
			diseases = append(diseases, d)
		}
	}

	return diseases, nil
}

func (it *Interactor) fetchSensors(ctx context.Context, diseasIDs []string) ([]domain.Sensor, error) {
	if len(diseasIDs) == 0 {
		return nil, nil
	}

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
				return nil
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

	if len(sensorIDSet) == 0 {
		return nil, nil
	}

	// Fetch all sensor details
	allSensorsReply, err := it.sensorService.SensorGet(ctx, &biometric_v1.SensorGetRequest{})
	if err != nil {
		return nil, err
	}

	sensors := make([]domain.Sensor, 0)
	for _, s := range allSensorsReply.GetSensors() {
		if _, ok := sensorIDSet[s.GetSensorId()]; ok {
			sensors = append(sensors, domain.Sensor{
				SensorID: s.GetSensorId(),
				Name:     s.GetName(),
				Code:     s.GetCode(),
			})
		}
	}

	return sensors, nil
}
