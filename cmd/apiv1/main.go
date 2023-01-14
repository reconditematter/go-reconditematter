package main

import (
	"context"
	"log"
	"net/http"

	"github.com/bufbuild/connect-go"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/reconditematter/go-reconditematter/gen/proto/go/reconditematter/v1"
	"github.com/reconditematter/go-reconditematter/gen/proto/go/reconditematter/v1/reconditematterv1connect"
	"github.com/reconditematter/go-reconditematter/pkg/randomnames"
)

func main() {
	listen := "127.0.0.1:8080"
	server := &ReconditeMatterServer{}
	mux := http.NewServeMux()
	path, handler := reconditematterv1connect.NewReconditeMatterServiceHandler(server)
	mux.Handle(path, handler)
	log.Println("apiv1 started:", listen, path)
	if err := http.ListenAndServe(listen, h2c.NewHandler(mux, &http2.Server{})); err != nil {
		log.Fatalf("failed ListenAndServe: %w\n", err)
	}
}

type ReconditeMatterServer struct {
	reconditematterv1connect.UnimplementedReconditeMatterServiceHandler
}

func (s *ReconditeMatterServer) RandomNames(
	ctx context.Context,
	req *connect.Request[reconditematterv1.RandomNamesRequest],
) (*connect.Response[reconditematterv1.RandomNamesResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	count := req.Msg.GetCount()
	result, err := randomnames.Generate(uint(count))
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
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
	return connect.NewResponse(resp), nil
}

func (s *ReconditeMatterServer) FibonacciPoints(
	ctx context.Context,
	req *connect.Request[reconditematterv1.FibonacciPointsRequest],
) (*connect.Response[reconditematterv1.FibonacciPointsResponse], error) {
	return nil, nil
}
