// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0--rc2
// source: mmesh/protobuf/resources/v1/events/source.proto

package events

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SourceType int32

const (
	SourceType_UNDEFINED_SOURCE SourceType = 0
	SourceType_MANAGER          SourceType = 10
	SourceType_CONTROLLER       SourceType = 20
	SourceType_ROUTER           SourceType = 30
	SourceType_NODE             SourceType = 100
	SourceType_CLI              SourceType = 200
	SourceType_WEBUI            SourceType = 210
)

// Enum value maps for SourceType.
var (
	SourceType_name = map[int32]string{
		0:   "UNDEFINED_SOURCE",
		10:  "MANAGER",
		20:  "CONTROLLER",
		30:  "ROUTER",
		100: "NODE",
		200: "CLI",
		210: "WEBUI",
	}
	SourceType_value = map[string]int32{
		"UNDEFINED_SOURCE": 0,
		"MANAGER":          10,
		"CONTROLLER":       20,
		"ROUTER":           30,
		"NODE":             100,
		"CLI":              200,
		"WEBUI":            210,
	}
)

func (x SourceType) Enum() *SourceType {
	p := new(SourceType)
	*p = x
	return p
}

func (x SourceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SourceType) Descriptor() protoreflect.EnumDescriptor {
	return file_mmesh_protobuf_resources_v1_events_source_proto_enumTypes[0].Descriptor()
}

func (SourceType) Type() protoreflect.EnumType {
	return &file_mmesh_protobuf_resources_v1_events_source_proto_enumTypes[0]
}

func (x SourceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SourceType.Descriptor instead.
func (SourceType) EnumDescriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_events_source_proto_rawDescGZIP(), []int{0}
}

type Source struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Source entity of the event or the affected system,
	// for example 'controller', 'manager', 'node', 'cli', 'app-stack'
	Type SourceType `protobuf:"varint,1,opt,name=type,proto3,enum=events.SourceType" json:"type,omitempty"`
	// group == CONTROLLER
	Controller *SourceController `protobuf:"bytes,21,opt,name=controller,proto3" json:"controller,omitempty"`
	// group == MANAGER
	Manager *SourceManager `protobuf:"bytes,31,opt,name=manager,proto3" json:"manager,omitempty"`
	// group == ROUTER
	Router *SourceRouter `protobuf:"bytes,41,opt,name=router,proto3" json:"router,omitempty"`
	// group == NODE
	Node *SourceNode `protobuf:"bytes,51,opt,name=node,proto3" json:"node,omitempty"`
	// group == CLI
	Cli *SourceCLI `protobuf:"bytes,101,opt,name=cli,proto3" json:"cli,omitempty"`
	// group == WEBUI
	WebUI *SourceWebUI `protobuf:"bytes,111,opt,name=webUI,proto3" json:"webUI,omitempty"`
}

func (x *Source) Reset() {
	*x = Source{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_events_source_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Source) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Source) ProtoMessage() {}

func (x *Source) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_events_source_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Source.ProtoReflect.Descriptor instead.
func (*Source) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_events_source_proto_rawDescGZIP(), []int{0}
}

func (x *Source) GetType() SourceType {
	if x != nil {
		return x.Type
	}
	return SourceType_UNDEFINED_SOURCE
}

func (x *Source) GetController() *SourceController {
	if x != nil {
		return x.Controller
	}
	return nil
}

func (x *Source) GetManager() *SourceManager {
	if x != nil {
		return x.Manager
	}
	return nil
}

func (x *Source) GetRouter() *SourceRouter {
	if x != nil {
		return x.Router
	}
	return nil
}

func (x *Source) GetNode() *SourceNode {
	if x != nil {
		return x.Node
	}
	return nil
}

func (x *Source) GetCli() *SourceCLI {
	if x != nil {
		return x.Cli
	}
	return nil
}

func (x *Source) GetWebUI() *SourceWebUI {
	if x != nil {
		return x.WebUI
	}
	return nil
}

type SourceManager struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ManagerID string `protobuf:"bytes,1,opt,name=managerID,proto3" json:"managerID,omitempty"`
}

func (x *SourceManager) Reset() {
	*x = SourceManager{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_events_source_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SourceManager) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SourceManager) ProtoMessage() {}

func (x *SourceManager) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_events_source_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SourceManager.ProtoReflect.Descriptor instead.
func (*SourceManager) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_events_source_proto_rawDescGZIP(), []int{1}
}

func (x *SourceManager) GetManagerID() string {
	if x != nil {
		return x.ManagerID
	}
	return ""
}

type SourceController struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ControllerID string `protobuf:"bytes,1,opt,name=controllerID,proto3" json:"controllerID,omitempty"`
}

func (x *SourceController) Reset() {
	*x = SourceController{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_events_source_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SourceController) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SourceController) ProtoMessage() {}

func (x *SourceController) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_events_source_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SourceController.ProtoReflect.Descriptor instead.
func (*SourceController) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_events_source_proto_rawDescGZIP(), []int{2}
}

func (x *SourceController) GetControllerID() string {
	if x != nil {
		return x.ControllerID
	}
	return ""
}

type SourceRouter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RouterID string `protobuf:"bytes,1,opt,name=routerID,proto3" json:"routerID,omitempty"`
}

func (x *SourceRouter) Reset() {
	*x = SourceRouter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_events_source_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SourceRouter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SourceRouter) ProtoMessage() {}

func (x *SourceRouter) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_events_source_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SourceRouter.ProtoReflect.Descriptor instead.
func (*SourceRouter) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_events_source_proto_rawDescGZIP(), []int{3}
}

func (x *SourceRouter) GetRouterID() string {
	if x != nil {
		return x.RouterID
	}
	return ""
}

type SourceNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountID string `protobuf:"bytes,1,opt,name=accountID,proto3" json:"accountID,omitempty"`
	TenantID  string `protobuf:"bytes,11,opt,name=tenantID,proto3" json:"tenantID,omitempty"`
	NetID     string `protobuf:"bytes,21,opt,name=netID,proto3" json:"netID,omitempty"`
	SubnetID  string `protobuf:"bytes,31,opt,name=subnetID,proto3" json:"subnetID,omitempty"`
	NodeID    string `protobuf:"bytes,41,opt,name=nodeID,proto3" json:"nodeID,omitempty"`
	NodeName  string `protobuf:"bytes,42,opt,name=nodeName,proto3" json:"nodeName,omitempty"`
}

func (x *SourceNode) Reset() {
	*x = SourceNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_events_source_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SourceNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SourceNode) ProtoMessage() {}

func (x *SourceNode) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_events_source_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SourceNode.ProtoReflect.Descriptor instead.
func (*SourceNode) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_events_source_proto_rawDescGZIP(), []int{4}
}

func (x *SourceNode) GetAccountID() string {
	if x != nil {
		return x.AccountID
	}
	return ""
}

func (x *SourceNode) GetTenantID() string {
	if x != nil {
		return x.TenantID
	}
	return ""
}

func (x *SourceNode) GetNetID() string {
	if x != nil {
		return x.NetID
	}
	return ""
}

func (x *SourceNode) GetSubnetID() string {
	if x != nil {
		return x.SubnetID
	}
	return ""
}

func (x *SourceNode) GetNodeID() string {
	if x != nil {
		return x.NodeID
	}
	return ""
}

func (x *SourceNode) GetNodeName() string {
	if x != nil {
		return x.NodeName
	}
	return ""
}

type SourceCLI struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientID string `protobuf:"bytes,1,opt,name=clientID,proto3" json:"clientID,omitempty"`
}

func (x *SourceCLI) Reset() {
	*x = SourceCLI{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_events_source_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SourceCLI) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SourceCLI) ProtoMessage() {}

func (x *SourceCLI) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_events_source_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SourceCLI.ProtoReflect.Descriptor instead.
func (*SourceCLI) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_events_source_proto_rawDescGZIP(), []int{5}
}

func (x *SourceCLI) GetClientID() string {
	if x != nil {
		return x.ClientID
	}
	return ""
}

type SourceWebUI struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientID string `protobuf:"bytes,1,opt,name=clientID,proto3" json:"clientID,omitempty"`
}

func (x *SourceWebUI) Reset() {
	*x = SourceWebUI{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_events_source_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SourceWebUI) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SourceWebUI) ProtoMessage() {}

func (x *SourceWebUI) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_events_source_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SourceWebUI.ProtoReflect.Descriptor instead.
func (*SourceWebUI) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_events_source_proto_rawDescGZIP(), []int{6}
}

func (x *SourceWebUI) GetClientID() string {
	if x != nil {
		return x.ClientID
	}
	return ""
}

var File_mmesh_protobuf_resources_v1_events_source_proto protoreflect.FileDescriptor

var file_mmesh_protobuf_resources_v1_events_source_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x22, 0xc1, 0x02, 0x0a, 0x06, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x12, 0x26, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x12, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x38, 0x0a, 0x0a,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x2f, 0x0a, 0x07, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x52, 0x07,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x72, 0x18, 0x29, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x52, 0x06, 0x72,
	0x6f, 0x75, 0x74, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x33, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x12, 0x23, 0x0a,
	0x03, 0x63, 0x6c, 0x69, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x4c, 0x49, 0x52, 0x03, 0x63,
	0x6c, 0x69, 0x12, 0x29, 0x0a, 0x05, 0x77, 0x65, 0x62, 0x55, 0x49, 0x18, 0x6f, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x57, 0x65, 0x62, 0x55, 0x49, 0x52, 0x05, 0x77, 0x65, 0x62, 0x55, 0x49, 0x22, 0x2d, 0x0a,
	0x0d, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x1c,
	0x0a, 0x09, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x49, 0x44, 0x22, 0x36, 0x0a, 0x10,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72,
	0x12, 0x22, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c,
	0x65, 0x72, 0x49, 0x44, 0x22, 0x2a, 0x0a, 0x0c, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x49, 0x44,
	0x22, 0xac, 0x01, 0x0a, 0x0a, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x1a, 0x0a,
	0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x65, 0x74,
	0x49, 0x44, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x65, 0x74, 0x49, 0x44, 0x12,
	0x1a, 0x0a, 0x08, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x49, 0x44, 0x18, 0x1f, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x6e,
	0x6f, 0x64, 0x65, 0x49, 0x44, 0x18, 0x29, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x64,
	0x65, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x2a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22,
	0x27, 0x0a, 0x09, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x4c, 0x49, 0x12, 0x1a, 0x0a, 0x08,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x22, 0x29, 0x0a, 0x0b, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x57, 0x65, 0x62, 0x55, 0x49, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x49, 0x44, 0x2a, 0x6b, 0x0a, 0x0a, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x14, 0x0a, 0x10, 0x55, 0x4e, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x45, 0x44, 0x5f, 0x53,
	0x4f, 0x55, 0x52, 0x43, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x4d, 0x41, 0x4e, 0x41, 0x47,
	0x45, 0x52, 0x10, 0x0a, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x4f, 0x4e, 0x54, 0x52, 0x4f, 0x4c, 0x4c,
	0x45, 0x52, 0x10, 0x14, 0x12, 0x0a, 0x0a, 0x06, 0x52, 0x4f, 0x55, 0x54, 0x45, 0x52, 0x10, 0x1e,
	0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x44, 0x45, 0x10, 0x64, 0x12, 0x08, 0x0a, 0x03, 0x43, 0x4c,
	0x49, 0x10, 0xc8, 0x01, 0x12, 0x0a, 0x0a, 0x05, 0x57, 0x45, 0x42, 0x55, 0x49, 0x10, 0xd2, 0x01,
	0x42, 0x2a, 0x5a, 0x28, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x6d, 0x2d,
	0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mmesh_protobuf_resources_v1_events_source_proto_rawDescOnce sync.Once
	file_mmesh_protobuf_resources_v1_events_source_proto_rawDescData = file_mmesh_protobuf_resources_v1_events_source_proto_rawDesc
)

func file_mmesh_protobuf_resources_v1_events_source_proto_rawDescGZIP() []byte {
	file_mmesh_protobuf_resources_v1_events_source_proto_rawDescOnce.Do(func() {
		file_mmesh_protobuf_resources_v1_events_source_proto_rawDescData = protoimpl.X.CompressGZIP(file_mmesh_protobuf_resources_v1_events_source_proto_rawDescData)
	})
	return file_mmesh_protobuf_resources_v1_events_source_proto_rawDescData
}

var file_mmesh_protobuf_resources_v1_events_source_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_mmesh_protobuf_resources_v1_events_source_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_mmesh_protobuf_resources_v1_events_source_proto_goTypes = []interface{}{
	(SourceType)(0),          // 0: events.SourceType
	(*Source)(nil),           // 1: events.Source
	(*SourceManager)(nil),    // 2: events.SourceManager
	(*SourceController)(nil), // 3: events.SourceController
	(*SourceRouter)(nil),     // 4: events.SourceRouter
	(*SourceNode)(nil),       // 5: events.SourceNode
	(*SourceCLI)(nil),        // 6: events.SourceCLI
	(*SourceWebUI)(nil),      // 7: events.SourceWebUI
}
var file_mmesh_protobuf_resources_v1_events_source_proto_depIdxs = []int32{
	0, // 0: events.Source.type:type_name -> events.SourceType
	3, // 1: events.Source.controller:type_name -> events.SourceController
	2, // 2: events.Source.manager:type_name -> events.SourceManager
	4, // 3: events.Source.router:type_name -> events.SourceRouter
	5, // 4: events.Source.node:type_name -> events.SourceNode
	6, // 5: events.Source.cli:type_name -> events.SourceCLI
	7, // 6: events.Source.webUI:type_name -> events.SourceWebUI
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_mmesh_protobuf_resources_v1_events_source_proto_init() }
func file_mmesh_protobuf_resources_v1_events_source_proto_init() {
	if File_mmesh_protobuf_resources_v1_events_source_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mmesh_protobuf_resources_v1_events_source_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Source); i {
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
		file_mmesh_protobuf_resources_v1_events_source_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SourceManager); i {
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
		file_mmesh_protobuf_resources_v1_events_source_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SourceController); i {
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
		file_mmesh_protobuf_resources_v1_events_source_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SourceRouter); i {
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
		file_mmesh_protobuf_resources_v1_events_source_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SourceNode); i {
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
		file_mmesh_protobuf_resources_v1_events_source_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SourceCLI); i {
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
		file_mmesh_protobuf_resources_v1_events_source_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SourceWebUI); i {
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
			RawDescriptor: file_mmesh_protobuf_resources_v1_events_source_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mmesh_protobuf_resources_v1_events_source_proto_goTypes,
		DependencyIndexes: file_mmesh_protobuf_resources_v1_events_source_proto_depIdxs,
		EnumInfos:         file_mmesh_protobuf_resources_v1_events_source_proto_enumTypes,
		MessageInfos:      file_mmesh_protobuf_resources_v1_events_source_proto_msgTypes,
	}.Build()
	File_mmesh_protobuf_resources_v1_events_source_proto = out.File
	file_mmesh_protobuf_resources_v1_events_source_proto_rawDesc = nil
	file_mmesh_protobuf_resources_v1_events_source_proto_goTypes = nil
	file_mmesh_protobuf_resources_v1_events_source_proto_depIdxs = nil
}
