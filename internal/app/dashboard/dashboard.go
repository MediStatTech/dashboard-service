package dashboard

import (
	"fmt"

	"github.com/MediStatTech/dashboard-service/internal/app/dashboard/contracts"
	"github.com/MediStatTech/dashboard-service/internal/app/dashboard/usecases"
	"github.com/MediStatTech/dashboard-service/internal/app/dashboard/usecases/uc_options"
	"github.com/MediStatTech/dashboard-service/internal/infra/client"
	"github.com/MediStatTech/dashboard-service/pkg"
)

type Facade struct {
	pkg        *pkg.Facade
	UseCases   *usecases.Facade
	JwtService contracts.JwtService
}

func New(pkg *pkg.Facade) (*Facade, error) {
	authClient, err := client.NewAuthClient(pkg.Logger)
	if err != nil {
		return nil, fmt.Errorf("init auth client: %w", err)
	}

	patientClient, err := client.NewPatientClient(pkg.Logger)
	if err != nil {
		return nil, fmt.Errorf("init patient client: %w", err)
	}

	biometricClient, err := client.NewBiometricClient(pkg.Logger)
	if err != nil {
		return nil, fmt.Errorf("init biometric client: %w", err)
	}

	useCasesInstance := usecases.New(&uc_options.Options{
		Logger: pkg.Logger,

		// Auth
		JwtService:      authClient.Jwt,
		StaffsService:   authClient.Staffs,
		PositionService: authClient.Position,

		// Patient
		PatientService:            patientClient.Patient,
		PatientDiseasService:      patientClient.PatientDiseas,
		PatientContactInfoService: patientClient.PatientContactInfo,
		PatientAddressService:     patientClient.PatientAddress,

		// Biometric
		DiseasSensorService:  biometricClient.DiseasSensor,
		DiseasService:        biometricClient.Diseas,
		SensorService:        biometricClient.Sensor,
		SensorPatientService: biometricClient.SensorPatient,
	})

	return &Facade{
		pkg:        pkg,
		UseCases:   useCasesInstance,
		JwtService: authClient.Jwt,
	}, nil
}
