// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: mmesh/protobuf/network/v1/rpc/network.proto

package rpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status1 "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	status "mmesh.dev/m-api-go/grpc/common/status"
	dns "mmesh.dev/m-api-go/grpc/network/dns"
	natProbe "mmesh.dev/m-api-go/grpc/network/mmnp/natProbe"
	register "mmesh.dev/m-api-go/grpc/network/mmnp/register"
	routing "mmesh.dev/m-api-go/grpc/network/mmnp/routing"
	mmsp "mmesh.dev/m-api-go/grpc/network/mmsp"
	controller "mmesh.dev/m-api-go/grpc/network/resources/controller"
	network "mmesh.dev/m-api-go/grpc/network/resources/network"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_mmesh_protobuf_network_v1_rpc_network_proto protoreflect.FileDescriptor

var file_mmesh_protobuf_network_v1_rpc_network_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x70, 0x63, 0x2f,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x1a, 0x27, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x76,
	0x31, 0x2f, 0x64, 0x6e, 0x73, 0x2f, 0x64, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x2d, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6d, 0x6e, 0x70, 0x2f,
	0x6e, 0x61, 0x74, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2d,
	0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6d, 0x6e, 0x70, 0x2f, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x6d,
	0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6d, 0x6e, 0x70, 0x2f, 0x72, 0x6f,
	0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x36, 0x6d, 0x6d, 0x65,
	0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x6d,
	0x6d, 0x73, 0x70, 0x2f, 0x6d, 0x6d, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3f,
	0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x66,
	0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x2c, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x9d, 0x04,
	0x0a, 0x09, 0x4e, 0x78, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x34, 0x0a, 0x08, 0x4e,
	0x41, 0x54, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x12, 0x12, 0x2e, 0x6e, 0x61, 0x74, 0x50, 0x72, 0x6f,
	0x62, 0x65, 0x2e, 0x4e, 0x41, 0x54, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x1a, 0x12, 0x2e, 0x6e, 0x61,
	0x74, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x2e, 0x4e, 0x41, 0x54, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x22,
	0x00, 0x12, 0x51, 0x0a, 0x10, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x45, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x1c, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x45,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x0e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x45, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x1c, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x4e,
	0x6f, 0x64, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x0c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x19, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52,
	0x65, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x1d, 0x0a, 0x03,
	0x44, 0x4e, 0x53, 0x12, 0x09, 0x2e, 0x64, 0x6e, 0x73, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x1a, 0x09,
	0x2e, 0x64, 0x6e, 0x73, 0x2e, 0x49, 0x50, 0x76, 0x34, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x07, 0x52,
	0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x12, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67,
	0x2e, 0x52, 0x54, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x72, 0x6f, 0x75,
	0x74, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x54, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x29, 0x0a, 0x03, 0x43, 0x6d, 0x64, 0x12, 0x0d, 0x2e, 0x6d,
	0x6d, 0x73, 0x70, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x1a, 0x0d, 0x2e, 0x6d, 0x6d,
	0x73, 0x70, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01,
	0x12, 0x32, 0x0a, 0x07, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x0d, 0x2e, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x1a, 0x16, 0x2e, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x13, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x0d, 0x2e, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x1a, 0x1f, 0x2e, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x22, 0x00, 0x42, 0x25, 0x5a,
	0x23, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x6d, 0x2d, 0x61, 0x70, 0x69,
	0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x2f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_mmesh_protobuf_network_v1_rpc_network_proto_goTypes = []interface{}{
	(*natProbe.NATProbe)(nil),              // 0: natProbe.NATProbe
	(*register.EndpointRegRequest)(nil),    // 1: register.EndpointRegRequest
	(*register.NodeRegRequest)(nil),        // 2: register.NodeRegRequest
	(*dns.Host)(nil),                       // 3: dns.Host
	(*routing.RTRequest)(nil),              // 4: routing.RTRequest
	(*mmsp.Payload)(nil),                   // 5: mmsp.Payload
	(*network.Node)(nil),                   // 6: network.Node
	(*register.EndpointRegResponse)(nil),   // 7: register.EndpointRegResponse
	(*register.NodeRegResponse)(nil),       // 8: register.NodeRegResponse
	(*dns.IPv4)(nil),                       // 9: dns.IPv4
	(*routing.RTResponse)(nil),             // 10: routing.RTResponse
	(*status.StatusResponse)(nil),          // 11: status.StatusResponse
	(*controller.FederationEndpoints)(nil), // 12: controller.FederationEndpoints
}
var file_mmesh_protobuf_network_v1_rpc_network_proto_depIdxs = []int32{
	0,  // 0: network.NxNetwork.NATProbe:input_type -> natProbe.NATProbe
	1,  // 1: network.NxNetwork.RegisterEndpoint:input_type -> register.EndpointRegRequest
	1,  // 2: network.NxNetwork.RemoveEndpoint:input_type -> register.EndpointRegRequest
	2,  // 3: network.NxNetwork.RegisterNode:input_type -> register.NodeRegRequest
	3,  // 4: network.NxNetwork.DNS:input_type -> dns.Host
	4,  // 5: network.NxNetwork.Routing:input_type -> routing.RTRequest
	5,  // 6: network.NxNetwork.Cmd:input_type -> mmsp.Payload
	6,  // 7: network.NxNetwork.Metrics:input_type -> network.Node
	6,  // 8: network.NxNetwork.FederationEndpoints:input_type -> network.Node
	0,  // 9: network.NxNetwork.NATProbe:output_type -> natProbe.NATProbe
	7,  // 10: network.NxNetwork.RegisterEndpoint:output_type -> register.EndpointRegResponse
	6,  // 11: network.NxNetwork.RemoveEndpoint:output_type -> network.Node
	8,  // 12: network.NxNetwork.RegisterNode:output_type -> register.NodeRegResponse
	9,  // 13: network.NxNetwork.DNS:output_type -> dns.IPv4
	10, // 14: network.NxNetwork.Routing:output_type -> routing.RTResponse
	5,  // 15: network.NxNetwork.Cmd:output_type -> mmsp.Payload
	11, // 16: network.NxNetwork.Metrics:output_type -> status.StatusResponse
	12, // 17: network.NxNetwork.FederationEndpoints:output_type -> controller.FederationEndpoints
	9,  // [9:18] is the sub-list for method output_type
	0,  // [0:9] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_mmesh_protobuf_network_v1_rpc_network_proto_init() }
func file_mmesh_protobuf_network_v1_rpc_network_proto_init() {
	if File_mmesh_protobuf_network_v1_rpc_network_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mmesh_protobuf_network_v1_rpc_network_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mmesh_protobuf_network_v1_rpc_network_proto_goTypes,
		DependencyIndexes: file_mmesh_protobuf_network_v1_rpc_network_proto_depIdxs,
	}.Build()
	File_mmesh_protobuf_network_v1_rpc_network_proto = out.File
	file_mmesh_protobuf_network_v1_rpc_network_proto_rawDesc = nil
	file_mmesh_protobuf_network_v1_rpc_network_proto_goTypes = nil
	file_mmesh_protobuf_network_v1_rpc_network_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// NxNetworkClient is the client API for NxNetwork service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NxNetworkClient interface {
	NATProbe(ctx context.Context, in *natProbe.NATProbe, opts ...grpc.CallOption) (*natProbe.NATProbe, error)
	RegisterEndpoint(ctx context.Context, in *register.EndpointRegRequest, opts ...grpc.CallOption) (*register.EndpointRegResponse, error)
	RemoveEndpoint(ctx context.Context, in *register.EndpointRegRequest, opts ...grpc.CallOption) (*network.Node, error)
	RegisterNode(ctx context.Context, in *register.NodeRegRequest, opts ...grpc.CallOption) (*register.NodeRegResponse, error)
	DNS(ctx context.Context, in *dns.Host, opts ...grpc.CallOption) (*dns.IPv4, error)
	Routing(ctx context.Context, opts ...grpc.CallOption) (NxNetwork_RoutingClient, error)
	Cmd(ctx context.Context, opts ...grpc.CallOption) (NxNetwork_CmdClient, error)
	Metrics(ctx context.Context, in *network.Node, opts ...grpc.CallOption) (*status.StatusResponse, error)
	FederationEndpoints(ctx context.Context, in *network.Node, opts ...grpc.CallOption) (*controller.FederationEndpoints, error)
}

type nxNetworkClient struct {
	cc grpc.ClientConnInterface
}

func NewNxNetworkClient(cc grpc.ClientConnInterface) NxNetworkClient {
	return &nxNetworkClient{cc}
}

func (c *nxNetworkClient) NATProbe(ctx context.Context, in *natProbe.NATProbe, opts ...grpc.CallOption) (*natProbe.NATProbe, error) {
	out := new(natProbe.NATProbe)
	err := c.cc.Invoke(ctx, "/network.NxNetwork/NATProbe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nxNetworkClient) RegisterEndpoint(ctx context.Context, in *register.EndpointRegRequest, opts ...grpc.CallOption) (*register.EndpointRegResponse, error) {
	out := new(register.EndpointRegResponse)
	err := c.cc.Invoke(ctx, "/network.NxNetwork/RegisterEndpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nxNetworkClient) RemoveEndpoint(ctx context.Context, in *register.EndpointRegRequest, opts ...grpc.CallOption) (*network.Node, error) {
	out := new(network.Node)
	err := c.cc.Invoke(ctx, "/network.NxNetwork/RemoveEndpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nxNetworkClient) RegisterNode(ctx context.Context, in *register.NodeRegRequest, opts ...grpc.CallOption) (*register.NodeRegResponse, error) {
	out := new(register.NodeRegResponse)
	err := c.cc.Invoke(ctx, "/network.NxNetwork/RegisterNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nxNetworkClient) DNS(ctx context.Context, in *dns.Host, opts ...grpc.CallOption) (*dns.IPv4, error) {
	out := new(dns.IPv4)
	err := c.cc.Invoke(ctx, "/network.NxNetwork/DNS", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nxNetworkClient) Routing(ctx context.Context, opts ...grpc.CallOption) (NxNetwork_RoutingClient, error) {
	stream, err := c.cc.NewStream(ctx, &_NxNetwork_serviceDesc.Streams[0], "/network.NxNetwork/Routing", opts...)
	if err != nil {
		return nil, err
	}
	x := &nxNetworkRoutingClient{stream}
	return x, nil
}

type NxNetwork_RoutingClient interface {
	Send(*routing.RTRequest) error
	Recv() (*routing.RTResponse, error)
	grpc.ClientStream
}

type nxNetworkRoutingClient struct {
	grpc.ClientStream
}

func (x *nxNetworkRoutingClient) Send(m *routing.RTRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *nxNetworkRoutingClient) Recv() (*routing.RTResponse, error) {
	m := new(routing.RTResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nxNetworkClient) Cmd(ctx context.Context, opts ...grpc.CallOption) (NxNetwork_CmdClient, error) {
	stream, err := c.cc.NewStream(ctx, &_NxNetwork_serviceDesc.Streams[1], "/network.NxNetwork/Cmd", opts...)
	if err != nil {
		return nil, err
	}
	x := &nxNetworkCmdClient{stream}
	return x, nil
}

type NxNetwork_CmdClient interface {
	Send(*mmsp.Payload) error
	Recv() (*mmsp.Payload, error)
	grpc.ClientStream
}

type nxNetworkCmdClient struct {
	grpc.ClientStream
}

func (x *nxNetworkCmdClient) Send(m *mmsp.Payload) error {
	return x.ClientStream.SendMsg(m)
}

func (x *nxNetworkCmdClient) Recv() (*mmsp.Payload, error) {
	m := new(mmsp.Payload)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nxNetworkClient) Metrics(ctx context.Context, in *network.Node, opts ...grpc.CallOption) (*status.StatusResponse, error) {
	out := new(status.StatusResponse)
	err := c.cc.Invoke(ctx, "/network.NxNetwork/Metrics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nxNetworkClient) FederationEndpoints(ctx context.Context, in *network.Node, opts ...grpc.CallOption) (*controller.FederationEndpoints, error) {
	out := new(controller.FederationEndpoints)
	err := c.cc.Invoke(ctx, "/network.NxNetwork/FederationEndpoints", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NxNetworkServer is the server API for NxNetwork service.
type NxNetworkServer interface {
	NATProbe(context.Context, *natProbe.NATProbe) (*natProbe.NATProbe, error)
	RegisterEndpoint(context.Context, *register.EndpointRegRequest) (*register.EndpointRegResponse, error)
	RemoveEndpoint(context.Context, *register.EndpointRegRequest) (*network.Node, error)
	RegisterNode(context.Context, *register.NodeRegRequest) (*register.NodeRegResponse, error)
	DNS(context.Context, *dns.Host) (*dns.IPv4, error)
	Routing(NxNetwork_RoutingServer) error
	Cmd(NxNetwork_CmdServer) error
	Metrics(context.Context, *network.Node) (*status.StatusResponse, error)
	FederationEndpoints(context.Context, *network.Node) (*controller.FederationEndpoints, error)
}

// UnimplementedNxNetworkServer can be embedded to have forward compatible implementations.
type UnimplementedNxNetworkServer struct {
}

func (*UnimplementedNxNetworkServer) NATProbe(context.Context, *natProbe.NATProbe) (*natProbe.NATProbe, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method NATProbe not implemented")
}
func (*UnimplementedNxNetworkServer) RegisterEndpoint(context.Context, *register.EndpointRegRequest) (*register.EndpointRegResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method RegisterEndpoint not implemented")
}
func (*UnimplementedNxNetworkServer) RemoveEndpoint(context.Context, *register.EndpointRegRequest) (*network.Node, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method RemoveEndpoint not implemented")
}
func (*UnimplementedNxNetworkServer) RegisterNode(context.Context, *register.NodeRegRequest) (*register.NodeRegResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method RegisterNode not implemented")
}
func (*UnimplementedNxNetworkServer) DNS(context.Context, *dns.Host) (*dns.IPv4, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method DNS not implemented")
}
func (*UnimplementedNxNetworkServer) Routing(NxNetwork_RoutingServer) error {
	return status1.Errorf(codes.Unimplemented, "method Routing not implemented")
}
func (*UnimplementedNxNetworkServer) Cmd(NxNetwork_CmdServer) error {
	return status1.Errorf(codes.Unimplemented, "method Cmd not implemented")
}
func (*UnimplementedNxNetworkServer) Metrics(context.Context, *network.Node) (*status.StatusResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method Metrics not implemented")
}
func (*UnimplementedNxNetworkServer) FederationEndpoints(context.Context, *network.Node) (*controller.FederationEndpoints, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method FederationEndpoints not implemented")
}

func RegisterNxNetworkServer(s *grpc.Server, srv NxNetworkServer) {
	s.RegisterService(&_NxNetwork_serviceDesc, srv)
}

func _NxNetwork_NATProbe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(natProbe.NATProbe)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NxNetworkServer).NATProbe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/network.NxNetwork/NATProbe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NxNetworkServer).NATProbe(ctx, req.(*natProbe.NATProbe))
	}
	return interceptor(ctx, in, info, handler)
}

func _NxNetwork_RegisterEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(register.EndpointRegRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NxNetworkServer).RegisterEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/network.NxNetwork/RegisterEndpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NxNetworkServer).RegisterEndpoint(ctx, req.(*register.EndpointRegRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NxNetwork_RemoveEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(register.EndpointRegRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NxNetworkServer).RemoveEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/network.NxNetwork/RemoveEndpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NxNetworkServer).RemoveEndpoint(ctx, req.(*register.EndpointRegRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NxNetwork_RegisterNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(register.NodeRegRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NxNetworkServer).RegisterNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/network.NxNetwork/RegisterNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NxNetworkServer).RegisterNode(ctx, req.(*register.NodeRegRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NxNetwork_DNS_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(dns.Host)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NxNetworkServer).DNS(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/network.NxNetwork/DNS",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NxNetworkServer).DNS(ctx, req.(*dns.Host))
	}
	return interceptor(ctx, in, info, handler)
}

func _NxNetwork_Routing_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(NxNetworkServer).Routing(&nxNetworkRoutingServer{stream})
}

type NxNetwork_RoutingServer interface {
	Send(*routing.RTResponse) error
	Recv() (*routing.RTRequest, error)
	grpc.ServerStream
}

type nxNetworkRoutingServer struct {
	grpc.ServerStream
}

func (x *nxNetworkRoutingServer) Send(m *routing.RTResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *nxNetworkRoutingServer) Recv() (*routing.RTRequest, error) {
	m := new(routing.RTRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _NxNetwork_Cmd_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(NxNetworkServer).Cmd(&nxNetworkCmdServer{stream})
}

type NxNetwork_CmdServer interface {
	Send(*mmsp.Payload) error
	Recv() (*mmsp.Payload, error)
	grpc.ServerStream
}

type nxNetworkCmdServer struct {
	grpc.ServerStream
}

func (x *nxNetworkCmdServer) Send(m *mmsp.Payload) error {
	return x.ServerStream.SendMsg(m)
}

func (x *nxNetworkCmdServer) Recv() (*mmsp.Payload, error) {
	m := new(mmsp.Payload)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _NxNetwork_Metrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(network.Node)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NxNetworkServer).Metrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/network.NxNetwork/Metrics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NxNetworkServer).Metrics(ctx, req.(*network.Node))
	}
	return interceptor(ctx, in, info, handler)
}

func _NxNetwork_FederationEndpoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(network.Node)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NxNetworkServer).FederationEndpoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/network.NxNetwork/FederationEndpoints",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NxNetworkServer).FederationEndpoints(ctx, req.(*network.Node))
	}
	return interceptor(ctx, in, info, handler)
}

var _NxNetwork_serviceDesc = grpc.ServiceDesc{
	ServiceName: "network.NxNetwork",
	HandlerType: (*NxNetworkServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NATProbe",
			Handler:    _NxNetwork_NATProbe_Handler,
		},
		{
			MethodName: "RegisterEndpoint",
			Handler:    _NxNetwork_RegisterEndpoint_Handler,
		},
		{
			MethodName: "RemoveEndpoint",
			Handler:    _NxNetwork_RemoveEndpoint_Handler,
		},
		{
			MethodName: "RegisterNode",
			Handler:    _NxNetwork_RegisterNode_Handler,
		},
		{
			MethodName: "DNS",
			Handler:    _NxNetwork_DNS_Handler,
		},
		{
			MethodName: "Metrics",
			Handler:    _NxNetwork_Metrics_Handler,
		},
		{
			MethodName: "FederationEndpoints",
			Handler:    _NxNetwork_FederationEndpoints_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Routing",
			Handler:       _NxNetwork_Routing_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Cmd",
			Handler:       _NxNetwork_Cmd_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "mmesh/protobuf/network/v1/rpc/network.proto",
}
