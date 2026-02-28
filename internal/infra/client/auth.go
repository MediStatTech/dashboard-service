package client

import (
	"fmt"

	auth_client "github.com/MediStatTech/auth-client"
	"github.com/MediStatTech/auth-client/client_options"
	log "github.com/MediStatTech/logger"
)

func NewAuthClient(l *log.Logger) (*auth_client.Facade, error) {
	client, err := auth_client.New(&client_options.Options{
		Log: l,
	})
	if err != nil {
		return nil, fmt.Errorf("auth client connect: %w", err)
	}

	return client, nil
}
