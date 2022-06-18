// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package position

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

// PositionServiceClient is the client API for PositionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PositionServiceClient interface {
	GetPosition(ctx context.Context, in *GetPositionRequest, opts ...grpc.CallOption) (*GetPositionResponse, error)
}

type positionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPositionServiceClient(cc grpc.ClientConnInterface) PositionServiceClient {
	return &positionServiceClient{cc}
}

func (c *positionServiceClient) GetPosition(ctx context.Context, in *GetPositionRequest, opts ...grpc.CallOption) (*GetPositionResponse, error) {
	out := new(GetPositionResponse)
	err := c.cc.Invoke(ctx, "/position.PositionService/GetPosition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PositionServiceServer is the server API for PositionService service.
// All implementations should embed UnimplementedPositionServiceServer
// for forward compatibility
type PositionServiceServer interface {
	GetPosition(context.Context, *GetPositionRequest) (*GetPositionResponse, error)
}

// UnimplementedPositionServiceServer should be embedded to have forward compatible implementations.
type UnimplementedPositionServiceServer struct {
}

func (UnimplementedPositionServiceServer) GetPosition(context.Context, *GetPositionRequest) (*GetPositionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPosition not implemented")
}

// UnsafePositionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PositionServiceServer will
// result in compilation errors.
type UnsafePositionServiceServer interface {
	mustEmbedUnimplementedPositionServiceServer()
}

func RegisterPositionServiceServer(s grpc.ServiceRegistrar, srv PositionServiceServer) {
	s.RegisterService(&PositionService_ServiceDesc, srv)
}

func _PositionService_GetPosition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPositionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PositionServiceServer).GetPosition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/position.PositionService/GetPosition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PositionServiceServer).GetPosition(ctx, req.(*GetPositionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PositionService_ServiceDesc is the grpc.ServiceDesc for PositionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PositionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "position.PositionService",
	HandlerType: (*PositionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPosition",
			Handler:    _PositionService_GetPosition_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/proto/position/position.proto",
}