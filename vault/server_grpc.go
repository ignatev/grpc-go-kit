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

func EncodeGRPCHashRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(hashRequest)
	return &pb.HashRequest{Password: req.Password}, nil
}

func DecodeGRPCHashRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.HashRequest)
	return hashRequest{Password: req.Password}, nil
}

func EncodeGRPCHashResponse(ctx context.Context, r interface{}) (interface{}, error) {
	res := r.(hashResponse)
	return &pb.HashResponse{Hash: res.Hash, Err: res.Err}, nil
}

func DecodeGRPCHashResponse(ctx context.Context, r interface{}) (interface{}, error) {
	res := r.(*pb.HashResponse)
	return hashResponse{Hash: res.Hash, Err: res.Err}, nil
}

func EncodeGRPCValidateRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(validateRequest)
	return &pb.ValidateRequest{Password: req.Password, Hash: req.Hash}, nil
}

func DecodeGRPCValidateRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.ValidateRequest)
	return validateRequest{Password: req.Password, Hash: req.Hash}, nil
}

func EncodeGRPCValidateResponse(ctx context.Context, r interface{}) (interface{}, error) {
	res := r.(validateResponse)
	return &pb.ValidateResponse{Valid: res.Valid}, nil
}

func DecodeGRPCValidateResponse(ctx context.Context, r interface{}) (interface{}, error) {
	res := r.(*pb.ValidateResponse)
	return validateResponse{Valid: res.Valid}, nil
}