package main

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/bufbuild/connect-go"
	"github.com/google/uuid"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/reconditematter/go-reconditematter/gen/proto/go/reconditematter/v1"
	"github.com/reconditematter/go-reconditematter/gen/proto/go/reconditematter/v1/reconditematterv1connect"
	"github.com/reconditematter/go-reconditematter/pkg/geo"
	"github.com/reconditematter/go-reconditematter/pkg/randomnames"
)

func main() {
	listen := "127.0.0.1:8080"
	server := &ReconditeMatterServer{}
	mux := http.NewServeMux()
	path, handler := reconditematterv1connect.NewReconditeMatterServiceHandler(server,
		connect.WithInterceptors(NewInterceptor()))
	mux.Handle(path, handler)
	log.Println("apiv1", listen, path)
	if err := http.ListenAndServe(listen, h2c.NewHandler(mux, &http2.Server{})); err != nil {
		log.Fatalf("failed ListenAndServe: %w\n", err)
	}
}

func NewInterceptor() connect.UnaryInterceptorFunc {
	const ReqID = "X-Request-Id"
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			reqid := req.Header().Get(ReqID)
			if reqid == "" {
				reqid = uuid.New().String()
				req.Header().Set(ReqID, reqid)
				log.Printf("%s %s %q\n", reqid, req.Spec().Procedure, req.Any())
			}

			resp, err := next(ctx, req)

			if err != nil {
				log.Printf("%s %q\n", reqid, err.Error())
				var cerr *connect.Error
				if errors.As(err, &cerr) {
					cerr.Meta().Set(ReqID, reqid)
					err = cerr
				}
			} else {
				log.Printf("%s OK\n", reqid)
				if resp != nil {
					resp.Header().Set(ReqID, reqid)
				}
			}
			return resp, err
		})
	}
	return connect.UnaryInterceptorFunc(interceptor)
}

type ReconditeMatterServer struct {
	reconditematterv1connect.UnimplementedReconditeMatterServiceHandler
}

func (s *ReconditeMatterServer) RandomNames(
	ctx context.Context,
	req *connect.Request[reconditematterv1.RandomNamesRequest],
) (*connect.Response[reconditematterv1.RandomNamesResponse], error) {
	resp := &reconditematterv1.RandomNamesResponse{}
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	const maxCount = 1000
	count := req.Msg.GetCount()
	if count > maxCount {
		return nil, connect.NewError(
			connect.CodeInvalidArgument,
			errors.New("RandomNames: invalid count: "+req.Msg.String()))
	}

	result := randomnames.Generate(uint(count))

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
	resp := &reconditematterv1.FibonacciPointsResponse{}
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	const maxCount = 10001
	count := req.Msg.GetCount()
	if count > maxCount {
		return nil, connect.NewError(
			connect.CodeInvalidArgument,
			errors.New("FibonacciPoints: invalid count: "+req.Msg.String()))
	}

	result := geo.FibonacciPoints(uint(count) / 2)

	resp.Count = uint32(len(result))
	resp.Points = make([]*reconditematterv1.GeoPoint, len(result))
	for k := range resp.Points {
		resp.Points[k] = &reconditematterv1.GeoPoint{Lat: result[k].Lat, Lon: result[k].Lon}
	}
	return connect.NewResponse(resp), nil
}

func (s *ReconditeMatterServer) FibonacciCell(
	ctx context.Context,
	req *connect.Request[reconditematterv1.FibonacciCellRequest],
) (*connect.Response[reconditematterv1.FibonacciCellResponse], error) {
	resp := &reconditematterv1.FibonacciCellResponse{}
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	const maxCount = 10000
	count := req.Msg.GetCount()
	if count > maxCount {
		return nil, connect.NewError(
			connect.CodeInvalidArgument,
			errors.New("FibonacciCell: invalid count: "+req.Msg.String()))
	}
	origin := req.Msg.GetOrigin()
	fibo := geo.Point{Lat: origin.GetLat(), Lon: origin.GetLon()}
	latok, lonok := fibo.Valid()
	if !(latok && lonok) {
		return nil, connect.NewError(
			connect.CodeInvalidArgument,
			errors.New("FibonacciCell: invalid origin: "+req.Msg.String()))
	}

	result := geo.FibonacciCell(uint(count), fibo)

	resp.Count = uint32(len(result))
	resp.Origin = origin
	resp.Points = make([]*reconditematterv1.GeoPoint, len(result))
	for k := range resp.Points {
		resp.Points[k] = &reconditematterv1.GeoPoint{Lat: result[k].Lat, Lon: result[k].Lon}
	}
	return connect.NewResponse(resp), nil
}
