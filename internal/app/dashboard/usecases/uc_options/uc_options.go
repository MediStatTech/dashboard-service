package uc_options

import "github.com/MediStatTech/dashboard-service/internal/app/dashboard/contracts"

type Options struct {
	Logger contracts.Logger

	// Auth service contracts
	JwtService      contracts.JwtService
	StaffsService   contracts.StaffsService
	PositionService contracts.PositionService

	// Patient service contracts
	PatientService            contracts.PatientService
	PatientDiseasService      contracts.PatientDiseasService
	PatientContactInfoService contracts.PatientContactInfoService
	PatientAddressService     contracts.PatientAddressService

	// Biometric service contracts
	DiseasSensorService        contracts.DiseasSensorService
	DiseasService              contracts.DiseasService
	SensorService              contracts.SensorService
	SensorPatientService       contracts.SensorPatientService
	SensorPatientMetricService contracts.SensorPatientMetricService
}
