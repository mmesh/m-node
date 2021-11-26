// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: mmesh/protobuf/resources/v1/metrics/node.proto

package metrics

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

type HostMetrics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OS             string  `protobuf:"bytes,1,opt,name=OS,proto3" json:"OS,omitempty"`
	Uptime         string  `protobuf:"bytes,11,opt,name=uptime,proto3" json:"uptime,omitempty"`
	LoadAvg        float64 `protobuf:"fixed64,21,opt,name=loadAvg,proto3" json:"loadAvg,omitempty"`
	CpuUsage       uint64  `protobuf:"varint,31,opt,name=cpuUsage,proto3" json:"cpuUsage,omitempty"`
	CpuPressure    bool    `protobuf:"varint,32,opt,name=cpuPressure,proto3" json:"cpuPressure,omitempty"`
	MemoryUsage    uint64  `protobuf:"varint,41,opt,name=memoryUsage,proto3" json:"memoryUsage,omitempty"`
	MemoryPressure bool    `protobuf:"varint,42,opt,name=memoryPressure,proto3" json:"memoryPressure,omitempty"`
	DiskUsage      uint64  `protobuf:"varint,51,opt,name=diskUsage,proto3" json:"diskUsage,omitempty"`
	DiskPressure   bool    `protobuf:"varint,52,opt,name=diskPressure,proto3" json:"diskPressure,omitempty"`
}

func (x *HostMetrics) Reset() {
	*x = HostMetrics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_metrics_node_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HostMetrics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HostMetrics) ProtoMessage() {}

func (x *HostMetrics) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_metrics_node_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HostMetrics.ProtoReflect.Descriptor instead.
func (*HostMetrics) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_metrics_node_proto_rawDescGZIP(), []int{0}
}

func (x *HostMetrics) GetOS() string {
	if x != nil {
		return x.OS
	}
	return ""
}

func (x *HostMetrics) GetUptime() string {
	if x != nil {
		return x.Uptime
	}
	return ""
}

func (x *HostMetrics) GetLoadAvg() float64 {
	if x != nil {
		return x.LoadAvg
	}
	return 0
}

func (x *HostMetrics) GetCpuUsage() uint64 {
	if x != nil {
		return x.CpuUsage
	}
	return 0
}

func (x *HostMetrics) GetCpuPressure() bool {
	if x != nil {
		return x.CpuPressure
	}
	return false
}

func (x *HostMetrics) GetMemoryUsage() uint64 {
	if x != nil {
		return x.MemoryUsage
	}
	return 0
}

func (x *HostMetrics) GetMemoryPressure() bool {
	if x != nil {
		return x.MemoryPressure
	}
	return false
}

func (x *HostMetrics) GetDiskUsage() uint64 {
	if x != nil {
		return x.DiskUsage
	}
	return 0
}

func (x *HostMetrics) GetDiskPressure() bool {
	if x != nil {
		return x.DiskPressure
	}
	return false
}

type NetworkMetrics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TxBps              float32 `protobuf:"fixed32,1,opt,name=txBps,proto3" json:"txBps,omitempty"` // last minute tx bps
	RxBps              float32 `protobuf:"fixed32,2,opt,name=rxBps,proto3" json:"rxBps,omitempty"` // last minute rx bps
	CountHourlyReset   bool    `protobuf:"varint,3,opt,name=countHourlyReset,proto3" json:"countHourlyReset,omitempty"`
	CountDailyReset    bool    `protobuf:"varint,4,opt,name=countDailyReset,proto3" json:"countDailyReset,omitempty"`
	CountMonthlyReset  bool    `protobuf:"varint,5,opt,name=countMonthlyReset,proto3" json:"countMonthlyReset,omitempty"`
	TxHourlyBytes      uint64  `protobuf:"varint,6,opt,name=txHourlyBytes,proto3" json:"txHourlyBytes,omitempty"` // total tx bytes
	RxHourlyBytes      uint64  `protobuf:"varint,7,opt,name=rxHourlyBytes,proto3" json:"rxHourlyBytes,omitempty"` // total rx bytes
	TxHourlyPkts       uint64  `protobuf:"varint,8,opt,name=txHourlyPkts,proto3" json:"txHourlyPkts,omitempty"`
	RxHourlyPkts       uint64  `protobuf:"varint,9,opt,name=rxHourlyPkts,proto3" json:"rxHourlyPkts,omitempty"`
	DroppedHourlyPkts  uint64  `protobuf:"varint,10,opt,name=droppedHourlyPkts,proto3" json:"droppedHourlyPkts,omitempty"`
	TxDailyBytes       uint64  `protobuf:"varint,11,opt,name=txDailyBytes,proto3" json:"txDailyBytes,omitempty"` // total tx bytes
	RxDailyBytes       uint64  `protobuf:"varint,12,opt,name=rxDailyBytes,proto3" json:"rxDailyBytes,omitempty"` // total rx bytes
	TxDailyPkts        uint64  `protobuf:"varint,13,opt,name=txDailyPkts,proto3" json:"txDailyPkts,omitempty"`
	RxDailyPkts        uint64  `protobuf:"varint,14,opt,name=rxDailyPkts,proto3" json:"rxDailyPkts,omitempty"`
	DroppedDailyPkts   uint64  `protobuf:"varint,15,opt,name=droppedDailyPkts,proto3" json:"droppedDailyPkts,omitempty"`
	TxMonthlyBytes     uint64  `protobuf:"varint,16,opt,name=txMonthlyBytes,proto3" json:"txMonthlyBytes,omitempty"` // total tx bytes
	RxMonthlyBytes     uint64  `protobuf:"varint,17,opt,name=rxMonthlyBytes,proto3" json:"rxMonthlyBytes,omitempty"` // total rx bytes
	TxMonthlyPkts      uint64  `protobuf:"varint,18,opt,name=txMonthlyPkts,proto3" json:"txMonthlyPkts,omitempty"`
	RxMonthlyPkts      uint64  `protobuf:"varint,19,opt,name=rxMonthlyPkts,proto3" json:"rxMonthlyPkts,omitempty"`
	DroppedMonthlyPkts uint64  `protobuf:"varint,20,opt,name=droppedMonthlyPkts,proto3" json:"droppedMonthlyPkts,omitempty"`
	TxTotalBytes       uint64  `protobuf:"varint,21,opt,name=txTotalBytes,proto3" json:"txTotalBytes,omitempty"` // monthly tx bytes
	RxTotalBytes       uint64  `protobuf:"varint,22,opt,name=rxTotalBytes,proto3" json:"rxTotalBytes,omitempty"` // monthly rx bytes
	TxTotalPkts        uint64  `protobuf:"varint,23,opt,name=txTotalPkts,proto3" json:"txTotalPkts,omitempty"`
	RxTotalPkts        uint64  `protobuf:"varint,24,opt,name=rxTotalPkts,proto3" json:"rxTotalPkts,omitempty"`
	DroppedPkts        uint64  `protobuf:"varint,25,opt,name=droppedPkts,proto3" json:"droppedPkts,omitempty"`
}

func (x *NetworkMetrics) Reset() {
	*x = NetworkMetrics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_metrics_node_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkMetrics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkMetrics) ProtoMessage() {}

func (x *NetworkMetrics) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_metrics_node_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkMetrics.ProtoReflect.Descriptor instead.
func (*NetworkMetrics) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_metrics_node_proto_rawDescGZIP(), []int{1}
}

func (x *NetworkMetrics) GetTxBps() float32 {
	if x != nil {
		return x.TxBps
	}
	return 0
}

func (x *NetworkMetrics) GetRxBps() float32 {
	if x != nil {
		return x.RxBps
	}
	return 0
}

func (x *NetworkMetrics) GetCountHourlyReset() bool {
	if x != nil {
		return x.CountHourlyReset
	}
	return false
}

func (x *NetworkMetrics) GetCountDailyReset() bool {
	if x != nil {
		return x.CountDailyReset
	}
	return false
}

func (x *NetworkMetrics) GetCountMonthlyReset() bool {
	if x != nil {
		return x.CountMonthlyReset
	}
	return false
}

func (x *NetworkMetrics) GetTxHourlyBytes() uint64 {
	if x != nil {
		return x.TxHourlyBytes
	}
	return 0
}

func (x *NetworkMetrics) GetRxHourlyBytes() uint64 {
	if x != nil {
		return x.RxHourlyBytes
	}
	return 0
}

func (x *NetworkMetrics) GetTxHourlyPkts() uint64 {
	if x != nil {
		return x.TxHourlyPkts
	}
	return 0
}

func (x *NetworkMetrics) GetRxHourlyPkts() uint64 {
	if x != nil {
		return x.RxHourlyPkts
	}
	return 0
}

func (x *NetworkMetrics) GetDroppedHourlyPkts() uint64 {
	if x != nil {
		return x.DroppedHourlyPkts
	}
	return 0
}

func (x *NetworkMetrics) GetTxDailyBytes() uint64 {
	if x != nil {
		return x.TxDailyBytes
	}
	return 0
}

func (x *NetworkMetrics) GetRxDailyBytes() uint64 {
	if x != nil {
		return x.RxDailyBytes
	}
	return 0
}

func (x *NetworkMetrics) GetTxDailyPkts() uint64 {
	if x != nil {
		return x.TxDailyPkts
	}
	return 0
}

func (x *NetworkMetrics) GetRxDailyPkts() uint64 {
	if x != nil {
		return x.RxDailyPkts
	}
	return 0
}

func (x *NetworkMetrics) GetDroppedDailyPkts() uint64 {
	if x != nil {
		return x.DroppedDailyPkts
	}
	return 0
}

func (x *NetworkMetrics) GetTxMonthlyBytes() uint64 {
	if x != nil {
		return x.TxMonthlyBytes
	}
	return 0
}

func (x *NetworkMetrics) GetRxMonthlyBytes() uint64 {
	if x != nil {
		return x.RxMonthlyBytes
	}
	return 0
}

func (x *NetworkMetrics) GetTxMonthlyPkts() uint64 {
	if x != nil {
		return x.TxMonthlyPkts
	}
	return 0
}

func (x *NetworkMetrics) GetRxMonthlyPkts() uint64 {
	if x != nil {
		return x.RxMonthlyPkts
	}
	return 0
}

func (x *NetworkMetrics) GetDroppedMonthlyPkts() uint64 {
	if x != nil {
		return x.DroppedMonthlyPkts
	}
	return 0
}

func (x *NetworkMetrics) GetTxTotalBytes() uint64 {
	if x != nil {
		return x.TxTotalBytes
	}
	return 0
}

func (x *NetworkMetrics) GetRxTotalBytes() uint64 {
	if x != nil {
		return x.RxTotalBytes
	}
	return 0
}

func (x *NetworkMetrics) GetTxTotalPkts() uint64 {
	if x != nil {
		return x.TxTotalPkts
	}
	return 0
}

func (x *NetworkMetrics) GetRxTotalPkts() uint64 {
	if x != nil {
		return x.RxTotalPkts
	}
	return 0
}

func (x *NetworkMetrics) GetDroppedPkts() uint64 {
	if x != nil {
		return x.DroppedPkts
	}
	return 0
}

type RelayMetrics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Connections int32 `protobuf:"varint,1,opt,name=connections,proto3" json:"connections,omitempty"`
}

func (x *RelayMetrics) Reset() {
	*x = RelayMetrics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_metrics_node_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelayMetrics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelayMetrics) ProtoMessage() {}

func (x *RelayMetrics) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_metrics_node_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelayMetrics.ProtoReflect.Descriptor instead.
func (*RelayMetrics) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_metrics_node_proto_rawDescGZIP(), []int{2}
}

func (x *RelayMetrics) GetConnections() int32 {
	if x != nil {
		return x.Connections
	}
	return 0
}

var File_mmesh_protobuf_resources_v1_metrics_node_proto protoreflect.FileDescriptor

var file_mmesh_protobuf_resources_v1_metrics_node_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x73, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x22, 0x99, 0x02, 0x0a, 0x0b, 0x48, 0x6f,
	0x73, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x4f, 0x53, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x4f, 0x53, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x70, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x70, 0x74, 0x69, 0x6d,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x6c, 0x6f, 0x61, 0x64, 0x41, 0x76, 0x67, 0x18, 0x15, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x07, 0x6c, 0x6f, 0x61, 0x64, 0x41, 0x76, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x63,
	0x70, 0x75, 0x55, 0x73, 0x61, 0x67, 0x65, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x63,
	0x70, 0x75, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x70, 0x75, 0x50, 0x72,
	0x65, 0x73, 0x73, 0x75, 0x72, 0x65, 0x18, 0x20, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x63, 0x70,
	0x75, 0x50, 0x72, 0x65, 0x73, 0x73, 0x75, 0x72, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x6d, 0x65, 0x6d,
	0x6f, 0x72, 0x79, 0x55, 0x73, 0x61, 0x67, 0x65, 0x18, 0x29, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b,
	0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x6d,
	0x65, 0x6d, 0x6f, 0x72, 0x79, 0x50, 0x72, 0x65, 0x73, 0x73, 0x75, 0x72, 0x65, 0x18, 0x2a, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0e, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x50, 0x72, 0x65, 0x73, 0x73,
	0x75, 0x72, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x69, 0x73, 0x6b, 0x55, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x33, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x64, 0x69, 0x73, 0x6b, 0x55, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x22, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x6b, 0x50, 0x72, 0x65, 0x73, 0x73, 0x75, 0x72,
	0x65, 0x18, 0x34, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x64, 0x69, 0x73, 0x6b, 0x50, 0x72, 0x65,
	0x73, 0x73, 0x75, 0x72, 0x65, 0x22, 0xb4, 0x07, 0x0a, 0x0e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x78, 0x42, 0x70,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x74, 0x78, 0x42, 0x70, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x72, 0x78, 0x42, 0x70, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x72,
	0x78, 0x42, 0x70, 0x73, 0x12, 0x2a, 0x0a, 0x10, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x48, 0x6f, 0x75,
	0x72, 0x6c, 0x79, 0x52, 0x65, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x48, 0x6f, 0x75, 0x72, 0x6c, 0x79, 0x52, 0x65, 0x73, 0x65, 0x74,
	0x12, 0x28, 0x0a, 0x0f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x52, 0x65,
	0x73, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x44, 0x61, 0x69, 0x6c, 0x79, 0x52, 0x65, 0x73, 0x65, 0x74, 0x12, 0x2c, 0x0a, 0x11, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x52, 0x65, 0x73, 0x65, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4d, 0x6f, 0x6e, 0x74,
	0x68, 0x6c, 0x79, 0x52, 0x65, 0x73, 0x65, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x74, 0x78, 0x48, 0x6f,
	0x75, 0x72, 0x6c, 0x79, 0x42, 0x79, 0x74, 0x65, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0d, 0x74, 0x78, 0x48, 0x6f, 0x75, 0x72, 0x6c, 0x79, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x24,
	0x0a, 0x0d, 0x72, 0x78, 0x48, 0x6f, 0x75, 0x72, 0x6c, 0x79, 0x42, 0x79, 0x74, 0x65, 0x73, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x72, 0x78, 0x48, 0x6f, 0x75, 0x72, 0x6c, 0x79, 0x42,
	0x79, 0x74, 0x65, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x78, 0x48, 0x6f, 0x75, 0x72, 0x6c, 0x79,
	0x50, 0x6b, 0x74, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x74, 0x78, 0x48, 0x6f,
	0x75, 0x72, 0x6c, 0x79, 0x50, 0x6b, 0x74, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x78, 0x48, 0x6f,
	0x75, 0x72, 0x6c, 0x79, 0x50, 0x6b, 0x74, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c,
	0x72, 0x78, 0x48, 0x6f, 0x75, 0x72, 0x6c, 0x79, 0x50, 0x6b, 0x74, 0x73, 0x12, 0x2c, 0x0a, 0x11,
	0x64, 0x72, 0x6f, 0x70, 0x70, 0x65, 0x64, 0x48, 0x6f, 0x75, 0x72, 0x6c, 0x79, 0x50, 0x6b, 0x74,
	0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x04, 0x52, 0x11, 0x64, 0x72, 0x6f, 0x70, 0x70, 0x65, 0x64,
	0x48, 0x6f, 0x75, 0x72, 0x6c, 0x79, 0x50, 0x6b, 0x74, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x78,
	0x44, 0x61, 0x69, 0x6c, 0x79, 0x42, 0x79, 0x74, 0x65, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0c, 0x74, 0x78, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x22,
	0x0a, 0x0c, 0x72, 0x78, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x42, 0x79, 0x74, 0x65, 0x73, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x72, 0x78, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x42, 0x79, 0x74,
	0x65, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x78, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x50, 0x6b, 0x74,
	0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x74, 0x78, 0x44, 0x61, 0x69, 0x6c, 0x79,
	0x50, 0x6b, 0x74, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x78, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x50,
	0x6b, 0x74, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x72, 0x78, 0x44, 0x61, 0x69,
	0x6c, 0x79, 0x50, 0x6b, 0x74, 0x73, 0x12, 0x2a, 0x0a, 0x10, 0x64, 0x72, 0x6f, 0x70, 0x70, 0x65,
	0x64, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x50, 0x6b, 0x74, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x10, 0x64, 0x72, 0x6f, 0x70, 0x70, 0x65, 0x64, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x50, 0x6b,
	0x74, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x74, 0x78, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x42,
	0x79, 0x74, 0x65, 0x73, 0x18, 0x10, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x74, 0x78, 0x4d, 0x6f,
	0x6e, 0x74, 0x68, 0x6c, 0x79, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x72, 0x78,
	0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x42, 0x79, 0x74, 0x65, 0x73, 0x18, 0x11, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x0e, 0x72, 0x78, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x42, 0x79, 0x74,
	0x65, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x74, 0x78, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x50,
	0x6b, 0x74, 0x73, 0x18, 0x12, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x74, 0x78, 0x4d, 0x6f, 0x6e,
	0x74, 0x68, 0x6c, 0x79, 0x50, 0x6b, 0x74, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x72, 0x78, 0x4d, 0x6f,
	0x6e, 0x74, 0x68, 0x6c, 0x79, 0x50, 0x6b, 0x74, 0x73, 0x18, 0x13, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0d, 0x72, 0x78, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x50, 0x6b, 0x74, 0x73, 0x12, 0x2e,
	0x0a, 0x12, 0x64, 0x72, 0x6f, 0x70, 0x70, 0x65, 0x64, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79,
	0x50, 0x6b, 0x74, 0x73, 0x18, 0x14, 0x20, 0x01, 0x28, 0x04, 0x52, 0x12, 0x64, 0x72, 0x6f, 0x70,
	0x70, 0x65, 0x64, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x50, 0x6b, 0x74, 0x73, 0x12, 0x22,
	0x0a, 0x0c, 0x74, 0x78, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x42, 0x79, 0x74, 0x65, 0x73, 0x18, 0x15,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x74, 0x78, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x42, 0x79, 0x74,
	0x65, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x78, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x42, 0x79, 0x74,
	0x65, 0x73, 0x18, 0x16, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x72, 0x78, 0x54, 0x6f, 0x74, 0x61,
	0x6c, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x78, 0x54, 0x6f, 0x74, 0x61,
	0x6c, 0x50, 0x6b, 0x74, 0x73, 0x18, 0x17, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x74, 0x78, 0x54,
	0x6f, 0x74, 0x61, 0x6c, 0x50, 0x6b, 0x74, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x78, 0x54, 0x6f,
	0x74, 0x61, 0x6c, 0x50, 0x6b, 0x74, 0x73, 0x18, 0x18, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x72,
	0x78, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x6b, 0x74, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x72,
	0x6f, 0x70, 0x70, 0x65, 0x64, 0x50, 0x6b, 0x74, 0x73, 0x18, 0x19, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0b, 0x64, 0x72, 0x6f, 0x70, 0x70, 0x65, 0x64, 0x50, 0x6b, 0x74, 0x73, 0x22, 0x30, 0x0a, 0x0c,
	0x52, 0x65, 0x6c, 0x61, 0x79, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x20, 0x0a, 0x0b,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x2b,
	0x5a, 0x29, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x6d, 0x2d, 0x61, 0x70,
	0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_mmesh_protobuf_resources_v1_metrics_node_proto_rawDescOnce sync.Once
	file_mmesh_protobuf_resources_v1_metrics_node_proto_rawDescData = file_mmesh_protobuf_resources_v1_metrics_node_proto_rawDesc
)

func file_mmesh_protobuf_resources_v1_metrics_node_proto_rawDescGZIP() []byte {
	file_mmesh_protobuf_resources_v1_metrics_node_proto_rawDescOnce.Do(func() {
		file_mmesh_protobuf_resources_v1_metrics_node_proto_rawDescData = protoimpl.X.CompressGZIP(file_mmesh_protobuf_resources_v1_metrics_node_proto_rawDescData)
	})
	return file_mmesh_protobuf_resources_v1_metrics_node_proto_rawDescData
}

var file_mmesh_protobuf_resources_v1_metrics_node_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_mmesh_protobuf_resources_v1_metrics_node_proto_goTypes = []interface{}{
	(*HostMetrics)(nil),    // 0: metrics.HostMetrics
	(*NetworkMetrics)(nil), // 1: metrics.NetworkMetrics
	(*RelayMetrics)(nil),   // 2: metrics.RelayMetrics
}
var file_mmesh_protobuf_resources_v1_metrics_node_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_mmesh_protobuf_resources_v1_metrics_node_proto_init() }
func file_mmesh_protobuf_resources_v1_metrics_node_proto_init() {
	if File_mmesh_protobuf_resources_v1_metrics_node_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mmesh_protobuf_resources_v1_metrics_node_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HostMetrics); i {
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
		file_mmesh_protobuf_resources_v1_metrics_node_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkMetrics); i {
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
		file_mmesh_protobuf_resources_v1_metrics_node_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelayMetrics); i {
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
			RawDescriptor: file_mmesh_protobuf_resources_v1_metrics_node_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mmesh_protobuf_resources_v1_metrics_node_proto_goTypes,
		DependencyIndexes: file_mmesh_protobuf_resources_v1_metrics_node_proto_depIdxs,
		MessageInfos:      file_mmesh_protobuf_resources_v1_metrics_node_proto_msgTypes,
	}.Build()
	File_mmesh_protobuf_resources_v1_metrics_node_proto = out.File
	file_mmesh_protobuf_resources_v1_metrics_node_proto_rawDesc = nil
	file_mmesh_protobuf_resources_v1_metrics_node_proto_goTypes = nil
	file_mmesh_protobuf_resources_v1_metrics_node_proto_depIdxs = nil
}
