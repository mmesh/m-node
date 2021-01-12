// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: mmesh/protobuf/network/v1/mmsp/portFwd.proto

package portFwd

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

type PortFwd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	PortFwdType int32  `protobuf:"varint,2,opt,name=portFwdType,proto3" json:"portFwdType,omitempty"`
	Link        *Link  `protobuf:"bytes,3,opt,name=link,proto3" json:"link,omitempty"`
	Data        []byte `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *PortFwd) Reset() {
	*x = PortFwd{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_network_v1_mmsp_portFwd_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PortFwd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PortFwd) ProtoMessage() {}

func (x *PortFwd) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_network_v1_mmsp_portFwd_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PortFwd.ProtoReflect.Descriptor instead.
func (*PortFwd) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_network_v1_mmsp_portFwd_proto_rawDescGZIP(), []int{0}
}

func (x *PortFwd) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *PortFwd) GetPortFwdType() int32 {
	if x != nil {
		return x.PortFwdType
	}
	return 0
}

func (x *PortFwd) GetLink() *Link {
	if x != nil {
		return x.Link
	}
	return nil
}

func (x *PortFwd) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type Link struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID           string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Proto        uint32 `protobuf:"varint,2,opt,name=proto,proto3" json:"proto,omitempty"`
	SrcNodeID    string `protobuf:"bytes,3,opt,name=srcNodeID,proto3" json:"srcNodeID,omitempty"`
	SrcPort      uint32 `protobuf:"varint,4,opt,name=srcPort,proto3" json:"srcPort,omitempty"`
	DstNodeID    string `protobuf:"bytes,5,opt,name=dstNodeID,proto3" json:"dstNodeID,omitempty"`
	DstPort      uint32 `protobuf:"varint,6,opt,name=dstPort,proto3" json:"dstPort,omitempty"`
	ConnectionID string `protobuf:"bytes,7,opt,name=connectionID,proto3" json:"connectionID,omitempty"`
}

func (x *Link) Reset() {
	*x = Link{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_network_v1_mmsp_portFwd_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Link) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Link) ProtoMessage() {}

func (x *Link) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_network_v1_mmsp_portFwd_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Link.ProtoReflect.Descriptor instead.
func (*Link) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_network_v1_mmsp_portFwd_proto_rawDescGZIP(), []int{1}
}

func (x *Link) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Link) GetProto() uint32 {
	if x != nil {
		return x.Proto
	}
	return 0
}

func (x *Link) GetSrcNodeID() string {
	if x != nil {
		return x.SrcNodeID
	}
	return ""
}

func (x *Link) GetSrcPort() uint32 {
	if x != nil {
		return x.SrcPort
	}
	return 0
}

func (x *Link) GetDstNodeID() string {
	if x != nil {
		return x.DstNodeID
	}
	return ""
}

func (x *Link) GetDstPort() uint32 {
	if x != nil {
		return x.DstPort
	}
	return 0
}

func (x *Link) GetConnectionID() string {
	if x != nil {
		return x.ConnectionID
	}
	return ""
}

var File_mmesh_protobuf_network_v1_mmsp_portFwd_proto protoreflect.FileDescriptor

var file_mmesh_protobuf_network_v1_mmsp_portFwd_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6d, 0x73, 0x70,
	0x2f, 0x70, 0x6f, 0x72, 0x74, 0x46, 0x77, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07,
	0x70, 0x6f, 0x72, 0x74, 0x46, 0x77, 0x64, 0x22, 0x72, 0x0a, 0x07, 0x50, 0x6f, 0x72, 0x74, 0x46,
	0x77, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x6f, 0x72, 0x74, 0x46, 0x77, 0x64, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x70, 0x6f, 0x72, 0x74, 0x46, 0x77, 0x64,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x46, 0x77, 0x64, 0x2e, 0x4c, 0x69, 0x6e,
	0x6b, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xc0, 0x01, 0x0a, 0x04,
	0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x72,
	0x63, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73,
	0x72, 0x63, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x72, 0x63, 0x50,
	0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x72, 0x63, 0x50, 0x6f,
	0x72, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x73, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x44, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x73, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x44,
	0x12, 0x18, 0x0a, 0x07, 0x64, 0x73, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x07, 0x64, 0x73, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x42, 0x2e,
	0x5a, 0x2c, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x6d, 0x2d, 0x61, 0x70,
	0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x2f, 0x6d, 0x6d, 0x73, 0x70, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x46, 0x77, 0x64, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mmesh_protobuf_network_v1_mmsp_portFwd_proto_rawDescOnce sync.Once
	file_mmesh_protobuf_network_v1_mmsp_portFwd_proto_rawDescData = file_mmesh_protobuf_network_v1_mmsp_portFwd_proto_rawDesc
)

func file_mmesh_protobuf_network_v1_mmsp_portFwd_proto_rawDescGZIP() []byte {
	file_mmesh_protobuf_network_v1_mmsp_portFwd_proto_rawDescOnce.Do(func() {
		file_mmesh_protobuf_network_v1_mmsp_portFwd_proto_rawDescData = protoimpl.X.CompressGZIP(file_mmesh_protobuf_network_v1_mmsp_portFwd_proto_rawDescData)
	})
	return file_mmesh_protobuf_network_v1_mmsp_portFwd_proto_rawDescData
}

var file_mmesh_protobuf_network_v1_mmsp_portFwd_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_mmesh_protobuf_network_v1_mmsp_portFwd_proto_goTypes = []interface{}{
	(*PortFwd)(nil), // 0: portFwd.PortFwd
	(*Link)(nil),    // 1: portFwd.Link
}
var file_mmesh_protobuf_network_v1_mmsp_portFwd_proto_depIdxs = []int32{
	1, // 0: portFwd.PortFwd.link:type_name -> portFwd.Link
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_mmesh_protobuf_network_v1_mmsp_portFwd_proto_init() }
func file_mmesh_protobuf_network_v1_mmsp_portFwd_proto_init() {
	if File_mmesh_protobuf_network_v1_mmsp_portFwd_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mmesh_protobuf_network_v1_mmsp_portFwd_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PortFwd); i {
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
		file_mmesh_protobuf_network_v1_mmsp_portFwd_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Link); i {
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
			RawDescriptor: file_mmesh_protobuf_network_v1_mmsp_portFwd_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mmesh_protobuf_network_v1_mmsp_portFwd_proto_goTypes,
		DependencyIndexes: file_mmesh_protobuf_network_v1_mmsp_portFwd_proto_depIdxs,
		MessageInfos:      file_mmesh_protobuf_network_v1_mmsp_portFwd_proto_msgTypes,
	}.Build()
	File_mmesh_protobuf_network_v1_mmsp_portFwd_proto = out.File
	file_mmesh_protobuf_network_v1_mmsp_portFwd_proto_rawDesc = nil
	file_mmesh_protobuf_network_v1_mmsp_portFwd_proto_goTypes = nil
	file_mmesh_protobuf_network_v1_mmsp_portFwd_proto_depIdxs = nil
}
