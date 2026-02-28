package patient

import (
	s_options "github.com/MediStatTech/dashboard-service/internal/app/options"
	"github.com/MediStatTech/dashboard-service/internal/app/dashboard/usecases/patient_create"
	"github.com/MediStatTech/dashboard-service/internal/app/dashboard/usecases/patient_get"
	"github.com/MediStatTech/dashboard-service/internal/app/dashboard/usecases/patient_retrieve"
	"github.com/MediStatTech/dashboard-service/pkg"

	pb "github.com/MediStatTech/dashboard-client/pb/go/services/v1"
)

type Handler struct {
	pb.UnimplementedPatientServiceServer
	pkg      *pkg.Facade
	commands *Commands
	queries  *Queries
}

type Commands struct {
	PatientCreate *patient_create.Interactor
}

type Queries struct {
	PatientGet      *patient_get.Interactor
	PatientRetrieve *patient_retrieve.Interactor
}

func New(opts *s_options.Options) *Handler {
	return &Handler{
		pkg: opts.PKG,
		commands: &Commands{
			PatientCreate: opts.App.Dashboard.PatientCreate,
		},
		queries: &Queries{
			PatientGet:      opts.App.Dashboard.PatientGet,
			PatientRetrieve: opts.App.Dashboard.PatientRetrieve,
		},
	}
}
