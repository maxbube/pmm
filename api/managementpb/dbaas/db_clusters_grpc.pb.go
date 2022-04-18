// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: managementpb/dbaas/db_clusters.proto

package dbaasv1beta1

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

// DBClustersClient is the client API for DBClusters service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DBClustersClient interface {
	// ListDBClusters returns a list of DB clusters.
	ListDBClusters(ctx context.Context, in *ListDBClustersRequest, opts ...grpc.CallOption) (*ListDBClustersResponse, error)
	// RestartDBCluster restarts DB cluster.
	RestartDBCluster(ctx context.Context, in *RestartDBClusterRequest, opts ...grpc.CallOption) (*RestartDBClusterResponse, error)
	// DeleteDBCluster deletes DB cluster.
	DeleteDBCluster(ctx context.Context, in *DeleteDBClusterRequest, opts ...grpc.CallOption) (*DeleteDBClusterResponse, error)
}

type dBClustersClient struct {
	cc grpc.ClientConnInterface
}

func NewDBClustersClient(cc grpc.ClientConnInterface) DBClustersClient {
	return &dBClustersClient{cc}
}

func (c *dBClustersClient) ListDBClusters(ctx context.Context, in *ListDBClustersRequest, opts ...grpc.CallOption) (*ListDBClustersResponse, error) {
	out := new(ListDBClustersResponse)
	err := c.cc.Invoke(ctx, "/dbaas.v1beta1.DBClusters/ListDBClusters", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBClustersClient) RestartDBCluster(ctx context.Context, in *RestartDBClusterRequest, opts ...grpc.CallOption) (*RestartDBClusterResponse, error) {
	out := new(RestartDBClusterResponse)
	err := c.cc.Invoke(ctx, "/dbaas.v1beta1.DBClusters/RestartDBCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBClustersClient) DeleteDBCluster(ctx context.Context, in *DeleteDBClusterRequest, opts ...grpc.CallOption) (*DeleteDBClusterResponse, error) {
	out := new(DeleteDBClusterResponse)
	err := c.cc.Invoke(ctx, "/dbaas.v1beta1.DBClusters/DeleteDBCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DBClustersServer is the server API for DBClusters service.
// All implementations must embed UnimplementedDBClustersServer
// for forward compatibility
type DBClustersServer interface {
	// ListDBClusters returns a list of DB clusters.
	ListDBClusters(context.Context, *ListDBClustersRequest) (*ListDBClustersResponse, error)
	// RestartDBCluster restarts DB cluster.
	RestartDBCluster(context.Context, *RestartDBClusterRequest) (*RestartDBClusterResponse, error)
	// DeleteDBCluster deletes DB cluster.
	DeleteDBCluster(context.Context, *DeleteDBClusterRequest) (*DeleteDBClusterResponse, error)
	mustEmbedUnimplementedDBClustersServer()
}

// UnimplementedDBClustersServer must be embedded to have forward compatible implementations.
type UnimplementedDBClustersServer struct {
}

func (UnimplementedDBClustersServer) ListDBClusters(context.Context, *ListDBClustersRequest) (*ListDBClustersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDBClusters not implemented")
}
func (UnimplementedDBClustersServer) RestartDBCluster(context.Context, *RestartDBClusterRequest) (*RestartDBClusterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RestartDBCluster not implemented")
}
func (UnimplementedDBClustersServer) DeleteDBCluster(context.Context, *DeleteDBClusterRequest) (*DeleteDBClusterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDBCluster not implemented")
}
func (UnimplementedDBClustersServer) mustEmbedUnimplementedDBClustersServer() {}

// UnsafeDBClustersServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DBClustersServer will
// result in compilation errors.
type UnsafeDBClustersServer interface {
	mustEmbedUnimplementedDBClustersServer()
}

func RegisterDBClustersServer(s grpc.ServiceRegistrar, srv DBClustersServer) {
	s.RegisterService(&DBClusters_ServiceDesc, srv)
}

func _DBClusters_ListDBClusters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDBClustersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBClustersServer).ListDBClusters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dbaas.v1beta1.DBClusters/ListDBClusters",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBClustersServer).ListDBClusters(ctx, req.(*ListDBClustersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DBClusters_RestartDBCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RestartDBClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBClustersServer).RestartDBCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dbaas.v1beta1.DBClusters/RestartDBCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBClustersServer).RestartDBCluster(ctx, req.(*RestartDBClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DBClusters_DeleteDBCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDBClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBClustersServer).DeleteDBCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dbaas.v1beta1.DBClusters/DeleteDBCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBClustersServer).DeleteDBCluster(ctx, req.(*DeleteDBClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DBClusters_ServiceDesc is the grpc.ServiceDesc for DBClusters service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DBClusters_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dbaas.v1beta1.DBClusters",
	HandlerType: (*DBClustersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListDBClusters",
			Handler:    _DBClusters_ListDBClusters_Handler,
		},
		{
			MethodName: "RestartDBCluster",
			Handler:    _DBClusters_RestartDBCluster_Handler,
		},
		{
			MethodName: "DeleteDBCluster",
			Handler:    _DBClusters_DeleteDBCluster_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "managementpb/dbaas/db_clusters.proto",
}
