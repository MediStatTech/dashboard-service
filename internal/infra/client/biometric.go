package client

import (
	"fmt"

	biometric_client "github.com/MediStatTech/biometric-client"
	"github.com/MediStatTech/biometric-client/client_options"
	log "github.com/MediStatTech/logger"
)

func NewBiometricClient(l *log.Logger) (*biometric_client.Facade, error) {
	client, err := biometric_client.New(&client_options.Options{
		Log: l,
	})
	if err != nil {
		return nil, fmt.Errorf("biometric client connect: %w", err)
	}

	return client, nil
}
