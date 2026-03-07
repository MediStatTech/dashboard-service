package usecases

import (
	"github.com/MediStatTech/dashboard-service/internal/app/dashboard/usecases/diseas_get"
	"github.com/MediStatTech/dashboard-service/internal/app/dashboard/usecases/patient_create"
	"github.com/MediStatTech/dashboard-service/internal/app/dashboard/usecases/patient_get"
	"github.com/MediStatTech/dashboard-service/internal/app/dashboard/usecases/patient_retrieve"
	"github.com/MediStatTech/dashboard-service/internal/app/dashboard/usecases/sign_in"
	"github.com/MediStatTech/dashboard-service/internal/app/dashboard/usecases/staff_retrieve"
	"github.com/MediStatTech/dashboard-service/internal/app/dashboard/usecases/uc_options"
)

type Facade struct {
	SignIn          *sign_in.Interactor
	StaffRetrieve   *staff_retrieve.Interactor
	PatientGet      *patient_get.Interactor
	PatientRetrieve *patient_retrieve.Interactor
	PatientCreate   *patient_create.Interactor
	DiseasGet       *diseas_get.Interactor
}

func New(o *uc_options.Options) *Facade {
	return &Facade{
		SignIn: sign_in.New(
			o.StaffsService,
			o.Logger,
		),
		StaffRetrieve: staff_retrieve.New(
			o.StaffsService,
			o.PositionService,
			o.Logger,
		),
		PatientGet: patient_get.New(
			o.PatientService,
			o.Logger,
		),
		PatientRetrieve: patient_retrieve.New(
			o.PatientService,
			o.PatientContactInfoService,
			o.PatientAddressService,
			o.PatientDiseasService,
			o.DiseasService,
			o.DiseasSensorService,
			o.SensorService,
			o.SensorPatientMetricService,
			o.Logger,
		),
		PatientCreate: patient_create.New(
			o.PatientService,
			o.PatientAddressService,
			o.PatientContactInfoService,
			o.PatientDiseasService,
			o.DiseasSensorService,
			o.SensorPatientService,
			o.Logger,
		),
		DiseasGet: diseas_get.New(
			o.DiseasService,
			o.Logger,
		),
	}
}
