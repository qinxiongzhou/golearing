package string_service

import (
	"ch45rpc/grpcstreamdemo/pb"
	"context"
	"errors"
	"io"
	"log"
	"strings"
)

const (
	StrMaxSize = 1024
)

// Service errors
var (
	ErrMaxSize = errors.New("maximum size of 1024 bytes exceeded")
	ErrStrValue = errors.New("maximum size of 1024 bytes exceeded")
)

type StringService struct{
	pb.UnsafeStringServiceServer
}

func (s *StringService) LotsOfServerStream(request *pb.StringRequest, server pb.StringService_LotsOfServerStreamServer) error {
	response := pb.StringResponse{Ret: request.A + request.B}
	for i := 0; i < 10; i++ {
		server.Send(&response)
	}
	return nil
}

func (s *StringService) LotsOfClientStream(server pb.StringService_LotsOfClientStreamServer) error {
	var params []string
	for {
		in, err := server.Recv()
		if err == io.EOF {
			server.SendAndClose(&pb.StringResponse{Ret: strings.Join(params, "")})
			return nil
		}
		if err != nil {
			log.Printf("failed to recv: %v", err)
			return err
		}
		params = append(params, in.A, in.B)
	}
}

func (s *StringService) LotsOfServerAndClientStream(server pb.StringService_LotsOfServerAndClientStreamServer) error {
	for {
		in, err := server.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Printf("failed to recv %v", err)
			return err
		}
		server.Send(&pb.StringResponse{Ret: in.A + in.B})
	}
	return nil
}

func (s *StringService) Concat(ctx context.Context, req *pb.StringRequest) (*pb.StringResponse, error) {
	if len(req.A)+len(req.B) > StrMaxSize {
		response := pb.StringResponse{Ret: ""}
		return &response, nil
	}
	response := pb.StringResponse{Ret: req.A + req.B}
	return &response, nil
}
