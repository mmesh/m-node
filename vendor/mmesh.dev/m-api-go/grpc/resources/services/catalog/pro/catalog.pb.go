// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: mmesh/protobuf/resources/v1/services/catalog/pro/catalog.proto

package pro

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

type ProductCatalog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	APIVersion          string         `protobuf:"bytes,1,opt,name=APIVersion,proto3" json:"APIVersion,omitempty"`
	Kind                string         `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	Enabled             bool           `protobuf:"varint,3,opt,name=enabled,proto3" json:"enabled,omitempty"`
	ServiceID           string         `protobuf:"bytes,11,opt,name=serviceID,proto3" json:"serviceID,omitempty"`
	ProviderID          string         `protobuf:"bytes,21,opt,name=providerID,proto3" json:"providerID,omitempty"`
	ProviderName        string         `protobuf:"bytes,22,opt,name=providerName,proto3" json:"providerName,omitempty"`
	ProviderDescription string         `protobuf:"bytes,23,opt,name=providerDescription,proto3" json:"providerDescription,omitempty"`
	ProviderWebsiteURL  string         `protobuf:"bytes,24,opt,name=providerWebsiteURL,proto3" json:"providerWebsiteURL,omitempty"`
	Products            []*ProductSpec `protobuf:"bytes,51,rep,name=products,proto3" json:"products,omitempty"`
}

func (x *ProductCatalog) Reset() {
	*x = ProductCatalog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_services_catalog_pro_catalog_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductCatalog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductCatalog) ProtoMessage() {}

func (x *ProductCatalog) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_services_catalog_pro_catalog_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductCatalog.ProtoReflect.Descriptor instead.
func (*ProductCatalog) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_services_catalog_pro_catalog_proto_rawDescGZIP(), []int{0}
}

func (x *ProductCatalog) GetAPIVersion() string {
	if x != nil {
		return x.APIVersion
	}
	return ""
}

func (x *ProductCatalog) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *ProductCatalog) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *ProductCatalog) GetServiceID() string {
	if x != nil {
		return x.ServiceID
	}
	return ""
}

func (x *ProductCatalog) GetProviderID() string {
	if x != nil {
		return x.ProviderID
	}
	return ""
}

func (x *ProductCatalog) GetProviderName() string {
	if x != nil {
		return x.ProviderName
	}
	return ""
}

func (x *ProductCatalog) GetProviderDescription() string {
	if x != nil {
		return x.ProviderDescription
	}
	return ""
}

func (x *ProductCatalog) GetProviderWebsiteURL() string {
	if x != nil {
		return x.ProviderWebsiteURL
	}
	return ""
}

func (x *ProductCatalog) GetProducts() []*ProductSpec {
	if x != nil {
		return x.Products
	}
	return nil
}

type ProductSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductID       string   `protobuf:"bytes,1,opt,name=productID,proto3" json:"productID,omitempty"`
	StripeProductID string   `protobuf:"bytes,5,opt,name=stripeProductID,proto3" json:"stripeProductID,omitempty"`
	Name            string   `protobuf:"bytes,11,opt,name=name,proto3" json:"name,omitempty"`
	Description     string   `protobuf:"bytes,12,opt,name=description,proto3" json:"description,omitempty"`
	UnitLabel       string   `protobuf:"bytes,15,opt,name=unitLabel,proto3" json:"unitLabel,omitempty"`
	Available       bool     `protobuf:"varint,21,opt,name=available,proto3" json:"available,omitempty"`
	PricingModel    string   `protobuf:"bytes,30,opt,name=pricingModel,proto3" json:"pricingModel,omitempty"`
	PriceFixed      float64  `protobuf:"fixed64,31,opt,name=priceFixed,proto3" json:"priceFixed,omitempty"`
	PriceHourly     float64  `protobuf:"fixed64,32,opt,name=priceHourly,proto3" json:"priceHourly,omitempty"`
	PriceDaily      float64  `protobuf:"fixed64,33,opt,name=priceDaily,proto3" json:"priceDaily,omitempty"`
	PriceWeekly     float64  `protobuf:"fixed64,34,opt,name=priceWeekly,proto3" json:"priceWeekly,omitempty"`
	PriceMonthly    float64  `protobuf:"fixed64,35,opt,name=priceMonthly,proto3" json:"priceMonthly,omitempty"`
	PriceYearly     float64  `protobuf:"fixed64,36,opt,name=priceYearly,proto3" json:"priceYearly,omitempty"`
	SLA             string   `protobuf:"bytes,41,opt,name=SLA,proto3" json:"SLA,omitempty"`
	Class           string   `protobuf:"bytes,51,opt,name=class,proto3" json:"class,omitempty"`
	CloudCategory   string   `protobuf:"bytes,61,opt,name=cloudCategory,proto3" json:"cloudCategory,omitempty"`
	CloudType       string   `protobuf:"bytes,62,opt,name=cloudType,proto3" json:"cloudType,omitempty"`
	ServiceCategory string   `protobuf:"bytes,71,opt,name=serviceCategory,proto3" json:"serviceCategory,omitempty"`
	ServiceType     string   `protobuf:"bytes,72,opt,name=serviceType,proto3" json:"serviceType,omitempty"`
	ServiceScope    string   `protobuf:"bytes,73,opt,name=serviceScope,proto3" json:"serviceScope,omitempty"`
	Locations       []string `protobuf:"bytes,101,rep,name=locations,proto3" json:"locations,omitempty"`
	Langs           []string `protobuf:"bytes,102,rep,name=langs,proto3" json:"langs,omitempty"`
	OpMgrs          []string `protobuf:"bytes,111,rep,name=opMgrs,proto3" json:"opMgrs,omitempty"`
	SalesReps       []string `protobuf:"bytes,112,rep,name=salesReps,proto3" json:"salesReps,omitempty"`
}

func (x *ProductSpec) Reset() {
	*x = ProductSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_services_catalog_pro_catalog_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductSpec) ProtoMessage() {}

func (x *ProductSpec) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_services_catalog_pro_catalog_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductSpec.ProtoReflect.Descriptor instead.
func (*ProductSpec) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_services_catalog_pro_catalog_proto_rawDescGZIP(), []int{1}
}

func (x *ProductSpec) GetProductID() string {
	if x != nil {
		return x.ProductID
	}
	return ""
}

func (x *ProductSpec) GetStripeProductID() string {
	if x != nil {
		return x.StripeProductID
	}
	return ""
}

func (x *ProductSpec) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProductSpec) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ProductSpec) GetUnitLabel() string {
	if x != nil {
		return x.UnitLabel
	}
	return ""
}

func (x *ProductSpec) GetAvailable() bool {
	if x != nil {
		return x.Available
	}
	return false
}

func (x *ProductSpec) GetPricingModel() string {
	if x != nil {
		return x.PricingModel
	}
	return ""
}

func (x *ProductSpec) GetPriceFixed() float64 {
	if x != nil {
		return x.PriceFixed
	}
	return 0
}

func (x *ProductSpec) GetPriceHourly() float64 {
	if x != nil {
		return x.PriceHourly
	}
	return 0
}

func (x *ProductSpec) GetPriceDaily() float64 {
	if x != nil {
		return x.PriceDaily
	}
	return 0
}

func (x *ProductSpec) GetPriceWeekly() float64 {
	if x != nil {
		return x.PriceWeekly
	}
	return 0
}

func (x *ProductSpec) GetPriceMonthly() float64 {
	if x != nil {
		return x.PriceMonthly
	}
	return 0
}

func (x *ProductSpec) GetPriceYearly() float64 {
	if x != nil {
		return x.PriceYearly
	}
	return 0
}

func (x *ProductSpec) GetSLA() string {
	if x != nil {
		return x.SLA
	}
	return ""
}

func (x *ProductSpec) GetClass() string {
	if x != nil {
		return x.Class
	}
	return ""
}

func (x *ProductSpec) GetCloudCategory() string {
	if x != nil {
		return x.CloudCategory
	}
	return ""
}

func (x *ProductSpec) GetCloudType() string {
	if x != nil {
		return x.CloudType
	}
	return ""
}

func (x *ProductSpec) GetServiceCategory() string {
	if x != nil {
		return x.ServiceCategory
	}
	return ""
}

func (x *ProductSpec) GetServiceType() string {
	if x != nil {
		return x.ServiceType
	}
	return ""
}

func (x *ProductSpec) GetServiceScope() string {
	if x != nil {
		return x.ServiceScope
	}
	return ""
}

func (x *ProductSpec) GetLocations() []string {
	if x != nil {
		return x.Locations
	}
	return nil
}

func (x *ProductSpec) GetLangs() []string {
	if x != nil {
		return x.Langs
	}
	return nil
}

func (x *ProductSpec) GetOpMgrs() []string {
	if x != nil {
		return x.OpMgrs
	}
	return nil
}

func (x *ProductSpec) GetSalesReps() []string {
	if x != nil {
		return x.SalesReps
	}
	return nil
}

var File_mmesh_protobuf_resources_v1_services_catalog_pro_catalog_proto protoreflect.FileDescriptor

var file_mmesh_protobuf_resources_v1_services_catalog_pro_catalog_proto_rawDesc = []byte{
	0x0a, 0x3e, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x70,
	0x72, 0x6f, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x03, 0x70, 0x72, 0x6f, 0x22, 0xd0, 0x02, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x12, 0x1e, 0x0a, 0x0a, 0x41, 0x50, 0x49, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x41, 0x50,
	0x49, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x49, 0x44, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x49, 0x44, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x49, 0x44, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x13, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x12, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x57, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x55, 0x52, 0x4c,
	0x18, 0x18, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x57, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x55, 0x52, 0x4c, 0x12, 0x2c, 0x0a, 0x08, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x33, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70,
	0x72, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x70, 0x65, 0x63, 0x52, 0x08,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x22, 0xfb, 0x05, 0x0a, 0x0b, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x53, 0x70, 0x65, 0x63, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x12, 0x28, 0x0a, 0x0f, 0x73, 0x74, 0x72, 0x69, 0x70, 0x65,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x73, 0x74, 0x72, 0x69, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x6e, 0x69, 0x74, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x6e, 0x69, 0x74, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c,
	0x65, 0x18, 0x15, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62,
	0x6c, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e,
	0x67, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x69, 0x63, 0x65, 0x46,
	0x69, 0x78, 0x65, 0x64, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x46, 0x69, 0x78, 0x65, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x69, 0x63, 0x65, 0x48,
	0x6f, 0x75, 0x72, 0x6c, 0x79, 0x18, 0x20, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x48, 0x6f, 0x75, 0x72, 0x6c, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x18, 0x21, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x57, 0x65, 0x65, 0x6b, 0x6c, 0x79, 0x18, 0x22, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x57, 0x65, 0x65, 0x6b, 0x6c, 0x79, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x18, 0x23, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x0c, 0x70, 0x72, 0x69, 0x63, 0x65, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x12, 0x20,
	0x0a, 0x0b, 0x70, 0x72, 0x69, 0x63, 0x65, 0x59, 0x65, 0x61, 0x72, 0x6c, 0x79, 0x18, 0x24, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x0b, 0x70, 0x72, 0x69, 0x63, 0x65, 0x59, 0x65, 0x61, 0x72, 0x6c, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x53, 0x4c, 0x41, 0x18, 0x29, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x53,
	0x4c, 0x41, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x33, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x3d, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x1c,
	0x0a, 0x09, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x54, 0x79, 0x70, 0x65, 0x18, 0x3e, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x28, 0x0a, 0x0f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18,
	0x47, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x48, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x49, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x65, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x09, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61,
	0x6e, 0x67, 0x73, 0x18, 0x66, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x6e, 0x67, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x6f, 0x70, 0x4d, 0x67, 0x72, 0x73, 0x18, 0x6f, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x06, 0x6f, 0x70, 0x4d, 0x67, 0x72, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x61, 0x6c, 0x65,
	0x73, 0x52, 0x65, 0x70, 0x73, 0x18, 0x70, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x73, 0x61, 0x6c,
	0x65, 0x73, 0x52, 0x65, 0x70, 0x73, 0x42, 0x38, 0x5a, 0x36, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2e,
	0x64, 0x65, 0x76, 0x2f, 0x6d, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x70, 0x72, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mmesh_protobuf_resources_v1_services_catalog_pro_catalog_proto_rawDescOnce sync.Once
	file_mmesh_protobuf_resources_v1_services_catalog_pro_catalog_proto_rawDescData = file_mmesh_protobuf_resources_v1_services_catalog_pro_catalog_proto_rawDesc
)

func file_mmesh_protobuf_resources_v1_services_catalog_pro_catalog_proto_rawDescGZIP() []byte {
	file_mmesh_protobuf_resources_v1_services_catalog_pro_catalog_proto_rawDescOnce.Do(func() {
		file_mmesh_protobuf_resources_v1_services_catalog_pro_catalog_proto_rawDescData = protoimpl.X.CompressGZIP(file_mmesh_protobuf_resources_v1_services_catalog_pro_catalog_proto_rawDescData)
	})
	return file_mmesh_protobuf_resources_v1_services_catalog_pro_catalog_proto_rawDescData
}

var file_mmesh_protobuf_resources_v1_services_catalog_pro_catalog_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_mmesh_protobuf_resources_v1_services_catalog_pro_catalog_proto_goTypes = []interface{}{
	(*ProductCatalog)(nil), // 0: pro.ProductCatalog
	(*ProductSpec)(nil),    // 1: pro.ProductSpec
}
var file_mmesh_protobuf_resources_v1_services_catalog_pro_catalog_proto_depIdxs = []int32{
	1, // 0: pro.ProductCatalog.products:type_name -> pro.ProductSpec
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_mmesh_protobuf_resources_v1_services_catalog_pro_catalog_proto_init() }
func file_mmesh_protobuf_resources_v1_services_catalog_pro_catalog_proto_init() {
	if File_mmesh_protobuf_resources_v1_services_catalog_pro_catalog_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mmesh_protobuf_resources_v1_services_catalog_pro_catalog_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductCatalog); i {
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
		file_mmesh_protobuf_resources_v1_services_catalog_pro_catalog_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductSpec); i {
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
			RawDescriptor: file_mmesh_protobuf_resources_v1_services_catalog_pro_catalog_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mmesh_protobuf_resources_v1_services_catalog_pro_catalog_proto_goTypes,
		DependencyIndexes: file_mmesh_protobuf_resources_v1_services_catalog_pro_catalog_proto_depIdxs,
		MessageInfos:      file_mmesh_protobuf_resources_v1_services_catalog_pro_catalog_proto_msgTypes,
	}.Build()
	File_mmesh_protobuf_resources_v1_services_catalog_pro_catalog_proto = out.File
	file_mmesh_protobuf_resources_v1_services_catalog_pro_catalog_proto_rawDesc = nil
	file_mmesh_protobuf_resources_v1_services_catalog_pro_catalog_proto_goTypes = nil
	file_mmesh_protobuf_resources_v1_services_catalog_pro_catalog_proto_depIdxs = nil
}
