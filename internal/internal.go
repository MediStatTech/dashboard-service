package internal

import (
	"context"
	"fmt"

	"github.com/MediStatTech/dashboard-service/internal/app"
	grpc "github.com/MediStatTech/dashboard-service/internal/transport/grpc"
	"github.com/MediStatTech/dashboard-service/pkg"
)

func New(_ context.Context, p *pkg.Facade) (*grpc.Server, error) {
	appInstance, err := app.New(p)
	if err != nil {
		return nil, fmt.Errorf("initialize app: %w", err)
	}

	grpcServer, err := grpc.New(p, appInstance)
	if err != nil {
		return nil, fmt.Errorf("initialize grpc: %w", err)
	}

	return grpcServer, nil
}
