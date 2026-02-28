package diseas

import (
	"context"

	pb_models "github.com/MediStatTech/dashboard-client/pb/go/models/v1"
	pb "github.com/MediStatTech/dashboard-client/pb/go/services/v1"
	"github.com/MediStatTech/dashboard-service/internal/app/dashboard/usecases/diseas_get"
)

func (h *Handler) DiseasGet(
	ctx context.Context,
	_ *pb.DiseasGetRequest,
) (*pb.DiseasGetReply, error) {
	resp, err := h.queries.DiseasGet.Execute(ctx, diseas_get.Request{})
	if err != nil {
		return nil, err
	}

	diseases := make([]*pb_models.Diseas, 0, len(resp.Diseases))
	for _, d := range resp.Diseases {
		diseases = append(diseases, diseasToPb(d))
	}

	return &pb.DiseasGetReply{
		Diseases: diseases,
	}, nil
}
