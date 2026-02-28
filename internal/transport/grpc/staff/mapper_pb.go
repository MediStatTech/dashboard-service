package staff

import (
	pb_models "github.com/MediStatTech/dashboard-client/pb/go/models/v1"
	"github.com/MediStatTech/dashboard-service/internal/app/dashboard/domain"
)

func staffToPb(s domain.Staff) *pb_models.Staff_Read {
	return &pb_models.Staff_Read{
		StaffId:        s.StaffID,
		FirstName:      s.FirstName,
		LastName:       s.LastName,
		SelfieUrl:      s.SelfieURL,
		SelfieThumbUrl: s.SelfieThumbURL,
		Status:         s.Status,
		Email:          s.Email,
		Position: &pb_models.Position{
			PositionId: s.Position.PositionID,
			Name:       s.Position.Name,
		},
	}
}
