// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.0
// source: mmesh/protobuf/resources/v1/services/thirdParty/cloud.proto

package thirdParty

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

type DigitalOcean struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enabled bool   `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Token   string `protobuf:"bytes,11,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *DigitalOcean) Reset() {
	*x = DigitalOcean{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_services_thirdParty_cloud_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DigitalOcean) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DigitalOcean) ProtoMessage() {}

func (x *DigitalOcean) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_services_thirdParty_cloud_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DigitalOcean.ProtoReflect.Descriptor instead.
func (*DigitalOcean) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_services_thirdParty_cloud_proto_rawDescGZIP(), []int{0}
}

func (x *DigitalOcean) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *DigitalOcean) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type GCP struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enabled   bool   `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	ProjectID string `protobuf:"bytes,11,opt,name=projectID,proto3" json:"projectID,omitempty"`
	JSONKey   []byte `protobuf:"bytes,21,opt,name=JSONKey,proto3" json:"JSONKey,omitempty"`
}

func (x *GCP) Reset() {
	*x = GCP{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_services_thirdParty_cloud_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GCP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GCP) ProtoMessage() {}

func (x *GCP) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_services_thirdParty_cloud_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GCP.ProtoReflect.Descriptor instead.
func (*GCP) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_services_thirdParty_cloud_proto_rawDescGZIP(), []int{1}
}

func (x *GCP) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *GCP) GetProjectID() string {
	if x != nil {
		return x.ProjectID
	}
	return ""
}

func (x *GCP) GetJSONKey() []byte {
	if x != nil {
		return x.JSONKey
	}
	return nil
}

type Scaleway struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enabled        bool   `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	OrganizationID string `protobuf:"bytes,11,opt,name=organizationID,proto3" json:"organizationID,omitempty"`
	ProjectID      string `protobuf:"bytes,12,opt,name=projectID,proto3" json:"projectID,omitempty"`
	AccessKey      string `protobuf:"bytes,21,opt,name=accessKey,proto3" json:"accessKey,omitempty"`
	SecretKey      string `protobuf:"bytes,22,opt,name=secretKey,proto3" json:"secretKey,omitempty"`
}

func (x *Scaleway) Reset() {
	*x = Scaleway{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_services_thirdParty_cloud_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Scaleway) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Scaleway) ProtoMessage() {}

func (x *Scaleway) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_services_thirdParty_cloud_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Scaleway.ProtoReflect.Descriptor instead.
func (*Scaleway) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_services_thirdParty_cloud_proto_rawDescGZIP(), []int{2}
}

func (x *Scaleway) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *Scaleway) GetOrganizationID() string {
	if x != nil {
		return x.OrganizationID
	}
	return ""
}

func (x *Scaleway) GetProjectID() string {
	if x != nil {
		return x.ProjectID
	}
	return ""
}

func (x *Scaleway) GetAccessKey() string {
	if x != nil {
		return x.AccessKey
	}
	return ""
}

func (x *Scaleway) GetSecretKey() string {
	if x != nil {
		return x.SecretKey
	}
	return ""
}

var File_mmesh_protobuf_resources_v1_services_thirdParty_cloud_proto protoreflect.FileDescriptor

var file_mmesh_protobuf_resources_v1_services_thirdParty_cloud_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x74, 0x68, 0x69, 0x72, 0x64, 0x50, 0x61, 0x72, 0x74,
	0x79, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x74,
	0x68, 0x69, 0x72, 0x64, 0x50, 0x61, 0x72, 0x74, 0x79, 0x22, 0x3e, 0x0a, 0x0c, 0x44, 0x69, 0x67,
	0x69, 0x74, 0x61, 0x6c, 0x4f, 0x63, 0x65, 0x61, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x57, 0x0a, 0x03, 0x47, 0x43, 0x50,
	0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x44, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x4a, 0x53, 0x4f, 0x4e,
	0x4b, 0x65, 0x79, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x4a, 0x53, 0x4f, 0x4e, 0x4b,
	0x65, 0x79, 0x22, 0xa6, 0x01, 0x0a, 0x08, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x77, 0x61, 0x79, 0x12,
	0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x6f, 0x72, 0x67,
	0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x44, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x44, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x44, 0x12,
	0x1c, 0x0a, 0x09, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x18, 0x15, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x12, 0x1c, 0x0a,
	0x09, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x42, 0x37, 0x5a, 0x35, 0x6d,
	0x6d, 0x65, 0x73, 0x68, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x6d, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x67,
	0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x74, 0x68, 0x69, 0x72, 0x64, 0x50,
	0x61, 0x72, 0x74, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mmesh_protobuf_resources_v1_services_thirdParty_cloud_proto_rawDescOnce sync.Once
	file_mmesh_protobuf_resources_v1_services_thirdParty_cloud_proto_rawDescData = file_mmesh_protobuf_resources_v1_services_thirdParty_cloud_proto_rawDesc
)

func file_mmesh_protobuf_resources_v1_services_thirdParty_cloud_proto_rawDescGZIP() []byte {
	file_mmesh_protobuf_resources_v1_services_thirdParty_cloud_proto_rawDescOnce.Do(func() {
		file_mmesh_protobuf_resources_v1_services_thirdParty_cloud_proto_rawDescData = protoimpl.X.CompressGZIP(file_mmesh_protobuf_resources_v1_services_thirdParty_cloud_proto_rawDescData)
	})
	return file_mmesh_protobuf_resources_v1_services_thirdParty_cloud_proto_rawDescData
}

var file_mmesh_protobuf_resources_v1_services_thirdParty_cloud_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_mmesh_protobuf_resources_v1_services_thirdParty_cloud_proto_goTypes = []interface{}{
	(*DigitalOcean)(nil), // 0: thirdParty.DigitalOcean
	(*GCP)(nil),          // 1: thirdParty.GCP
	(*Scaleway)(nil),     // 2: thirdParty.Scaleway
}
var file_mmesh_protobuf_resources_v1_services_thirdParty_cloud_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_mmesh_protobuf_resources_v1_services_thirdParty_cloud_proto_init() }
func file_mmesh_protobuf_resources_v1_services_thirdParty_cloud_proto_init() {
	if File_mmesh_protobuf_resources_v1_services_thirdParty_cloud_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mmesh_protobuf_resources_v1_services_thirdParty_cloud_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DigitalOcean); i {
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
		file_mmesh_protobuf_resources_v1_services_thirdParty_cloud_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GCP); i {
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
		file_mmesh_protobuf_resources_v1_services_thirdParty_cloud_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Scaleway); i {
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
			RawDescriptor: file_mmesh_protobuf_resources_v1_services_thirdParty_cloud_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mmesh_protobuf_resources_v1_services_thirdParty_cloud_proto_goTypes,
		DependencyIndexes: file_mmesh_protobuf_resources_v1_services_thirdParty_cloud_proto_depIdxs,
		MessageInfos:      file_mmesh_protobuf_resources_v1_services_thirdParty_cloud_proto_msgTypes,
	}.Build()
	File_mmesh_protobuf_resources_v1_services_thirdParty_cloud_proto = out.File
	file_mmesh_protobuf_resources_v1_services_thirdParty_cloud_proto_rawDesc = nil
	file_mmesh_protobuf_resources_v1_services_thirdParty_cloud_proto_goTypes = nil
	file_mmesh_protobuf_resources_v1_services_thirdParty_cloud_proto_depIdxs = nil
}
