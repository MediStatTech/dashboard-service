package client

import (
	"fmt"

	patient_client "github.com/MediStatTech/patient-client"
	"github.com/MediStatTech/patient-client/client_options"
	log "github.com/MediStatTech/logger"
)

func NewPatientClient(l *log.Logger) (*patient_client.Facade, error) {
	client, err := patient_client.New(&client_options.Options{
		Log: l,
		AddressName: "localhost:50054",
	})
	if err != nil {
		return nil, fmt.Errorf("patient client connect: %w", err)
	}

	return client, nil
}
