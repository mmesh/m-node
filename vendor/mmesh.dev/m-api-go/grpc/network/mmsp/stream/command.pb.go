// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.2
// source: mmesh/protobuf/network/v1/mmsp/stream/command.proto

package stream

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	datetime "mmesh.dev/m-api-go/grpc/common/datetime"
	status "mmesh.dev/m-api-go/grpc/common/status"
	ops "mmesh.dev/m-api-go/grpc/resources/ops"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Command struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// string ID = 1;
	// CommandType type = 11;
	Request  *CommandRequest  `protobuf:"bytes,21,opt,name=request,proto3" json:"request,omitempty"`
	Response *CommandResponse `protobuf:"bytes,31,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *Command) Reset() {
	*x = Command{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_network_v1_mmsp_stream_command_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Command) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Command) ProtoMessage() {}

func (x *Command) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_network_v1_mmsp_stream_command_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Command.ProtoReflect.Descriptor instead.
func (*Command) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_network_v1_mmsp_stream_command_proto_rawDescGZIP(), []int{0}
}

func (x *Command) GetRequest() *CommandRequest {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *Command) GetResponse() *CommandResponse {
	if x != nil {
		return x.Response
	}
	return nil
}

type CommandRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exec     *ops.CommandExec   `protobuf:"bytes,1,opt,name=exec,proto3" json:"exec,omitempty"`
	Schedule *datetime.DateTime `protobuf:"bytes,11,opt,name=schedule,proto3" json:"schedule,omitempty"`
	Stdin    []byte             `protobuf:"bytes,21,opt,name=stdin,proto3" json:"stdin,omitempty"`
}

func (x *CommandRequest) Reset() {
	*x = CommandRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_network_v1_mmsp_stream_command_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommandRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandRequest) ProtoMessage() {}

func (x *CommandRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_network_v1_mmsp_stream_command_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandRequest.ProtoReflect.Descriptor instead.
func (*CommandRequest) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_network_v1_mmsp_stream_command_proto_rawDescGZIP(), []int{1}
}

func (x *CommandRequest) GetExec() *ops.CommandExec {
	if x != nil {
		return x.Exec
	}
	return nil
}

func (x *CommandRequest) GetSchedule() *datetime.DateTime {
	if x != nil {
		return x.Schedule
	}
	return nil
}

func (x *CommandRequest) GetStdin() []byte {
	if x != nil {
		return x.Stdin
	}
	return nil
}

type CommandResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestedCommand *CommandRequest        `protobuf:"bytes,1,opt,name=requestedCommand,proto3" json:"requestedCommand,omitempty"`
	Result           *ops.CommandResult     `protobuf:"bytes,11,opt,name=result,proto3" json:"result,omitempty"`
	Stdout           []byte                 `protobuf:"bytes,21,opt,name=stdout,proto3" json:"stdout,omitempty"`
	Stderr           []byte                 `protobuf:"bytes,22,opt,name=stderr,proto3" json:"stderr,omitempty"`
	StdoutStderr     []byte                 `protobuf:"bytes,23,opt,name=stdoutStderr,proto3" json:"stdoutStderr,omitempty"`
	Status           *status.StatusResponse `protobuf:"bytes,101,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *CommandResponse) Reset() {
	*x = CommandResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_network_v1_mmsp_stream_command_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommandResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandResponse) ProtoMessage() {}

func (x *CommandResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_network_v1_mmsp_stream_command_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandResponse.ProtoReflect.Descriptor instead.
func (*CommandResponse) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_network_v1_mmsp_stream_command_proto_rawDescGZIP(), []int{2}
}

func (x *CommandResponse) GetRequestedCommand() *CommandRequest {
	if x != nil {
		return x.RequestedCommand
	}
	return nil
}

func (x *CommandResponse) GetResult() *ops.CommandResult {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *CommandResponse) GetStdout() []byte {
	if x != nil {
		return x.Stdout
	}
	return nil
}

func (x *CommandResponse) GetStderr() []byte {
	if x != nil {
		return x.Stderr
	}
	return nil
}

func (x *CommandResponse) GetStdoutStderr() []byte {
	if x != nil {
		return x.StdoutStderr
	}
	return nil
}

func (x *CommandResponse) GetStatus() *status.StatusResponse {
	if x != nil {
		return x.Status
	}
	return nil
}

var File_mmesh_protobuf_network_v1_mmsp_stream_command_proto protoreflect.FileDescriptor

var file_mmesh_protobuf_network_v1_mmsp_stream_command_proto_rawDesc = []byte{
	0x0a, 0x33, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6d, 0x73, 0x70,
	0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x1a, 0x2e, 0x6d,
	0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x73, 0x2f, 0x77,
	0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x6d,
	0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x30, 0x6d, 0x6d, 0x65,
	0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x64,
	0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x70, 0x0a,
	0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x30, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x08, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x7c, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x24, 0x0a, 0x04, 0x65, 0x78, 0x65, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x6f, 0x70, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x45, 0x78, 0x65,
	0x63, 0x52, 0x04, 0x65, 0x78, 0x65, 0x63, 0x12, 0x2e, 0x0a, 0x08, 0x73, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x64, 0x61, 0x74, 0x65,
	0x74, 0x69, 0x6d, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x08, 0x73,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x64, 0x69, 0x6e,
	0x18, 0x15, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x73, 0x74, 0x64, 0x69, 0x6e, 0x22, 0x85, 0x02,
	0x0a, 0x0f, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x42, 0x0a, 0x10, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x43, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x52, 0x10, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x43, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x2a, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6f, 0x70, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x64, 0x6f, 0x75, 0x74, 0x18, 0x15, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x06, 0x73, 0x74, 0x64, 0x6f, 0x75, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x64,
	0x65, 0x72, 0x72, 0x18, 0x16, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x73, 0x74, 0x64, 0x65, 0x72,
	0x72, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x74, 0x64, 0x6f, 0x75, 0x74, 0x53, 0x74, 0x64, 0x65, 0x72,
	0x72, 0x18, 0x17, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x73, 0x74, 0x64, 0x6f, 0x75, 0x74, 0x53,
	0x74, 0x64, 0x65, 0x72, 0x72, 0x12, 0x2e, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x2d, 0x5a, 0x2b, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x64,
	0x65, 0x76, 0x2f, 0x6d, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x6d, 0x6d, 0x73, 0x70, 0x2f, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mmesh_protobuf_network_v1_mmsp_stream_command_proto_rawDescOnce sync.Once
	file_mmesh_protobuf_network_v1_mmsp_stream_command_proto_rawDescData = file_mmesh_protobuf_network_v1_mmsp_stream_command_proto_rawDesc
)

func file_mmesh_protobuf_network_v1_mmsp_stream_command_proto_rawDescGZIP() []byte {
	file_mmesh_protobuf_network_v1_mmsp_stream_command_proto_rawDescOnce.Do(func() {
		file_mmesh_protobuf_network_v1_mmsp_stream_command_proto_rawDescData = protoimpl.X.CompressGZIP(file_mmesh_protobuf_network_v1_mmsp_stream_command_proto_rawDescData)
	})
	return file_mmesh_protobuf_network_v1_mmsp_stream_command_proto_rawDescData
}

var file_mmesh_protobuf_network_v1_mmsp_stream_command_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_mmesh_protobuf_network_v1_mmsp_stream_command_proto_goTypes = []interface{}{
	(*Command)(nil),               // 0: stream.Command
	(*CommandRequest)(nil),        // 1: stream.CommandRequest
	(*CommandResponse)(nil),       // 2: stream.CommandResponse
	(*ops.CommandExec)(nil),       // 3: ops.CommandExec
	(*datetime.DateTime)(nil),     // 4: datetime.DateTime
	(*ops.CommandResult)(nil),     // 5: ops.CommandResult
	(*status.StatusResponse)(nil), // 6: status.StatusResponse
}
var file_mmesh_protobuf_network_v1_mmsp_stream_command_proto_depIdxs = []int32{
	1, // 0: stream.Command.request:type_name -> stream.CommandRequest
	2, // 1: stream.Command.response:type_name -> stream.CommandResponse
	3, // 2: stream.CommandRequest.exec:type_name -> ops.CommandExec
	4, // 3: stream.CommandRequest.schedule:type_name -> datetime.DateTime
	1, // 4: stream.CommandResponse.requestedCommand:type_name -> stream.CommandRequest
	5, // 5: stream.CommandResponse.result:type_name -> ops.CommandResult
	6, // 6: stream.CommandResponse.status:type_name -> status.StatusResponse
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_mmesh_protobuf_network_v1_mmsp_stream_command_proto_init() }
func file_mmesh_protobuf_network_v1_mmsp_stream_command_proto_init() {
	if File_mmesh_protobuf_network_v1_mmsp_stream_command_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mmesh_protobuf_network_v1_mmsp_stream_command_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Command); i {
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
		file_mmesh_protobuf_network_v1_mmsp_stream_command_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommandRequest); i {
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
		file_mmesh_protobuf_network_v1_mmsp_stream_command_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommandResponse); i {
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
			RawDescriptor: file_mmesh_protobuf_network_v1_mmsp_stream_command_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mmesh_protobuf_network_v1_mmsp_stream_command_proto_goTypes,
		DependencyIndexes: file_mmesh_protobuf_network_v1_mmsp_stream_command_proto_depIdxs,
		MessageInfos:      file_mmesh_protobuf_network_v1_mmsp_stream_command_proto_msgTypes,
	}.Build()
	File_mmesh_protobuf_network_v1_mmsp_stream_command_proto = out.File
	file_mmesh_protobuf_network_v1_mmsp_stream_command_proto_rawDesc = nil
	file_mmesh_protobuf_network_v1_mmsp_stream_command_proto_goTypes = nil
	file_mmesh_protobuf_network_v1_mmsp_stream_command_proto_depIdxs = nil
}
