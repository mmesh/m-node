// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.2
// source: mmesh/protobuf/resources/v1/resource/mongo.proto

package resource

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

type MongoQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountID            string         `protobuf:"bytes,1,opt,name=accountID,proto3" json:"accountID,omitempty"`
	DataSourceInstanceID string         `protobuf:"bytes,11,opt,name=dataSourceInstanceID,proto3" json:"dataSourceInstanceID,omitempty"`
	Table                string         `protobuf:"bytes,21,opt,name=table,proto3" json:"table,omitempty"`
	Fields               []string       `protobuf:"bytes,31,rep,name=fields,proto3" json:"fields,omitempty"`
	Filters              []*MongoFilter `protobuf:"bytes,41,rep,name=filters,proto3" json:"filters,omitempty"`
}

func (x *MongoQuery) Reset() {
	*x = MongoQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_resource_mongo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MongoQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MongoQuery) ProtoMessage() {}

func (x *MongoQuery) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_resource_mongo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MongoQuery.ProtoReflect.Descriptor instead.
func (*MongoQuery) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_resource_mongo_proto_rawDescGZIP(), []int{0}
}

func (x *MongoQuery) GetAccountID() string {
	if x != nil {
		return x.AccountID
	}
	return ""
}

func (x *MongoQuery) GetDataSourceInstanceID() string {
	if x != nil {
		return x.DataSourceInstanceID
	}
	return ""
}

func (x *MongoQuery) GetTable() string {
	if x != nil {
		return x.Table
	}
	return ""
}

func (x *MongoQuery) GetFields() []string {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *MongoQuery) GetFilters() []*MongoFilter {
	if x != nil {
		return x.Filters
	}
	return nil
}

type MongoQueryResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *MongoQueryResult) Reset() {
	*x = MongoQueryResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_resource_mongo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MongoQueryResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MongoQueryResult) ProtoMessage() {}

func (x *MongoQueryResult) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_resource_mongo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MongoQueryResult.ProtoReflect.Descriptor instead.
func (*MongoQueryResult) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_resource_mongo_proto_rawDescGZIP(), []int{1}
}

func (x *MongoQueryResult) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type MongoFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,11,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,21,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *MongoFilter) Reset() {
	*x = MongoFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_resource_mongo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MongoFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MongoFilter) ProtoMessage() {}

func (x *MongoFilter) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_resource_mongo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MongoFilter.ProtoReflect.Descriptor instead.
func (*MongoFilter) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_resource_mongo_proto_rawDescGZIP(), []int{2}
}

func (x *MongoFilter) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *MongoFilter) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_mmesh_protobuf_resources_v1_resource_mongo_proto protoreflect.FileDescriptor

var file_mmesh_protobuf_resources_v1_resource_mongo_proto_rawDesc = []byte{
	0x0a, 0x30, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0xbd, 0x01, 0x0a,
	0x0a, 0x4d, 0x6f, 0x6e, 0x67, 0x6f, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x32, 0x0a, 0x14, 0x64, 0x61, 0x74,
	0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49,
	0x44, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x64, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x44, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x1f, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x2f, 0x0a, 0x07, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x29, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x4d, 0x6f, 0x6e, 0x67, 0x6f, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x22, 0x26, 0x0a, 0x10,
	0x4d, 0x6f, 0x6e, 0x67, 0x6f, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0x35, 0x0a, 0x0b, 0x4d, 0x6f, 0x6e, 0x67, 0x6f, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x15,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x2c, 0x5a, 0x2a, 0x6d,
	0x6d, 0x65, 0x73, 0x68, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x6d, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x67,
	0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_mmesh_protobuf_resources_v1_resource_mongo_proto_rawDescOnce sync.Once
	file_mmesh_protobuf_resources_v1_resource_mongo_proto_rawDescData = file_mmesh_protobuf_resources_v1_resource_mongo_proto_rawDesc
)

func file_mmesh_protobuf_resources_v1_resource_mongo_proto_rawDescGZIP() []byte {
	file_mmesh_protobuf_resources_v1_resource_mongo_proto_rawDescOnce.Do(func() {
		file_mmesh_protobuf_resources_v1_resource_mongo_proto_rawDescData = protoimpl.X.CompressGZIP(file_mmesh_protobuf_resources_v1_resource_mongo_proto_rawDescData)
	})
	return file_mmesh_protobuf_resources_v1_resource_mongo_proto_rawDescData
}

var file_mmesh_protobuf_resources_v1_resource_mongo_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_mmesh_protobuf_resources_v1_resource_mongo_proto_goTypes = []interface{}{
	(*MongoQuery)(nil),       // 0: resource.MongoQuery
	(*MongoQueryResult)(nil), // 1: resource.MongoQueryResult
	(*MongoFilter)(nil),      // 2: resource.MongoFilter
}
var file_mmesh_protobuf_resources_v1_resource_mongo_proto_depIdxs = []int32{
	2, // 0: resource.MongoQuery.filters:type_name -> resource.MongoFilter
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_mmesh_protobuf_resources_v1_resource_mongo_proto_init() }
func file_mmesh_protobuf_resources_v1_resource_mongo_proto_init() {
	if File_mmesh_protobuf_resources_v1_resource_mongo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mmesh_protobuf_resources_v1_resource_mongo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MongoQuery); i {
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
		file_mmesh_protobuf_resources_v1_resource_mongo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MongoQueryResult); i {
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
		file_mmesh_protobuf_resources_v1_resource_mongo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MongoFilter); i {
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
			RawDescriptor: file_mmesh_protobuf_resources_v1_resource_mongo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mmesh_protobuf_resources_v1_resource_mongo_proto_goTypes,
		DependencyIndexes: file_mmesh_protobuf_resources_v1_resource_mongo_proto_depIdxs,
		MessageInfos:      file_mmesh_protobuf_resources_v1_resource_mongo_proto_msgTypes,
	}.Build()
	File_mmesh_protobuf_resources_v1_resource_mongo_proto = out.File
	file_mmesh_protobuf_resources_v1_resource_mongo_proto_rawDesc = nil
	file_mmesh_protobuf_resources_v1_resource_mongo_proto_goTypes = nil
	file_mmesh_protobuf_resources_v1_resource_mongo_proto_depIdxs = nil
}
