// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.0
// source: mmesh/protobuf/network/v1/mmsp/command.proto

package command

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	datetime "mmesh.dev/m-api-go/grpc/common/datetime"
	status "mmesh.dev/m-api-go/grpc/common/status"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CommandResultStatus int32

const (
	CommandResultStatus_EXECUTED CommandResultStatus = 0
	CommandResultStatus_RUNNING  CommandResultStatus = 1
	CommandResultStatus_FAILED   CommandResultStatus = 2
)

// Enum value maps for CommandResultStatus.
var (
	CommandResultStatus_name = map[int32]string{
		0: "EXECUTED",
		1: "RUNNING",
		2: "FAILED",
	}
	CommandResultStatus_value = map[string]int32{
		"EXECUTED": 0,
		"RUNNING":  1,
		"FAILED":   2,
	}
)

func (x CommandResultStatus) Enum() *CommandResultStatus {
	p := new(CommandResultStatus)
	*p = x
	return p
}

func (x CommandResultStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CommandResultStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_mmesh_protobuf_network_v1_mmsp_command_proto_enumTypes[0].Descriptor()
}

func (CommandResultStatus) Type() protoreflect.EnumType {
	return &file_mmesh_protobuf_network_v1_mmsp_command_proto_enumTypes[0]
}

func (x CommandResultStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CommandResultStatus.Descriptor instead.
func (CommandResultStatus) EnumDescriptor() ([]byte, []int) {
	return file_mmesh_protobuf_network_v1_mmsp_command_proto_rawDescGZIP(), []int{0}
}

type CommandType int32

const (
	CommandType_SHELL  CommandType = 0
	CommandType_ACTION CommandType = 1
)

// Enum value maps for CommandType.
var (
	CommandType_name = map[int32]string{
		0: "SHELL",
		1: "ACTION",
	}
	CommandType_value = map[string]int32{
		"SHELL":  0,
		"ACTION": 1,
	}
)

func (x CommandType) Enum() *CommandType {
	p := new(CommandType)
	*p = x
	return p
}

func (x CommandType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CommandType) Descriptor() protoreflect.EnumDescriptor {
	return file_mmesh_protobuf_network_v1_mmsp_command_proto_enumTypes[1].Descriptor()
}

func (CommandType) Type() protoreflect.EnumType {
	return &file_mmesh_protobuf_network_v1_mmsp_command_proto_enumTypes[1]
}

func (x CommandType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CommandType.Descriptor instead.
func (CommandType) EnumDescriptor() ([]byte, []int) {
	return file_mmesh_protobuf_network_v1_mmsp_command_proto_rawDescGZIP(), []int{1}
}

type Command struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID              string           `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	CommandType     CommandType      `protobuf:"varint,2,opt,name=commandType,proto3,enum=command.CommandType" json:"commandType,omitempty"`
	CommandRequest  *CommandRequest  `protobuf:"bytes,3,opt,name=commandRequest,proto3" json:"commandRequest,omitempty"`
	CommandResponse *CommandResponse `protobuf:"bytes,4,opt,name=commandResponse,proto3" json:"commandResponse,omitempty"`
}

func (x *Command) Reset() {
	*x = Command{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_network_v1_mmsp_command_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Command) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Command) ProtoMessage() {}

func (x *Command) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_network_v1_mmsp_command_proto_msgTypes[0]
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
	return file_mmesh_protobuf_network_v1_mmsp_command_proto_rawDescGZIP(), []int{0}
}

func (x *Command) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Command) GetCommandType() CommandType {
	if x != nil {
		return x.CommandType
	}
	return CommandType_SHELL
}

func (x *Command) GetCommandRequest() *CommandRequest {
	if x != nil {
		return x.CommandRequest
	}
	return nil
}

func (x *Command) GetCommandResponse() *CommandResponse {
	if x != nil {
		return x.CommandResponse
	}
	return nil
}

type CommandRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Command  *CommandExec       `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"`
	Schedule *datetime.DateTime `protobuf:"bytes,2,opt,name=schedule,proto3" json:"schedule,omitempty"`
	Stdin    []byte             `protobuf:"bytes,3,opt,name=stdin,proto3" json:"stdin,omitempty"`
}

func (x *CommandRequest) Reset() {
	*x = CommandRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_network_v1_mmsp_command_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommandRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandRequest) ProtoMessage() {}

func (x *CommandRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_network_v1_mmsp_command_proto_msgTypes[1]
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
	return file_mmesh_protobuf_network_v1_mmsp_command_proto_rawDescGZIP(), []int{1}
}

func (x *CommandRequest) GetCommand() *CommandExec {
	if x != nil {
		return x.Command
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
	Result           *CommandResult         `protobuf:"bytes,2,opt,name=result,proto3" json:"result,omitempty"`
	Stdout           []byte                 `protobuf:"bytes,3,opt,name=stdout,proto3" json:"stdout,omitempty"`
	Stderr           []byte                 `protobuf:"bytes,4,opt,name=stderr,proto3" json:"stderr,omitempty"`
	StdoutStderr     []byte                 `protobuf:"bytes,5,opt,name=stdoutStderr,proto3" json:"stdoutStderr,omitempty"`
	Status           *status.StatusResponse `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *CommandResponse) Reset() {
	*x = CommandResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_network_v1_mmsp_command_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommandResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandResponse) ProtoMessage() {}

func (x *CommandResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_network_v1_mmsp_command_proto_msgTypes[2]
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
	return file_mmesh_protobuf_network_v1_mmsp_command_proto_rawDescGZIP(), []int{2}
}

func (x *CommandResponse) GetRequestedCommand() *CommandRequest {
	if x != nil {
		return x.RequestedCommand
	}
	return nil
}

func (x *CommandResponse) GetResult() *CommandResult {
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

type CommandExec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cmd  string   `protobuf:"bytes,1,opt,name=cmd,proto3" json:"cmd,omitempty"`
	Args []string `protobuf:"bytes,2,rep,name=args,proto3" json:"args,omitempty"`
	UID  int32    `protobuf:"varint,3,opt,name=UID,proto3" json:"UID,omitempty"`
	GID  int32    `protobuf:"varint,4,opt,name=GID,proto3" json:"GID,omitempty"`
}

func (x *CommandExec) Reset() {
	*x = CommandExec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_network_v1_mmsp_command_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommandExec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandExec) ProtoMessage() {}

func (x *CommandExec) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_network_v1_mmsp_command_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandExec.ProtoReflect.Descriptor instead.
func (*CommandExec) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_network_v1_mmsp_command_proto_rawDescGZIP(), []int{3}
}

func (x *CommandExec) GetCmd() string {
	if x != nil {
		return x.Cmd
	}
	return ""
}

func (x *CommandExec) GetArgs() []string {
	if x != nil {
		return x.Args
	}
	return nil
}

func (x *CommandExec) GetUID() int32 {
	if x != nil {
		return x.UID
	}
	return 0
}

func (x *CommandExec) GetGID() int32 {
	if x != nil {
		return x.GID
	}
	return 0
}

type CommandResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status   CommandResultStatus `protobuf:"varint,1,opt,name=status,proto3,enum=command.CommandResultStatus" json:"status,omitempty"` // executed, canceled...
	Duration int64               `protobuf:"varint,2,opt,name=duration,proto3" json:"duration,omitempty"`
}

func (x *CommandResult) Reset() {
	*x = CommandResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_network_v1_mmsp_command_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommandResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandResult) ProtoMessage() {}

func (x *CommandResult) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_network_v1_mmsp_command_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandResult.ProtoReflect.Descriptor instead.
func (*CommandResult) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_network_v1_mmsp_command_proto_rawDescGZIP(), []int{4}
}

func (x *CommandResult) GetStatus() CommandResultStatus {
	if x != nil {
		return x.Status
	}
	return CommandResultStatus_EXECUTED
}

func (x *CommandResult) GetDuration() int64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

var File_mmesh_protobuf_network_v1_mmsp_command_proto protoreflect.FileDescriptor

var file_mmesh_protobuf_network_v1_mmsp_command_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6d, 0x73, 0x70,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07,
	0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x1a, 0x2c, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76,
	0x31, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x30, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f,
	0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd6, 0x01, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x49, 0x44, 0x12, 0x36, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b,
	0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x3f, 0x0a, 0x0e, 0x63,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x43, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0e, 0x63, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x42, 0x0a, 0x0f,
	0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e,
	0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52,
	0x0f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x86, 0x01, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x45, 0x78, 0x65, 0x63, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x12, 0x2e, 0x0a, 0x08, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65,
	0x2e, 0x44, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x08, 0x73, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x64, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x05, 0x73, 0x74, 0x64, 0x69, 0x6e, 0x22, 0x8a, 0x02, 0x0a, 0x0f, 0x43, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a,
	0x10, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x52, 0x10, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x12, 0x2e, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x64, 0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x06, 0x73, 0x74, 0x64, 0x6f, 0x75, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x64, 0x65, 0x72, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x73, 0x74, 0x64, 0x65,
	0x72, 0x72, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x74, 0x64, 0x6f, 0x75, 0x74, 0x53, 0x74, 0x64, 0x65,
	0x72, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x73, 0x74, 0x64, 0x6f, 0x75, 0x74,
	0x53, 0x74, 0x64, 0x65, 0x72, 0x72, 0x12, 0x2e, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x57, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x45, 0x78, 0x65, 0x63, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x6d, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x63, 0x6d, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x67, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x61, 0x72, 0x67, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x55,
	0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x55, 0x49, 0x44, 0x12, 0x10, 0x0a,
	0x03, 0x47, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x47, 0x49, 0x44, 0x22,
	0x61, 0x0a, 0x0d, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x34, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2a, 0x3c, 0x0a, 0x13, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0c, 0x0a, 0x08, 0x45, 0x58, 0x45,
	0x43, 0x55, 0x54, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x55, 0x4e, 0x4e, 0x49,
	0x4e, 0x47, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x02,
	0x2a, 0x24, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x09, 0x0a, 0x05, 0x53, 0x48, 0x45, 0x4c, 0x4c, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x43,
	0x54, 0x49, 0x4f, 0x4e, 0x10, 0x01, 0x42, 0x2e, 0x5a, 0x2c, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2e,
	0x64, 0x65, 0x76, 0x2f, 0x6d, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x6d, 0x6d, 0x73, 0x70, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mmesh_protobuf_network_v1_mmsp_command_proto_rawDescOnce sync.Once
	file_mmesh_protobuf_network_v1_mmsp_command_proto_rawDescData = file_mmesh_protobuf_network_v1_mmsp_command_proto_rawDesc
)

func file_mmesh_protobuf_network_v1_mmsp_command_proto_rawDescGZIP() []byte {
	file_mmesh_protobuf_network_v1_mmsp_command_proto_rawDescOnce.Do(func() {
		file_mmesh_protobuf_network_v1_mmsp_command_proto_rawDescData = protoimpl.X.CompressGZIP(file_mmesh_protobuf_network_v1_mmsp_command_proto_rawDescData)
	})
	return file_mmesh_protobuf_network_v1_mmsp_command_proto_rawDescData
}

var file_mmesh_protobuf_network_v1_mmsp_command_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_mmesh_protobuf_network_v1_mmsp_command_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_mmesh_protobuf_network_v1_mmsp_command_proto_goTypes = []interface{}{
	(CommandResultStatus)(0),      // 0: command.CommandResultStatus
	(CommandType)(0),              // 1: command.CommandType
	(*Command)(nil),               // 2: command.Command
	(*CommandRequest)(nil),        // 3: command.CommandRequest
	(*CommandResponse)(nil),       // 4: command.CommandResponse
	(*CommandExec)(nil),           // 5: command.CommandExec
	(*CommandResult)(nil),         // 6: command.CommandResult
	(*datetime.DateTime)(nil),     // 7: datetime.DateTime
	(*status.StatusResponse)(nil), // 8: status.StatusResponse
}
var file_mmesh_protobuf_network_v1_mmsp_command_proto_depIdxs = []int32{
	1, // 0: command.Command.commandType:type_name -> command.CommandType
	3, // 1: command.Command.commandRequest:type_name -> command.CommandRequest
	4, // 2: command.Command.commandResponse:type_name -> command.CommandResponse
	5, // 3: command.CommandRequest.command:type_name -> command.CommandExec
	7, // 4: command.CommandRequest.schedule:type_name -> datetime.DateTime
	3, // 5: command.CommandResponse.requestedCommand:type_name -> command.CommandRequest
	6, // 6: command.CommandResponse.result:type_name -> command.CommandResult
	8, // 7: command.CommandResponse.status:type_name -> status.StatusResponse
	0, // 8: command.CommandResult.status:type_name -> command.CommandResultStatus
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_mmesh_protobuf_network_v1_mmsp_command_proto_init() }
func file_mmesh_protobuf_network_v1_mmsp_command_proto_init() {
	if File_mmesh_protobuf_network_v1_mmsp_command_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mmesh_protobuf_network_v1_mmsp_command_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_mmesh_protobuf_network_v1_mmsp_command_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_mmesh_protobuf_network_v1_mmsp_command_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_mmesh_protobuf_network_v1_mmsp_command_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommandExec); i {
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
		file_mmesh_protobuf_network_v1_mmsp_command_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommandResult); i {
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
			RawDescriptor: file_mmesh_protobuf_network_v1_mmsp_command_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mmesh_protobuf_network_v1_mmsp_command_proto_goTypes,
		DependencyIndexes: file_mmesh_protobuf_network_v1_mmsp_command_proto_depIdxs,
		EnumInfos:         file_mmesh_protobuf_network_v1_mmsp_command_proto_enumTypes,
		MessageInfos:      file_mmesh_protobuf_network_v1_mmsp_command_proto_msgTypes,
	}.Build()
	File_mmesh_protobuf_network_v1_mmsp_command_proto = out.File
	file_mmesh_protobuf_network_v1_mmsp_command_proto_rawDesc = nil
	file_mmesh_protobuf_network_v1_mmsp_command_proto_goTypes = nil
	file_mmesh_protobuf_network_v1_mmsp_command_proto_depIdxs = nil
}
