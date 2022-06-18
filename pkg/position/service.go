package position

import (
	"context"

	pb "github.com/cory-evans/Ichnaea/pkg/proto/position"
	"google.golang.org/grpc"
)

type Service struct {
	pb.UnimplementedPositionServiceServer
}

func (s *Service) GetPosition(ctx context.Context, req *pb.GetPositionRequest) (*pb.GetPositionResponse, error) {
	return &pb.GetPositionResponse{}, nil
}

func RegisterService(s *grpc.Server) {
	pb.RegisterPositionServiceServer(s, &Service{})
}
