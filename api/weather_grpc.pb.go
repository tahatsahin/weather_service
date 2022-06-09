// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.1
// source: weather.proto

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// WeatherServiceClient is the client API for WeatherService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WeatherServiceClient interface {
	QueryWeather(ctx context.Context, in *WeatherRequest, opts ...grpc.CallOption) (WeatherService_QueryWeatherClient, error)
	ListCities(ctx context.Context, in *ListCitiesRequest, opts ...grpc.CallOption) (*ListCitiesResponse, error)
}

type weatherServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWeatherServiceClient(cc grpc.ClientConnInterface) WeatherServiceClient {
	return &weatherServiceClient{cc}
}

func (c *weatherServiceClient) QueryWeather(ctx context.Context, in *WeatherRequest, opts ...grpc.CallOption) (WeatherService_QueryWeatherClient, error) {
	stream, err := c.cc.NewStream(ctx, &WeatherService_ServiceDesc.Streams[0], "/WeatherService/QueryWeather", opts...)
	if err != nil {
		return nil, err
	}
	x := &weatherServiceQueryWeatherClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WeatherService_QueryWeatherClient interface {
	Recv() (*WeatherResponse, error)
	grpc.ClientStream
}

type weatherServiceQueryWeatherClient struct {
	grpc.ClientStream
}

func (x *weatherServiceQueryWeatherClient) Recv() (*WeatherResponse, error) {
	m := new(WeatherResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *weatherServiceClient) ListCities(ctx context.Context, in *ListCitiesRequest, opts ...grpc.CallOption) (*ListCitiesResponse, error) {
	out := new(ListCitiesResponse)
	err := c.cc.Invoke(ctx, "/WeatherService/ListCities", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WeatherServiceServer is the server API for WeatherService service.
// All implementations must embed UnimplementedWeatherServiceServer
// for forward compatibility
type WeatherServiceServer interface {
	QueryWeather(*WeatherRequest, WeatherService_QueryWeatherServer) error
	ListCities(context.Context, *ListCitiesRequest) (*ListCitiesResponse, error)
	mustEmbedUnimplementedWeatherServiceServer()
}

// UnimplementedWeatherServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWeatherServiceServer struct {
}

func (UnimplementedWeatherServiceServer) QueryWeather(*WeatherRequest, WeatherService_QueryWeatherServer) error {
	return status.Errorf(codes.Unimplemented, "method QueryWeather not implemented")
}
func (UnimplementedWeatherServiceServer) ListCities(context.Context, *ListCitiesRequest) (*ListCitiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCities not implemented")
}
func (UnimplementedWeatherServiceServer) mustEmbedUnimplementedWeatherServiceServer() {}

// UnsafeWeatherServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WeatherServiceServer will
// result in compilation errors.
type UnsafeWeatherServiceServer interface {
	mustEmbedUnimplementedWeatherServiceServer()
}

func RegisterWeatherServiceServer(s grpc.ServiceRegistrar, srv WeatherServiceServer) {
	s.RegisterService(&WeatherService_ServiceDesc, srv)
}

func _WeatherService_QueryWeather_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WeatherRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(WeatherServiceServer).QueryWeather(m, &weatherServiceQueryWeatherServer{stream})
}

type WeatherService_QueryWeatherServer interface {
	Send(*WeatherResponse) error
	grpc.ServerStream
}

type weatherServiceQueryWeatherServer struct {
	grpc.ServerStream
}

func (x *weatherServiceQueryWeatherServer) Send(m *WeatherResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _WeatherService_ListCities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCitiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WeatherServiceServer).ListCities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/WeatherService/ListCities",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WeatherServiceServer).ListCities(ctx, req.(*ListCitiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WeatherService_ServiceDesc is the grpc.ServiceDesc for WeatherService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WeatherService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "WeatherService",
	HandlerType: (*WeatherServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListCities",
			Handler:    _WeatherService_ListCities_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "QueryWeather",
			Handler:       _WeatherService_QueryWeather_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "weather.proto",
}
