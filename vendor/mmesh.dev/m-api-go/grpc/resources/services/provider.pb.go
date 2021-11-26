// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: mmesh/protobuf/resources/v1/services/provider.proto

package services

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	resource "mmesh.dev/m-api-go/grpc/resources/resource"
	catalog "mmesh.dev/m-api-go/grpc/resources/services/catalog"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ProviderType int32

const (
	ProviderType_UNSPECIFIED           ProviderType = 0
	ProviderType_CLOUD                 ProviderType = 11
	ProviderType_PROFESSIONAL_SERVICES ProviderType = 21
)

// Enum value maps for ProviderType.
var (
	ProviderType_name = map[int32]string{
		0:  "UNSPECIFIED",
		11: "CLOUD",
		21: "PROFESSIONAL_SERVICES",
	}
	ProviderType_value = map[string]int32{
		"UNSPECIFIED":           0,
		"CLOUD":                 11,
		"PROFESSIONAL_SERVICES": 21,
	}
)

func (x ProviderType) Enum() *ProviderType {
	p := new(ProviderType)
	*p = x
	return p
}

func (x ProviderType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProviderType) Descriptor() protoreflect.EnumDescriptor {
	return file_mmesh_protobuf_resources_v1_services_provider_proto_enumTypes[0].Descriptor()
}

func (ProviderType) Type() protoreflect.EnumType {
	return &file_mmesh_protobuf_resources_v1_services_provider_proto_enumTypes[0]
}

func (x ProviderType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProviderType.Descriptor instead.
func (ProviderType) EnumDescriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_services_provider_proto_rawDescGZIP(), []int{0}
}

type Provider struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceID  string `protobuf:"bytes,1,opt,name=serviceID,proto3" json:"serviceID,omitempty"`
	ProviderID string `protobuf:"bytes,2,opt,name=providerID,proto3" json:"providerID,omitempty"` // == accountID since mmesh customers can be external providers
	// 3-letter prefix:
	// mcp (mmesh cloud platform), scw (scaleway), lnd (linode), dgo (digitalOcean)...
	Prefix          string           `protobuf:"bytes,3,opt,name=prefix,proto3" json:"prefix,omitempty"`
	IsCustomer      bool             `protobuf:"varint,4,opt,name=isCustomer,proto3" json:"isCustomer,omitempty"` // mmesh customers can be external service providers
	Type            ProviderType     `protobuf:"varint,5,opt,name=type,proto3,enum=services.ProviderType" json:"type,omitempty"`
	Name            string           `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
	Description     string           `protobuf:"bytes,8,opt,name=description,proto3" json:"description,omitempty"`
	WebsiteURL      string           `protobuf:"bytes,9,opt,name=websiteURL,proto3" json:"websiteURL,omitempty"`
	Catalog         *catalog.Catalog `protobuf:"bytes,11,opt,name=catalog,proto3" json:"catalog,omitempty"`
	LastUpdated     int64            `protobuf:"varint,31,opt,name=lastUpdated,proto3" json:"lastUpdated,omitempty"`
	SupportChannels *SupportChannels `protobuf:"bytes,41,opt,name=supportChannels,proto3" json:"supportChannels,omitempty"`
	Rating          uint32           `protobuf:"varint,91,opt,name=rating,proto3" json:"rating,omitempty"`
}

func (x *Provider) Reset() {
	*x = Provider{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_services_provider_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Provider) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Provider) ProtoMessage() {}

func (x *Provider) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_services_provider_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Provider.ProtoReflect.Descriptor instead.
func (*Provider) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_services_provider_proto_rawDescGZIP(), []int{0}
}

func (x *Provider) GetServiceID() string {
	if x != nil {
		return x.ServiceID
	}
	return ""
}

func (x *Provider) GetProviderID() string {
	if x != nil {
		return x.ProviderID
	}
	return ""
}

func (x *Provider) GetPrefix() string {
	if x != nil {
		return x.Prefix
	}
	return ""
}

func (x *Provider) GetIsCustomer() bool {
	if x != nil {
		return x.IsCustomer
	}
	return false
}

func (x *Provider) GetType() ProviderType {
	if x != nil {
		return x.Type
	}
	return ProviderType_UNSPECIFIED
}

func (x *Provider) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Provider) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Provider) GetWebsiteURL() string {
	if x != nil {
		return x.WebsiteURL
	}
	return ""
}

func (x *Provider) GetCatalog() *catalog.Catalog {
	if x != nil {
		return x.Catalog
	}
	return nil
}

func (x *Provider) GetLastUpdated() int64 {
	if x != nil {
		return x.LastUpdated
	}
	return 0
}

func (x *Provider) GetSupportChannels() *SupportChannels {
	if x != nil {
		return x.SupportChannels
	}
	return nil
}

func (x *Provider) GetRating() uint32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

type Providers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta      *resource.ListResponse `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	Providers []*Provider            `protobuf:"bytes,2,rep,name=providers,proto3" json:"providers,omitempty"`
}

func (x *Providers) Reset() {
	*x = Providers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_services_provider_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Providers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Providers) ProtoMessage() {}

func (x *Providers) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_services_provider_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Providers.ProtoReflect.Descriptor instead.
func (*Providers) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_services_provider_proto_rawDescGZIP(), []int{1}
}

func (x *Providers) GetMeta() *resource.ListResponse {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *Providers) GetProviders() []*Provider {
	if x != nil {
		return x.Providers
	}
	return nil
}

type ListProvidersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta      *resource.ListRequest `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	ServiceID string                `protobuf:"bytes,2,opt,name=serviceID,proto3" json:"serviceID,omitempty"`
}

func (x *ListProvidersRequest) Reset() {
	*x = ListProvidersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_services_provider_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListProvidersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProvidersRequest) ProtoMessage() {}

func (x *ListProvidersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_services_provider_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProvidersRequest.ProtoReflect.Descriptor instead.
func (*ListProvidersRequest) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_services_provider_proto_rawDescGZIP(), []int{2}
}

func (x *ListProvidersRequest) GetMeta() *resource.ListRequest {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *ListProvidersRequest) GetServiceID() string {
	if x != nil {
		return x.ServiceID
	}
	return ""
}

type SupportChannels struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Crisp        bool   `protobuf:"varint,1,opt,name=crisp,proto3" json:"crisp,omitempty"`
	Slack        bool   `protobuf:"varint,10,opt,name=slack,proto3" json:"slack,omitempty"`
	SlackChannel string `protobuf:"bytes,11,opt,name=slackChannel,proto3" json:"slackChannel,omitempty"` // slack channel at m-gate org
}

func (x *SupportChannels) Reset() {
	*x = SupportChannels{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_services_provider_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SupportChannels) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SupportChannels) ProtoMessage() {}

func (x *SupportChannels) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_services_provider_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SupportChannels.ProtoReflect.Descriptor instead.
func (*SupportChannels) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_services_provider_proto_rawDescGZIP(), []int{3}
}

func (x *SupportChannels) GetCrisp() bool {
	if x != nil {
		return x.Crisp
	}
	return false
}

func (x *SupportChannels) GetSlack() bool {
	if x != nil {
		return x.Slack
	}
	return false
}

func (x *SupportChannels) GetSlackChannel() string {
	if x != nil {
		return x.SlackChannel
	}
	return ""
}

var File_mmesh_protobuf_resources_v1_services_provider_proto protoreflect.FileDescriptor

var file_mmesh_protobuf_resources_v1_services_provider_proto_rawDesc = []byte{
	0x0a, 0x33, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a,
	0x2f, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x3a, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x63,
	0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xad, 0x03, 0x0a,
	0x08, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69,
	0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12,
	0x1e, 0x0a, 0x0a, 0x69, 0x73, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12,
	0x2a, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x55, 0x52, 0x4c, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x55, 0x52,
	0x4c, 0x12, 0x2a, 0x0a, 0x07, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x43, 0x61, 0x74,
	0x61, 0x6c, 0x6f, 0x67, 0x52, 0x07, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x12, 0x20, 0x0a,
	0x0b, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x1f, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12,
	0x43, 0x0a, 0x0f, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x73, 0x18, 0x29, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x73, 0x52, 0x0f, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x5b,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x69, 0x0a, 0x09,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x12, 0x2a, 0x0a, 0x04, 0x6d, 0x65, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52,
	0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x30, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x09, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x22, 0x5f, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x29, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x44, 0x22, 0x61, 0x0a, 0x0f, 0x53, 0x75, 0x70, 0x70,
	0x6f, 0x72, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x72, 0x69, 0x73, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x63, 0x72, 0x69, 0x73,
	0x70, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x6c, 0x61, 0x63, 0x6b, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x05, 0x73, 0x6c, 0x61, 0x63, 0x6b, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x6c, 0x61, 0x63, 0x6b,
	0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73,
	0x6c, 0x61, 0x63, 0x6b, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2a, 0x45, 0x0a, 0x0c, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05,
	0x43, 0x4c, 0x4f, 0x55, 0x44, 0x10, 0x0b, 0x12, 0x19, 0x0a, 0x15, 0x50, 0x52, 0x4f, 0x46, 0x45,
	0x53, 0x53, 0x49, 0x4f, 0x4e, 0x41, 0x4c, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x53,
	0x10, 0x15, 0x42, 0x2c, 0x5a, 0x2a, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x64, 0x65, 0x76, 0x2f,
	0x6d, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mmesh_protobuf_resources_v1_services_provider_proto_rawDescOnce sync.Once
	file_mmesh_protobuf_resources_v1_services_provider_proto_rawDescData = file_mmesh_protobuf_resources_v1_services_provider_proto_rawDesc
)

func file_mmesh_protobuf_resources_v1_services_provider_proto_rawDescGZIP() []byte {
	file_mmesh_protobuf_resources_v1_services_provider_proto_rawDescOnce.Do(func() {
		file_mmesh_protobuf_resources_v1_services_provider_proto_rawDescData = protoimpl.X.CompressGZIP(file_mmesh_protobuf_resources_v1_services_provider_proto_rawDescData)
	})
	return file_mmesh_protobuf_resources_v1_services_provider_proto_rawDescData
}

var file_mmesh_protobuf_resources_v1_services_provider_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_mmesh_protobuf_resources_v1_services_provider_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_mmesh_protobuf_resources_v1_services_provider_proto_goTypes = []interface{}{
	(ProviderType)(0),             // 0: services.ProviderType
	(*Provider)(nil),              // 1: services.Provider
	(*Providers)(nil),             // 2: services.Providers
	(*ListProvidersRequest)(nil),  // 3: services.ListProvidersRequest
	(*SupportChannels)(nil),       // 4: services.SupportChannels
	(*catalog.Catalog)(nil),       // 5: catalog.Catalog
	(*resource.ListResponse)(nil), // 6: resource.ListResponse
	(*resource.ListRequest)(nil),  // 7: resource.ListRequest
}
var file_mmesh_protobuf_resources_v1_services_provider_proto_depIdxs = []int32{
	0, // 0: services.Provider.type:type_name -> services.ProviderType
	5, // 1: services.Provider.catalog:type_name -> catalog.Catalog
	4, // 2: services.Provider.supportChannels:type_name -> services.SupportChannels
	6, // 3: services.Providers.meta:type_name -> resource.ListResponse
	1, // 4: services.Providers.providers:type_name -> services.Provider
	7, // 5: services.ListProvidersRequest.meta:type_name -> resource.ListRequest
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_mmesh_protobuf_resources_v1_services_provider_proto_init() }
func file_mmesh_protobuf_resources_v1_services_provider_proto_init() {
	if File_mmesh_protobuf_resources_v1_services_provider_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mmesh_protobuf_resources_v1_services_provider_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Provider); i {
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
		file_mmesh_protobuf_resources_v1_services_provider_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Providers); i {
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
		file_mmesh_protobuf_resources_v1_services_provider_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListProvidersRequest); i {
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
		file_mmesh_protobuf_resources_v1_services_provider_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SupportChannels); i {
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
			RawDescriptor: file_mmesh_protobuf_resources_v1_services_provider_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mmesh_protobuf_resources_v1_services_provider_proto_goTypes,
		DependencyIndexes: file_mmesh_protobuf_resources_v1_services_provider_proto_depIdxs,
		EnumInfos:         file_mmesh_protobuf_resources_v1_services_provider_proto_enumTypes,
		MessageInfos:      file_mmesh_protobuf_resources_v1_services_provider_proto_msgTypes,
	}.Build()
	File_mmesh_protobuf_resources_v1_services_provider_proto = out.File
	file_mmesh_protobuf_resources_v1_services_provider_proto_rawDesc = nil
	file_mmesh_protobuf_resources_v1_services_provider_proto_goTypes = nil
	file_mmesh_protobuf_resources_v1_services_provider_proto_depIdxs = nil
}
