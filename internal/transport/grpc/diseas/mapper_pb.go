package diseas

import (
	pb_models "github.com/MediStatTech/dashboard-client/pb/go/models/v1"
	"github.com/MediStatTech/dashboard-service/internal/app/dashboard/domain"
)

func diseasToPb(d domain.Diseas) *pb_models.Diseas {
	return &pb_models.Diseas{
		DiseasId: d.DiseasID,
		Name:     d.Name,
		Code:     d.Code,
	}
}
