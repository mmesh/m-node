// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: mmesh/protobuf/resources/v1/services/service.proto

package services

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

type Service struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceID   string `protobuf:"bytes,1,opt,name=serviceID,proto3" json:"serviceID,omitempty"`
	LastUpdated int64  `protobuf:"varint,21,opt,name=lastUpdated,proto3" json:"lastUpdated,omitempty"`
}

func (x *Service) Reset() {
	*x = Service{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_services_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Service) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Service) ProtoMessage() {}

func (x *Service) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_services_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Service.ProtoReflect.Descriptor instead.
func (*Service) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_services_service_proto_rawDescGZIP(), []int{0}
}

func (x *Service) GetServiceID() string {
	if x != nil {
		return x.ServiceID
	}
	return ""
}

func (x *Service) GetLastUpdated() int64 {
	if x != nil {
		return x.LastUpdated
	}
	return 0
}

type Services struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta     *resource.ListResponse `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	Services []*Service             `protobuf:"bytes,2,rep,name=services,proto3" json:"services,omitempty"`
}

func (x *Services) Reset() {
	*x = Services{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_services_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Services) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Services) ProtoMessage() {}

func (x *Services) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_services_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Services.ProtoReflect.Descriptor instead.
func (*Services) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_services_service_proto_rawDescGZIP(), []int{1}
}

func (x *Services) GetMeta() *resource.ListResponse {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *Services) GetServices() []*Service {
	if x != nil {
		return x.Services
	}
	return nil
}

type ListServicesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta *resource.ListRequest `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
}

func (x *ListServicesRequest) Reset() {
	*x = ListServicesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_services_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListServicesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListServicesRequest) ProtoMessage() {}

func (x *ListServicesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_services_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListServicesRequest.ProtoReflect.Descriptor instead.
func (*ListServicesRequest) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_services_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListServicesRequest) GetMeta() *resource.ListRequest {
	if x != nil {
		return x.Meta
	}
	return nil
}

var File_mmesh_protobuf_resources_v1_services_service_proto protoreflect.FileDescriptor

var file_mmesh_protobuf_resources_v1_services_service_proto_rawDesc = []byte{
	0x0a, 0x32, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x2f,
	0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x49, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x6c, 0x61, 0x73, 0x74,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x15, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6c,
	0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x22, 0x65, 0x0a, 0x08, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x2a, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04, 0x6d, 0x65,
	0x74, 0x61, 0x12, 0x2d, 0x0a, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x22, 0x40, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x6d,
	0x65, 0x74, 0x61, 0x42, 0x2c, 0x5a, 0x2a, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x64, 0x65, 0x76,
	0x2f, 0x6d, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mmesh_protobuf_resources_v1_services_service_proto_rawDescOnce sync.Once
	file_mmesh_protobuf_resources_v1_services_service_proto_rawDescData = file_mmesh_protobuf_resources_v1_services_service_proto_rawDesc
)

func file_mmesh_protobuf_resources_v1_services_service_proto_rawDescGZIP() []byte {
	file_mmesh_protobuf_resources_v1_services_service_proto_rawDescOnce.Do(func() {
		file_mmesh_protobuf_resources_v1_services_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_mmesh_protobuf_resources_v1_services_service_proto_rawDescData)
	})
	return file_mmesh_protobuf_resources_v1_services_service_proto_rawDescData
}

var file_mmesh_protobuf_resources_v1_services_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_mmesh_protobuf_resources_v1_services_service_proto_goTypes = []interface{}{
	(*Service)(nil),               // 0: services.Service
	(*Services)(nil),              // 1: services.Services
	(*ListServicesRequest)(nil),   // 2: services.ListServicesRequest
	(*resource.ListResponse)(nil), // 3: resource.ListResponse
	(*resource.ListRequest)(nil),  // 4: resource.ListRequest
}
var file_mmesh_protobuf_resources_v1_services_service_proto_depIdxs = []int32{
	3, // 0: services.Services.meta:type_name -> resource.ListResponse
	0, // 1: services.Services.services:type_name -> services.Service
	4, // 2: services.ListServicesRequest.meta:type_name -> resource.ListRequest
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_mmesh_protobuf_resources_v1_services_service_proto_init() }
func file_mmesh_protobuf_resources_v1_services_service_proto_init() {
	if File_mmesh_protobuf_resources_v1_services_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mmesh_protobuf_resources_v1_services_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Service); i {
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
		file_mmesh_protobuf_resources_v1_services_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Services); i {
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
		file_mmesh_protobuf_resources_v1_services_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListServicesRequest); i {
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
			RawDescriptor: file_mmesh_protobuf_resources_v1_services_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mmesh_protobuf_resources_v1_services_service_proto_goTypes,
		DependencyIndexes: file_mmesh_protobuf_resources_v1_services_service_proto_depIdxs,
		MessageInfos:      file_mmesh_protobuf_resources_v1_services_service_proto_msgTypes,
	}.Build()
	File_mmesh_protobuf_resources_v1_services_service_proto = out.File
	file_mmesh_protobuf_resources_v1_services_service_proto_rawDesc = nil
	file_mmesh_protobuf_resources_v1_services_service_proto_goTypes = nil
	file_mmesh_protobuf_resources_v1_services_service_proto_depIdxs = nil
}
