// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: managementpb/service/service.proto

package servicev1beta1

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
	MgmtService_ListServices_FullMethodName = "/service.v1beta1.MgmtService/ListServices"
)

// MgmtServiceClient is the client API for MgmtService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MgmtServiceClient interface {
	// ListServices returns a list of Services with a rich set of properties.
	ListServices(ctx context.Context, in *ListServiceRequest, opts ...grpc.CallOption) (*ListServiceResponse, error)
}

type mgmtServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMgmtServiceClient(cc grpc.ClientConnInterface) MgmtServiceClient {
	return &mgmtServiceClient{cc}
}

func (c *mgmtServiceClient) ListServices(ctx context.Context, in *ListServiceRequest, opts ...grpc.CallOption) (*ListServiceResponse, error) {
	out := new(ListServiceResponse)
	err := c.cc.Invoke(ctx, MgmtService_ListServices_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MgmtServiceServer is the server API for MgmtService service.
// All implementations must embed UnimplementedMgmtServiceServer
// for forward compatibility
type MgmtServiceServer interface {
	// ListServices returns a list of Services with a rich set of properties.
	ListServices(context.Context, *ListServiceRequest) (*ListServiceResponse, error)
	mustEmbedUnimplementedMgmtServiceServer()
}

// UnimplementedMgmtServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMgmtServiceServer struct{}

func (UnimplementedMgmtServiceServer) ListServices(context.Context, *ListServiceRequest) (*ListServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListServices not implemented")
}
func (UnimplementedMgmtServiceServer) mustEmbedUnimplementedMgmtServiceServer() {}

// UnsafeMgmtServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MgmtServiceServer will
// result in compilation errors.
type UnsafeMgmtServiceServer interface {
	mustEmbedUnimplementedMgmtServiceServer()
}

func RegisterMgmtServiceServer(s grpc.ServiceRegistrar, srv MgmtServiceServer) {
	s.RegisterService(&MgmtService_ServiceDesc, srv)
}

func _MgmtService_ListServices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MgmtServiceServer).ListServices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MgmtService_ListServices_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MgmtServiceServer).ListServices(ctx, req.(*ListServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MgmtService_ServiceDesc is the grpc.ServiceDesc for MgmtService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MgmtService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.v1beta1.MgmtService",
	HandlerType: (*MgmtServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListServices",
			Handler:    _MgmtService_ListServices_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "managementpb/service/service.proto",
}