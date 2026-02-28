package usecases

import "github.com/MediStatTech/dashboard-service/internal/app/dashboard/usecases/uc_options"

type Facade struct {
}

func New(o *uc_options.Options) *Facade {
	return &Facade{}
}
