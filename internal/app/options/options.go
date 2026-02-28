package options

import (
	"github.com/MediStatTech/dashboard-service/internal/app"
	"github.com/MediStatTech/dashboard-service/pkg"
)

type Options struct {
	App *app.Facade
	PKG *pkg.Facade
}
