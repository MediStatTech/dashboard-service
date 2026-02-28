package staff

import (
	s_options "github.com/MediStatTech/dashboard-service/internal/app/options"
	"github.com/MediStatTech/dashboard-service/internal/app/dashboard/usecases/staff_retrieve"
	"github.com/MediStatTech/dashboard-service/pkg"

	pb "github.com/MediStatTech/dashboard-client/pb/go/services/v1"
)

type Handler struct {
	pb.UnimplementedStaffServiceServer
	pkg     *pkg.Facade
	queries *Queries
}

type Queries struct {
	StaffRetrieve *staff_retrieve.Interactor
}

func New(opts *s_options.Options) *Handler {
	return &Handler{
		pkg: opts.PKG,
		queries: &Queries{
			StaffRetrieve: opts.App.Dashboard.StaffRetrieve,
		},
	}
}
