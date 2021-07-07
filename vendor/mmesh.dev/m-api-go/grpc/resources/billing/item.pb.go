// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.1
// source: mmesh/protobuf/resources/v1/billing/item.proto

package billing

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	resource "mmesh.dev/m-api-go/grpc/resources/resource"
	services "mmesh.dev/m-api-go/grpc/resources/services"
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

type BilledItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceID          string                `protobuf:"bytes,1,opt,name=serviceID,proto3" json:"serviceID,omitempty"`
	ProviderID         string                `protobuf:"bytes,2,opt,name=providerID,proto3" json:"providerID,omitempty"`
	ProductID          string                `protobuf:"bytes,3,opt,name=productID,proto3" json:"productID,omitempty"` // naming convention: providerPrefix_providerProductName
	PriceID            string                `protobuf:"bytes,4,opt,name=priceID,proto3" json:"priceID,omitempty"`     // naming convention: providerPrefix_providerPriceName
	StripeCustomerID   string                `protobuf:"bytes,11,opt,name=stripeCustomerID,proto3" json:"stripeCustomerID,omitempty"`
	StripeProductID    string                `protobuf:"bytes,12,opt,name=stripeProductID,proto3" json:"stripeProductID,omitempty"`
	StripePriceID      string                `protobuf:"bytes,13,opt,name=stripePriceID,proto3" json:"stripePriceID,omitempty"`
	AccountID          string                `protobuf:"bytes,21,opt,name=accountID,proto3" json:"accountID,omitempty"`
	BilledItemID       string                `protobuf:"bytes,22,opt,name=billedItemID,proto3" json:"billedItemID,omitempty"`             // UUID
	ResourceID         string                `protobuf:"bytes,23,opt,name=resourceID,proto3" json:"resourceID,omitempty"`                 // instanceID, contractID, dbID...
	ProductName        string                `protobuf:"bytes,24,opt,name=productName,proto3" json:"productName,omitempty"`               // kubernetes cluster, vm instance...
	Description        string                `protobuf:"bytes,25,opt,name=description,proto3" json:"description,omitempty"`               // instance.Description, serviceContract.Description...
	ProviderResourceID string                `protobuf:"bytes,31,opt,name=providerResourceID,proto3" json:"providerResourceID,omitempty"` // providerInstanceID, providerContractID...
	Class              services.ProductClass `protobuf:"varint,37,opt,name=class,proto3,enum=services.ProductClass" json:"class,omitempty"`
	CreationDate       int64                 `protobuf:"varint,41,opt,name=creationDate,proto3" json:"creationDate,omitempty"`
	LastModified       int64                 `protobuf:"varint,42,opt,name=lastModified,proto3" json:"lastModified,omitempty"`
	StartDate          int64                 `protobuf:"varint,51,opt,name=startDate,proto3" json:"startDate,omitempty"`
	EndDate            int64                 `protobuf:"varint,52,opt,name=endDate,proto3" json:"endDate,omitempty"`
	TrialStartDate     int64                 `protobuf:"varint,56,opt,name=trialStartDate,proto3" json:"trialStartDate,omitempty"`
	TrialEndDate       int64                 `protobuf:"varint,57,opt,name=trialEndDate,proto3" json:"trialEndDate,omitempty"`
	CancelAtPeriodEnd  bool                  `protobuf:"varint,61,opt,name=cancelAtPeriodEnd,proto3" json:"cancelAtPeriodEnd,omitempty"`
	CancelationDate    int64                 `protobuf:"varint,62,opt,name=cancelationDate,proto3" json:"cancelationDate,omitempty"`
	CurrentPeriodStart int64                 `protobuf:"varint,71,opt,name=currentPeriodStart,proto3" json:"currentPeriodStart,omitempty"`
	CurrentPeriodEnd   int64                 `protobuf:"varint,72,opt,name=currentPeriodEnd,proto3" json:"currentPeriodEnd,omitempty"`
}

func (x *BilledItem) Reset() {
	*x = BilledItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_billing_item_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BilledItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BilledItem) ProtoMessage() {}

func (x *BilledItem) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_billing_item_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BilledItem.ProtoReflect.Descriptor instead.
func (*BilledItem) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_billing_item_proto_rawDescGZIP(), []int{0}
}

func (x *BilledItem) GetServiceID() string {
	if x != nil {
		return x.ServiceID
	}
	return ""
}

func (x *BilledItem) GetProviderID() string {
	if x != nil {
		return x.ProviderID
	}
	return ""
}

func (x *BilledItem) GetProductID() string {
	if x != nil {
		return x.ProductID
	}
	return ""
}

func (x *BilledItem) GetPriceID() string {
	if x != nil {
		return x.PriceID
	}
	return ""
}

func (x *BilledItem) GetStripeCustomerID() string {
	if x != nil {
		return x.StripeCustomerID
	}
	return ""
}

func (x *BilledItem) GetStripeProductID() string {
	if x != nil {
		return x.StripeProductID
	}
	return ""
}

func (x *BilledItem) GetStripePriceID() string {
	if x != nil {
		return x.StripePriceID
	}
	return ""
}

func (x *BilledItem) GetAccountID() string {
	if x != nil {
		return x.AccountID
	}
	return ""
}

func (x *BilledItem) GetBilledItemID() string {
	if x != nil {
		return x.BilledItemID
	}
	return ""
}

func (x *BilledItem) GetResourceID() string {
	if x != nil {
		return x.ResourceID
	}
	return ""
}

func (x *BilledItem) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

func (x *BilledItem) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *BilledItem) GetProviderResourceID() string {
	if x != nil {
		return x.ProviderResourceID
	}
	return ""
}

func (x *BilledItem) GetClass() services.ProductClass {
	if x != nil {
		return x.Class
	}
	return services.ProductClass_NETWORK_BANDWIDTH
}

func (x *BilledItem) GetCreationDate() int64 {
	if x != nil {
		return x.CreationDate
	}
	return 0
}

func (x *BilledItem) GetLastModified() int64 {
	if x != nil {
		return x.LastModified
	}
	return 0
}

func (x *BilledItem) GetStartDate() int64 {
	if x != nil {
		return x.StartDate
	}
	return 0
}

func (x *BilledItem) GetEndDate() int64 {
	if x != nil {
		return x.EndDate
	}
	return 0
}

func (x *BilledItem) GetTrialStartDate() int64 {
	if x != nil {
		return x.TrialStartDate
	}
	return 0
}

func (x *BilledItem) GetTrialEndDate() int64 {
	if x != nil {
		return x.TrialEndDate
	}
	return 0
}

func (x *BilledItem) GetCancelAtPeriodEnd() bool {
	if x != nil {
		return x.CancelAtPeriodEnd
	}
	return false
}

func (x *BilledItem) GetCancelationDate() int64 {
	if x != nil {
		return x.CancelationDate
	}
	return 0
}

func (x *BilledItem) GetCurrentPeriodStart() int64 {
	if x != nil {
		return x.CurrentPeriodStart
	}
	return 0
}

func (x *BilledItem) GetCurrentPeriodEnd() int64 {
	if x != nil {
		return x.CurrentPeriodEnd
	}
	return 0
}

type BilledItems struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta        *resource.ListResponse `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	BilledItems []*BilledItem          `protobuf:"bytes,2,rep,name=billedItems,proto3" json:"billedItems,omitempty"`
}

func (x *BilledItems) Reset() {
	*x = BilledItems{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_billing_item_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BilledItems) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BilledItems) ProtoMessage() {}

func (x *BilledItems) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_billing_item_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BilledItems.ProtoReflect.Descriptor instead.
func (*BilledItems) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_billing_item_proto_rawDescGZIP(), []int{1}
}

func (x *BilledItems) GetMeta() *resource.ListResponse {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *BilledItems) GetBilledItems() []*BilledItem {
	if x != nil {
		return x.BilledItems
	}
	return nil
}

type ListBilledItemsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta      *resource.ListRequest `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	AccountID string                `protobuf:"bytes,2,opt,name=accountID,proto3" json:"accountID,omitempty"`
}

func (x *ListBilledItemsRequest) Reset() {
	*x = ListBilledItemsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_billing_item_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBilledItemsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBilledItemsRequest) ProtoMessage() {}

func (x *ListBilledItemsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_billing_item_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBilledItemsRequest.ProtoReflect.Descriptor instead.
func (*ListBilledItemsRequest) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_billing_item_proto_rawDescGZIP(), []int{2}
}

func (x *ListBilledItemsRequest) GetMeta() *resource.ListRequest {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *ListBilledItemsRequest) GetAccountID() string {
	if x != nil {
		return x.AccountID
	}
	return ""
}

var File_mmesh_protobuf_resources_v1_billing_item_proto protoreflect.FileDescriptor

var file_mmesh_protobuf_resources_v1_billing_item_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x69,
	0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x1a, 0x2f, 0x6d, 0x6d, 0x65, 0x73, 0x68,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f,
	0x6c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x32, 0x6d, 0x6d, 0x65, 0x73,
	0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x82,
	0x07, 0x0a, 0x0a, 0x42, 0x69, 0x6c, 0x6c, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1c, 0x0a,
	0x09, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x49, 0x44, 0x12, 0x2a, 0x0a, 0x10, 0x73, 0x74, 0x72, 0x69, 0x70, 0x65, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x44, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x73,
	0x74, 0x72, 0x69, 0x70, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x44, 0x12,
	0x28, 0x0a, 0x0f, 0x73, 0x74, 0x72, 0x69, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x49, 0x44, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x74, 0x72, 0x69, 0x70, 0x65,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x74, 0x72,
	0x69, 0x70, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x49, 0x44, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x73, 0x74, 0x72, 0x69, 0x70, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x49, 0x44, 0x12,
	0x1c, 0x0a, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x15, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x22, 0x0a,
	0x0c, 0x62, 0x69, 0x6c, 0x6c, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x44, 0x18, 0x16, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x62, 0x69, 0x6c, 0x6c, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x49,
	0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x44, 0x18,
	0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49,
	0x44, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x18, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x19, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x12, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x44, 0x18, 0x1f, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x12, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x49, 0x44, 0x12, 0x2c, 0x0a, 0x05, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x25,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x05, 0x63, 0x6c,
	0x61, 0x73, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44,
	0x61, 0x74, 0x65, 0x18, 0x29, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x4d,
	0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x2a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x6c,
	0x61, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x18, 0x33, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x64,
	0x44, 0x61, 0x74, 0x65, 0x18, 0x34, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44,
	0x61, 0x74, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x74, 0x72, 0x69, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x44, 0x61, 0x74, 0x65, 0x18, 0x38, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x74, 0x72, 0x69,
	0x61, 0x6c, 0x53, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x74,
	0x72, 0x69, 0x61, 0x6c, 0x45, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x18, 0x39, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0c, 0x74, 0x72, 0x69, 0x61, 0x6c, 0x45, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12,
	0x2c, 0x0a, 0x11, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x41, 0x74, 0x50, 0x65, 0x72, 0x69, 0x6f,
	0x64, 0x45, 0x6e, 0x64, 0x18, 0x3d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x63, 0x61, 0x6e, 0x63,
	0x65, 0x6c, 0x41, 0x74, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x45, 0x6e, 0x64, 0x12, 0x28, 0x0a,
	0x0f, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65,
	0x18, 0x3e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x2e, 0x0a, 0x12, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x53, 0x74, 0x61, 0x72, 0x74, 0x18, 0x47, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x12, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x65, 0x72, 0x69,
	0x6f, 0x64, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x2a, 0x0a, 0x10, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x45, 0x6e, 0x64, 0x18, 0x48, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x10, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64,
	0x45, 0x6e, 0x64, 0x22, 0x70, 0x0a, 0x0b, 0x42, 0x69, 0x6c, 0x6c, 0x65, 0x64, 0x49, 0x74, 0x65,
	0x6d, 0x73, 0x12, 0x2a, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x35,
	0x0a, 0x0b, 0x62, 0x69, 0x6c, 0x6c, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x42, 0x69,
	0x6c, 0x6c, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x0b, 0x62, 0x69, 0x6c, 0x6c, 0x65, 0x64,
	0x49, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x61, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x69, 0x6c,
	0x6c, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x29, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x42, 0x2b, 0x5a, 0x29, 0x6d, 0x6d, 0x65, 0x73,
	0x68, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x6d, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x62, 0x69,
	0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mmesh_protobuf_resources_v1_billing_item_proto_rawDescOnce sync.Once
	file_mmesh_protobuf_resources_v1_billing_item_proto_rawDescData = file_mmesh_protobuf_resources_v1_billing_item_proto_rawDesc
)

func file_mmesh_protobuf_resources_v1_billing_item_proto_rawDescGZIP() []byte {
	file_mmesh_protobuf_resources_v1_billing_item_proto_rawDescOnce.Do(func() {
		file_mmesh_protobuf_resources_v1_billing_item_proto_rawDescData = protoimpl.X.CompressGZIP(file_mmesh_protobuf_resources_v1_billing_item_proto_rawDescData)
	})
	return file_mmesh_protobuf_resources_v1_billing_item_proto_rawDescData
}

var file_mmesh_protobuf_resources_v1_billing_item_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_mmesh_protobuf_resources_v1_billing_item_proto_goTypes = []interface{}{
	(*BilledItem)(nil),             // 0: billing.BilledItem
	(*BilledItems)(nil),            // 1: billing.BilledItems
	(*ListBilledItemsRequest)(nil), // 2: billing.ListBilledItemsRequest
	(services.ProductClass)(0),     // 3: services.ProductClass
	(*resource.ListResponse)(nil),  // 4: resource.ListResponse
	(*resource.ListRequest)(nil),   // 5: resource.ListRequest
}
var file_mmesh_protobuf_resources_v1_billing_item_proto_depIdxs = []int32{
	3, // 0: billing.BilledItem.class:type_name -> services.ProductClass
	4, // 1: billing.BilledItems.meta:type_name -> resource.ListResponse
	0, // 2: billing.BilledItems.billedItems:type_name -> billing.BilledItem
	5, // 3: billing.ListBilledItemsRequest.meta:type_name -> resource.ListRequest
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_mmesh_protobuf_resources_v1_billing_item_proto_init() }
func file_mmesh_protobuf_resources_v1_billing_item_proto_init() {
	if File_mmesh_protobuf_resources_v1_billing_item_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mmesh_protobuf_resources_v1_billing_item_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BilledItem); i {
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
		file_mmesh_protobuf_resources_v1_billing_item_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BilledItems); i {
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
		file_mmesh_protobuf_resources_v1_billing_item_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBilledItemsRequest); i {
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
			RawDescriptor: file_mmesh_protobuf_resources_v1_billing_item_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mmesh_protobuf_resources_v1_billing_item_proto_goTypes,
		DependencyIndexes: file_mmesh_protobuf_resources_v1_billing_item_proto_depIdxs,
		MessageInfos:      file_mmesh_protobuf_resources_v1_billing_item_proto_msgTypes,
	}.Build()
	File_mmesh_protobuf_resources_v1_billing_item_proto = out.File
	file_mmesh_protobuf_resources_v1_billing_item_proto_rawDesc = nil
	file_mmesh_protobuf_resources_v1_billing_item_proto_goTypes = nil
	file_mmesh_protobuf_resources_v1_billing_item_proto_depIdxs = nil
}