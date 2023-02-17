// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: reconditematter/v1/reconditematter.proto

package reconditematterv1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/reconditematter/go-reconditematter/gen/proto/go/reconditematter/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// ReconditeMatterServiceName is the fully-qualified name of the ReconditeMatterService service.
	ReconditeMatterServiceName = "reconditematter.v1.ReconditeMatterService"
)

// ReconditeMatterServiceClient is a client for the reconditematter.v1.ReconditeMatterService
// service.
type ReconditeMatterServiceClient interface {
	RandomNames(context.Context, *connect_go.Request[v1.RandomNamesRequest]) (*connect_go.Response[v1.RandomNamesResponse], error)
	FibonacciPoints(context.Context, *connect_go.Request[v1.FibonacciPointsRequest]) (*connect_go.Response[v1.FibonacciPointsResponse], error)
	FibonacciCell(context.Context, *connect_go.Request[v1.FibonacciCellRequest]) (*connect_go.Response[v1.FibonacciCellResponse], error)
	GeoHash(context.Context, *connect_go.Request[v1.GeoHashRequest]) (*connect_go.Response[v1.GeoHashResponse], error)
}

// NewReconditeMatterServiceClient constructs a client for the
// reconditematter.v1.ReconditeMatterService service. By default, it uses the Connect protocol with
// the binary Protobuf Codec, asks for gzipped responses, and sends uncompressed requests. To use
// the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewReconditeMatterServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) ReconditeMatterServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &reconditeMatterServiceClient{
		randomNames: connect_go.NewClient[v1.RandomNamesRequest, v1.RandomNamesResponse](
			httpClient,
			baseURL+"/reconditematter.v1.ReconditeMatterService/RandomNames",
			opts...,
		),
		fibonacciPoints: connect_go.NewClient[v1.FibonacciPointsRequest, v1.FibonacciPointsResponse](
			httpClient,
			baseURL+"/reconditematter.v1.ReconditeMatterService/FibonacciPoints",
			opts...,
		),
		fibonacciCell: connect_go.NewClient[v1.FibonacciCellRequest, v1.FibonacciCellResponse](
			httpClient,
			baseURL+"/reconditematter.v1.ReconditeMatterService/FibonacciCell",
			opts...,
		),
		geoHash: connect_go.NewClient[v1.GeoHashRequest, v1.GeoHashResponse](
			httpClient,
			baseURL+"/reconditematter.v1.ReconditeMatterService/GeoHash",
			opts...,
		),
	}
}

// reconditeMatterServiceClient implements ReconditeMatterServiceClient.
type reconditeMatterServiceClient struct {
	randomNames     *connect_go.Client[v1.RandomNamesRequest, v1.RandomNamesResponse]
	fibonacciPoints *connect_go.Client[v1.FibonacciPointsRequest, v1.FibonacciPointsResponse]
	fibonacciCell   *connect_go.Client[v1.FibonacciCellRequest, v1.FibonacciCellResponse]
	geoHash         *connect_go.Client[v1.GeoHashRequest, v1.GeoHashResponse]
}

// RandomNames calls reconditematter.v1.ReconditeMatterService.RandomNames.
func (c *reconditeMatterServiceClient) RandomNames(ctx context.Context, req *connect_go.Request[v1.RandomNamesRequest]) (*connect_go.Response[v1.RandomNamesResponse], error) {
	return c.randomNames.CallUnary(ctx, req)
}

// FibonacciPoints calls reconditematter.v1.ReconditeMatterService.FibonacciPoints.
func (c *reconditeMatterServiceClient) FibonacciPoints(ctx context.Context, req *connect_go.Request[v1.FibonacciPointsRequest]) (*connect_go.Response[v1.FibonacciPointsResponse], error) {
	return c.fibonacciPoints.CallUnary(ctx, req)
}

// FibonacciCell calls reconditematter.v1.ReconditeMatterService.FibonacciCell.
func (c *reconditeMatterServiceClient) FibonacciCell(ctx context.Context, req *connect_go.Request[v1.FibonacciCellRequest]) (*connect_go.Response[v1.FibonacciCellResponse], error) {
	return c.fibonacciCell.CallUnary(ctx, req)
}

// GeoHash calls reconditematter.v1.ReconditeMatterService.GeoHash.
func (c *reconditeMatterServiceClient) GeoHash(ctx context.Context, req *connect_go.Request[v1.GeoHashRequest]) (*connect_go.Response[v1.GeoHashResponse], error) {
	return c.geoHash.CallUnary(ctx, req)
}

// ReconditeMatterServiceHandler is an implementation of the
// reconditematter.v1.ReconditeMatterService service.
type ReconditeMatterServiceHandler interface {
	RandomNames(context.Context, *connect_go.Request[v1.RandomNamesRequest]) (*connect_go.Response[v1.RandomNamesResponse], error)
	FibonacciPoints(context.Context, *connect_go.Request[v1.FibonacciPointsRequest]) (*connect_go.Response[v1.FibonacciPointsResponse], error)
	FibonacciCell(context.Context, *connect_go.Request[v1.FibonacciCellRequest]) (*connect_go.Response[v1.FibonacciCellResponse], error)
	GeoHash(context.Context, *connect_go.Request[v1.GeoHashRequest]) (*connect_go.Response[v1.GeoHashResponse], error)
}

// NewReconditeMatterServiceHandler builds an HTTP handler from the service implementation. It
// returns the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewReconditeMatterServiceHandler(svc ReconditeMatterServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/reconditematter.v1.ReconditeMatterService/RandomNames", connect_go.NewUnaryHandler(
		"/reconditematter.v1.ReconditeMatterService/RandomNames",
		svc.RandomNames,
		opts...,
	))
	mux.Handle("/reconditematter.v1.ReconditeMatterService/FibonacciPoints", connect_go.NewUnaryHandler(
		"/reconditematter.v1.ReconditeMatterService/FibonacciPoints",
		svc.FibonacciPoints,
		opts...,
	))
	mux.Handle("/reconditematter.v1.ReconditeMatterService/FibonacciCell", connect_go.NewUnaryHandler(
		"/reconditematter.v1.ReconditeMatterService/FibonacciCell",
		svc.FibonacciCell,
		opts...,
	))
	mux.Handle("/reconditematter.v1.ReconditeMatterService/GeoHash", connect_go.NewUnaryHandler(
		"/reconditematter.v1.ReconditeMatterService/GeoHash",
		svc.GeoHash,
		opts...,
	))
	return "/reconditematter.v1.ReconditeMatterService/", mux
}

// UnimplementedReconditeMatterServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedReconditeMatterServiceHandler struct{}

func (UnimplementedReconditeMatterServiceHandler) RandomNames(context.Context, *connect_go.Request[v1.RandomNamesRequest]) (*connect_go.Response[v1.RandomNamesResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("reconditematter.v1.ReconditeMatterService.RandomNames is not implemented"))
}

func (UnimplementedReconditeMatterServiceHandler) FibonacciPoints(context.Context, *connect_go.Request[v1.FibonacciPointsRequest]) (*connect_go.Response[v1.FibonacciPointsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("reconditematter.v1.ReconditeMatterService.FibonacciPoints is not implemented"))
}

func (UnimplementedReconditeMatterServiceHandler) FibonacciCell(context.Context, *connect_go.Request[v1.FibonacciCellRequest]) (*connect_go.Response[v1.FibonacciCellResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("reconditematter.v1.ReconditeMatterService.FibonacciCell is not implemented"))
}

func (UnimplementedReconditeMatterServiceHandler) GeoHash(context.Context, *connect_go.Request[v1.GeoHashRequest]) (*connect_go.Response[v1.GeoHashResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("reconditematter.v1.ReconditeMatterService.GeoHash is not implemented"))
}
