// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.0
// source: mmesh/protobuf/resources/v1/ops/ops.proto

package ops

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

type OpsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountID string `protobuf:"bytes,1,opt,name=accountID,proto3" json:"accountID,omitempty"`
	TenantID  string `protobuf:"bytes,11,opt,name=tenantID,proto3" json:"tenantID,omitempty"`
}

func (x *OpsRequest) Reset() {
	*x = OpsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_ops_ops_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpsRequest) ProtoMessage() {}

func (x *OpsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_ops_ops_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpsRequest.ProtoReflect.Descriptor instead.
func (*OpsRequest) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_ops_ops_proto_rawDescGZIP(), []int{0}
}

func (x *OpsRequest) GetAccountID() string {
	if x != nil {
		return x.AccountID
	}
	return ""
}

func (x *OpsRequest) GetTenantID() string {
	if x != nil {
		return x.TenantID
	}
	return ""
}

type Ops struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stats    *Stats        `protobuf:"bytes,1,opt,name=stats,proto3" json:"stats,omitempty"`
	Projects []*OpsProject `protobuf:"bytes,11,rep,name=projects,proto3" json:"projects,omitempty"`
}

func (x *Ops) Reset() {
	*x = Ops{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_ops_ops_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ops) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ops) ProtoMessage() {}

func (x *Ops) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_ops_ops_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ops.ProtoReflect.Descriptor instead.
func (*Ops) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_ops_ops_proto_rawDescGZIP(), []int{1}
}

func (x *Ops) GetStats() *Stats {
	if x != nil {
		return x.Stats
	}
	return nil
}

func (x *Ops) GetProjects() []*OpsProject {
	if x != nil {
		return x.Projects
	}
	return nil
}

type OpsProject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Project   *Project        `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	Workflows []*Workflow     `protobuf:"bytes,11,rep,name=workflows,proto3" json:"workflows,omitempty"`
	Metrics   *ProjectMetrics `protobuf:"bytes,101,opt,name=metrics,proto3" json:"metrics,omitempty"`
}

func (x *OpsProject) Reset() {
	*x = OpsProject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_ops_ops_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpsProject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpsProject) ProtoMessage() {}

func (x *OpsProject) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_ops_ops_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpsProject.ProtoReflect.Descriptor instead.
func (*OpsProject) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_ops_ops_proto_rawDescGZIP(), []int{2}
}

func (x *OpsProject) GetProject() *Project {
	if x != nil {
		return x.Project
	}
	return nil
}

func (x *OpsProject) GetWorkflows() []*Workflow {
	if x != nil {
		return x.Workflows
	}
	return nil
}

func (x *OpsProject) GetMetrics() *ProjectMetrics {
	if x != nil {
		return x.Metrics
	}
	return nil
}

type Stats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalProjects  int32            `protobuf:"varint,11,opt,name=totalProjects,proto3" json:"totalProjects,omitempty"`
	TotalWorkflows int32            `protobuf:"varint,21,opt,name=totalWorkflows,proto3" json:"totalWorkflows,omitempty"`
	TotalTasks     int32            `protobuf:"varint,31,opt,name=totalTasks,proto3" json:"totalTasks,omitempty"`
	LastResult     *WorkflowMetrics `protobuf:"bytes,101,opt,name=lastResult,proto3" json:"lastResult,omitempty"`
}

func (x *Stats) Reset() {
	*x = Stats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_ops_ops_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stats) ProtoMessage() {}

func (x *Stats) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_ops_ops_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stats.ProtoReflect.Descriptor instead.
func (*Stats) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_ops_ops_proto_rawDescGZIP(), []int{3}
}

func (x *Stats) GetTotalProjects() int32 {
	if x != nil {
		return x.TotalProjects
	}
	return 0
}

func (x *Stats) GetTotalWorkflows() int32 {
	if x != nil {
		return x.TotalWorkflows
	}
	return 0
}

func (x *Stats) GetTotalTasks() int32 {
	if x != nil {
		return x.TotalTasks
	}
	return 0
}

func (x *Stats) GetLastResult() *WorkflowMetrics {
	if x != nil {
		return x.LastResult
	}
	return nil
}

type ProjectMetrics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalWorkflows int32            `protobuf:"varint,21,opt,name=totalWorkflows,proto3" json:"totalWorkflows,omitempty"`
	TotalTasks     int32            `protobuf:"varint,31,opt,name=totalTasks,proto3" json:"totalTasks,omitempty"`
	LastResult     *WorkflowMetrics `protobuf:"bytes,101,opt,name=lastResult,proto3" json:"lastResult,omitempty"`
}

func (x *ProjectMetrics) Reset() {
	*x = ProjectMetrics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_ops_ops_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectMetrics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectMetrics) ProtoMessage() {}

func (x *ProjectMetrics) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_ops_ops_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectMetrics.ProtoReflect.Descriptor instead.
func (*ProjectMetrics) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_ops_ops_proto_rawDescGZIP(), []int{4}
}

func (x *ProjectMetrics) GetTotalWorkflows() int32 {
	if x != nil {
		return x.TotalWorkflows
	}
	return 0
}

func (x *ProjectMetrics) GetTotalTasks() int32 {
	if x != nil {
		return x.TotalTasks
	}
	return 0
}

func (x *ProjectMetrics) GetLastResult() *WorkflowMetrics {
	if x != nil {
		return x.LastResult
	}
	return nil
}

type WorkflowMetrics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UnknownResultWorkflows int32 `protobuf:"varint,1,opt,name=unknownResultWorkflows,proto3" json:"unknownResultWorkflows,omitempty"`
	SuccessfulWorkflows    int32 `protobuf:"varint,11,opt,name=successfulWorkflows,proto3" json:"successfulWorkflows,omitempty"`
	FailedWorkflows        int32 `protobuf:"varint,21,opt,name=failedWorkflows,proto3" json:"failedWorkflows,omitempty"`
}

func (x *WorkflowMetrics) Reset() {
	*x = WorkflowMetrics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_ops_ops_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkflowMetrics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkflowMetrics) ProtoMessage() {}

func (x *WorkflowMetrics) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_ops_ops_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkflowMetrics.ProtoReflect.Descriptor instead.
func (*WorkflowMetrics) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_ops_ops_proto_rawDescGZIP(), []int{5}
}

func (x *WorkflowMetrics) GetUnknownResultWorkflows() int32 {
	if x != nil {
		return x.UnknownResultWorkflows
	}
	return 0
}

func (x *WorkflowMetrics) GetSuccessfulWorkflows() int32 {
	if x != nil {
		return x.SuccessfulWorkflows
	}
	return 0
}

func (x *WorkflowMetrics) GetFailedWorkflows() int32 {
	if x != nil {
		return x.FailedWorkflows
	}
	return 0
}

var File_mmesh_protobuf_resources_v1_ops_ops_proto protoreflect.FileDescriptor

var file_mmesh_protobuf_resources_v1_ops_ops_proto_rawDesc = []byte{
	0x0a, 0x29, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70,
	0x73, 0x2f, 0x6f, 0x70, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6f, 0x70, 0x73,
	0x1a, 0x2d, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70,
	0x73, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x2e, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x73,
	0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x46, 0x0a, 0x0a, 0x4f, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x74,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x22, 0x54, 0x0a, 0x03, 0x4f, 0x70, 0x73, 0x12, 0x20,
	0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x6f, 0x70, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73,
	0x12, 0x2b, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x18, 0x0b, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6f, 0x70, 0x73, 0x2e, 0x4f, 0x70, 0x73, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x22, 0x90, 0x01,
	0x0a, 0x0a, 0x4f, 0x70, 0x73, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x26, 0x0a, 0x07,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x6f, 0x70, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x07, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x12, 0x2b, 0x0a, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6f, 0x70, 0x73, 0x2e, 0x57, 0x6f,
	0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x73, 0x12, 0x2d, 0x0a, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18, 0x65, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6f, 0x70, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x22, 0xab, 0x01, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x12, 0x26, 0x0a, 0x0e, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f,
	0x77, 0x73, 0x18, 0x15, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x57,
	0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x12, 0x34, 0x0a, 0x0a, 0x6c, 0x61, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6f,
	0x70, 0x73, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x73, 0x52, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x8e,
	0x01, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x73, 0x12, 0x26, 0x0a, 0x0e, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c,
	0x6f, 0x77, 0x73, 0x18, 0x15, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x12, 0x34, 0x0a, 0x0a, 0x6c, 0x61, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x6f, 0x70, 0x73, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x4d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x73, 0x52, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22,
	0xa5, 0x01, 0x0a, 0x0f, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x4d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x73, 0x12, 0x36, 0x0a, 0x16, 0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x16, 0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x12, 0x30, 0x0a, 0x13, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f,
	0x77, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x13, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x66, 0x75, 0x6c, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x12, 0x28, 0x0a,
	0x0f, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73,
	0x18, 0x15, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x57, 0x6f,
	0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x42, 0x27, 0x5a, 0x25, 0x6d, 0x6d, 0x65, 0x73, 0x68,
	0x2e, 0x64, 0x65, 0x76, 0x2f, 0x6d, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x72,
	0x70, 0x63, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x6f, 0x70, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mmesh_protobuf_resources_v1_ops_ops_proto_rawDescOnce sync.Once
	file_mmesh_protobuf_resources_v1_ops_ops_proto_rawDescData = file_mmesh_protobuf_resources_v1_ops_ops_proto_rawDesc
)

func file_mmesh_protobuf_resources_v1_ops_ops_proto_rawDescGZIP() []byte {
	file_mmesh_protobuf_resources_v1_ops_ops_proto_rawDescOnce.Do(func() {
		file_mmesh_protobuf_resources_v1_ops_ops_proto_rawDescData = protoimpl.X.CompressGZIP(file_mmesh_protobuf_resources_v1_ops_ops_proto_rawDescData)
	})
	return file_mmesh_protobuf_resources_v1_ops_ops_proto_rawDescData
}

var file_mmesh_protobuf_resources_v1_ops_ops_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_mmesh_protobuf_resources_v1_ops_ops_proto_goTypes = []interface{}{
	(*OpsRequest)(nil),      // 0: ops.OpsRequest
	(*Ops)(nil),             // 1: ops.Ops
	(*OpsProject)(nil),      // 2: ops.OpsProject
	(*Stats)(nil),           // 3: ops.Stats
	(*ProjectMetrics)(nil),  // 4: ops.ProjectMetrics
	(*WorkflowMetrics)(nil), // 5: ops.WorkflowMetrics
	(*Project)(nil),         // 6: ops.Project
	(*Workflow)(nil),        // 7: ops.Workflow
}
var file_mmesh_protobuf_resources_v1_ops_ops_proto_depIdxs = []int32{
	3, // 0: ops.Ops.stats:type_name -> ops.Stats
	2, // 1: ops.Ops.projects:type_name -> ops.OpsProject
	6, // 2: ops.OpsProject.project:type_name -> ops.Project
	7, // 3: ops.OpsProject.workflows:type_name -> ops.Workflow
	4, // 4: ops.OpsProject.metrics:type_name -> ops.ProjectMetrics
	5, // 5: ops.Stats.lastResult:type_name -> ops.WorkflowMetrics
	5, // 6: ops.ProjectMetrics.lastResult:type_name -> ops.WorkflowMetrics
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_mmesh_protobuf_resources_v1_ops_ops_proto_init() }
func file_mmesh_protobuf_resources_v1_ops_ops_proto_init() {
	if File_mmesh_protobuf_resources_v1_ops_ops_proto != nil {
		return
	}
	file_mmesh_protobuf_resources_v1_ops_project_proto_init()
	file_mmesh_protobuf_resources_v1_ops_workflow_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_mmesh_protobuf_resources_v1_ops_ops_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpsRequest); i {
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
		file_mmesh_protobuf_resources_v1_ops_ops_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ops); i {
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
		file_mmesh_protobuf_resources_v1_ops_ops_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpsProject); i {
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
		file_mmesh_protobuf_resources_v1_ops_ops_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Stats); i {
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
		file_mmesh_protobuf_resources_v1_ops_ops_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectMetrics); i {
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
		file_mmesh_protobuf_resources_v1_ops_ops_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkflowMetrics); i {
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
			RawDescriptor: file_mmesh_protobuf_resources_v1_ops_ops_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mmesh_protobuf_resources_v1_ops_ops_proto_goTypes,
		DependencyIndexes: file_mmesh_protobuf_resources_v1_ops_ops_proto_depIdxs,
		MessageInfos:      file_mmesh_protobuf_resources_v1_ops_ops_proto_msgTypes,
	}.Build()
	File_mmesh_protobuf_resources_v1_ops_ops_proto = out.File
	file_mmesh_protobuf_resources_v1_ops_ops_proto_rawDesc = nil
	file_mmesh_protobuf_resources_v1_ops_ops_proto_goTypes = nil
	file_mmesh_protobuf_resources_v1_ops_ops_proto_depIdxs = nil
}
