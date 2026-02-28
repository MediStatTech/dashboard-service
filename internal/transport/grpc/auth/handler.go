package auth

import (
	s_options "github.com/MediStatTech/dashboard-service/internal/app/options"
	"github.com/MediStatTech/dashboard-service/internal/app/dashboard/usecases/sign_in"
	"github.com/MediStatTech/dashboard-service/pkg"

	pb "github.com/MediStatTech/dashboard-client/pb/go/services/v1"
)

type Handler struct {
	pb.UnimplementedAuthServiceServer
	pkg      *pkg.Facade
	commands *Commands
}

type Commands struct {
	SignIn *sign_in.Interactor
}

func New(opts *s_options.Options) *Handler {
	return &Handler{
		pkg: opts.PKG,
		commands: &Commands{
			SignIn: opts.App.Dashboard.SignIn,
		},
	}
}
