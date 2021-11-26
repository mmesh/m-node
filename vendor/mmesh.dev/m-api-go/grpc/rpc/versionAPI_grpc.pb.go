// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package rpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	empty "mmesh.dev/m-api-go/grpc/common/empty"
	version "mmesh.dev/m-api-go/grpc/common/version"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// VersionAPIClient is the client API for VersionAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VersionAPIClient interface {
	Version(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*version.VersionResponse, error)
}

type versionAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewVersionAPIClient(cc grpc.ClientConnInterface) VersionAPIClient {
	return &versionAPIClient{cc}
}

func (c *versionAPIClient) Version(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*version.VersionResponse, error) {
	out := new(version.VersionResponse)
	err := c.cc.Invoke(ctx, "/version.VersionAPI/Version", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VersionAPIServer is the server API for VersionAPI service.
// All implementations must embed UnimplementedVersionAPIServer
// for forward compatibility
type VersionAPIServer interface {
	Version(context.Context, *empty.Empty) (*version.VersionResponse, error)
	mustEmbedUnimplementedVersionAPIServer()
}

// UnimplementedVersionAPIServer must be embedded to have forward compatible implementations.
type UnimplementedVersionAPIServer struct {
}

func (UnimplementedVersionAPIServer) Version(context.Context, *empty.Empty) (*version.VersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}
func (UnimplementedVersionAPIServer) mustEmbedUnimplementedVersionAPIServer() {}

// UnsafeVersionAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VersionAPIServer will
// result in compilation errors.
type UnsafeVersionAPIServer interface {
	mustEmbedUnimplementedVersionAPIServer()
}

func RegisterVersionAPIServer(s grpc.ServiceRegistrar, srv VersionAPIServer) {
	s.RegisterService(&VersionAPI_ServiceDesc, srv)
}

func _VersionAPI_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VersionAPIServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/version.VersionAPI/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VersionAPIServer).Version(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// VersionAPI_ServiceDesc is the grpc.ServiceDesc for VersionAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VersionAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "version.VersionAPI",
	HandlerType: (*VersionAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Version",
			Handler:    _VersionAPI_Version_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mmesh/protobuf/rpc/v1/versionAPI.proto",
}