// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: mmesh/protobuf/network/v1/resources/ae/workflow/job.proto

package workflow

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	command "mmesh.dev/m-api-go/grpc/network/mmsp/command"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Job struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	OnSuccess   []string `protobuf:"bytes,3,rep,name=onSuccess,proto3" json:"onSuccess,omitempty"` // array of workflow names
	OnFailure   []string `protobuf:"bytes,4,rep,name=onFailure,proto3" json:"onFailure,omitempty"` // array of workflow names
	Tasks       []*Task  `protobuf:"bytes,6,rep,name=tasks,proto3" json:"tasks,omitempty"`         // job steps
}

func (x *Job) Reset() {
	*x = Job{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_network_v1_resources_ae_workflow_job_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Job) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Job) ProtoMessage() {}

func (x *Job) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_network_v1_resources_ae_workflow_job_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Job.ProtoReflect.Descriptor instead.
func (*Job) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_network_v1_resources_ae_workflow_job_proto_rawDescGZIP(), []int{0}
}

func (x *Job) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Job) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Job) GetOnSuccess() []string {
	if x != nil {
		return x.OnSuccess
	}
	return nil
}

func (x *Job) GetOnFailure() []string {
	if x != nil {
		return x.OnFailure
	}
	return nil
}

func (x *Job) GetTasks() []*Task {
	if x != nil {
		return x.Tasks
	}
	return nil
}

type Task struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string    `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Actions     []*Action `protobuf:"bytes,4,rep,name=actions,proto3" json:"actions,omitempty"`
}

func (x *Task) Reset() {
	*x = Task{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_network_v1_resources_ae_workflow_job_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_network_v1_resources_ae_workflow_job_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_network_v1_resources_ae_workflow_job_proto_rawDescGZIP(), []int{1}
}

func (x *Task) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Task) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Task) GetActions() []*Action {
	if x != nil {
		return x.Actions
	}
	return nil
}

type Action struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type    string               `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`       // "command" | "docker"
	Command *command.CommandExec `protobuf:"bytes,3,opt,name=command,proto3" json:"command,omitempty"` // command.DockerRun dockerImg = 4;
}

func (x *Action) Reset() {
	*x = Action{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_network_v1_resources_ae_workflow_job_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Action) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Action) ProtoMessage() {}

func (x *Action) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_network_v1_resources_ae_workflow_job_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Action.ProtoReflect.Descriptor instead.
func (*Action) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_network_v1_resources_ae_workflow_job_proto_rawDescGZIP(), []int{2}
}

func (x *Action) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Action) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Action) GetCommand() *command.CommandExec {
	if x != nil {
		return x.Command
	}
	return nil
}

var File_mmesh_protobuf_network_v1_resources_ae_workflow_job_proto protoreflect.FileDescriptor

var file_mmesh_protobuf_network_v1_resources_ae_workflow_job_proto_rawDesc = []byte{
	0x0a, 0x39, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x61, 0x65, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f,
	0x77, 0x2f, 0x6a, 0x6f, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x77, 0x6f, 0x72,
	0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x1a, 0x2c, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x76, 0x31,
	0x2f, 0x6d, 0x6d, 0x73, 0x70, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x9d, 0x01, 0x0a, 0x03, 0x4a, 0x6f, 0x62, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x6e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x6e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12,
	0x1c, 0x0a, 0x09, 0x6f, 0x6e, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x09, 0x6f, 0x6e, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x12, 0x24, 0x0a,
	0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x77,
	0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x05, 0x74, 0x61,
	0x73, 0x6b, 0x73, 0x22, 0x68, 0x0a, 0x04, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x2a, 0x0a, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x60, 0x0a,
	0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x2e, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x45, 0x78, 0x65, 0x63, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x42,
	0x37, 0x5a, 0x35, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x6d, 0x2d, 0x61,
	0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x61, 0x65, 0x2f,
	0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mmesh_protobuf_network_v1_resources_ae_workflow_job_proto_rawDescOnce sync.Once
	file_mmesh_protobuf_network_v1_resources_ae_workflow_job_proto_rawDescData = file_mmesh_protobuf_network_v1_resources_ae_workflow_job_proto_rawDesc
)

func file_mmesh_protobuf_network_v1_resources_ae_workflow_job_proto_rawDescGZIP() []byte {
	file_mmesh_protobuf_network_v1_resources_ae_workflow_job_proto_rawDescOnce.Do(func() {
		file_mmesh_protobuf_network_v1_resources_ae_workflow_job_proto_rawDescData = protoimpl.X.CompressGZIP(file_mmesh_protobuf_network_v1_resources_ae_workflow_job_proto_rawDescData)
	})
	return file_mmesh_protobuf_network_v1_resources_ae_workflow_job_proto_rawDescData
}

var file_mmesh_protobuf_network_v1_resources_ae_workflow_job_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_mmesh_protobuf_network_v1_resources_ae_workflow_job_proto_goTypes = []interface{}{
	(*Job)(nil),                 // 0: workflow.Job
	(*Task)(nil),                // 1: workflow.Task
	(*Action)(nil),              // 2: workflow.Action
	(*command.CommandExec)(nil), // 3: command.CommandExec
}
var file_mmesh_protobuf_network_v1_resources_ae_workflow_job_proto_depIdxs = []int32{
	1, // 0: workflow.Job.tasks:type_name -> workflow.Task
	2, // 1: workflow.Task.actions:type_name -> workflow.Action
	3, // 2: workflow.Action.command:type_name -> command.CommandExec
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_mmesh_protobuf_network_v1_resources_ae_workflow_job_proto_init() }
func file_mmesh_protobuf_network_v1_resources_ae_workflow_job_proto_init() {
	if File_mmesh_protobuf_network_v1_resources_ae_workflow_job_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mmesh_protobuf_network_v1_resources_ae_workflow_job_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Job); i {
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
		file_mmesh_protobuf_network_v1_resources_ae_workflow_job_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Task); i {
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
		file_mmesh_protobuf_network_v1_resources_ae_workflow_job_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Action); i {
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
			RawDescriptor: file_mmesh_protobuf_network_v1_resources_ae_workflow_job_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mmesh_protobuf_network_v1_resources_ae_workflow_job_proto_goTypes,
		DependencyIndexes: file_mmesh_protobuf_network_v1_resources_ae_workflow_job_proto_depIdxs,
		MessageInfos:      file_mmesh_protobuf_network_v1_resources_ae_workflow_job_proto_msgTypes,
	}.Build()
	File_mmesh_protobuf_network_v1_resources_ae_workflow_job_proto = out.File
	file_mmesh_protobuf_network_v1_resources_ae_workflow_job_proto_rawDesc = nil
	file_mmesh_protobuf_network_v1_resources_ae_workflow_job_proto_goTypes = nil
	file_mmesh_protobuf_network_v1_resources_ae_workflow_job_proto_depIdxs = nil
}
