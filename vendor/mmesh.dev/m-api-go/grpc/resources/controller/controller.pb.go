// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0--rc2
// source: mmesh/protobuf/resources/v1/controller/controller.proto

package controller

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	auth "mmesh.dev/m-api-go/grpc/resources/iam/auth"
	resource "mmesh.dev/m-api-go/grpc/resources/resource"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Federation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LocationID   string                    `protobuf:"bytes,1,opt,name=locationID,proto3" json:"locationID,omitempty"`
	FederationID string                    `protobuf:"bytes,2,opt,name=federationID,proto3" json:"federationID,omitempty"`                                                                                        // string federationToken = 3;
	Controllers  map[string]*Controller    `protobuf:"bytes,21,rep,name=controllers,proto3" json:"controllers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // map[controllerID]*Controller
	Usage        *FederationUsage          `protobuf:"bytes,91,opt,name=usage,proto3" json:"usage,omitempty"`
	Private      bool                      `protobuf:"varint,101,opt,name=private,proto3" json:"private,omitempty"`
	Options      *PrivateFederationOptions `protobuf:"bytes,105,opt,name=options,proto3" json:"options,omitempty"`
}

func (x *Federation) Reset() {
	*x = Federation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_controller_controller_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Federation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Federation) ProtoMessage() {}

func (x *Federation) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_controller_controller_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Federation.ProtoReflect.Descriptor instead.
func (*Federation) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_controller_controller_proto_rawDescGZIP(), []int{0}
}

func (x *Federation) GetLocationID() string {
	if x != nil {
		return x.LocationID
	}
	return ""
}

func (x *Federation) GetFederationID() string {
	if x != nil {
		return x.FederationID
	}
	return ""
}

func (x *Federation) GetControllers() map[string]*Controller {
	if x != nil {
		return x.Controllers
	}
	return nil
}

func (x *Federation) GetUsage() *FederationUsage {
	if x != nil {
		return x.Usage
	}
	return nil
}

func (x *Federation) GetPrivate() bool {
	if x != nil {
		return x.Private
	}
	return false
}

func (x *Federation) GetOptions() *PrivateFederationOptions {
	if x != nil {
		return x.Options
	}
	return nil
}

type Federations struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta        *resource.ListResponse `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	Federations []*Federation          `protobuf:"bytes,2,rep,name=federations,proto3" json:"federations,omitempty"`
}

func (x *Federations) Reset() {
	*x = Federations{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_controller_controller_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Federations) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Federations) ProtoMessage() {}

func (x *Federations) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_controller_controller_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Federations.ProtoReflect.Descriptor instead.
func (*Federations) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_controller_controller_proto_rawDescGZIP(), []int{1}
}

func (x *Federations) GetMeta() *resource.ListResponse {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *Federations) GetFederations() []*Federation {
	if x != nil {
		return x.Federations
	}
	return nil
}

type ListFederationsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta       *resource.ListRequest `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	LocationID string                `protobuf:"bytes,2,opt,name=locationID,proto3" json:"locationID,omitempty"`
}

func (x *ListFederationsRequest) Reset() {
	*x = ListFederationsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_controller_controller_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFederationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFederationsRequest) ProtoMessage() {}

func (x *ListFederationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_controller_controller_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFederationsRequest.ProtoReflect.Descriptor instead.
func (*ListFederationsRequest) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_controller_controller_proto_rawDescGZIP(), []int{2}
}

func (x *ListFederationsRequest) GetMeta() *resource.ListRequest {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *ListFederationsRequest) GetLocationID() string {
	if x != nil {
		return x.LocationID
	}
	return ""
}

type FederationUsage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LastUpdated int64 `protobuf:"varint,1,opt,name=lastUpdated,proto3" json:"lastUpdated,omitempty"`
	Accounts    int32 `protobuf:"varint,11,opt,name=accounts,proto3" json:"accounts,omitempty"`
	Tenants     int32 `protobuf:"varint,21,opt,name=tenants,proto3" json:"tenants,omitempty"`
	Networks    int32 `protobuf:"varint,101,opt,name=networks,proto3" json:"networks,omitempty"`
	Subnets     int32 `protobuf:"varint,111,opt,name=subnets,proto3" json:"subnets,omitempty"`
	Nodes       int32 `protobuf:"varint,121,opt,name=nodes,proto3" json:"nodes,omitempty"`
	Endpoints   int32 `protobuf:"varint,122,opt,name=endpoints,proto3" json:"endpoints,omitempty"`
	Relays      int32 `protobuf:"varint,125,opt,name=relays,proto3" json:"relays,omitempty"`
}

func (x *FederationUsage) Reset() {
	*x = FederationUsage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_controller_controller_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FederationUsage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FederationUsage) ProtoMessage() {}

func (x *FederationUsage) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_controller_controller_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FederationUsage.ProtoReflect.Descriptor instead.
func (*FederationUsage) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_controller_controller_proto_rawDescGZIP(), []int{3}
}

func (x *FederationUsage) GetLastUpdated() int64 {
	if x != nil {
		return x.LastUpdated
	}
	return 0
}

func (x *FederationUsage) GetAccounts() int32 {
	if x != nil {
		return x.Accounts
	}
	return 0
}

func (x *FederationUsage) GetTenants() int32 {
	if x != nil {
		return x.Tenants
	}
	return 0
}

func (x *FederationUsage) GetNetworks() int32 {
	if x != nil {
		return x.Networks
	}
	return 0
}

func (x *FederationUsage) GetSubnets() int32 {
	if x != nil {
		return x.Subnets
	}
	return 0
}

func (x *FederationUsage) GetNodes() int32 {
	if x != nil {
		return x.Nodes
	}
	return 0
}

func (x *FederationUsage) GetEndpoints() int32 {
	if x != nil {
		return x.Endpoints
	}
	return 0
}

func (x *FederationUsage) GetRelays() int32 {
	if x != nil {
		return x.Relays
	}
	return 0
}

type PrivateFederationOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountID string `protobuf:"bytes,1,opt,name=accountID,proto3" json:"accountID,omitempty"` // for account-dedicated private federations (private == true)
}

func (x *PrivateFederationOptions) Reset() {
	*x = PrivateFederationOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_controller_controller_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrivateFederationOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrivateFederationOptions) ProtoMessage() {}

func (x *PrivateFederationOptions) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_controller_controller_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrivateFederationOptions.ProtoReflect.Descriptor instead.
func (*PrivateFederationOptions) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_controller_controller_proto_rawDescGZIP(), []int{4}
}

func (x *PrivateFederationOptions) GetAccountID() string {
	if x != nil {
		return x.AccountID
	}
	return ""
}

type FederationEndpoints struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FederationID string                 `protobuf:"bytes,1,opt,name=federationID,proto3" json:"federationID,omitempty"`
	Controllers  map[string]*Controller `protobuf:"bytes,3,rep,name=controllers,proto3" json:"controllers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // map[controllerID]*Controller
}

func (x *FederationEndpoints) Reset() {
	*x = FederationEndpoints{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_controller_controller_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FederationEndpoints) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FederationEndpoints) ProtoMessage() {}

func (x *FederationEndpoints) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_controller_controller_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FederationEndpoints.ProtoReflect.Descriptor instead.
func (*FederationEndpoints) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_controller_controller_proto_rawDescGZIP(), []int{5}
}

func (x *FederationEndpoints) GetFederationID() string {
	if x != nil {
		return x.FederationID
	}
	return ""
}

func (x *FederationEndpoints) GetControllers() map[string]*Controller {
	if x != nil {
		return x.Controllers
	}
	return nil
}

type FederationSelected struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LocationID   string `protobuf:"bytes,1,opt,name=locationID,proto3" json:"locationID,omitempty"`
	FederationID string `protobuf:"bytes,2,opt,name=federationID,proto3" json:"federationID,omitempty"`
	VirtualHost  string `protobuf:"bytes,11,opt,name=virtualHost,proto3" json:"virtualHost,omitempty"`
	Port         int32  `protobuf:"varint,12,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *FederationSelected) Reset() {
	*x = FederationSelected{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_controller_controller_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FederationSelected) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FederationSelected) ProtoMessage() {}

func (x *FederationSelected) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_controller_controller_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FederationSelected.ProtoReflect.Descriptor instead.
func (*FederationSelected) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_controller_controller_proto_rawDescGZIP(), []int{6}
}

func (x *FederationSelected) GetLocationID() string {
	if x != nil {
		return x.LocationID
	}
	return ""
}

func (x *FederationSelected) GetFederationID() string {
	if x != nil {
		return x.FederationID
	}
	return ""
}

func (x *FederationSelected) GetVirtualHost() string {
	if x != nil {
		return x.VirtualHost
	}
	return ""
}

func (x *FederationSelected) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

type Controller struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LocationID   string `protobuf:"bytes,1,opt,name=locationID,proto3" json:"locationID,omitempty"`
	FederationID string `protobuf:"bytes,2,opt,name=federationID,proto3" json:"federationID,omitempty"`
	// string federationToken = 3;
	ControllerID string        `protobuf:"bytes,11,opt,name=controllerID,proto3" json:"controllerID,omitempty"`
	AuthKey      *auth.AuthKey `protobuf:"bytes,21,opt,name=authKey,proto3" json:"authKey,omitempty"`
	VirtualHost  string        `protobuf:"bytes,31,opt,name=virtualHost,proto3" json:"virtualHost,omitempty"`
	Host         string        `protobuf:"bytes,32,opt,name=host,proto3" json:"host,omitempty"`
	// string privateHost = 33;
	Port   int32   `protobuf:"varint,34,opt,name=port,proto3" json:"port,omitempty"`
	Status *Status `protobuf:"bytes,15,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Controller) Reset() {
	*x = Controller{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_controller_controller_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Controller) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Controller) ProtoMessage() {}

func (x *Controller) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_controller_controller_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Controller.ProtoReflect.Descriptor instead.
func (*Controller) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_controller_controller_proto_rawDescGZIP(), []int{7}
}

func (x *Controller) GetLocationID() string {
	if x != nil {
		return x.LocationID
	}
	return ""
}

func (x *Controller) GetFederationID() string {
	if x != nil {
		return x.FederationID
	}
	return ""
}

func (x *Controller) GetControllerID() string {
	if x != nil {
		return x.ControllerID
	}
	return ""
}

func (x *Controller) GetAuthKey() *auth.AuthKey {
	if x != nil {
		return x.AuthKey
	}
	return nil
}

func (x *Controller) GetVirtualHost() string {
	if x != nil {
		return x.VirtualHost
	}
	return ""
}

func (x *Controller) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *Controller) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *Controller) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

type Controllers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta        *resource.ListResponse `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	Controllers []*Controller          `protobuf:"bytes,2,rep,name=controllers,proto3" json:"controllers,omitempty"`
}

func (x *Controllers) Reset() {
	*x = Controllers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_controller_controller_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Controllers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Controllers) ProtoMessage() {}

func (x *Controllers) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_controller_controller_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Controllers.ProtoReflect.Descriptor instead.
func (*Controllers) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_controller_controller_proto_rawDescGZIP(), []int{8}
}

func (x *Controllers) GetMeta() *resource.ListResponse {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *Controllers) GetControllers() []*Controller {
	if x != nil {
		return x.Controllers
	}
	return nil
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LastUpdated int64 `protobuf:"varint,1,opt,name=lastUpdated,proto3" json:"lastUpdated,omitempty"`
	Connections int32 `protobuf:"varint,11,opt,name=connections,proto3" json:"connections,omitempty"` // int32 connectedUsers = 3;
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_controller_controller_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_controller_controller_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_controller_controller_proto_rawDescGZIP(), []int{9}
}

func (x *Status) GetLastUpdated() int64 {
	if x != nil {
		return x.LastUpdated
	}
	return 0
}

func (x *Status) GetConnections() int32 {
	if x != nil {
		return x.Connections
	}
	return 0
}

var File_mmesh_protobuf_resources_v1_controller_controller_proto protoreflect.FileDescriptor

var file_mmesh_protobuf_resources_v1_controller_controller_proto_rawDesc = []byte{
	0x0a, 0x37, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x1a, 0x2f, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x6c, 0x69, 0x73, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2f, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x80, 0x03, 0x0a, 0x0a, 0x46, 0x65, 0x64, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x22, 0x0a, 0x0c, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x65,
	0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x49, 0x0a, 0x0b, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x73, 0x18, 0x15, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x27, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x46, 0x65, 0x64,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c,
	0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x6c, 0x65, 0x72, 0x73, 0x12, 0x31, 0x0a, 0x05, 0x75, 0x73, 0x61, 0x67, 0x65, 0x18, 0x5b,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65,
	0x72, 0x2e, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x05, 0x75, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x18, 0x65, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x70, 0x72, 0x69, 0x76, 0x61,
	0x74, 0x65, 0x12, 0x3e, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x69, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72,
	0x2e, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x1a, 0x56, 0x0a, 0x10, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x73, 0x0a, 0x0b, 0x46, 0x65,
	0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2a, 0x0a, 0x04, 0x6d, 0x65, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52,
	0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x38, 0x0a, 0x0b, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0b, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22,
	0x63, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x04, 0x6d, 0x65, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04,
	0x6d, 0x65, 0x74, 0x61, 0x12, 0x1e, 0x0a, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x44, 0x22, 0xeb, 0x01, 0x0a, 0x0f, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x6c, 0x61, 0x73, 0x74,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6c,
	0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x73, 0x18, 0x15, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x73,
	0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x18, 0x65, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x73, 0x18, 0x6f, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x73,
	0x75, 0x62, 0x6e, 0x65, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18,
	0x79, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09,
	0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x7a, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65,
	0x6c, 0x61, 0x79, 0x73, 0x18, 0x7d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x65, 0x6c, 0x61,
	0x79, 0x73, 0x22, 0x38, 0x0a, 0x18, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x46, 0x65, 0x64,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1c,
	0x0a, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x22, 0xe5, 0x01, 0x0a,
	0x13, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x65, 0x64, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x52, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x46, 0x65, 0x64, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x73, 0x1a, 0x56, 0x0a, 0x10,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x2c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0x8e, 0x01, 0x0a, 0x12, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x22, 0x0a, 0x0c, 0x66,
	0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12,
	0x20, 0x0a, 0x0b, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x48, 0x6f, 0x73, 0x74, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x48, 0x6f, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x93, 0x02, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x6c, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x44, 0x12, 0x22, 0x0a, 0x0c, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x65, 0x64, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x49, 0x44, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x49, 0x44, 0x12, 0x27, 0x0a, 0x07,
	0x61, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x52, 0x07, 0x61, 0x75,
	0x74, 0x68, 0x4b, 0x65, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c,
	0x48, 0x6f, 0x73, 0x74, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x76, 0x69, 0x72, 0x74,
	0x75, 0x61, 0x6c, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18,
	0x20, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x6f, 0x72, 0x74, 0x18, 0x22, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12,
	0x2a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x73, 0x0a, 0x0b, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x73, 0x12, 0x2a, 0x0a, 0x04, 0x6d, 0x65,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x38, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x6c, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x6c, 0x65, 0x72, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x73,
	0x22, 0x4c, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x6c, 0x61,
	0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0b, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x20, 0x0a, 0x0b,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x2e,
	0x5a, 0x2c, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x6d, 0x2d, 0x61, 0x70,
	0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mmesh_protobuf_resources_v1_controller_controller_proto_rawDescOnce sync.Once
	file_mmesh_protobuf_resources_v1_controller_controller_proto_rawDescData = file_mmesh_protobuf_resources_v1_controller_controller_proto_rawDesc
)

func file_mmesh_protobuf_resources_v1_controller_controller_proto_rawDescGZIP() []byte {
	file_mmesh_protobuf_resources_v1_controller_controller_proto_rawDescOnce.Do(func() {
		file_mmesh_protobuf_resources_v1_controller_controller_proto_rawDescData = protoimpl.X.CompressGZIP(file_mmesh_protobuf_resources_v1_controller_controller_proto_rawDescData)
	})
	return file_mmesh_protobuf_resources_v1_controller_controller_proto_rawDescData
}

var file_mmesh_protobuf_resources_v1_controller_controller_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_mmesh_protobuf_resources_v1_controller_controller_proto_goTypes = []interface{}{
	(*Federation)(nil),               // 0: controller.Federation
	(*Federations)(nil),              // 1: controller.Federations
	(*ListFederationsRequest)(nil),   // 2: controller.ListFederationsRequest
	(*FederationUsage)(nil),          // 3: controller.FederationUsage
	(*PrivateFederationOptions)(nil), // 4: controller.PrivateFederationOptions
	(*FederationEndpoints)(nil),      // 5: controller.FederationEndpoints
	(*FederationSelected)(nil),       // 6: controller.FederationSelected
	(*Controller)(nil),               // 7: controller.Controller
	(*Controllers)(nil),              // 8: controller.Controllers
	(*Status)(nil),                   // 9: controller.Status
	nil,                              // 10: controller.Federation.ControllersEntry
	nil,                              // 11: controller.FederationEndpoints.ControllersEntry
	(*resource.ListResponse)(nil),    // 12: resource.ListResponse
	(*resource.ListRequest)(nil),     // 13: resource.ListRequest
	(*auth.AuthKey)(nil),             // 14: auth.AuthKey
}
var file_mmesh_protobuf_resources_v1_controller_controller_proto_depIdxs = []int32{
	10, // 0: controller.Federation.controllers:type_name -> controller.Federation.ControllersEntry
	3,  // 1: controller.Federation.usage:type_name -> controller.FederationUsage
	4,  // 2: controller.Federation.options:type_name -> controller.PrivateFederationOptions
	12, // 3: controller.Federations.meta:type_name -> resource.ListResponse
	0,  // 4: controller.Federations.federations:type_name -> controller.Federation
	13, // 5: controller.ListFederationsRequest.meta:type_name -> resource.ListRequest
	11, // 6: controller.FederationEndpoints.controllers:type_name -> controller.FederationEndpoints.ControllersEntry
	14, // 7: controller.Controller.authKey:type_name -> auth.AuthKey
	9,  // 8: controller.Controller.status:type_name -> controller.Status
	12, // 9: controller.Controllers.meta:type_name -> resource.ListResponse
	7,  // 10: controller.Controllers.controllers:type_name -> controller.Controller
	7,  // 11: controller.Federation.ControllersEntry.value:type_name -> controller.Controller
	7,  // 12: controller.FederationEndpoints.ControllersEntry.value:type_name -> controller.Controller
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_mmesh_protobuf_resources_v1_controller_controller_proto_init() }
func file_mmesh_protobuf_resources_v1_controller_controller_proto_init() {
	if File_mmesh_protobuf_resources_v1_controller_controller_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mmesh_protobuf_resources_v1_controller_controller_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Federation); i {
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
		file_mmesh_protobuf_resources_v1_controller_controller_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Federations); i {
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
		file_mmesh_protobuf_resources_v1_controller_controller_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFederationsRequest); i {
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
		file_mmesh_protobuf_resources_v1_controller_controller_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FederationUsage); i {
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
		file_mmesh_protobuf_resources_v1_controller_controller_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrivateFederationOptions); i {
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
		file_mmesh_protobuf_resources_v1_controller_controller_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FederationEndpoints); i {
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
		file_mmesh_protobuf_resources_v1_controller_controller_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FederationSelected); i {
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
		file_mmesh_protobuf_resources_v1_controller_controller_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Controller); i {
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
		file_mmesh_protobuf_resources_v1_controller_controller_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Controllers); i {
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
		file_mmesh_protobuf_resources_v1_controller_controller_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
			RawDescriptor: file_mmesh_protobuf_resources_v1_controller_controller_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mmesh_protobuf_resources_v1_controller_controller_proto_goTypes,
		DependencyIndexes: file_mmesh_protobuf_resources_v1_controller_controller_proto_depIdxs,
		MessageInfos:      file_mmesh_protobuf_resources_v1_controller_controller_proto_msgTypes,
	}.Build()
	File_mmesh_protobuf_resources_v1_controller_controller_proto = out.File
	file_mmesh_protobuf_resources_v1_controller_controller_proto_rawDesc = nil
	file_mmesh_protobuf_resources_v1_controller_controller_proto_goTypes = nil
	file_mmesh_protobuf_resources_v1_controller_controller_proto_depIdxs = nil
}
