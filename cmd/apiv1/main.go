package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/reconditematter/go-reconditematter/gen/proto/go/reconditematter/v1"
	"github.com/reconditematter/go-reconditematter/pkg/randomnames"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	listen := "127.0.0.1:8080"
	listener, err := net.Listen("tcp", listen)
	if err != nil {
		return fmt.Errorf("failed to listen on %q: %w", listen, err)
	}
	log.Println("listening on:", listen)

	server := grpc.NewServer()
	reconditematterv1.RegisterReconditeMatterServiceServer(server, &reconditeMatterServiceServer{})
	if err := server.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve gRPC server: %w", err)
	}

	return nil
}

type reconditeMatterServiceServer struct {
	reconditematterv1.UnimplementedReconditeMatterServiceServer
}

func (s *reconditeMatterServiceServer) RandomNames(
	ctx context.Context,
	req *reconditematterv1.RandomNamesRequest) (*reconditematterv1.RandomNamesResponse, error) {
	count := req.GetCount()
	result, err := randomnames.Generate(uint(count))
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "`count` cannot be greater than 1000")
	}

	resp := &reconditematterv1.RandomNamesResponse{}
	resp.Count = uint32(len(result))
	resp.Names = make([]*reconditematterv1.HumanName, len(result))
	for k := range resp.Names {
		resp.Names[k] = &reconditematterv1.HumanName{
			Family: result[k].Family,
			Given:  result[k].Given,
			Gender: result[k].Gender,
		}
	}
	return resp, nil
}
