// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.2
// source: mmesh/protobuf/resources/v1/location/location.proto

package location

import (
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

type NewLocationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Region           string       `protobuf:"bytes,11,opt,name=region,proto3" json:"region,omitempty"`
	Zone             string       `protobuf:"bytes,12,opt,name=zone,proto3" json:"zone,omitempty"`
	ControlZone      *FeatureZone `protobuf:"bytes,21,opt,name=controlZone,proto3" json:"controlZone,omitempty"`           // controllers available
	ConnectivityZone *FeatureZone `protobuf:"bytes,31,opt,name=connectivityZone,proto3" json:"connectivityZone,omitempty"` // routers available
}

func (x *NewLocationRequest) Reset() {
	*x = NewLocationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_location_location_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewLocationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewLocationRequest) ProtoMessage() {}

func (x *NewLocationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_location_location_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewLocationRequest.ProtoReflect.Descriptor instead.
func (*NewLocationRequest) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_location_location_proto_rawDescGZIP(), []int{0}
}

func (x *NewLocationRequest) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *NewLocationRequest) GetZone() string {
	if x != nil {
		return x.Zone
	}
	return ""
}

func (x *NewLocationRequest) GetControlZone() *FeatureZone {
	if x != nil {
		return x.ControlZone
	}
	return nil
}

func (x *NewLocationRequest) GetConnectivityZone() *FeatureZone {
	if x != nil {
		return x.ConnectivityZone
	}
	return nil
}

type UpdateLocationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LocationID       string       `protobuf:"bytes,1,opt,name=locationID,proto3" json:"locationID,omitempty"`
	ControlZone      *FeatureZone `protobuf:"bytes,21,opt,name=controlZone,proto3" json:"controlZone,omitempty"`           // controllers available
	ConnectivityZone *FeatureZone `protobuf:"bytes,31,opt,name=connectivityZone,proto3" json:"connectivityZone,omitempty"` // routers available
}

func (x *UpdateLocationRequest) Reset() {
	*x = UpdateLocationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_location_location_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateLocationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateLocationRequest) ProtoMessage() {}

func (x *UpdateLocationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_location_location_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateLocationRequest.ProtoReflect.Descriptor instead.
func (*UpdateLocationRequest) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_location_location_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateLocationRequest) GetLocationID() string {
	if x != nil {
		return x.LocationID
	}
	return ""
}

func (x *UpdateLocationRequest) GetControlZone() *FeatureZone {
	if x != nil {
		return x.ControlZone
	}
	return nil
}

func (x *UpdateLocationRequest) GetConnectivityZone() *FeatureZone {
	if x != nil {
		return x.ConnectivityZone
	}
	return nil
}

type Location struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LocationID       string       `protobuf:"bytes,1,opt,name=locationID,proto3" json:"locationID,omitempty"`
	Region           string       `protobuf:"bytes,11,opt,name=region,proto3" json:"region,omitempty"`
	Zone             string       `protobuf:"bytes,12,opt,name=zone,proto3" json:"zone,omitempty"`
	ControlZone      *FeatureZone `protobuf:"bytes,21,opt,name=controlZone,proto3" json:"controlZone,omitempty"`           // controllers available
	ConnectivityZone *FeatureZone `protobuf:"bytes,31,opt,name=connectivityZone,proto3" json:"connectivityZone,omitempty"` // routers available
}

func (x *Location) Reset() {
	*x = Location{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_location_location_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Location) ProtoMessage() {}

func (x *Location) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_location_location_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Location.ProtoReflect.Descriptor instead.
func (*Location) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_location_location_proto_rawDescGZIP(), []int{2}
}

func (x *Location) GetLocationID() string {
	if x != nil {
		return x.LocationID
	}
	return ""
}

func (x *Location) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *Location) GetZone() string {
	if x != nil {
		return x.Zone
	}
	return ""
}

func (x *Location) GetControlZone() *FeatureZone {
	if x != nil {
		return x.ControlZone
	}
	return nil
}

func (x *Location) GetConnectivityZone() *FeatureZone {
	if x != nil {
		return x.ConnectivityZone
	}
	return nil
}

type Locations struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta      *resource.ListResponse `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	Locations []*Location            `protobuf:"bytes,2,rep,name=locations,proto3" json:"locations,omitempty"`
}

func (x *Locations) Reset() {
	*x = Locations{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_location_location_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Locations) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Locations) ProtoMessage() {}

func (x *Locations) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_location_location_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Locations.ProtoReflect.Descriptor instead.
func (*Locations) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_location_location_proto_rawDescGZIP(), []int{3}
}

func (x *Locations) GetMeta() *resource.ListResponse {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *Locations) GetLocations() []*Location {
	if x != nil {
		return x.Locations
	}
	return nil
}

type ListLocationsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta *resource.ListRequest `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
}

func (x *ListLocationsRequest) Reset() {
	*x = ListLocationsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_location_location_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListLocationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListLocationsRequest) ProtoMessage() {}

func (x *ListLocationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_location_location_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListLocationsRequest.ProtoReflect.Descriptor instead.
func (*ListLocationsRequest) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_location_location_proto_rawDescGZIP(), []int{4}
}

func (x *ListLocationsRequest) GetMeta() *resource.ListRequest {
	if x != nil {
		return x.Meta
	}
	return nil
}

type FeatureZone struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Active bool `protobuf:"varint,1,opt,name=active,proto3" json:"active,omitempty"`
}

func (x *FeatureZone) Reset() {
	*x = FeatureZone{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_location_location_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeatureZone) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeatureZone) ProtoMessage() {}

func (x *FeatureZone) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_location_location_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeatureZone.ProtoReflect.Descriptor instead.
func (*FeatureZone) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_location_location_proto_rawDescGZIP(), []int{5}
}

func (x *FeatureZone) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

type LocationReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LocationID string `protobuf:"bytes,1,opt,name=locationID,proto3" json:"locationID,omitempty"`
}

func (x *LocationReq) Reset() {
	*x = LocationReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_location_location_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocationReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocationReq) ProtoMessage() {}

func (x *LocationReq) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_location_location_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocationReq.ProtoReflect.Descriptor instead.
func (*LocationReq) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_location_location_proto_rawDescGZIP(), []int{6}
}

func (x *LocationReq) GetLocationID() string {
	if x != nil {
		return x.LocationID
	}
	return ""
}

var File_mmesh_protobuf_resources_v1_location_location_proto protoreflect.FileDescriptor

var file_mmesh_protobuf_resources_v1_location_location_proto_rawDesc = []byte{
	0x0a, 0x33, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a,
	0x2f, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xbc, 0x01, 0x0a, 0x12, 0x4e, 0x65, 0x77, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12,
	0x12, 0x0a, 0x04, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x7a,
	0x6f, 0x6e, 0x65, 0x12, 0x37, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5a, 0x6f,
	0x6e, 0x65, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5a, 0x6f, 0x6e, 0x65, 0x52,
	0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5a, 0x6f, 0x6e, 0x65, 0x12, 0x41, 0x0a, 0x10,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5a, 0x6f, 0x6e, 0x65,
	0x18, 0x1f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5a, 0x6f, 0x6e, 0x65, 0x52, 0x10, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5a, 0x6f, 0x6e, 0x65, 0x22,
	0xb3, 0x01, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x37, 0x0a, 0x0b, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x5a, 0x6f, 0x6e, 0x65, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x5a, 0x6f, 0x6e, 0x65, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5a, 0x6f,
	0x6e, 0x65, 0x12, 0x41, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x5a, 0x6f, 0x6e, 0x65, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5a,
	0x6f, 0x6e, 0x65, 0x52, 0x10, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x5a, 0x6f, 0x6e, 0x65, 0x22, 0xd2, 0x01, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x7a, 0x6f,
	0x6e, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x7a, 0x6f, 0x6e, 0x65, 0x12, 0x37,
	0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5a, 0x6f, 0x6e, 0x65, 0x18, 0x15, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x46,
	0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5a, 0x6f, 0x6e, 0x65, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x5a, 0x6f, 0x6e, 0x65, 0x12, 0x41, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5a, 0x6f, 0x6e, 0x65, 0x18, 0x1f, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x46, 0x65, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x5a, 0x6f, 0x6e, 0x65, 0x52, 0x10, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5a, 0x6f, 0x6e, 0x65, 0x22, 0x69, 0x0a, 0x09, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2a, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04, 0x6d,
	0x65, 0x74, 0x61, 0x12, 0x30, 0x0a, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x41, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a,
	0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x22, 0x25, 0x0a, 0x0b, 0x46, 0x65, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x5a, 0x6f, 0x6e, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x22,
	0x2d, 0x0a, 0x0b, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x1e,
	0x0a, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x42, 0x2c,
	0x5a, 0x2a, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x6d, 0x2d, 0x61, 0x70,
	0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mmesh_protobuf_resources_v1_location_location_proto_rawDescOnce sync.Once
	file_mmesh_protobuf_resources_v1_location_location_proto_rawDescData = file_mmesh_protobuf_resources_v1_location_location_proto_rawDesc
)

func file_mmesh_protobuf_resources_v1_location_location_proto_rawDescGZIP() []byte {
	file_mmesh_protobuf_resources_v1_location_location_proto_rawDescOnce.Do(func() {
		file_mmesh_protobuf_resources_v1_location_location_proto_rawDescData = protoimpl.X.CompressGZIP(file_mmesh_protobuf_resources_v1_location_location_proto_rawDescData)
	})
	return file_mmesh_protobuf_resources_v1_location_location_proto_rawDescData
}

var file_mmesh_protobuf_resources_v1_location_location_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_mmesh_protobuf_resources_v1_location_location_proto_goTypes = []interface{}{
	(*NewLocationRequest)(nil),    // 0: location.NewLocationRequest
	(*UpdateLocationRequest)(nil), // 1: location.UpdateLocationRequest
	(*Location)(nil),              // 2: location.Location
	(*Locations)(nil),             // 3: location.Locations
	(*ListLocationsRequest)(nil),  // 4: location.ListLocationsRequest
	(*FeatureZone)(nil),           // 5: location.FeatureZone
	(*LocationReq)(nil),           // 6: location.LocationReq
	(*resource.ListResponse)(nil), // 7: resource.ListResponse
	(*resource.ListRequest)(nil),  // 8: resource.ListRequest
}
var file_mmesh_protobuf_resources_v1_location_location_proto_depIdxs = []int32{
	5, // 0: location.NewLocationRequest.controlZone:type_name -> location.FeatureZone
	5, // 1: location.NewLocationRequest.connectivityZone:type_name -> location.FeatureZone
	5, // 2: location.UpdateLocationRequest.controlZone:type_name -> location.FeatureZone
	5, // 3: location.UpdateLocationRequest.connectivityZone:type_name -> location.FeatureZone
	5, // 4: location.Location.controlZone:type_name -> location.FeatureZone
	5, // 5: location.Location.connectivityZone:type_name -> location.FeatureZone
	7, // 6: location.Locations.meta:type_name -> resource.ListResponse
	2, // 7: location.Locations.locations:type_name -> location.Location
	8, // 8: location.ListLocationsRequest.meta:type_name -> resource.ListRequest
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_mmesh_protobuf_resources_v1_location_location_proto_init() }
func file_mmesh_protobuf_resources_v1_location_location_proto_init() {
	if File_mmesh_protobuf_resources_v1_location_location_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mmesh_protobuf_resources_v1_location_location_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewLocationRequest); i {
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
		file_mmesh_protobuf_resources_v1_location_location_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateLocationRequest); i {
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
		file_mmesh_protobuf_resources_v1_location_location_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Location); i {
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
		file_mmesh_protobuf_resources_v1_location_location_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Locations); i {
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
		file_mmesh_protobuf_resources_v1_location_location_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListLocationsRequest); i {
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
		file_mmesh_protobuf_resources_v1_location_location_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeatureZone); i {
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
		file_mmesh_protobuf_resources_v1_location_location_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocationReq); i {
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
			RawDescriptor: file_mmesh_protobuf_resources_v1_location_location_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mmesh_protobuf_resources_v1_location_location_proto_goTypes,
		DependencyIndexes: file_mmesh_protobuf_resources_v1_location_location_proto_depIdxs,
		MessageInfos:      file_mmesh_protobuf_resources_v1_location_location_proto_msgTypes,
	}.Build()
	File_mmesh_protobuf_resources_v1_location_location_proto = out.File
	file_mmesh_protobuf_resources_v1_location_location_proto_rawDesc = nil
	file_mmesh_protobuf_resources_v1_location_location_proto_goTypes = nil
	file_mmesh_protobuf_resources_v1_location_location_proto_depIdxs = nil
}
