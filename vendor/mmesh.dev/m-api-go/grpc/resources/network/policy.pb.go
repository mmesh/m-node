// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: mmesh/protobuf/resources/v1/network/policy.proto

package network

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

type Policy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DefaultPolicy string `protobuf:"bytes,1,opt,name=defaultPolicy,proto3" json:"defaultPolicy,omitempty"` // ACCEPT, DROP
	//map<string, Filter> networkFilters = 2;  // map[nfID]*NetworkFilter
	NetworkFilters []*Filter `protobuf:"bytes,2,rep,name=networkFilters,proto3" json:"networkFilters,omitempty"`
}

func (x *Policy) Reset() {
	*x = Policy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_network_policy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Policy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Policy) ProtoMessage() {}

func (x *Policy) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_network_policy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Policy.ProtoReflect.Descriptor instead.
func (*Policy) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_network_policy_proto_rawDescGZIP(), []int{0}
}

func (x *Policy) GetDefaultPolicy() string {
	if x != nil {
		return x.DefaultPolicy
	}
	return ""
}

func (x *Policy) GetNetworkFilters() []*Filter {
	if x != nil {
		return x.NetworkFilters
	}
	return nil
}

type Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index        uint32 `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Description  string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	SrcIPNet     string `protobuf:"bytes,3,opt,name=srcIPNet,proto3" json:"srcIPNet,omitempty"`
	DstIPNet     string `protobuf:"bytes,4,opt,name=dstIPNet,proto3" json:"dstIPNet,omitempty"` // usually, the endpoint addr
	Proto        string `protobuf:"bytes,5,opt,name=proto,proto3" json:"proto,omitempty"`       // TCP, UDP, ICMP
	DstPort      int32  `protobuf:"varint,6,opt,name=dstPort,proto3" json:"dstPort,omitempty"`
	Policy       string `protobuf:"bytes,7,opt,name=policy,proto3" json:"policy,omitempty"` // ACCEPT, DROP
	LastModified int64  `protobuf:"varint,8,opt,name=lastModified,proto3" json:"lastModified,omitempty"`
}

func (x *Filter) Reset() {
	*x = Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_network_policy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Filter) ProtoMessage() {}

func (x *Filter) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_network_policy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Filter.ProtoReflect.Descriptor instead.
func (*Filter) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_network_policy_proto_rawDescGZIP(), []int{1}
}

func (x *Filter) GetIndex() uint32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Filter) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Filter) GetSrcIPNet() string {
	if x != nil {
		return x.SrcIPNet
	}
	return ""
}

func (x *Filter) GetDstIPNet() string {
	if x != nil {
		return x.DstIPNet
	}
	return ""
}

func (x *Filter) GetProto() string {
	if x != nil {
		return x.Proto
	}
	return ""
}

func (x *Filter) GetDstPort() int32 {
	if x != nil {
		return x.DstPort
	}
	return 0
}

func (x *Filter) GetPolicy() string {
	if x != nil {
		return x.Policy
	}
	return ""
}

func (x *Filter) GetLastModified() int64 {
	if x != nil {
		return x.LastModified
	}
	return 0
}

var File_mmesh_protobuf_resources_v1_network_policy_proto protoreflect.FileDescriptor

var file_mmesh_protobuf_resources_v1_network_policy_proto_rawDesc = []byte{
	0x0a, 0x30, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x22, 0x67, 0x0a, 0x06, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x24, 0x0a, 0x0d, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x64, 0x65,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x37, 0x0a, 0x0e, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x52, 0x0e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x73, 0x22, 0xe4, 0x01, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12,
	0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x72, 0x63, 0x49, 0x50,
	0x4e, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x72, 0x63, 0x49, 0x50,
	0x4e, 0x65, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x73, 0x74, 0x49, 0x50, 0x4e, 0x65, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x73, 0x74, 0x49, 0x50, 0x4e, 0x65, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x73, 0x74, 0x50, 0x6f, 0x72, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x64, 0x73, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x22, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x4d,
	0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x6c,
	0x61, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x42, 0x2b, 0x5a, 0x29, 0x6d,
	0x6d, 0x65, 0x73, 0x68, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x6d, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x67,
	0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mmesh_protobuf_resources_v1_network_policy_proto_rawDescOnce sync.Once
	file_mmesh_protobuf_resources_v1_network_policy_proto_rawDescData = file_mmesh_protobuf_resources_v1_network_policy_proto_rawDesc
)

func file_mmesh_protobuf_resources_v1_network_policy_proto_rawDescGZIP() []byte {
	file_mmesh_protobuf_resources_v1_network_policy_proto_rawDescOnce.Do(func() {
		file_mmesh_protobuf_resources_v1_network_policy_proto_rawDescData = protoimpl.X.CompressGZIP(file_mmesh_protobuf_resources_v1_network_policy_proto_rawDescData)
	})
	return file_mmesh_protobuf_resources_v1_network_policy_proto_rawDescData
}

var file_mmesh_protobuf_resources_v1_network_policy_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_mmesh_protobuf_resources_v1_network_policy_proto_goTypes = []interface{}{
	(*Policy)(nil), // 0: network.Policy
	(*Filter)(nil), // 1: network.Filter
}
var file_mmesh_protobuf_resources_v1_network_policy_proto_depIdxs = []int32{
	1, // 0: network.Policy.networkFilters:type_name -> network.Filter
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_mmesh_protobuf_resources_v1_network_policy_proto_init() }
func file_mmesh_protobuf_resources_v1_network_policy_proto_init() {
	if File_mmesh_protobuf_resources_v1_network_policy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mmesh_protobuf_resources_v1_network_policy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Policy); i {
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
		file_mmesh_protobuf_resources_v1_network_policy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Filter); i {
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
			RawDescriptor: file_mmesh_protobuf_resources_v1_network_policy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mmesh_protobuf_resources_v1_network_policy_proto_goTypes,
		DependencyIndexes: file_mmesh_protobuf_resources_v1_network_policy_proto_depIdxs,
		MessageInfos:      file_mmesh_protobuf_resources_v1_network_policy_proto_msgTypes,
	}.Build()
	File_mmesh_protobuf_resources_v1_network_policy_proto = out.File
	file_mmesh_protobuf_resources_v1_network_policy_proto_rawDesc = nil
	file_mmesh_protobuf_resources_v1_network_policy_proto_goTypes = nil
	file_mmesh_protobuf_resources_v1_network_policy_proto_depIdxs = nil
}
