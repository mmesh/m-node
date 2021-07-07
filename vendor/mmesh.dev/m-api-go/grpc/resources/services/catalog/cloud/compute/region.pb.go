// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.1
// source: mmesh/protobuf/resources/v1/services/catalog/cloud/compute/region.proto

package compute

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Region struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RegionID        string   `protobuf:"bytes,1,opt,name=regionID,proto3" json:"regionID,omitempty"`
	Name            string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	InstanceTypeIDs []string `protobuf:"bytes,5,rep,name=instanceTypeIDs,proto3" json:"instanceTypeIDs,omitempty"`
	Available       bool     `protobuf:"varint,11,opt,name=available,proto3" json:"available,omitempty"`
	Features        []string `protobuf:"bytes,21,rep,name=features,proto3" json:"features,omitempty"`
}

func (x *Region) Reset() {
	*x = Region{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_services_catalog_cloud_compute_region_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Region) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Region) ProtoMessage() {}

func (x *Region) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_services_catalog_cloud_compute_region_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Region.ProtoReflect.Descriptor instead.
func (*Region) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_services_catalog_cloud_compute_region_proto_rawDescGZIP(), []int{0}
}

func (x *Region) GetRegionID() string {
	if x != nil {
		return x.RegionID
	}
	return ""
}

func (x *Region) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Region) GetInstanceTypeIDs() []string {
	if x != nil {
		return x.InstanceTypeIDs
	}
	return nil
}

func (x *Region) GetAvailable() bool {
	if x != nil {
		return x.Available
	}
	return false
}

func (x *Region) GetFeatures() []string {
	if x != nil {
		return x.Features
	}
	return nil
}

var File_mmesh_protobuf_resources_v1_services_catalog_cloud_compute_region_proto protoreflect.FileDescriptor

var file_mmesh_protobuf_resources_v1_services_catalog_cloud_compute_region_proto_rawDesc = []byte{
	0x0a, 0x47, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2f, 0x72, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x75,
	0x74, 0x65, 0x22, 0x9c, 0x01, 0x0a, 0x06, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a,
	0x08, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a,
	0x0f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x73,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x76, 0x61, 0x69, 0x6c,
	0x61, 0x62, 0x6c, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x61, 0x76, 0x61, 0x69,
	0x6c, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x73, 0x18, 0x15, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x73, 0x42, 0x42, 0x5a, 0x40, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x6d,
	0x2d, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f,
	0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x63, 0x6f,
	0x6d, 0x70, 0x75, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mmesh_protobuf_resources_v1_services_catalog_cloud_compute_region_proto_rawDescOnce sync.Once
	file_mmesh_protobuf_resources_v1_services_catalog_cloud_compute_region_proto_rawDescData = file_mmesh_protobuf_resources_v1_services_catalog_cloud_compute_region_proto_rawDesc
)

func file_mmesh_protobuf_resources_v1_services_catalog_cloud_compute_region_proto_rawDescGZIP() []byte {
	file_mmesh_protobuf_resources_v1_services_catalog_cloud_compute_region_proto_rawDescOnce.Do(func() {
		file_mmesh_protobuf_resources_v1_services_catalog_cloud_compute_region_proto_rawDescData = protoimpl.X.CompressGZIP(file_mmesh_protobuf_resources_v1_services_catalog_cloud_compute_region_proto_rawDescData)
	})
	return file_mmesh_protobuf_resources_v1_services_catalog_cloud_compute_region_proto_rawDescData
}

var file_mmesh_protobuf_resources_v1_services_catalog_cloud_compute_region_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_mmesh_protobuf_resources_v1_services_catalog_cloud_compute_region_proto_goTypes = []interface{}{
	(*Region)(nil), // 0: compute.Region
}
var file_mmesh_protobuf_resources_v1_services_catalog_cloud_compute_region_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_mmesh_protobuf_resources_v1_services_catalog_cloud_compute_region_proto_init() }
func file_mmesh_protobuf_resources_v1_services_catalog_cloud_compute_region_proto_init() {
	if File_mmesh_protobuf_resources_v1_services_catalog_cloud_compute_region_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mmesh_protobuf_resources_v1_services_catalog_cloud_compute_region_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Region); i {
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
			RawDescriptor: file_mmesh_protobuf_resources_v1_services_catalog_cloud_compute_region_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mmesh_protobuf_resources_v1_services_catalog_cloud_compute_region_proto_goTypes,
		DependencyIndexes: file_mmesh_protobuf_resources_v1_services_catalog_cloud_compute_region_proto_depIdxs,
		MessageInfos:      file_mmesh_protobuf_resources_v1_services_catalog_cloud_compute_region_proto_msgTypes,
	}.Build()
	File_mmesh_protobuf_resources_v1_services_catalog_cloud_compute_region_proto = out.File
	file_mmesh_protobuf_resources_v1_services_catalog_cloud_compute_region_proto_rawDesc = nil
	file_mmesh_protobuf_resources_v1_services_catalog_cloud_compute_region_proto_goTypes = nil
	file_mmesh_protobuf_resources_v1_services_catalog_cloud_compute_region_proto_depIdxs = nil
}
