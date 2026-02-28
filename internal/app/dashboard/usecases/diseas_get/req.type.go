package diseas_get

import "github.com/MediStatTech/dashboard-service/internal/app/dashboard/domain"

type Request struct{}

type Response struct {
	Diseases []domain.Diseas
}
