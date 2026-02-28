package app

import (
	"fmt"

	"github.com/MediStatTech/dashboard-service/internal/app/dashboard"
	"github.com/MediStatTech/dashboard-service/internal/app/dashboard/contracts"
	"github.com/MediStatTech/dashboard-service/internal/app/dashboard/usecases"
	"github.com/MediStatTech/dashboard-service/pkg"
)

type Facade struct {
	Dashboard  *usecases.Facade
	JwtService contracts.JwtService
}

func New(pkg *pkg.Facade) (*Facade, error) {
	dashboardFacade, err := dashboard.New(pkg)
	if err != nil {
		return nil, fmt.Errorf("failed to create dashboard: %w", err)
	}

	return &Facade{
		Dashboard:  dashboardFacade.UseCases,
		JwtService: dashboardFacade.JwtService,
	}, nil
}
