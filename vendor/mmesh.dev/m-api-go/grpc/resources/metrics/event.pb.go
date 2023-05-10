// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.0
// source: mmesh/protobuf/resources/v1/metrics/event.proto

package metrics

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	events "mmesh.dev/m-api-go/grpc/resources/events"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type EventMetrics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SuccessCount       float32            `protobuf:"fixed32,1,opt,name=successCount,proto3" json:"successCount,omitempty"`
	FailCount          float32            `protobuf:"fixed32,2,opt,name=failCount,proto3" json:"failCount,omitempty"`
	SuccessProbability float32            `protobuf:"fixed32,3,opt,name=successProbability,proto3" json:"successProbability,omitempty"`
	FailProbability    float32            `protobuf:"fixed32,4,opt,name=failProbability,proto3" json:"failProbability,omitempty"`
	FirstActivity      int64              `protobuf:"varint,5,opt,name=firstActivity,proto3" json:"firstActivity,omitempty"`
	LastActivity       int64              `protobuf:"varint,6,opt,name=lastActivity,proto3" json:"lastActivity,omitempty"`
	ActivityIndex      float32            `protobuf:"fixed32,7,opt,name=activityIndex,proto3" json:"activityIndex,omitempty"`
	Score              float32            `protobuf:"fixed32,8,opt,name=score,proto3" json:"score,omitempty"`
	Rating             string             `protobuf:"bytes,9,opt,name=rating,proto3" json:"rating,omitempty"`
	LastResult         events.EventResult `protobuf:"varint,10,opt,name=lastResult,proto3,enum=events.EventResult" json:"lastResult,omitempty"`
}

func (x *EventMetrics) Reset() {
	*x = EventMetrics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_metrics_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventMetrics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventMetrics) ProtoMessage() {}

func (x *EventMetrics) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_metrics_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventMetrics.ProtoReflect.Descriptor instead.
func (*EventMetrics) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_metrics_event_proto_rawDescGZIP(), []int{0}
}

func (x *EventMetrics) GetSuccessCount() float32 {
	if x != nil {
		return x.SuccessCount
	}
	return 0
}

func (x *EventMetrics) GetFailCount() float32 {
	if x != nil {
		return x.FailCount
	}
	return 0
}

func (x *EventMetrics) GetSuccessProbability() float32 {
	if x != nil {
		return x.SuccessProbability
	}
	return 0
}

func (x *EventMetrics) GetFailProbability() float32 {
	if x != nil {
		return x.FailProbability
	}
	return 0
}

func (x *EventMetrics) GetFirstActivity() int64 {
	if x != nil {
		return x.FirstActivity
	}
	return 0
}

func (x *EventMetrics) GetLastActivity() int64 {
	if x != nil {
		return x.LastActivity
	}
	return 0
}

func (x *EventMetrics) GetActivityIndex() float32 {
	if x != nil {
		return x.ActivityIndex
	}
	return 0
}

func (x *EventMetrics) GetScore() float32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *EventMetrics) GetRating() string {
	if x != nil {
		return x.Rating
	}
	return ""
}

func (x *EventMetrics) GetLastResult() events.EventResult {
	if x != nil {
		return x.LastResult
	}
	return events.EventResult(0)
}

var File_mmesh_protobuf_resources_v1_metrics_event_proto protoreflect.FileDescriptor

var file_mmesh_protobuf_resources_v1_metrics_event_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x73, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x1a, 0x2e, 0x6d, 0x6d, 0x65, 0x73,
	0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfd, 0x02, 0x0a, 0x0c, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x0c, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x1c, 0x0a, 0x09, 0x66, 0x61, 0x69, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x09, 0x66, 0x61, 0x69, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2e, 0x0a,
	0x12, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x50, 0x72, 0x6f, 0x62, 0x61, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x12, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x50, 0x72, 0x6f, 0x62, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x28, 0x0a,
	0x0f, 0x66, 0x61, 0x69, 0x6c, 0x50, 0x72, 0x6f, 0x62, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0f, 0x66, 0x61, 0x69, 0x6c, 0x50, 0x72, 0x6f, 0x62,
	0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x24, 0x0a, 0x0d, 0x66, 0x69, 0x72, 0x73, 0x74,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x12, 0x22, 0x0a,
	0x0c, 0x6c, 0x61, 0x73, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x12, 0x24, 0x0a, 0x0d, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0d, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72,
	0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x33, 0x0a, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x0a,
	0x6c, 0x61, 0x73, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x42, 0x2b, 0x5a, 0x29, 0x6d, 0x6d,
	0x65, 0x73, 0x68, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x6d, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mmesh_protobuf_resources_v1_metrics_event_proto_rawDescOnce sync.Once
	file_mmesh_protobuf_resources_v1_metrics_event_proto_rawDescData = file_mmesh_protobuf_resources_v1_metrics_event_proto_rawDesc
)

func file_mmesh_protobuf_resources_v1_metrics_event_proto_rawDescGZIP() []byte {
	file_mmesh_protobuf_resources_v1_metrics_event_proto_rawDescOnce.Do(func() {
		file_mmesh_protobuf_resources_v1_metrics_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_mmesh_protobuf_resources_v1_metrics_event_proto_rawDescData)
	})
	return file_mmesh_protobuf_resources_v1_metrics_event_proto_rawDescData
}

var file_mmesh_protobuf_resources_v1_metrics_event_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_mmesh_protobuf_resources_v1_metrics_event_proto_goTypes = []interface{}{
	(*EventMetrics)(nil),    // 0: metrics.EventMetrics
	(events.EventResult)(0), // 1: events.EventResult
}
var file_mmesh_protobuf_resources_v1_metrics_event_proto_depIdxs = []int32{
	1, // 0: metrics.EventMetrics.lastResult:type_name -> events.EventResult
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_mmesh_protobuf_resources_v1_metrics_event_proto_init() }
func file_mmesh_protobuf_resources_v1_metrics_event_proto_init() {
	if File_mmesh_protobuf_resources_v1_metrics_event_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mmesh_protobuf_resources_v1_metrics_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventMetrics); i {
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
			RawDescriptor: file_mmesh_protobuf_resources_v1_metrics_event_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mmesh_protobuf_resources_v1_metrics_event_proto_goTypes,
		DependencyIndexes: file_mmesh_protobuf_resources_v1_metrics_event_proto_depIdxs,
		MessageInfos:      file_mmesh_protobuf_resources_v1_metrics_event_proto_msgTypes,
	}.Build()
	File_mmesh_protobuf_resources_v1_metrics_event_proto = out.File
	file_mmesh_protobuf_resources_v1_metrics_event_proto_rawDesc = nil
	file_mmesh_protobuf_resources_v1_metrics_event_proto_goTypes = nil
	file_mmesh_protobuf_resources_v1_metrics_event_proto_depIdxs = nil
}
