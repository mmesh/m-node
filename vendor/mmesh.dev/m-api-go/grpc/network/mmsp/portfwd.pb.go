// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.2
// source: mmesh/protobuf/network/v1/mmsp/portfwd.proto

package mmsp

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	stream "mmesh.dev/m-api-go/grpc/network/mmsp/stream"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PortFwdMsgType int32

const (
	PortFwdMsgType_UNDEFINED_PORTFWD_MSG PortFwdMsgType = 0
	PortFwdMsgType_PORTFWD_LISTEN        PortFwdMsgType = 11
	PortFwdMsgType_PORTFWD_DIAL          PortFwdMsgType = 21
	PortFwdMsgType_PORTFWD_DIALACK       PortFwdMsgType = 31
	PortFwdMsgType_PORTFWD_END           PortFwdMsgType = 41
	PortFwdMsgType_PORTFWD_DATA          PortFwdMsgType = 51
	PortFwdMsgType_PORTFWD_DISABLED      PortFwdMsgType = 61
)

// Enum value maps for PortFwdMsgType.
var (
	PortFwdMsgType_name = map[int32]string{
		0:  "UNDEFINED_PORTFWD_MSG",
		11: "PORTFWD_LISTEN",
		21: "PORTFWD_DIAL",
		31: "PORTFWD_DIALACK",
		41: "PORTFWD_END",
		51: "PORTFWD_DATA",
		61: "PORTFWD_DISABLED",
	}
	PortFwdMsgType_value = map[string]int32{
		"UNDEFINED_PORTFWD_MSG": 0,
		"PORTFWD_LISTEN":        11,
		"PORTFWD_DIAL":          21,
		"PORTFWD_DIALACK":       31,
		"PORTFWD_END":           41,
		"PORTFWD_DATA":          51,
		"PORTFWD_DISABLED":      61,
	}
)

func (x PortFwdMsgType) Enum() *PortFwdMsgType {
	p := new(PortFwdMsgType)
	*p = x
	return p
}

func (x PortFwdMsgType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PortFwdMsgType) Descriptor() protoreflect.EnumDescriptor {
	return file_mmesh_protobuf_network_v1_mmsp_portfwd_proto_enumTypes[0].Descriptor()
}

func (PortFwdMsgType) Type() protoreflect.EnumType {
	return &file_mmesh_protobuf_network_v1_mmsp_portfwd_proto_enumTypes[0]
}

func (x PortFwdMsgType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PortFwdMsgType.Descriptor instead.
func (PortFwdMsgType) EnumDescriptor() ([]byte, []int) {
	return file_mmesh_protobuf_network_v1_mmsp_portfwd_proto_rawDescGZIP(), []int{0}
}

type PortFwdPDU struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata *StreamMetadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Type     PortFwdMsgType  `protobuf:"varint,11,opt,name=type,proto3,enum=mmsp.PortFwdMsgType" json:"type,omitempty"`
	PortFwd  *stream.PortFwd `protobuf:"bytes,21,opt,name=portFwd,proto3" json:"portFwd,omitempty"`
}

func (x *PortFwdPDU) Reset() {
	*x = PortFwdPDU{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_network_v1_mmsp_portfwd_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PortFwdPDU) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PortFwdPDU) ProtoMessage() {}

func (x *PortFwdPDU) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_network_v1_mmsp_portfwd_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PortFwdPDU.ProtoReflect.Descriptor instead.
func (*PortFwdPDU) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_network_v1_mmsp_portfwd_proto_rawDescGZIP(), []int{0}
}

func (x *PortFwdPDU) GetMetadata() *StreamMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *PortFwdPDU) GetType() PortFwdMsgType {
	if x != nil {
		return x.Type
	}
	return PortFwdMsgType_UNDEFINED_PORTFWD_MSG
}

func (x *PortFwdPDU) GetPortFwd() *stream.PortFwd {
	if x != nil {
		return x.PortFwd
	}
	return nil
}

var File_mmesh_protobuf_network_v1_mmsp_portfwd_proto protoreflect.FileDescriptor

var file_mmesh_protobuf_network_v1_mmsp_portfwd_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6d, 0x73, 0x70,
	0x2f, 0x70, 0x6f, 0x72, 0x74, 0x66, 0x77, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04,
	0x6d, 0x6d, 0x73, 0x70, 0x1a, 0x2d, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x76, 0x31, 0x2f,
	0x6d, 0x6d, 0x73, 0x70, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x33, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x6d,
	0x6d, 0x73, 0x70, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x66,
	0x77, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x93, 0x01, 0x0a, 0x0a, 0x50, 0x6f, 0x72,
	0x74, 0x46, 0x77, 0x64, 0x50, 0x44, 0x55, 0x12, 0x30, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x6d, 0x73, 0x70,
	0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x28, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x6d, 0x6d, 0x73, 0x70, 0x2e, 0x50,
	0x6f, 0x72, 0x74, 0x46, 0x77, 0x64, 0x4d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x29, 0x0a, 0x07, 0x70, 0x6f, 0x72, 0x74, 0x46, 0x77, 0x64, 0x18, 0x15,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x50, 0x6f,
	0x72, 0x74, 0x46, 0x77, 0x64, 0x52, 0x07, 0x70, 0x6f, 0x72, 0x74, 0x46, 0x77, 0x64, 0x2a, 0x9f,
	0x01, 0x0a, 0x0e, 0x50, 0x6f, 0x72, 0x74, 0x46, 0x77, 0x64, 0x4d, 0x73, 0x67, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x19, 0x0a, 0x15, 0x55, 0x4e, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x45, 0x44, 0x5f, 0x50,
	0x4f, 0x52, 0x54, 0x46, 0x57, 0x44, 0x5f, 0x4d, 0x53, 0x47, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e,
	0x50, 0x4f, 0x52, 0x54, 0x46, 0x57, 0x44, 0x5f, 0x4c, 0x49, 0x53, 0x54, 0x45, 0x4e, 0x10, 0x0b,
	0x12, 0x10, 0x0a, 0x0c, 0x50, 0x4f, 0x52, 0x54, 0x46, 0x57, 0x44, 0x5f, 0x44, 0x49, 0x41, 0x4c,
	0x10, 0x15, 0x12, 0x13, 0x0a, 0x0f, 0x50, 0x4f, 0x52, 0x54, 0x46, 0x57, 0x44, 0x5f, 0x44, 0x49,
	0x41, 0x4c, 0x41, 0x43, 0x4b, 0x10, 0x1f, 0x12, 0x0f, 0x0a, 0x0b, 0x50, 0x4f, 0x52, 0x54, 0x46,
	0x57, 0x44, 0x5f, 0x45, 0x4e, 0x44, 0x10, 0x29, 0x12, 0x10, 0x0a, 0x0c, 0x50, 0x4f, 0x52, 0x54,
	0x46, 0x57, 0x44, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x10, 0x33, 0x12, 0x14, 0x0a, 0x10, 0x50, 0x4f,
	0x52, 0x54, 0x46, 0x57, 0x44, 0x5f, 0x44, 0x49, 0x53, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x3d,
	0x42, 0x26, 0x5a, 0x24, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x6d, 0x2d,
	0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x2f, 0x6d, 0x6d, 0x73, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mmesh_protobuf_network_v1_mmsp_portfwd_proto_rawDescOnce sync.Once
	file_mmesh_protobuf_network_v1_mmsp_portfwd_proto_rawDescData = file_mmesh_protobuf_network_v1_mmsp_portfwd_proto_rawDesc
)

func file_mmesh_protobuf_network_v1_mmsp_portfwd_proto_rawDescGZIP() []byte {
	file_mmesh_protobuf_network_v1_mmsp_portfwd_proto_rawDescOnce.Do(func() {
		file_mmesh_protobuf_network_v1_mmsp_portfwd_proto_rawDescData = protoimpl.X.CompressGZIP(file_mmesh_protobuf_network_v1_mmsp_portfwd_proto_rawDescData)
	})
	return file_mmesh_protobuf_network_v1_mmsp_portfwd_proto_rawDescData
}

var file_mmesh_protobuf_network_v1_mmsp_portfwd_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_mmesh_protobuf_network_v1_mmsp_portfwd_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_mmesh_protobuf_network_v1_mmsp_portfwd_proto_goTypes = []interface{}{
	(PortFwdMsgType)(0),    // 0: mmsp.PortFwdMsgType
	(*PortFwdPDU)(nil),     // 1: mmsp.PortFwdPDU
	(*StreamMetadata)(nil), // 2: mmsp.StreamMetadata
	(*stream.PortFwd)(nil), // 3: stream.PortFwd
}
var file_mmesh_protobuf_network_v1_mmsp_portfwd_proto_depIdxs = []int32{
	2, // 0: mmsp.PortFwdPDU.metadata:type_name -> mmsp.StreamMetadata
	0, // 1: mmsp.PortFwdPDU.type:type_name -> mmsp.PortFwdMsgType
	3, // 2: mmsp.PortFwdPDU.portFwd:type_name -> stream.PortFwd
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_mmesh_protobuf_network_v1_mmsp_portfwd_proto_init() }
func file_mmesh_protobuf_network_v1_mmsp_portfwd_proto_init() {
	if File_mmesh_protobuf_network_v1_mmsp_portfwd_proto != nil {
		return
	}
	file_mmesh_protobuf_network_v1_mmsp_metadata_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_mmesh_protobuf_network_v1_mmsp_portfwd_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PortFwdPDU); i {
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
			RawDescriptor: file_mmesh_protobuf_network_v1_mmsp_portfwd_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mmesh_protobuf_network_v1_mmsp_portfwd_proto_goTypes,
		DependencyIndexes: file_mmesh_protobuf_network_v1_mmsp_portfwd_proto_depIdxs,
		EnumInfos:         file_mmesh_protobuf_network_v1_mmsp_portfwd_proto_enumTypes,
		MessageInfos:      file_mmesh_protobuf_network_v1_mmsp_portfwd_proto_msgTypes,
	}.Build()
	File_mmesh_protobuf_network_v1_mmsp_portfwd_proto = out.File
	file_mmesh_protobuf_network_v1_mmsp_portfwd_proto_rawDesc = nil
	file_mmesh_protobuf_network_v1_mmsp_portfwd_proto_goTypes = nil
	file_mmesh_protobuf_network_v1_mmsp_portfwd_proto_depIdxs = nil
}
