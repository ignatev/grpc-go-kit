package vault

import (
	"golang.org/x/net/context"
	grpctransport "github.com/go-kit/kit/transport/grpc"
)

type grpcServer struct {
	hash grpctransport.Handler
	validate grpctransport.Handler
}

func (s *grpcServer) Hash(сtx context.Context, r *pb.HashRequest) (*pb.HashResponse, error) {
	_, resp, err := s.hash.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.HashResponse), nil
}
func (s *grpcServer) Validate(ctx context.Context, r *pb.ValidateRequest) (*p.ValidateResponse, error) {
	_, resp, err := s.validate.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ValidateREsponse), nil
}