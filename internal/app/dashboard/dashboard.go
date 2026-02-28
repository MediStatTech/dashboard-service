package dashboard

import (
	"github.com/MediStatTech/dashboard-service/internal/app/dashboard/usecases"
	"github.com/MediStatTech/dashboard-service/internal/app/dashboard/usecases/uc_options"
	"github.com/MediStatTech/dashboard-service/pkg"
)

type Facade struct {
	pkg      *pkg.Facade
	UseCases *usecases.Facade
}

func New(pkg *pkg.Facade) (*Facade, error) {
	useCasesInstance := usecases.New(&uc_options.Options{
		Logger:    pkg.Logger,
	})

	return &Facade{
		pkg:      pkg,
		UseCases: useCasesInstance,
	}, nil
}
