package measurement

import (
	pb "github.com/MediStatTech/dashboard-client/pb/go/services/v1"
	s_options "github.com/MediStatTech/dashboard-service/internal/app/options"
	"github.com/MediStatTech/dashboard-service/internal/app/dashboard/usecases/measurement_get"
	"github.com/MediStatTech/dashboard-service/internal/app/dashboard/usecases/measurement_history_get"
	"github.com/MediStatTech/dashboard-service/pkg"
)

type Handler struct {
	pb.UnimplementedMeasurementServiceServer
	pkg     *pkg.Facade
	queries *Queries
}

type Queries struct {
	MeasurementGet        *measurement_get.Interactor
	MeasurementHistoryGet *measurement_history_get.Interactor
}

func New(opts *s_options.Options) *Handler {
	return &Handler{
		pkg: opts.PKG,
		queries: &Queries{
			MeasurementGet:        opts.App.Dashboard.MeasurementGet,
			MeasurementHistoryGet: opts.App.Dashboard.MeasurementHistoryGet,
		},
	}
}
