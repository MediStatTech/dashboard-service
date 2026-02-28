package diseas

import (
	s_options "github.com/MediStatTech/dashboard-service/internal/app/options"
	"github.com/MediStatTech/dashboard-service/internal/app/dashboard/usecases/diseas_get"
	"github.com/MediStatTech/dashboard-service/pkg"

	pb "github.com/MediStatTech/dashboard-client/pb/go/services/v1"
)

type Handler struct {
	pb.UnimplementedDiseasServiceServer
	pkg     *pkg.Facade
	queries *Queries
}

type Queries struct {
	DiseasGet *diseas_get.Interactor
}

func New(opts *s_options.Options) *Handler {
	return &Handler{
		pkg: opts.PKG,
		queries: &Queries{
			DiseasGet: opts.App.Dashboard.DiseasGet,
		},
	}
}
