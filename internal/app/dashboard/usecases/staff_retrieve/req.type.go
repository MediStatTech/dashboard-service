package staff_retrieve

import "github.com/MediStatTech/dashboard-service/internal/app/dashboard/domain"

type Request struct {
	StaffID string
}

type Response struct {
	Staff domain.Staff
}
