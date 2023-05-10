// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.0
// source: mmesh/protobuf/network/v1/routing/rib.proto

package routing

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	nac "mmesh.dev/m-api-go/grpc/network/nac"
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

type RIBDataType int32

const (
	RIBDataType_ROUTING_SCOPE RIBDataType = 0
	RIBDataType_ROUTER        RIBDataType = 21
	RIBDataType_RELAY         RIBDataType = 31
	RIBDataType_ROUTING_TABLE RIBDataType = 41
	RIBDataType_POLICY        RIBDataType = 101
)

// Enum value maps for RIBDataType.
var (
	RIBDataType_name = map[int32]string{
		0:   "ROUTING_SCOPE",
		21:  "ROUTER",
		31:  "RELAY",
		41:  "ROUTING_TABLE",
		101: "POLICY",
	}
	RIBDataType_value = map[string]int32{
		"ROUTING_SCOPE": 0,
		"ROUTER":        21,
		"RELAY":         31,
		"ROUTING_TABLE": 41,
		"POLICY":        101,
	}
)

func (x RIBDataType) Enum() *RIBDataType {
	p := new(RIBDataType)
	*p = x
	return p
}

func (x RIBDataType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RIBDataType) Descriptor() protoreflect.EnumDescriptor {
	return file_mmesh_protobuf_network_v1_routing_rib_proto_enumTypes[0].Descriptor()
}

func (RIBDataType) Type() protoreflect.EnumType {
	return &file_mmesh_protobuf_network_v1_routing_rib_proto_enumTypes[0]
}

func (x RIBDataType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RIBDataType.Descriptor instead.
func (RIBDataType) EnumDescriptor() ([]byte, []int) {
	return file_mmesh_protobuf_network_v1_routing_rib_proto_rawDescGZIP(), []int{0}
}

type AddressFamily int32

const (
	AddressFamily_IP4 AddressFamily = 0
	AddressFamily_IP6 AddressFamily = 1
)

// Enum value maps for AddressFamily.
var (
	AddressFamily_name = map[int32]string{
		0: "IP4",
		1: "IP6",
	}
	AddressFamily_value = map[string]int32{
		"IP4": 0,
		"IP6": 1,
	}
)

func (x AddressFamily) Enum() *AddressFamily {
	p := new(AddressFamily)
	*p = x
	return p
}

func (x AddressFamily) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AddressFamily) Descriptor() protoreflect.EnumDescriptor {
	return file_mmesh_protobuf_network_v1_routing_rib_proto_enumTypes[1].Descriptor()
}

func (AddressFamily) Type() protoreflect.EnumType {
	return &file_mmesh_protobuf_network_v1_routing_rib_proto_enumTypes[1]
}

func (x AddressFamily) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AddressFamily.Descriptor instead.
func (AddressFamily) EnumDescriptor() ([]byte, []int) {
	return file_mmesh_protobuf_network_v1_routing_rib_proto_rawDescGZIP(), []int{1}
}

type RouteType int32

const (
	RouteType_CONNECTED RouteType = 0
	RouteType_STATIC    RouteType = 1
	RouteType_DYNAMIC   RouteType = 2
)

// Enum value maps for RouteType.
var (
	RouteType_name = map[int32]string{
		0: "CONNECTED",
		1: "STATIC",
		2: "DYNAMIC",
	}
	RouteType_value = map[string]int32{
		"CONNECTED": 0,
		"STATIC":    1,
		"DYNAMIC":   2,
	}
)

func (x RouteType) Enum() *RouteType {
	p := new(RouteType)
	*p = x
	return p
}

func (x RouteType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RouteType) Descriptor() protoreflect.EnumDescriptor {
	return file_mmesh_protobuf_network_v1_routing_rib_proto_enumTypes[2].Descriptor()
}

func (RouteType) Type() protoreflect.EnumType {
	return &file_mmesh_protobuf_network_v1_routing_rib_proto_enumTypes[2]
}

func (x RouteType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RouteType.Descriptor instead.
func (RouteType) EnumDescriptor() ([]byte, []int) {
	return file_mmesh_protobuf_network_v1_routing_rib_proto_rawDescGZIP(), []int{2}
}

// RIBData is received by nodes by mqtt
type RIBData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type RIBDataType `protobuf:"varint,1,opt,name=type,proto3,enum=routing.RIBDataType" json:"type,omitempty"`
	// type == ROUTING_SCOPE
	Scope nac.RoutingScope `protobuf:"varint,11,opt,name=scope,proto3,enum=nac.RoutingScope" json:"scope,omitempty"`
	// type == ROUTER
	Router *NetHop `protobuf:"bytes,21,opt,name=router,proto3" json:"router,omitempty"`
	// type == RELAY
	Relay *NetHop `protobuf:"bytes,31,opt,name=relay,proto3" json:"relay,omitempty"`
	// type == ROUTING_TABLE
	RoutingTable map[string]*RoutingEntry `protobuf:"bytes,41,rep,name=routingTable,proto3" json:"routingTable,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // map[addr]*RoutingEntry
	// type == POLICY
	Policy map[string]*topology.Policy `protobuf:"bytes,101,rep,name=policy,proto3" json:"policy,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // map[subnetID]*Policy
}

func (x *RIBData) Reset() {
	*x = RIBData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_network_v1_routing_rib_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RIBData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RIBData) ProtoMessage() {}

func (x *RIBData) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_network_v1_routing_rib_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RIBData.ProtoReflect.Descriptor instead.
func (*RIBData) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_network_v1_routing_rib_proto_rawDescGZIP(), []int{0}
}

func (x *RIBData) GetType() RIBDataType {
	if x != nil {
		return x.Type
	}
	return RIBDataType_ROUTING_SCOPE
}

func (x *RIBData) GetScope() nac.RoutingScope {
	if x != nil {
		return x.Scope
	}
	return nac.RoutingScope(0)
}

func (x *RIBData) GetRouter() *NetHop {
	if x != nil {
		return x.Router
	}
	return nil
}

func (x *RIBData) GetRelay() *NetHop {
	if x != nil {
		return x.Relay
	}
	return nil
}

func (x *RIBData) GetRoutingTable() map[string]*RoutingEntry {
	if x != nil {
		return x.RoutingTable
	}
	return nil
}

func (x *RIBData) GetPolicy() map[string]*topology.Policy {
	if x != nil {
		return x.Policy
	}
	return nil
}

type RoutingEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubnetID      string             `protobuf:"bytes,1,opt,name=subnetID,proto3" json:"subnetID,omitempty"`
	AddressFamily AddressFamily      `protobuf:"varint,11,opt,name=addressFamily,proto3,enum=routing.AddressFamily" json:"addressFamily,omitempty"`
	Type          RouteType          `protobuf:"varint,21,opt,name=type,proto3,enum=routing.RouteType" json:"type,omitempty"`
	Gw            map[string]*NetHop `protobuf:"bytes,31,rep,name=gw,proto3" json:"gw,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // map[p2pHostID]*NetHop
	DNSName       string             `protobuf:"bytes,41,opt,name=DNSName,proto3" json:"DNSName,omitempty"`
}

func (x *RoutingEntry) Reset() {
	*x = RoutingEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_network_v1_routing_rib_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoutingEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoutingEntry) ProtoMessage() {}

func (x *RoutingEntry) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_network_v1_routing_rib_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoutingEntry.ProtoReflect.Descriptor instead.
func (*RoutingEntry) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_network_v1_routing_rib_proto_rawDescGZIP(), []int{1}
}

func (x *RoutingEntry) GetSubnetID() string {
	if x != nil {
		return x.SubnetID
	}
	return ""
}

func (x *RoutingEntry) GetAddressFamily() AddressFamily {
	if x != nil {
		return x.AddressFamily
	}
	return AddressFamily_IP4
}

func (x *RoutingEntry) GetType() RouteType {
	if x != nil {
		return x.Type
	}
	return RouteType_CONNECTED
}

func (x *RoutingEntry) GetGw() map[string]*NetHop {
	if x != nil {
		return x.Gw
	}
	return nil
}

func (x *RoutingEntry) GetDNSName() string {
	if x != nil {
		return x.DNSName
	}
	return ""
}

type NetHop struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubnetID    string   `protobuf:"bytes,1,opt,name=subnetID,proto3" json:"subnetID,omitempty"`
	P2PHostID   string   `protobuf:"bytes,11,opt,name=P2PHostID,proto3" json:"P2PHostID,omitempty"`
	MAddrs      []string `protobuf:"bytes,12,rep,name=MAddrs,proto3" json:"MAddrs,omitempty"`
	Priority    int32    `protobuf:"varint,21,opt,name=priority,proto3" json:"priority,omitempty"`
	Connections int32    `protobuf:"varint,31,opt,name=connections,proto3" json:"connections,omitempty"`
	LastSeen    int64    `protobuf:"varint,101,opt,name=lastSeen,proto3" json:"lastSeen,omitempty"`
}

func (x *NetHop) Reset() {
	*x = NetHop{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_network_v1_routing_rib_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetHop) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetHop) ProtoMessage() {}

func (x *NetHop) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_network_v1_routing_rib_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetHop.ProtoReflect.Descriptor instead.
func (*NetHop) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_network_v1_routing_rib_proto_rawDescGZIP(), []int{2}
}

func (x *NetHop) GetSubnetID() string {
	if x != nil {
		return x.SubnetID
	}
	return ""
}

func (x *NetHop) GetP2PHostID() string {
	if x != nil {
		return x.P2PHostID
	}
	return ""
}

func (x *NetHop) GetMAddrs() []string {
	if x != nil {
		return x.MAddrs
	}
	return nil
}

func (x *NetHop) GetPriority() int32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

func (x *NetHop) GetConnections() int32 {
	if x != nil {
		return x.Connections
	}
	return 0
}

func (x *NetHop) GetLastSeen() int64 {
	if x != nil {
		return x.LastSeen
	}
	return 0
}

// RIB is the structure nodes manage in memory to route traffic
type RIB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountID     string                   `protobuf:"bytes,1,opt,name=accountID,proto3" json:"accountID,omitempty"`
	TenantID      string                   `protobuf:"bytes,2,opt,name=tenantID,proto3" json:"tenantID,omitempty"`
	NetID         string                   `protobuf:"bytes,3,opt,name=netID,proto3" json:"netID,omitempty"`
	RoutingDomain *nac.RoutingDomain       `protobuf:"bytes,11,opt,name=routingDomain,proto3" json:"routingDomain,omitempty"`
	Routers       map[string]*NetHop       `protobuf:"bytes,21,rep,name=routers,proto3" json:"routers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`           // map[p2pHostID]*NetHop
	Relays        map[string]*NetHop       `protobuf:"bytes,31,rep,name=relays,proto3" json:"relays,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`             // map[p2pHostID]*NetHop
	RoutingTable  map[string]*RoutingEntry `protobuf:"bytes,41,rep,name=routingTable,proto3" json:"routingTable,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // map[addr]*RoutingEntry
	// Policy policy = 101;
	Policy map[string]*topology.Policy `protobuf:"bytes,101,rep,name=policy,proto3" json:"policy,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // map[subnetID]*Policy
}

func (x *RIB) Reset() {
	*x = RIB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_network_v1_routing_rib_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RIB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RIB) ProtoMessage() {}

func (x *RIB) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_network_v1_routing_rib_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RIB.ProtoReflect.Descriptor instead.
func (*RIB) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_network_v1_routing_rib_proto_rawDescGZIP(), []int{3}
}

func (x *RIB) GetAccountID() string {
	if x != nil {
		return x.AccountID
	}
	return ""
}

func (x *RIB) GetTenantID() string {
	if x != nil {
		return x.TenantID
	}
	return ""
}

func (x *RIB) GetNetID() string {
	if x != nil {
		return x.NetID
	}
	return ""
}

func (x *RIB) GetRoutingDomain() *nac.RoutingDomain {
	if x != nil {
		return x.RoutingDomain
	}
	return nil
}

func (x *RIB) GetRouters() map[string]*NetHop {
	if x != nil {
		return x.Routers
	}
	return nil
}

func (x *RIB) GetRelays() map[string]*NetHop {
	if x != nil {
		return x.Relays
	}
	return nil
}

func (x *RIB) GetRoutingTable() map[string]*RoutingEntry {
	if x != nil {
		return x.RoutingTable
	}
	return nil
}

func (x *RIB) GetPolicy() map[string]*topology.Policy {
	if x != nil {
		return x.Policy
	}
	return nil
}

var File_mmesh_protobuf_network_v1_routing_rib_proto protoreflect.FileDescriptor

var file_mmesh_protobuf_network_v1_routing_rib_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x75, 0x74,
	0x69, 0x6e, 0x67, 0x2f, 0x72, 0x69, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x72,
	0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x1a, 0x31, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x2f, 0x73, 0x75, 0x62,
	0x6e, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x6d, 0x6d, 0x65, 0x73, 0x68,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x61, 0x63, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcf, 0x03, 0x0a, 0x07, 0x52, 0x49, 0x42, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x28, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x14, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x49, 0x42, 0x44, 0x61,
	0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x27, 0x0a, 0x05,
	0x73, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x6e, 0x61,
	0x63, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x52, 0x05,
	0x73, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x27, 0x0a, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x18,
	0x15, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e,
	0x4e, 0x65, 0x74, 0x48, 0x6f, 0x70, 0x52, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x12, 0x25,
	0x0a, 0x05, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x4e, 0x65, 0x74, 0x48, 0x6f, 0x70, 0x52, 0x05,
	0x72, 0x65, 0x6c, 0x61, 0x79, 0x12, 0x46, 0x0a, 0x0c, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67,
	0x54, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x29, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x72, 0x6f,
	0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x49, 0x42, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x52, 0x6f,
	0x75, 0x74, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x0c, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x34, 0x0a,
	0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x65, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x49, 0x42, 0x44, 0x61, 0x74, 0x61, 0x2e,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x1a, 0x56, 0x0a, 0x11, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x54, 0x61,
	0x62, 0x6c, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2b, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x72, 0x6f, 0x75, 0x74,
	0x69, 0x6e, 0x67, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x4b, 0x0a, 0x0b, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x26, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x74, 0x6f,
	0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xa1, 0x02, 0x0a, 0x0c, 0x52, 0x6f, 0x75,
	0x74, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x75, 0x62,
	0x6e, 0x65, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x75, 0x62,
	0x6e, 0x65, 0x74, 0x49, 0x44, 0x12, 0x3c, 0x0a, 0x0d, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x72,
	0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x46, 0x61,
	0x6d, 0x69, 0x6c, 0x79, 0x52, 0x0d, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x46, 0x61, 0x6d,
	0x69, 0x6c, 0x79, 0x12, 0x26, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x15, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x12, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x2d, 0x0a, 0x02, 0x67,
	0x77, 0x18, 0x1f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e,
	0x67, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x47,
	0x77, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x02, 0x67, 0x77, 0x12, 0x18, 0x0a, 0x07, 0x44, 0x4e,
	0x53, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x29, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x44, 0x4e, 0x53,
	0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x46, 0x0a, 0x07, 0x47, 0x77, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x25, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x4e, 0x65, 0x74, 0x48, 0x6f,
	0x70, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xb4, 0x01, 0x0a,
	0x06, 0x4e, 0x65, 0x74, 0x48, 0x6f, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x75, 0x62, 0x6e, 0x65,
	0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x75, 0x62, 0x6e, 0x65,
	0x74, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x32, 0x50, 0x48, 0x6f, 0x73, 0x74, 0x49, 0x44,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x50, 0x32, 0x50, 0x48, 0x6f, 0x73, 0x74, 0x49,
	0x44, 0x12, 0x16, 0x0a, 0x06, 0x4d, 0x41, 0x64, 0x64, 0x72, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x06, 0x4d, 0x41, 0x64, 0x64, 0x72, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x69,
	0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x15, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x72, 0x69,
	0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x53,
	0x65, 0x65, 0x6e, 0x18, 0x65, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x53,
	0x65, 0x65, 0x6e, 0x22, 0xaa, 0x05, 0x0a, 0x03, 0x52, 0x49, 0x42, 0x12, 0x1c, 0x0a, 0x09, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x65, 0x74, 0x49, 0x44, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x65, 0x74, 0x49, 0x44, 0x12, 0x38, 0x0a, 0x0d, 0x72,
	0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6e, 0x61, 0x63, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67,
	0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x0d, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x44,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x33, 0x0a, 0x07, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x73,
	0x18, 0x15, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67,
	0x2e, 0x52, 0x49, 0x42, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x07, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x73, 0x12, 0x30, 0x0a, 0x06, 0x72, 0x65,
	0x6c, 0x61, 0x79, 0x73, 0x18, 0x1f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x72, 0x6f, 0x75,
	0x74, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x49, 0x42, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x73, 0x12, 0x42, 0x0a, 0x0c,
	0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x29, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x49, 0x42,
	0x2e, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x0c, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x62, 0x6c, 0x65,
	0x12, 0x30, 0x0a, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x65, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x49, 0x42, 0x2e, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x1a, 0x4b, 0x0a, 0x0c, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x25, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x4e, 0x65,
	0x74, 0x48, 0x6f, 0x70, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a,
	0x4a, 0x0a, 0x0b, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x25, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x4e, 0x65, 0x74, 0x48, 0x6f, 0x70,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x56, 0x0a, 0x11, 0x52,
	0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x2b, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x6f, 0x75, 0x74,
	0x69, 0x6e, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x1a, 0x4b, 0x0a, 0x0b, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x26, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x74, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x2e, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x2a, 0x56, 0x0a, 0x0b, 0x52, 0x49, 0x42, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x11, 0x0a, 0x0d, 0x52, 0x4f, 0x55, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x43, 0x4f, 0x50, 0x45,
	0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x52, 0x4f, 0x55, 0x54, 0x45, 0x52, 0x10, 0x15, 0x12, 0x09,
	0x0a, 0x05, 0x52, 0x45, 0x4c, 0x41, 0x59, 0x10, 0x1f, 0x12, 0x11, 0x0a, 0x0d, 0x52, 0x4f, 0x55,
	0x54, 0x49, 0x4e, 0x47, 0x5f, 0x54, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x29, 0x12, 0x0a, 0x0a, 0x06,
	0x50, 0x4f, 0x4c, 0x49, 0x43, 0x59, 0x10, 0x65, 0x2a, 0x21, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x12, 0x07, 0x0a, 0x03, 0x49, 0x50, 0x34,
	0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x49, 0x50, 0x36, 0x10, 0x01, 0x2a, 0x33, 0x0a, 0x09, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x4f, 0x4e, 0x4e,
	0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x54, 0x41, 0x54, 0x49,
	0x43, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x59, 0x4e, 0x41, 0x4d, 0x49, 0x43, 0x10, 0x02,
	0x42, 0x29, 0x5a, 0x27, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x6d, 0x2d,
	0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_mmesh_protobuf_network_v1_routing_rib_proto_rawDescOnce sync.Once
	file_mmesh_protobuf_network_v1_routing_rib_proto_rawDescData = file_mmesh_protobuf_network_v1_routing_rib_proto_rawDesc
)

func file_mmesh_protobuf_network_v1_routing_rib_proto_rawDescGZIP() []byte {
	file_mmesh_protobuf_network_v1_routing_rib_proto_rawDescOnce.Do(func() {
		file_mmesh_protobuf_network_v1_routing_rib_proto_rawDescData = protoimpl.X.CompressGZIP(file_mmesh_protobuf_network_v1_routing_rib_proto_rawDescData)
	})
	return file_mmesh_protobuf_network_v1_routing_rib_proto_rawDescData
}

var file_mmesh_protobuf_network_v1_routing_rib_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_mmesh_protobuf_network_v1_routing_rib_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_mmesh_protobuf_network_v1_routing_rib_proto_goTypes = []interface{}{
	(RIBDataType)(0),          // 0: routing.RIBDataType
	(AddressFamily)(0),        // 1: routing.AddressFamily
	(RouteType)(0),            // 2: routing.RouteType
	(*RIBData)(nil),           // 3: routing.RIBData
	(*RoutingEntry)(nil),      // 4: routing.RoutingEntry
	(*NetHop)(nil),            // 5: routing.NetHop
	(*RIB)(nil),               // 6: routing.RIB
	nil,                       // 7: routing.RIBData.RoutingTableEntry
	nil,                       // 8: routing.RIBData.PolicyEntry
	nil,                       // 9: routing.RoutingEntry.GwEntry
	nil,                       // 10: routing.RIB.RoutersEntry
	nil,                       // 11: routing.RIB.RelaysEntry
	nil,                       // 12: routing.RIB.RoutingTableEntry
	nil,                       // 13: routing.RIB.PolicyEntry
	(nac.RoutingScope)(0),     // 14: nac.RoutingScope
	(*nac.RoutingDomain)(nil), // 15: nac.RoutingDomain
	(*topology.Policy)(nil),   // 16: topology.Policy
}
var file_mmesh_protobuf_network_v1_routing_rib_proto_depIdxs = []int32{
	0,  // 0: routing.RIBData.type:type_name -> routing.RIBDataType
	14, // 1: routing.RIBData.scope:type_name -> nac.RoutingScope
	5,  // 2: routing.RIBData.router:type_name -> routing.NetHop
	5,  // 3: routing.RIBData.relay:type_name -> routing.NetHop
	7,  // 4: routing.RIBData.routingTable:type_name -> routing.RIBData.RoutingTableEntry
	8,  // 5: routing.RIBData.policy:type_name -> routing.RIBData.PolicyEntry
	1,  // 6: routing.RoutingEntry.addressFamily:type_name -> routing.AddressFamily
	2,  // 7: routing.RoutingEntry.type:type_name -> routing.RouteType
	9,  // 8: routing.RoutingEntry.gw:type_name -> routing.RoutingEntry.GwEntry
	15, // 9: routing.RIB.routingDomain:type_name -> nac.RoutingDomain
	10, // 10: routing.RIB.routers:type_name -> routing.RIB.RoutersEntry
	11, // 11: routing.RIB.relays:type_name -> routing.RIB.RelaysEntry
	12, // 12: routing.RIB.routingTable:type_name -> routing.RIB.RoutingTableEntry
	13, // 13: routing.RIB.policy:type_name -> routing.RIB.PolicyEntry
	4,  // 14: routing.RIBData.RoutingTableEntry.value:type_name -> routing.RoutingEntry
	16, // 15: routing.RIBData.PolicyEntry.value:type_name -> topology.Policy
	5,  // 16: routing.RoutingEntry.GwEntry.value:type_name -> routing.NetHop
	5,  // 17: routing.RIB.RoutersEntry.value:type_name -> routing.NetHop
	5,  // 18: routing.RIB.RelaysEntry.value:type_name -> routing.NetHop
	4,  // 19: routing.RIB.RoutingTableEntry.value:type_name -> routing.RoutingEntry
	16, // 20: routing.RIB.PolicyEntry.value:type_name -> topology.Policy
	21, // [21:21] is the sub-list for method output_type
	21, // [21:21] is the sub-list for method input_type
	21, // [21:21] is the sub-list for extension type_name
	21, // [21:21] is the sub-list for extension extendee
	0,  // [0:21] is the sub-list for field type_name
}

func init() { file_mmesh_protobuf_network_v1_routing_rib_proto_init() }
func file_mmesh_protobuf_network_v1_routing_rib_proto_init() {
	if File_mmesh_protobuf_network_v1_routing_rib_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mmesh_protobuf_network_v1_routing_rib_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RIBData); i {
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
		file_mmesh_protobuf_network_v1_routing_rib_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoutingEntry); i {
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
		file_mmesh_protobuf_network_v1_routing_rib_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetHop); i {
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
		file_mmesh_protobuf_network_v1_routing_rib_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RIB); i {
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
			RawDescriptor: file_mmesh_protobuf_network_v1_routing_rib_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mmesh_protobuf_network_v1_routing_rib_proto_goTypes,
		DependencyIndexes: file_mmesh_protobuf_network_v1_routing_rib_proto_depIdxs,
		EnumInfos:         file_mmesh_protobuf_network_v1_routing_rib_proto_enumTypes,
		MessageInfos:      file_mmesh_protobuf_network_v1_routing_rib_proto_msgTypes,
	}.Build()
	File_mmesh_protobuf_network_v1_routing_rib_proto = out.File
	file_mmesh_protobuf_network_v1_routing_rib_proto_rawDesc = nil
	file_mmesh_protobuf_network_v1_routing_rib_proto_goTypes = nil
	file_mmesh_protobuf_network_v1_routing_rib_proto_depIdxs = nil
}
