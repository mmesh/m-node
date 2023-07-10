// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.2
// source: mmesh/protobuf/network/v1/nac/nac.proto

package nac

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	auth "mmesh.dev/m-api-go/grpc/resources/iam/auth"
	topology "mmesh.dev/m-api-go/grpc/resources/topology"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RouterAdmissionResult int32

const (
	RouterAdmissionResult_ROUTER_UNAUTHORIZED RouterAdmissionResult = 0
	RouterAdmissionResult_ROUTER_AUTHORIZED   RouterAdmissionResult = 1
)

// Enum value maps for RouterAdmissionResult.
var (
	RouterAdmissionResult_name = map[int32]string{
		0: "ROUTER_UNAUTHORIZED",
		1: "ROUTER_AUTHORIZED",
	}
	RouterAdmissionResult_value = map[string]int32{
		"ROUTER_UNAUTHORIZED": 0,
		"ROUTER_AUTHORIZED":   1,
	}
)

func (x RouterAdmissionResult) Enum() *RouterAdmissionResult {
	p := new(RouterAdmissionResult)
	*p = x
	return p
}

func (x RouterAdmissionResult) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RouterAdmissionResult) Descriptor() protoreflect.EnumDescriptor {
	return file_mmesh_protobuf_network_v1_nac_nac_proto_enumTypes[0].Descriptor()
}

func (RouterAdmissionResult) Type() protoreflect.EnumType {
	return &file_mmesh_protobuf_network_v1_nac_nac_proto_enumTypes[0]
}

func (x RouterAdmissionResult) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RouterAdmissionResult.Descriptor instead.
func (RouterAdmissionResult) EnumDescriptor() ([]byte, []int) {
	return file_mmesh_protobuf_network_v1_nac_nac_proto_rawDescGZIP(), []int{0}
}

type NetworkAdmissionResult int32

const (
	NetworkAdmissionResult_UNAUTHORIZED     NetworkAdmissionResult = 0
	NetworkAdmissionResult_AUTHORIZED       NetworkAdmissionResult = 1
	NetworkAdmissionResult_SERVICE_DISABLED NetworkAdmissionResult = 2
)

// Enum value maps for NetworkAdmissionResult.
var (
	NetworkAdmissionResult_name = map[int32]string{
		0: "UNAUTHORIZED",
		1: "AUTHORIZED",
		2: "SERVICE_DISABLED",
	}
	NetworkAdmissionResult_value = map[string]int32{
		"UNAUTHORIZED":     0,
		"AUTHORIZED":       1,
		"SERVICE_DISABLED": 2,
	}
)

func (x NetworkAdmissionResult) Enum() *NetworkAdmissionResult {
	p := new(NetworkAdmissionResult)
	*p = x
	return p
}

func (x NetworkAdmissionResult) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NetworkAdmissionResult) Descriptor() protoreflect.EnumDescriptor {
	return file_mmesh_protobuf_network_v1_nac_nac_proto_enumTypes[1].Descriptor()
}

func (NetworkAdmissionResult) Type() protoreflect.EnumType {
	return &file_mmesh_protobuf_network_v1_nac_nac_proto_enumTypes[1]
}

func (x NetworkAdmissionResult) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NetworkAdmissionResult.Descriptor instead.
func (NetworkAdmissionResult) EnumDescriptor() ([]byte, []int) {
	return file_mmesh_protobuf_network_v1_nac_nac_proto_rawDescGZIP(), []int{1}
}

type RouterAdmissionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RouterToken string `protobuf:"bytes,1,opt,name=routerToken,proto3" json:"routerToken,omitempty"` // b64-encoded RouterTokenPayload struct
}

func (x *RouterAdmissionRequest) Reset() {
	*x = RouterAdmissionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_network_v1_nac_nac_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouterAdmissionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouterAdmissionRequest) ProtoMessage() {}

func (x *RouterAdmissionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_network_v1_nac_nac_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouterAdmissionRequest.ProtoReflect.Descriptor instead.
func (*RouterAdmissionRequest) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_network_v1_nac_nac_proto_rawDescGZIP(), []int{0}
}

func (x *RouterAdmissionRequest) GetRouterToken() string {
	if x != nil {
		return x.RouterToken
	}
	return ""
}

type RouterAdmissionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result  RouterAdmissionResult `protobuf:"varint,1,opt,name=result,proto3,enum=nac.RouterAdmissionResult" json:"result,omitempty"`
	AuthKey *auth.AuthKey         `protobuf:"bytes,11,opt,name=authKey,proto3" json:"authKey,omitempty"`
	// string authSecret = 12;
	Router *topology.Router `protobuf:"bytes,21,opt,name=router,proto3" json:"router,omitempty"`
}

func (x *RouterAdmissionResponse) Reset() {
	*x = RouterAdmissionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_network_v1_nac_nac_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouterAdmissionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouterAdmissionResponse) ProtoMessage() {}

func (x *RouterAdmissionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_network_v1_nac_nac_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouterAdmissionResponse.ProtoReflect.Descriptor instead.
func (*RouterAdmissionResponse) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_network_v1_nac_nac_proto_rawDescGZIP(), []int{1}
}

func (x *RouterAdmissionResponse) GetResult() RouterAdmissionResult {
	if x != nil {
		return x.Result
	}
	return RouterAdmissionResult_ROUTER_UNAUTHORIZED
}

func (x *RouterAdmissionResponse) GetAuthKey() *auth.AuthKey {
	if x != nil {
		return x.AuthKey
	}
	return nil
}

func (x *RouterAdmissionResponse) GetRouter() *topology.Router {
	if x != nil {
		return x.Router
	}
	return nil
}

type RouterRegRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Router *topology.Router `protobuf:"bytes,1,opt,name=router,proto3" json:"router,omitempty"`
}

func (x *RouterRegRequest) Reset() {
	*x = RouterRegRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_network_v1_nac_nac_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouterRegRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouterRegRequest) ProtoMessage() {}

func (x *RouterRegRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_network_v1_nac_nac_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouterRegRequest.ProtoReflect.Descriptor instead.
func (*RouterRegRequest) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_network_v1_nac_nac_proto_rawDescGZIP(), []int{2}
}

func (x *RouterRegRequest) GetRouter() *topology.Router {
	if x != nil {
		return x.Router
	}
	return nil
}

type RouterRegResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LocationID string `protobuf:"bytes,11,opt,name=locationID,proto3" json:"locationID,omitempty"`
}

func (x *RouterRegResponse) Reset() {
	*x = RouterRegResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_network_v1_nac_nac_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouterRegResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouterRegResponse) ProtoMessage() {}

func (x *RouterRegResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_network_v1_nac_nac_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouterRegResponse.ProtoReflect.Descriptor instead.
func (*RouterRegResponse) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_network_v1_nac_nac_proto_rawDescGZIP(), []int{3}
}

func (x *RouterRegResponse) GetLocationID() string {
	if x != nil {
		return x.LocationID
	}
	return ""
}

type NetworkAdmissionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeToken string `protobuf:"bytes,1,opt,name=nodeToken,proto3" json:"nodeToken,omitempty"` // b64-encoded NodeTokenPayload struct
}

func (x *NetworkAdmissionRequest) Reset() {
	*x = NetworkAdmissionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_network_v1_nac_nac_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkAdmissionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkAdmissionRequest) ProtoMessage() {}

func (x *NetworkAdmissionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_network_v1_nac_nac_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkAdmissionRequest.ProtoReflect.Descriptor instead.
func (*NetworkAdmissionRequest) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_network_v1_nac_nac_proto_rawDescGZIP(), []int{4}
}

func (x *NetworkAdmissionRequest) GetNodeToken() string {
	if x != nil {
		return x.NodeToken
	}
	return ""
}

type NetworkAdmissionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result  NetworkAdmissionResult `protobuf:"varint,1,opt,name=result,proto3,enum=nac.NetworkAdmissionResult" json:"result,omitempty"`
	AuthKey *auth.AuthKey          `protobuf:"bytes,11,opt,name=authKey,proto3" json:"authKey,omitempty"`
	// string authSecret = 12;
	Node *topology.Node `protobuf:"bytes,21,opt,name=node,proto3" json:"node,omitempty"`
}

func (x *NetworkAdmissionResponse) Reset() {
	*x = NetworkAdmissionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_network_v1_nac_nac_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkAdmissionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkAdmissionResponse) ProtoMessage() {}

func (x *NetworkAdmissionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_network_v1_nac_nac_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkAdmissionResponse.ProtoReflect.Descriptor instead.
func (*NetworkAdmissionResponse) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_network_v1_nac_nac_proto_rawDescGZIP(), []int{5}
}

func (x *NetworkAdmissionResponse) GetResult() NetworkAdmissionResult {
	if x != nil {
		return x.Result
	}
	return NetworkAdmissionResult_UNAUTHORIZED
}

func (x *NetworkAdmissionResponse) GetAuthKey() *auth.AuthKey {
	if x != nil {
		return x.AuthKey
	}
	return nil
}

func (x *NetworkAdmissionResponse) GetNode() *topology.Node {
	if x != nil {
		return x.Node
	}
	return nil
}

type NodeRegRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Node *topology.Node `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
}

func (x *NodeRegRequest) Reset() {
	*x = NodeRegRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_network_v1_nac_nac_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeRegRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeRegRequest) ProtoMessage() {}

func (x *NodeRegRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_network_v1_nac_nac_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeRegRequest.ProtoReflect.Descriptor instead.
func (*NodeRegRequest) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_network_v1_nac_nac_proto_rawDescGZIP(), []int{6}
}

func (x *NodeRegRequest) GetNode() *topology.Node {
	if x != nil {
		return x.Node
	}
	return nil
}

type NodeRegResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoutingDomain *RoutingDomain   `protobuf:"bytes,11,opt,name=routingDomain,proto3" json:"routingDomain,omitempty"`
	NetworkPolicy *topology.Policy `protobuf:"bytes,41,opt,name=networkPolicy,proto3" json:"networkPolicy,omitempty"`
}

func (x *NodeRegResponse) Reset() {
	*x = NodeRegResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_network_v1_nac_nac_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeRegResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeRegResponse) ProtoMessage() {}

func (x *NodeRegResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_network_v1_nac_nac_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeRegResponse.ProtoReflect.Descriptor instead.
func (*NodeRegResponse) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_network_v1_nac_nac_proto_rawDescGZIP(), []int{7}
}

func (x *NodeRegResponse) GetRoutingDomain() *RoutingDomain {
	if x != nil {
		return x.RoutingDomain
	}
	return nil
}

func (x *NodeRegResponse) GetNetworkPolicy() *topology.Policy {
	if x != nil {
		return x.NetworkPolicy
	}
	return nil
}

type EndpointRegRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeReq  *topology.NodeReq  `protobuf:"bytes,11,opt,name=nodeReq,proto3" json:"nodeReq,omitempty"`
	Endpoint *topology.Endpoint `protobuf:"bytes,21,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	Priority int32              `protobuf:"varint,41,opt,name=priority,proto3" json:"priority,omitempty"`
}

func (x *EndpointRegRequest) Reset() {
	*x = EndpointRegRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_network_v1_nac_nac_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndpointRegRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndpointRegRequest) ProtoMessage() {}

func (x *EndpointRegRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_network_v1_nac_nac_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndpointRegRequest.ProtoReflect.Descriptor instead.
func (*EndpointRegRequest) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_network_v1_nac_nac_proto_rawDescGZIP(), []int{8}
}

func (x *EndpointRegRequest) GetNodeReq() *topology.NodeReq {
	if x != nil {
		return x.NodeReq
	}
	return nil
}

func (x *EndpointRegRequest) GetEndpoint() *topology.Endpoint {
	if x != nil {
		return x.Endpoint
	}
	return nil
}

func (x *EndpointRegRequest) GetPriority() int32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

type EndpointRegResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IPv4 string `protobuf:"bytes,1,opt,name=IPv4,proto3" json:"IPv4,omitempty"`
	IPv6 string `protobuf:"bytes,2,opt,name=IPv6,proto3" json:"IPv6,omitempty"`
}

func (x *EndpointRegResponse) Reset() {
	*x = EndpointRegResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_network_v1_nac_nac_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndpointRegResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndpointRegResponse) ProtoMessage() {}

func (x *EndpointRegResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_network_v1_nac_nac_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndpointRegResponse.ProtoReflect.Descriptor instead.
func (*EndpointRegResponse) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_network_v1_nac_nac_proto_rawDescGZIP(), []int{9}
}

func (x *EndpointRegResponse) GetIPv4() string {
	if x != nil {
		return x.IPv4
	}
	return ""
}

func (x *EndpointRegResponse) GetIPv6() string {
	if x != nil {
		return x.IPv6
	}
	return ""
}

var File_mmesh_protobuf_network_v1_nac_nac_proto protoreflect.FileDescriptor

var file_mmesh_protobuf_network_v1_nac_nac_proto_rawDesc = []byte{
	0x0a, 0x27, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x61, 0x63, 0x2f,
	0x6e, 0x61, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6e, 0x61, 0x63, 0x1a, 0x2f,
	0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x61, 0x6d, 0x2f,
	0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x31, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x70,
	0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x2f, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x74, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x31, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x74, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x2f, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x76,
	0x31, 0x2f, 0x6e, 0x61, 0x63, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x3a, 0x0a, 0x16, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x41, 0x64, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a,
	0x0b, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22,
	0xa0, 0x01, 0x0a, 0x17, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x6e, 0x61,
	0x63, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12,
	0x27, 0x0a, 0x07, 0x61, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x52,
	0x07, 0x61, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x12, 0x28, 0x0a, 0x06, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x72, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x74, 0x6f, 0x70, 0x6f, 0x6c,
	0x6f, 0x67, 0x79, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x52, 0x06, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x72, 0x22, 0x3c, 0x0a, 0x10, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x52, 0x65, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x74, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67,
	0x79, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x52, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72,
	0x22, 0x33, 0x0a, 0x11, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x52, 0x65, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x44, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x44, 0x22, 0x37, 0x0a, 0x17, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x41, 0x64, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x9c,
	0x01, 0x0a, 0x18, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x41, 0x64, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x6e, 0x61,
	0x63, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x41, 0x64, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x27, 0x0a, 0x07, 0x61, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79,
	0x52, 0x07, 0x61, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x12, 0x22, 0x0a, 0x04, 0x6e, 0x6f, 0x64,
	0x65, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x74, 0x6f, 0x70, 0x6f, 0x6c, 0x6f,
	0x67, 0x79, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x22, 0x34, 0x0a,
	0x0e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x22, 0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x74, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x6e,
	0x6f, 0x64, 0x65, 0x22, 0x83, 0x01, 0x0a, 0x0f, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x0d, 0x72, 0x6f, 0x75, 0x74, 0x69,
	0x6e, 0x67, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x6e, 0x61, 0x63, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x44, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x52, 0x0d, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x44, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x12, 0x36, 0x0a, 0x0d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x18, 0x29, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x74, 0x6f, 0x70, 0x6f, 0x6c,
	0x6f, 0x67, 0x79, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x0d, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x22, 0x8d, 0x01, 0x0a, 0x12, 0x45, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x2b, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x74, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x2e, 0x4e, 0x6f, 0x64,
	0x65, 0x52, 0x65, 0x71, 0x52, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x12, 0x2e, 0x0a,
	0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x74, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x29, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x22, 0x3d, 0x0a, 0x13, 0x45, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x49, 0x50, 0x76, 0x34, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x49, 0x50, 0x76, 0x34, 0x12, 0x12, 0x0a, 0x04, 0x49, 0x50, 0x76, 0x36, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x49, 0x50, 0x76, 0x36, 0x2a, 0x47, 0x0a, 0x15, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x12, 0x17, 0x0a, 0x13, 0x52, 0x4f, 0x55, 0x54, 0x45, 0x52, 0x5f, 0x55, 0x4e, 0x41, 0x55,
	0x54, 0x48, 0x4f, 0x52, 0x49, 0x5a, 0x45, 0x44, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x52, 0x4f,
	0x55, 0x54, 0x45, 0x52, 0x5f, 0x41, 0x55, 0x54, 0x48, 0x4f, 0x52, 0x49, 0x5a, 0x45, 0x44, 0x10,
	0x01, 0x2a, 0x50, 0x0a, 0x16, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x41, 0x64, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x10, 0x0a, 0x0c, 0x55,
	0x4e, 0x41, 0x55, 0x54, 0x48, 0x4f, 0x52, 0x49, 0x5a, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0e, 0x0a,
	0x0a, 0x41, 0x55, 0x54, 0x48, 0x4f, 0x52, 0x49, 0x5a, 0x45, 0x44, 0x10, 0x01, 0x12, 0x14, 0x0a,
	0x10, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x44, 0x49, 0x53, 0x41, 0x42, 0x4c, 0x45,
	0x44, 0x10, 0x02, 0x42, 0x25, 0x5a, 0x23, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x64, 0x65, 0x76,
	0x2f, 0x6d, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x6e, 0x61, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_mmesh_protobuf_network_v1_nac_nac_proto_rawDescOnce sync.Once
	file_mmesh_protobuf_network_v1_nac_nac_proto_rawDescData = file_mmesh_protobuf_network_v1_nac_nac_proto_rawDesc
)

func file_mmesh_protobuf_network_v1_nac_nac_proto_rawDescGZIP() []byte {
	file_mmesh_protobuf_network_v1_nac_nac_proto_rawDescOnce.Do(func() {
		file_mmesh_protobuf_network_v1_nac_nac_proto_rawDescData = protoimpl.X.CompressGZIP(file_mmesh_protobuf_network_v1_nac_nac_proto_rawDescData)
	})
	return file_mmesh_protobuf_network_v1_nac_nac_proto_rawDescData
}

var file_mmesh_protobuf_network_v1_nac_nac_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_mmesh_protobuf_network_v1_nac_nac_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_mmesh_protobuf_network_v1_nac_nac_proto_goTypes = []interface{}{
	(RouterAdmissionResult)(0),       // 0: nac.RouterAdmissionResult
	(NetworkAdmissionResult)(0),      // 1: nac.NetworkAdmissionResult
	(*RouterAdmissionRequest)(nil),   // 2: nac.RouterAdmissionRequest
	(*RouterAdmissionResponse)(nil),  // 3: nac.RouterAdmissionResponse
	(*RouterRegRequest)(nil),         // 4: nac.RouterRegRequest
	(*RouterRegResponse)(nil),        // 5: nac.RouterRegResponse
	(*NetworkAdmissionRequest)(nil),  // 6: nac.NetworkAdmissionRequest
	(*NetworkAdmissionResponse)(nil), // 7: nac.NetworkAdmissionResponse
	(*NodeRegRequest)(nil),           // 8: nac.NodeRegRequest
	(*NodeRegResponse)(nil),          // 9: nac.NodeRegResponse
	(*EndpointRegRequest)(nil),       // 10: nac.EndpointRegRequest
	(*EndpointRegResponse)(nil),      // 11: nac.EndpointRegResponse
	(*auth.AuthKey)(nil),             // 12: auth.AuthKey
	(*topology.Router)(nil),          // 13: topology.Router
	(*topology.Node)(nil),            // 14: topology.Node
	(*RoutingDomain)(nil),            // 15: nac.RoutingDomain
	(*topology.Policy)(nil),          // 16: topology.Policy
	(*topology.NodeReq)(nil),         // 17: topology.NodeReq
	(*topology.Endpoint)(nil),        // 18: topology.Endpoint
}
var file_mmesh_protobuf_network_v1_nac_nac_proto_depIdxs = []int32{
	0,  // 0: nac.RouterAdmissionResponse.result:type_name -> nac.RouterAdmissionResult
	12, // 1: nac.RouterAdmissionResponse.authKey:type_name -> auth.AuthKey
	13, // 2: nac.RouterAdmissionResponse.router:type_name -> topology.Router
	13, // 3: nac.RouterRegRequest.router:type_name -> topology.Router
	1,  // 4: nac.NetworkAdmissionResponse.result:type_name -> nac.NetworkAdmissionResult
	12, // 5: nac.NetworkAdmissionResponse.authKey:type_name -> auth.AuthKey
	14, // 6: nac.NetworkAdmissionResponse.node:type_name -> topology.Node
	14, // 7: nac.NodeRegRequest.node:type_name -> topology.Node
	15, // 8: nac.NodeRegResponse.routingDomain:type_name -> nac.RoutingDomain
	16, // 9: nac.NodeRegResponse.networkPolicy:type_name -> topology.Policy
	17, // 10: nac.EndpointRegRequest.nodeReq:type_name -> topology.NodeReq
	18, // 11: nac.EndpointRegRequest.endpoint:type_name -> topology.Endpoint
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_mmesh_protobuf_network_v1_nac_nac_proto_init() }
func file_mmesh_protobuf_network_v1_nac_nac_proto_init() {
	if File_mmesh_protobuf_network_v1_nac_nac_proto != nil {
		return
	}
	file_mmesh_protobuf_network_v1_nac_routing_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_mmesh_protobuf_network_v1_nac_nac_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouterAdmissionRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mmesh_protobuf_network_v1_nac_nac_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouterAdmissionResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mmesh_protobuf_network_v1_nac_nac_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouterRegRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mmesh_protobuf_network_v1_nac_nac_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouterRegResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mmesh_protobuf_network_v1_nac_nac_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkAdmissionRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mmesh_protobuf_network_v1_nac_nac_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkAdmissionResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mmesh_protobuf_network_v1_nac_nac_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeRegRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mmesh_protobuf_network_v1_nac_nac_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeRegResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mmesh_protobuf_network_v1_nac_nac_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndpointRegRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mmesh_protobuf_network_v1_nac_nac_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndpointRegResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mmesh_protobuf_network_v1_nac_nac_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mmesh_protobuf_network_v1_nac_nac_proto_goTypes,
		DependencyIndexes: file_mmesh_protobuf_network_v1_nac_nac_proto_depIdxs,
		EnumInfos:         file_mmesh_protobuf_network_v1_nac_nac_proto_enumTypes,
		MessageInfos:      file_mmesh_protobuf_network_v1_nac_nac_proto_msgTypes,
	}.Build()
	File_mmesh_protobuf_network_v1_nac_nac_proto = out.File
	file_mmesh_protobuf_network_v1_nac_nac_proto_rawDesc = nil
	file_mmesh_protobuf_network_v1_nac_nac_proto_goTypes = nil
	file_mmesh_protobuf_network_v1_nac_nac_proto_depIdxs = nil
}
