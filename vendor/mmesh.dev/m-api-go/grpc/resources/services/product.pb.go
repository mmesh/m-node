// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.1
// source: mmesh/protobuf/resources/v1/services/product.proto

package services

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ProductClass int32

const (
	ProductClass_NETWORK_BANDWIDTH                         ProductClass = 0
	ProductClass_CLOUD_COMPUTE_INSTANCE                    ProductClass = 101
	ProductClass_CLOUD_COMPUTE_KUBERNETES_CLUSTER          ProductClass = 111
	ProductClass_CLOUD_COMPUTE_KUBERNETES_NODE             ProductClass = 112
	ProductClass_CLOUD_COMPUTE_NOMAD_CLUSTER               ProductClass = 121
	ProductClass_CLOUD_APP                                 ProductClass = 131
	ProductClass_CLOUD_STORAGE_VOLUME                      ProductClass = 141
	ProductClass_CLOUD_STORAGE_DATABASE                    ProductClass = 151
	ProductClass_SERVICE_MANAGED_SERVICE                   ProductClass = 201 // monitoring, alerting, backup..
	ProductClass_SERVICE_SPECIAL_PROJECT                   ProductClass = 202
	ProductClass_SERVICE_CLOUD_MIGRATION                   ProductClass = 203
	ProductClass_SERVICE_SUPPORT                           ProductClass = 204
	ProductClass_SERVICE_INFRASTRUCTURE_MAINTENANCE        ProductClass = 205 // upgrades, patches..
	ProductClass_SERVICE_INFRASTRUCTURE_BUILDING           ProductClass = 206 // VM, container builds (mongodb, jenkins..)
	ProductClass_SERVICE_INCIDENT_RESPONSE                 ProductClass = 207
	ProductClass_SERVICE_SECURITY_ASSESMENT                ProductClass = 208 // security & virus scans, reports..
	ProductClass_SERVICE_HEALTH_CHECKING                   ProductClass = 209
	ProductClass_SERVICE_COMPLIANCE_ASSESMENT              ProductClass = 210
	ProductClass_SOFTWARE_DEVELOPMENT_PROJECT              ProductClass = 301
	ProductClass_SOFTWARE_DEVELOPMENT_SUPPORT              ProductClass = 302
	ProductClass_SOFTWARE_DEVELOPMENT_MAINTENANCE          ProductClass = 303 // upgrades, patches..
	ProductClass_SOFTWARE_DEVELOPMENT_FEATURE_REQUEST      ProductClass = 304
	ProductClass_SOFTWARE_DEVELOPMENT_INCIDENT_RESPONSE    ProductClass = 305 // bugs, performance issues...
	ProductClass_SOFTWARE_DEVELOPMENT_SECURITY_ASSESMENT   ProductClass = 306
	ProductClass_SOFTWARE_DEVELOPMENT_HEALTH_CHECKING      ProductClass = 307
	ProductClass_SOFTWARE_DEVELOPMENT_COMPLIANCE_ASSESMENT ProductClass = 308
)

// Enum value maps for ProductClass.
var (
	ProductClass_name = map[int32]string{
		0:   "NETWORK_BANDWIDTH",
		101: "CLOUD_COMPUTE_INSTANCE",
		111: "CLOUD_COMPUTE_KUBERNETES_CLUSTER",
		112: "CLOUD_COMPUTE_KUBERNETES_NODE",
		121: "CLOUD_COMPUTE_NOMAD_CLUSTER",
		131: "CLOUD_APP",
		141: "CLOUD_STORAGE_VOLUME",
		151: "CLOUD_STORAGE_DATABASE",
		201: "SERVICE_MANAGED_SERVICE",
		202: "SERVICE_SPECIAL_PROJECT",
		203: "SERVICE_CLOUD_MIGRATION",
		204: "SERVICE_SUPPORT",
		205: "SERVICE_INFRASTRUCTURE_MAINTENANCE",
		206: "SERVICE_INFRASTRUCTURE_BUILDING",
		207: "SERVICE_INCIDENT_RESPONSE",
		208: "SERVICE_SECURITY_ASSESMENT",
		209: "SERVICE_HEALTH_CHECKING",
		210: "SERVICE_COMPLIANCE_ASSESMENT",
		301: "SOFTWARE_DEVELOPMENT_PROJECT",
		302: "SOFTWARE_DEVELOPMENT_SUPPORT",
		303: "SOFTWARE_DEVELOPMENT_MAINTENANCE",
		304: "SOFTWARE_DEVELOPMENT_FEATURE_REQUEST",
		305: "SOFTWARE_DEVELOPMENT_INCIDENT_RESPONSE",
		306: "SOFTWARE_DEVELOPMENT_SECURITY_ASSESMENT",
		307: "SOFTWARE_DEVELOPMENT_HEALTH_CHECKING",
		308: "SOFTWARE_DEVELOPMENT_COMPLIANCE_ASSESMENT",
	}
	ProductClass_value = map[string]int32{
		"NETWORK_BANDWIDTH":                         0,
		"CLOUD_COMPUTE_INSTANCE":                    101,
		"CLOUD_COMPUTE_KUBERNETES_CLUSTER":          111,
		"CLOUD_COMPUTE_KUBERNETES_NODE":             112,
		"CLOUD_COMPUTE_NOMAD_CLUSTER":               121,
		"CLOUD_APP":                                 131,
		"CLOUD_STORAGE_VOLUME":                      141,
		"CLOUD_STORAGE_DATABASE":                    151,
		"SERVICE_MANAGED_SERVICE":                   201,
		"SERVICE_SPECIAL_PROJECT":                   202,
		"SERVICE_CLOUD_MIGRATION":                   203,
		"SERVICE_SUPPORT":                           204,
		"SERVICE_INFRASTRUCTURE_MAINTENANCE":        205,
		"SERVICE_INFRASTRUCTURE_BUILDING":           206,
		"SERVICE_INCIDENT_RESPONSE":                 207,
		"SERVICE_SECURITY_ASSESMENT":                208,
		"SERVICE_HEALTH_CHECKING":                   209,
		"SERVICE_COMPLIANCE_ASSESMENT":              210,
		"SOFTWARE_DEVELOPMENT_PROJECT":              301,
		"SOFTWARE_DEVELOPMENT_SUPPORT":              302,
		"SOFTWARE_DEVELOPMENT_MAINTENANCE":          303,
		"SOFTWARE_DEVELOPMENT_FEATURE_REQUEST":      304,
		"SOFTWARE_DEVELOPMENT_INCIDENT_RESPONSE":    305,
		"SOFTWARE_DEVELOPMENT_SECURITY_ASSESMENT":   306,
		"SOFTWARE_DEVELOPMENT_HEALTH_CHECKING":      307,
		"SOFTWARE_DEVELOPMENT_COMPLIANCE_ASSESMENT": 308,
	}
)

func (x ProductClass) Enum() *ProductClass {
	p := new(ProductClass)
	*p = x
	return p
}

func (x ProductClass) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProductClass) Descriptor() protoreflect.EnumDescriptor {
	return file_mmesh_protobuf_resources_v1_services_product_proto_enumTypes[0].Descriptor()
}

func (ProductClass) Type() protoreflect.EnumType {
	return &file_mmesh_protobuf_resources_v1_services_product_proto_enumTypes[0]
}

func (x ProductClass) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProductClass.Descriptor instead.
func (ProductClass) EnumDescriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_services_product_proto_rawDescGZIP(), []int{0}
}

type Product struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceID       string       `protobuf:"bytes,1,opt,name=serviceID,proto3" json:"serviceID,omitempty"`
	ProviderID      string       `protobuf:"bytes,2,opt,name=providerID,proto3" json:"providerID,omitempty"`
	ProductID       string       `protobuf:"bytes,3,opt,name=productID,proto3" json:"productID,omitempty"` // naming convention: providerPrefix_providerProductName
	StripeProductID string       `protobuf:"bytes,11,opt,name=stripeProductID,proto3" json:"stripeProductID,omitempty"`
	Name            string       `protobuf:"bytes,21,opt,name=name,proto3" json:"name,omitempty"`
	Description     string       `protobuf:"bytes,22,opt,name=description,proto3" json:"description,omitempty"`
	UnitLabel       string       `protobuf:"bytes,25,opt,name=unitLabel,proto3" json:"unitLabel,omitempty"`
	Class           ProductClass `protobuf:"varint,31,opt,name=class,proto3,enum=services.ProductClass" json:"class,omitempty"`
	CreationDate    int64        `protobuf:"varint,61,opt,name=creationDate,proto3" json:"creationDate,omitempty"`
	LastModified    int64        `protobuf:"varint,62,opt,name=lastModified,proto3" json:"lastModified,omitempty"`
}

func (x *Product) Reset() {
	*x = Product{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_services_product_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_services_product_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product.ProtoReflect.Descriptor instead.
func (*Product) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_services_product_proto_rawDescGZIP(), []int{0}
}

func (x *Product) GetServiceID() string {
	if x != nil {
		return x.ServiceID
	}
	return ""
}

func (x *Product) GetProviderID() string {
	if x != nil {
		return x.ProviderID
	}
	return ""
}

func (x *Product) GetProductID() string {
	if x != nil {
		return x.ProductID
	}
	return ""
}

func (x *Product) GetStripeProductID() string {
	if x != nil {
		return x.StripeProductID
	}
	return ""
}

func (x *Product) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Product) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Product) GetUnitLabel() string {
	if x != nil {
		return x.UnitLabel
	}
	return ""
}

func (x *Product) GetClass() ProductClass {
	if x != nil {
		return x.Class
	}
	return ProductClass_NETWORK_BANDWIDTH
}

func (x *Product) GetCreationDate() int64 {
	if x != nil {
		return x.CreationDate
	}
	return 0
}

func (x *Product) GetLastModified() int64 {
	if x != nil {
		return x.LastModified
	}
	return 0
}

type Products struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta     *resource.ListResponse `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	Products []*Product             `protobuf:"bytes,2,rep,name=products,proto3" json:"products,omitempty"`
}

func (x *Products) Reset() {
	*x = Products{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_services_product_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Products) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Products) ProtoMessage() {}

func (x *Products) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_services_product_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Products.ProtoReflect.Descriptor instead.
func (*Products) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_services_product_proto_rawDescGZIP(), []int{1}
}

func (x *Products) GetMeta() *resource.ListResponse {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *Products) GetProducts() []*Product {
	if x != nil {
		return x.Products
	}
	return nil
}

type ListProductsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta     *resource.ListRequest `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	Provider *Provider             `protobuf:"bytes,2,opt,name=provider,proto3" json:"provider,omitempty"`
}

func (x *ListProductsRequest) Reset() {
	*x = ListProductsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_services_product_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListProductsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProductsRequest) ProtoMessage() {}

func (x *ListProductsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_services_product_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProductsRequest.ProtoReflect.Descriptor instead.
func (*ListProductsRequest) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_services_product_proto_rawDescGZIP(), []int{2}
}

func (x *ListProductsRequest) GetMeta() *resource.ListRequest {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *ListProductsRequest) GetProvider() *Provider {
	if x != nil {
		return x.Provider
	}
	return nil
}

var File_mmesh_protobuf_resources_v1_services_product_proto protoreflect.FileDescriptor

var file_mmesh_protobuf_resources_v1_services_product_proto_rawDesc = []byte{
	0x0a, 0x32, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x2f,
	0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x33, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd9, 0x02, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x44, 0x12, 0x1e,
	0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1c,
	0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x12, 0x28, 0x0a, 0x0f,
	0x73, 0x74, 0x72, 0x69, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x74, 0x72, 0x69, 0x70, 0x65, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x15,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09,
	0x75, 0x6e, 0x69, 0x74, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x19, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x75, 0x6e, 0x69, 0x74, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x2c, 0x0a, 0x05, 0x63, 0x6c,
	0x61, 0x73, 0x73, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x6c, 0x61, 0x73,
	0x73, 0x52, 0x05, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x18, 0x3d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x22, 0x0a, 0x0c,
	0x6c, 0x61, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x3e, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64,
	0x22, 0x65, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x2a, 0x0a, 0x04,
	0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x2d, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x08, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x22, 0x70, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29,
	0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x2e, 0x0a, 0x08, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52,
	0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2a, 0x82, 0x07, 0x0a, 0x0c, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x15, 0x0a, 0x11, 0x4e, 0x45,
	0x54, 0x57, 0x4f, 0x52, 0x4b, 0x5f, 0x42, 0x41, 0x4e, 0x44, 0x57, 0x49, 0x44, 0x54, 0x48, 0x10,
	0x00, 0x12, 0x1a, 0x0a, 0x16, 0x43, 0x4c, 0x4f, 0x55, 0x44, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x55,
	0x54, 0x45, 0x5f, 0x49, 0x4e, 0x53, 0x54, 0x41, 0x4e, 0x43, 0x45, 0x10, 0x65, 0x12, 0x24, 0x0a,
	0x20, 0x43, 0x4c, 0x4f, 0x55, 0x44, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x55, 0x54, 0x45, 0x5f, 0x4b,
	0x55, 0x42, 0x45, 0x52, 0x4e, 0x45, 0x54, 0x45, 0x53, 0x5f, 0x43, 0x4c, 0x55, 0x53, 0x54, 0x45,
	0x52, 0x10, 0x6f, 0x12, 0x21, 0x0a, 0x1d, 0x43, 0x4c, 0x4f, 0x55, 0x44, 0x5f, 0x43, 0x4f, 0x4d,
	0x50, 0x55, 0x54, 0x45, 0x5f, 0x4b, 0x55, 0x42, 0x45, 0x52, 0x4e, 0x45, 0x54, 0x45, 0x53, 0x5f,
	0x4e, 0x4f, 0x44, 0x45, 0x10, 0x70, 0x12, 0x1f, 0x0a, 0x1b, 0x43, 0x4c, 0x4f, 0x55, 0x44, 0x5f,
	0x43, 0x4f, 0x4d, 0x50, 0x55, 0x54, 0x45, 0x5f, 0x4e, 0x4f, 0x4d, 0x41, 0x44, 0x5f, 0x43, 0x4c,
	0x55, 0x53, 0x54, 0x45, 0x52, 0x10, 0x79, 0x12, 0x0e, 0x0a, 0x09, 0x43, 0x4c, 0x4f, 0x55, 0x44,
	0x5f, 0x41, 0x50, 0x50, 0x10, 0x83, 0x01, 0x12, 0x19, 0x0a, 0x14, 0x43, 0x4c, 0x4f, 0x55, 0x44,
	0x5f, 0x53, 0x54, 0x4f, 0x52, 0x41, 0x47, 0x45, 0x5f, 0x56, 0x4f, 0x4c, 0x55, 0x4d, 0x45, 0x10,
	0x8d, 0x01, 0x12, 0x1b, 0x0a, 0x16, 0x43, 0x4c, 0x4f, 0x55, 0x44, 0x5f, 0x53, 0x54, 0x4f, 0x52,
	0x41, 0x47, 0x45, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x42, 0x41, 0x53, 0x45, 0x10, 0x97, 0x01, 0x12,
	0x1c, 0x0a, 0x17, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x4d, 0x41, 0x4e, 0x41, 0x47,
	0x45, 0x44, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x10, 0xc9, 0x01, 0x12, 0x1c, 0x0a,
	0x17, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x53, 0x50, 0x45, 0x43, 0x49, 0x41, 0x4c,
	0x5f, 0x50, 0x52, 0x4f, 0x4a, 0x45, 0x43, 0x54, 0x10, 0xca, 0x01, 0x12, 0x1c, 0x0a, 0x17, 0x53,
	0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x43, 0x4c, 0x4f, 0x55, 0x44, 0x5f, 0x4d, 0x49, 0x47,
	0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0xcb, 0x01, 0x12, 0x14, 0x0a, 0x0f, 0x53, 0x45, 0x52,
	0x56, 0x49, 0x43, 0x45, 0x5f, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x10, 0xcc, 0x01, 0x12,
	0x27, 0x0a, 0x22, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x49, 0x4e, 0x46, 0x52, 0x41,
	0x53, 0x54, 0x52, 0x55, 0x43, 0x54, 0x55, 0x52, 0x45, 0x5f, 0x4d, 0x41, 0x49, 0x4e, 0x54, 0x45,
	0x4e, 0x41, 0x4e, 0x43, 0x45, 0x10, 0xcd, 0x01, 0x12, 0x24, 0x0a, 0x1f, 0x53, 0x45, 0x52, 0x56,
	0x49, 0x43, 0x45, 0x5f, 0x49, 0x4e, 0x46, 0x52, 0x41, 0x53, 0x54, 0x52, 0x55, 0x43, 0x54, 0x55,
	0x52, 0x45, 0x5f, 0x42, 0x55, 0x49, 0x4c, 0x44, 0x49, 0x4e, 0x47, 0x10, 0xce, 0x01, 0x12, 0x1e,
	0x0a, 0x19, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x49, 0x4e, 0x43, 0x49, 0x44, 0x45,
	0x4e, 0x54, 0x5f, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x10, 0xcf, 0x01, 0x12, 0x1f,
	0x0a, 0x1a, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x53, 0x45, 0x43, 0x55, 0x52, 0x49,
	0x54, 0x59, 0x5f, 0x41, 0x53, 0x53, 0x45, 0x53, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0xd0, 0x01, 0x12,
	0x1c, 0x0a, 0x17, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x48, 0x45, 0x41, 0x4c, 0x54,
	0x48, 0x5f, 0x43, 0x48, 0x45, 0x43, 0x4b, 0x49, 0x4e, 0x47, 0x10, 0xd1, 0x01, 0x12, 0x21, 0x0a,
	0x1c, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x49, 0x41,
	0x4e, 0x43, 0x45, 0x5f, 0x41, 0x53, 0x53, 0x45, 0x53, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0xd2, 0x01,
	0x12, 0x21, 0x0a, 0x1c, 0x53, 0x4f, 0x46, 0x54, 0x57, 0x41, 0x52, 0x45, 0x5f, 0x44, 0x45, 0x56,
	0x45, 0x4c, 0x4f, 0x50, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x50, 0x52, 0x4f, 0x4a, 0x45, 0x43, 0x54,
	0x10, 0xad, 0x02, 0x12, 0x21, 0x0a, 0x1c, 0x53, 0x4f, 0x46, 0x54, 0x57, 0x41, 0x52, 0x45, 0x5f,
	0x44, 0x45, 0x56, 0x45, 0x4c, 0x4f, 0x50, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x55, 0x50, 0x50,
	0x4f, 0x52, 0x54, 0x10, 0xae, 0x02, 0x12, 0x25, 0x0a, 0x20, 0x53, 0x4f, 0x46, 0x54, 0x57, 0x41,
	0x52, 0x45, 0x5f, 0x44, 0x45, 0x56, 0x45, 0x4c, 0x4f, 0x50, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x4d,
	0x41, 0x49, 0x4e, 0x54, 0x45, 0x4e, 0x41, 0x4e, 0x43, 0x45, 0x10, 0xaf, 0x02, 0x12, 0x29, 0x0a,
	0x24, 0x53, 0x4f, 0x46, 0x54, 0x57, 0x41, 0x52, 0x45, 0x5f, 0x44, 0x45, 0x56, 0x45, 0x4c, 0x4f,
	0x50, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x46, 0x45, 0x41, 0x54, 0x55, 0x52, 0x45, 0x5f, 0x52, 0x45,
	0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0xb0, 0x02, 0x12, 0x2b, 0x0a, 0x26, 0x53, 0x4f, 0x46, 0x54,
	0x57, 0x41, 0x52, 0x45, 0x5f, 0x44, 0x45, 0x56, 0x45, 0x4c, 0x4f, 0x50, 0x4d, 0x45, 0x4e, 0x54,
	0x5f, 0x49, 0x4e, 0x43, 0x49, 0x44, 0x45, 0x4e, 0x54, 0x5f, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e,
	0x53, 0x45, 0x10, 0xb1, 0x02, 0x12, 0x2c, 0x0a, 0x27, 0x53, 0x4f, 0x46, 0x54, 0x57, 0x41, 0x52,
	0x45, 0x5f, 0x44, 0x45, 0x56, 0x45, 0x4c, 0x4f, 0x50, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x45,
	0x43, 0x55, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x41, 0x53, 0x53, 0x45, 0x53, 0x4d, 0x45, 0x4e, 0x54,
	0x10, 0xb2, 0x02, 0x12, 0x29, 0x0a, 0x24, 0x53, 0x4f, 0x46, 0x54, 0x57, 0x41, 0x52, 0x45, 0x5f,
	0x44, 0x45, 0x56, 0x45, 0x4c, 0x4f, 0x50, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x48, 0x45, 0x41, 0x4c,
	0x54, 0x48, 0x5f, 0x43, 0x48, 0x45, 0x43, 0x4b, 0x49, 0x4e, 0x47, 0x10, 0xb3, 0x02, 0x12, 0x2e,
	0x0a, 0x29, 0x53, 0x4f, 0x46, 0x54, 0x57, 0x41, 0x52, 0x45, 0x5f, 0x44, 0x45, 0x56, 0x45, 0x4c,
	0x4f, 0x50, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x49, 0x41, 0x4e, 0x43,
	0x45, 0x5f, 0x41, 0x53, 0x53, 0x45, 0x53, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0xb4, 0x02, 0x42, 0x2c,
	0x5a, 0x2a, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x6d, 0x2d, 0x61, 0x70,
	0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mmesh_protobuf_resources_v1_services_product_proto_rawDescOnce sync.Once
	file_mmesh_protobuf_resources_v1_services_product_proto_rawDescData = file_mmesh_protobuf_resources_v1_services_product_proto_rawDesc
)

func file_mmesh_protobuf_resources_v1_services_product_proto_rawDescGZIP() []byte {
	file_mmesh_protobuf_resources_v1_services_product_proto_rawDescOnce.Do(func() {
		file_mmesh_protobuf_resources_v1_services_product_proto_rawDescData = protoimpl.X.CompressGZIP(file_mmesh_protobuf_resources_v1_services_product_proto_rawDescData)
	})
	return file_mmesh_protobuf_resources_v1_services_product_proto_rawDescData
}

var file_mmesh_protobuf_resources_v1_services_product_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_mmesh_protobuf_resources_v1_services_product_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_mmesh_protobuf_resources_v1_services_product_proto_goTypes = []interface{}{
	(ProductClass)(0),             // 0: services.ProductClass
	(*Product)(nil),               // 1: services.Product
	(*Products)(nil),              // 2: services.Products
	(*ListProductsRequest)(nil),   // 3: services.ListProductsRequest
	(*resource.ListResponse)(nil), // 4: resource.ListResponse
	(*resource.ListRequest)(nil),  // 5: resource.ListRequest
	(*Provider)(nil),              // 6: services.Provider
}
var file_mmesh_protobuf_resources_v1_services_product_proto_depIdxs = []int32{
	0, // 0: services.Product.class:type_name -> services.ProductClass
	4, // 1: services.Products.meta:type_name -> resource.ListResponse
	1, // 2: services.Products.products:type_name -> services.Product
	5, // 3: services.ListProductsRequest.meta:type_name -> resource.ListRequest
	6, // 4: services.ListProductsRequest.provider:type_name -> services.Provider
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_mmesh_protobuf_resources_v1_services_product_proto_init() }
func file_mmesh_protobuf_resources_v1_services_product_proto_init() {
	if File_mmesh_protobuf_resources_v1_services_product_proto != nil {
		return
	}
	file_mmesh_protobuf_resources_v1_services_provider_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_mmesh_protobuf_resources_v1_services_product_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Product); i {
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
		file_mmesh_protobuf_resources_v1_services_product_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Products); i {
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
		file_mmesh_protobuf_resources_v1_services_product_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListProductsRequest); i {
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
			RawDescriptor: file_mmesh_protobuf_resources_v1_services_product_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mmesh_protobuf_resources_v1_services_product_proto_goTypes,
		DependencyIndexes: file_mmesh_protobuf_resources_v1_services_product_proto_depIdxs,
		EnumInfos:         file_mmesh_protobuf_resources_v1_services_product_proto_enumTypes,
		MessageInfos:      file_mmesh_protobuf_resources_v1_services_product_proto_msgTypes,
	}.Build()
	File_mmesh_protobuf_resources_v1_services_product_proto = out.File
	file_mmesh_protobuf_resources_v1_services_product_proto_rawDesc = nil
	file_mmesh_protobuf_resources_v1_services_product_proto_goTypes = nil
	file_mmesh_protobuf_resources_v1_services_product_proto_depIdxs = nil
}
