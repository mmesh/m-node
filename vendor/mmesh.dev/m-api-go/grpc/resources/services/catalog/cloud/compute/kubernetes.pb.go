// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: mmesh/protobuf/resources/v1/services/catalog/cloud/compute/kubernetes.proto

package compute

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

type KubernetesOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Versions          []*KubernetesVersion          `protobuf:"bytes,11,rep,name=versions,proto3" json:"versions,omitempty"`
	Regions           []*KubernetesRegion           `protobuf:"bytes,21,rep,name=regions,proto3" json:"regions,omitempty"`
	NodeInstanceTypes []*KubernetesNodeInstanceType `protobuf:"bytes,31,rep,name=nodeInstanceTypes,proto3" json:"nodeInstanceTypes,omitempty"`
}

func (x *KubernetesOptions) Reset() {
	*x = KubernetesOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_services_catalog_cloud_compute_kubernetes_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KubernetesOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KubernetesOptions) ProtoMessage() {}

func (x *KubernetesOptions) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_services_catalog_cloud_compute_kubernetes_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KubernetesOptions.ProtoReflect.Descriptor instead.
func (*KubernetesOptions) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_services_catalog_cloud_compute_kubernetes_proto_rawDescGZIP(), []int{0}
}

func (x *KubernetesOptions) GetVersions() []*KubernetesVersion {
	if x != nil {
		return x.Versions
	}
	return nil
}

func (x *KubernetesOptions) GetRegions() []*KubernetesRegion {
	if x != nil {
		return x.Regions
	}
	return nil
}

func (x *KubernetesOptions) GetNodeInstanceTypes() []*KubernetesNodeInstanceType {
	if x != nil {
		return x.NodeInstanceTypes
	}
	return nil
}

type KubernetesVersion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VersionID     string `protobuf:"bytes,1,opt,name=versionID,proto3" json:"versionID,omitempty"`
	VersionNumber string `protobuf:"bytes,11,opt,name=versionNumber,proto3" json:"versionNumber,omitempty"`
}

func (x *KubernetesVersion) Reset() {
	*x = KubernetesVersion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_services_catalog_cloud_compute_kubernetes_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KubernetesVersion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KubernetesVersion) ProtoMessage() {}

func (x *KubernetesVersion) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_services_catalog_cloud_compute_kubernetes_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KubernetesVersion.ProtoReflect.Descriptor instead.
func (*KubernetesVersion) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_services_catalog_cloud_compute_kubernetes_proto_rawDescGZIP(), []int{1}
}

func (x *KubernetesVersion) GetVersionID() string {
	if x != nil {
		return x.VersionID
	}
	return ""
}

func (x *KubernetesVersion) GetVersionNumber() string {
	if x != nil {
		return x.VersionNumber
	}
	return ""
}

type KubernetesRegion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RegionID   string `protobuf:"bytes,1,opt,name=regionID,proto3" json:"regionID,omitempty"`
	RegionName string `protobuf:"bytes,11,opt,name=regionName,proto3" json:"regionName,omitempty"`
}

func (x *KubernetesRegion) Reset() {
	*x = KubernetesRegion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_services_catalog_cloud_compute_kubernetes_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KubernetesRegion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KubernetesRegion) ProtoMessage() {}

func (x *KubernetesRegion) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_services_catalog_cloud_compute_kubernetes_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KubernetesRegion.ProtoReflect.Descriptor instead.
func (*KubernetesRegion) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_services_catalog_cloud_compute_kubernetes_proto_rawDescGZIP(), []int{2}
}

func (x *KubernetesRegion) GetRegionID() string {
	if x != nil {
		return x.RegionID
	}
	return ""
}

func (x *KubernetesRegion) GetRegionName() string {
	if x != nil {
		return x.RegionName
	}
	return ""
}

type KubernetesNodeInstanceType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InstanceTypeID   string `protobuf:"bytes,1,opt,name=instanceTypeID,proto3" json:"instanceTypeID,omitempty"`
	InstanceTypeName string `protobuf:"bytes,11,opt,name=instanceTypeName,proto3" json:"instanceTypeName,omitempty"`
}

func (x *KubernetesNodeInstanceType) Reset() {
	*x = KubernetesNodeInstanceType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_services_catalog_cloud_compute_kubernetes_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KubernetesNodeInstanceType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KubernetesNodeInstanceType) ProtoMessage() {}

func (x *KubernetesNodeInstanceType) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_services_catalog_cloud_compute_kubernetes_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KubernetesNodeInstanceType.ProtoReflect.Descriptor instead.
func (*KubernetesNodeInstanceType) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_services_catalog_cloud_compute_kubernetes_proto_rawDescGZIP(), []int{3}
}

func (x *KubernetesNodeInstanceType) GetInstanceTypeID() string {
	if x != nil {
		return x.InstanceTypeID
	}
	return ""
}

func (x *KubernetesNodeInstanceType) GetInstanceTypeName() string {
	if x != nil {
		return x.InstanceTypeName
	}
	return ""
}

var File_mmesh_protobuf_resources_v1_services_catalog_cloud_compute_kubernetes_proto protoreflect.FileDescriptor

var file_mmesh_protobuf_resources_v1_services_catalog_cloud_compute_kubernetes_proto_rawDesc = []byte{
	0x0a, 0x4b, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2f, 0x6b, 0x75, 0x62,
	0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63,
	0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x22, 0xd3, 0x01, 0x0a, 0x11, 0x4b, 0x75, 0x62, 0x65, 0x72,
	0x6e, 0x65, 0x74, 0x65, 0x73, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x36, 0x0a, 0x08,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2e, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x65, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x33, 0x0a, 0x07, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x15, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2e,
	0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x52, 0x07, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x51, 0x0a, 0x11, 0x6e, 0x6f, 0x64,
	0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x73, 0x18, 0x1f,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2e, 0x4b,
	0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x11, 0x6e, 0x6f, 0x64, 0x65, 0x49,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x73, 0x22, 0x57, 0x0a, 0x11,
	0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12,
	0x24, 0x0a, 0x0d, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x4e, 0x0a, 0x10, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x65, 0x73, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x70, 0x0a, 0x1a, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x65, 0x73, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x69, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x12, 0x2a, 0x0a, 0x10, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x42, 0x5a, 0x40, 0x6d, 0x6d, 0x65, 0x73, 0x68,
	0x2e, 0x64, 0x65, 0x76, 0x2f, 0x6d, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x72,
	0x70, 0x63, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_mmesh_protobuf_resources_v1_services_catalog_cloud_compute_kubernetes_proto_rawDescOnce sync.Once
	file_mmesh_protobuf_resources_v1_services_catalog_cloud_compute_kubernetes_proto_rawDescData = file_mmesh_protobuf_resources_v1_services_catalog_cloud_compute_kubernetes_proto_rawDesc
)

func file_mmesh_protobuf_resources_v1_services_catalog_cloud_compute_kubernetes_proto_rawDescGZIP() []byte {
	file_mmesh_protobuf_resources_v1_services_catalog_cloud_compute_kubernetes_proto_rawDescOnce.Do(func() {
		file_mmesh_protobuf_resources_v1_services_catalog_cloud_compute_kubernetes_proto_rawDescData = protoimpl.X.CompressGZIP(file_mmesh_protobuf_resources_v1_services_catalog_cloud_compute_kubernetes_proto_rawDescData)
	})
	return file_mmesh_protobuf_resources_v1_services_catalog_cloud_compute_kubernetes_proto_rawDescData
}

var file_mmesh_protobuf_resources_v1_services_catalog_cloud_compute_kubernetes_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_mmesh_protobuf_resources_v1_services_catalog_cloud_compute_kubernetes_proto_goTypes = []interface{}{
	(*KubernetesOptions)(nil),          // 0: compute.KubernetesOptions
	(*KubernetesVersion)(nil),          // 1: compute.KubernetesVersion
	(*KubernetesRegion)(nil),           // 2: compute.KubernetesRegion
	(*KubernetesNodeInstanceType)(nil), // 3: compute.KubernetesNodeInstanceType
}
var file_mmesh_protobuf_resources_v1_services_catalog_cloud_compute_kubernetes_proto_depIdxs = []int32{
	1, // 0: compute.KubernetesOptions.versions:type_name -> compute.KubernetesVersion
	2, // 1: compute.KubernetesOptions.regions:type_name -> compute.KubernetesRegion
	3, // 2: compute.KubernetesOptions.nodeInstanceTypes:type_name -> compute.KubernetesNodeInstanceType
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_mmesh_protobuf_resources_v1_services_catalog_cloud_compute_kubernetes_proto_init() }
func file_mmesh_protobuf_resources_v1_services_catalog_cloud_compute_kubernetes_proto_init() {
	if File_mmesh_protobuf_resources_v1_services_catalog_cloud_compute_kubernetes_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mmesh_protobuf_resources_v1_services_catalog_cloud_compute_kubernetes_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KubernetesOptions); i {
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
		file_mmesh_protobuf_resources_v1_services_catalog_cloud_compute_kubernetes_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KubernetesVersion); i {
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
		file_mmesh_protobuf_resources_v1_services_catalog_cloud_compute_kubernetes_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KubernetesRegion); i {
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
		file_mmesh_protobuf_resources_v1_services_catalog_cloud_compute_kubernetes_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KubernetesNodeInstanceType); i {
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
			RawDescriptor: file_mmesh_protobuf_resources_v1_services_catalog_cloud_compute_kubernetes_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mmesh_protobuf_resources_v1_services_catalog_cloud_compute_kubernetes_proto_goTypes,
		DependencyIndexes: file_mmesh_protobuf_resources_v1_services_catalog_cloud_compute_kubernetes_proto_depIdxs,
		MessageInfos:      file_mmesh_protobuf_resources_v1_services_catalog_cloud_compute_kubernetes_proto_msgTypes,
	}.Build()
	File_mmesh_protobuf_resources_v1_services_catalog_cloud_compute_kubernetes_proto = out.File
	file_mmesh_protobuf_resources_v1_services_catalog_cloud_compute_kubernetes_proto_rawDesc = nil
	file_mmesh_protobuf_resources_v1_services_catalog_cloud_compute_kubernetes_proto_goTypes = nil
	file_mmesh_protobuf_resources_v1_services_catalog_cloud_compute_kubernetes_proto_depIdxs = nil
}
