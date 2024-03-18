// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: demo.proto

package pb

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

const (
	DemoService_Hi_FullMethodName    = "/demo.pb.DemoService/Hi"
	DemoService_Watch_FullMethodName = "/demo.pb.DemoService/Watch"
)

// DemoServiceClient is the client API for DemoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DemoServiceClient interface {
	Hi(ctx context.Context, in *HiRequest, opts ...grpc.CallOption) (*HiResponse, error)
	Watch(ctx context.Context, in *WatchRequest, opts ...grpc.CallOption) (DemoService_WatchClient, error)
}

type demoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDemoServiceClient(cc grpc.ClientConnInterface) DemoServiceClient {
	return &demoServiceClient{cc}
}

func (c *demoServiceClient) Hi(ctx context.Context, in *HiRequest, opts ...grpc.CallOption) (*HiResponse, error) {
	out := new(HiResponse)
	err := c.cc.Invoke(ctx, DemoService_Hi_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *demoServiceClient) Watch(ctx context.Context, in *WatchRequest, opts ...grpc.CallOption) (DemoService_WatchClient, error) {
	stream, err := c.cc.NewStream(ctx, &DemoService_ServiceDesc.Streams[0], DemoService_Watch_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &demoServiceWatchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DemoService_WatchClient interface {
	Recv() (*WatchResponse, error)
	grpc.ClientStream
}

type demoServiceWatchClient struct {
	grpc.ClientStream
}

func (x *demoServiceWatchClient) Recv() (*WatchResponse, error) {
	m := new(WatchResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DemoServiceServer is the server API for DemoService service.
// All implementations should embed UnimplementedDemoServiceServer
// for forward compatibility
type DemoServiceServer interface {
	Hi(context.Context, *HiRequest) (*HiResponse, error)
	Watch(*WatchRequest, DemoService_WatchServer) error
}

// UnimplementedDemoServiceServer should be embedded to have forward compatible implementations.
type UnimplementedDemoServiceServer struct {
}

func (UnimplementedDemoServiceServer) Hi(context.Context, *HiRequest) (*HiResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hi not implemented")
}
func (UnimplementedDemoServiceServer) Watch(*WatchRequest, DemoService_WatchServer) error {
	return status.Errorf(codes.Unimplemented, "method Watch not implemented")
}

// UnsafeDemoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DemoServiceServer will
// result in compilation errors.
type UnsafeDemoServiceServer interface {
	mustEmbedUnimplementedDemoServiceServer()
}

func RegisterDemoServiceServer(s grpc.ServiceRegistrar, srv DemoServiceServer) {
	s.RegisterService(&DemoService_ServiceDesc, srv)
}

func _DemoService_Hi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DemoServiceServer).Hi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DemoService_Hi_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DemoServiceServer).Hi(ctx, req.(*HiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DemoService_Watch_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WatchRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DemoServiceServer).Watch(m, &demoServiceWatchServer{stream})
}

type DemoService_WatchServer interface {
	Send(*WatchResponse) error
	grpc.ServerStream
}

type demoServiceWatchServer struct {
	grpc.ServerStream
}

func (x *demoServiceWatchServer) Send(m *WatchResponse) error {
	return x.ServerStream.SendMsg(m)
}

// DemoService_ServiceDesc is the grpc.ServiceDesc for DemoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DemoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "demo.pb.DemoService",
	HandlerType: (*DemoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hi",
			Handler:    _DemoService_Hi_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Watch",
			Handler:       _DemoService_Watch_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "demo.proto",
}
