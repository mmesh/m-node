// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0--rc2
// source: mmesh/protobuf/resources/v1/tss/tss.proto

package tss

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

type TimeRange int32

const (
	TimeRange_TTL_UNDEFINED TimeRange = 0
	TimeRange_TTL_1H        TimeRange = 11 // 1H - 60min       / 30 samples - 1 point per 2min
	TimeRange_TTL_6H        TimeRange = 21 // 6H - 360min      / 30 samples - 1 point per 12min
	TimeRange_TTL_12H       TimeRange = 31 // 12H - 720min     / 30 samples - 1 point per 24min
	TimeRange_TTL_24H       TimeRange = 41 // 24H - 1440min    / 30 samples - 1 point per 48min
	TimeRange_TTL_7D        TimeRange = 51 // 7D - 10080min    / 30 samples - 1 point per 336min
	TimeRange_TTL_14D       TimeRange = 61 // 14D - 20160min   / 30 samples - 1 point per 672min
	TimeRange_TTL_30D       TimeRange = 71 // 30D - 40320min   / 30 samples - 1 point per 1344min
	TimeRange_TTL_365D      TimeRange = 81 // 365D - 525600min / 30 samples - 1 point per 12days
)

// Enum value maps for TimeRange.
var (
	TimeRange_name = map[int32]string{
		0:  "TTL_UNDEFINED",
		11: "TTL_1H",
		21: "TTL_6H",
		31: "TTL_12H",
		41: "TTL_24H",
		51: "TTL_7D",
		61: "TTL_14D",
		71: "TTL_30D",
		81: "TTL_365D",
	}
	TimeRange_value = map[string]int32{
		"TTL_UNDEFINED": 0,
		"TTL_1H":        11,
		"TTL_6H":        21,
		"TTL_12H":       31,
		"TTL_24H":       41,
		"TTL_7D":        51,
		"TTL_14D":       61,
		"TTL_30D":       71,
		"TTL_365D":      81,
	}
)

func (x TimeRange) Enum() *TimeRange {
	p := new(TimeRange)
	*p = x
	return p
}

func (x TimeRange) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TimeRange) Descriptor() protoreflect.EnumDescriptor {
	return file_mmesh_protobuf_resources_v1_tss_tss_proto_enumTypes[0].Descriptor()
}

func (TimeRange) Type() protoreflect.EnumType {
	return &file_mmesh_protobuf_resources_v1_tss_tss_proto_enumTypes[0]
}

func (x TimeRange) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TimeRange.Descriptor instead.
func (TimeRange) EnumDescriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_tss_tss_proto_rawDescGZIP(), []int{0}
}

type MetricType int32

const (
	MetricType_UNDEFINED       MetricType = 0
	MetricType_NET_RX_BYTES    MetricType = 11
	MetricType_NET_TX_BYTES    MetricType = 12
	MetricType_HOST_LOAD_AVG   MetricType = 111
	MetricType_HOST_CPU_USAGE  MetricType = 121
	MetricType_HOST_MEM_USAGE  MetricType = 131
	MetricType_HOST_DISK_USAGE MetricType = 141
)

// Enum value maps for MetricType.
var (
	MetricType_name = map[int32]string{
		0:   "UNDEFINED",
		11:  "NET_RX_BYTES",
		12:  "NET_TX_BYTES",
		111: "HOST_LOAD_AVG",
		121: "HOST_CPU_USAGE",
		131: "HOST_MEM_USAGE",
		141: "HOST_DISK_USAGE",
	}
	MetricType_value = map[string]int32{
		"UNDEFINED":       0,
		"NET_RX_BYTES":    11,
		"NET_TX_BYTES":    12,
		"HOST_LOAD_AVG":   111,
		"HOST_CPU_USAGE":  121,
		"HOST_MEM_USAGE":  131,
		"HOST_DISK_USAGE": 141,
	}
)

func (x MetricType) Enum() *MetricType {
	p := new(MetricType)
	*p = x
	return p
}

func (x MetricType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MetricType) Descriptor() protoreflect.EnumDescriptor {
	return file_mmesh_protobuf_resources_v1_tss_tss_proto_enumTypes[1].Descriptor()
}

func (MetricType) Type() protoreflect.EnumType {
	return &file_mmesh_protobuf_resources_v1_tss_tss_proto_enumTypes[1]
}

func (x MetricType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MetricType.Descriptor instead.
func (MetricType) EnumDescriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_tss_tss_proto_rawDescGZIP(), []int{1}
}

type DataPoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp int64      `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	TimeRange TimeRange  `protobuf:"varint,11,opt,name=timeRange,proto3,enum=tss.TimeRange" json:"timeRange,omitempty"` // namespace
	Metric    MetricType `protobuf:"varint,21,opt,name=metric,proto3,enum=tss.MetricType" json:"metric,omitempty"`      // key
	Value     float64    `protobuf:"fixed64,101,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *DataPoint) Reset() {
	*x = DataPoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_tss_tss_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataPoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataPoint) ProtoMessage() {}

func (x *DataPoint) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_tss_tss_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataPoint.ProtoReflect.Descriptor instead.
func (*DataPoint) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_tss_tss_proto_rawDescGZIP(), []int{0}
}

func (x *DataPoint) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *DataPoint) GetTimeRange() TimeRange {
	if x != nil {
		return x.TimeRange
	}
	return TimeRange_TTL_UNDEFINED
}

func (x *DataPoint) GetMetric() MetricType {
	if x != nil {
		return x.Metric
	}
	return MetricType_UNDEFINED
}

func (x *DataPoint) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type MetricsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountID    string `protobuf:"bytes,1,opt,name=accountID,proto3" json:"accountID,omitempty"`
	TenantID     string `protobuf:"bytes,2,opt,name=tenantID,proto3" json:"tenantID,omitempty"`
	NetID        string `protobuf:"bytes,3,opt,name=netID,proto3" json:"netID,omitempty"`
	SubnetID     string `protobuf:"bytes,4,opt,name=subnetID,proto3" json:"subnetID,omitempty"`
	NodeID       string `protobuf:"bytes,5,opt,name=nodeID,proto3" json:"nodeID,omitempty"`
	ControllerID string `protobuf:"bytes,11,opt,name=controllerID,proto3" json:"controllerID,omitempty"`
	QueryID      string `protobuf:"bytes,21,opt,name=queryID,proto3" json:"queryID,omitempty"`
}

func (x *MetricsRequest) Reset() {
	*x = MetricsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_tss_tss_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricsRequest) ProtoMessage() {}

func (x *MetricsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_tss_tss_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricsRequest.ProtoReflect.Descriptor instead.
func (*MetricsRequest) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_tss_tss_proto_rawDescGZIP(), []int{1}
}

func (x *MetricsRequest) GetAccountID() string {
	if x != nil {
		return x.AccountID
	}
	return ""
}

func (x *MetricsRequest) GetTenantID() string {
	if x != nil {
		return x.TenantID
	}
	return ""
}

func (x *MetricsRequest) GetNetID() string {
	if x != nil {
		return x.NetID
	}
	return ""
}

func (x *MetricsRequest) GetSubnetID() string {
	if x != nil {
		return x.SubnetID
	}
	return ""
}

func (x *MetricsRequest) GetNodeID() string {
	if x != nil {
		return x.NodeID
	}
	return ""
}

func (x *MetricsRequest) GetControllerID() string {
	if x != nil {
		return x.ControllerID
	}
	return ""
}

func (x *MetricsRequest) GetQueryID() string {
	if x != nil {
		return x.QueryID
	}
	return ""
}

type NodeMetrics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountID string                  `protobuf:"bytes,1,opt,name=accountID,proto3" json:"accountID,omitempty"`
	TenantID  string                  `protobuf:"bytes,2,opt,name=tenantID,proto3" json:"tenantID,omitempty"`
	NetID     string                  `protobuf:"bytes,3,opt,name=netID,proto3" json:"netID,omitempty"`
	SubnetID  string                  `protobuf:"bytes,4,opt,name=subnetID,proto3" json:"subnetID,omitempty"`
	NodeID    string                  `protobuf:"bytes,5,opt,name=nodeID,proto3" json:"nodeID,omitempty"`
	QueryID   string                  `protobuf:"bytes,11,opt,name=queryID,proto3" json:"queryID,omitempty"`
	Metrics   map[string]*HostMetrics `protobuf:"bytes,101,rep,name=metrics,proto3" json:"metrics,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // map[TimeRange.String()]*HostMetrics
	Timestamp int64                   `protobuf:"varint,1001,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *NodeMetrics) Reset() {
	*x = NodeMetrics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_tss_tss_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeMetrics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeMetrics) ProtoMessage() {}

func (x *NodeMetrics) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_tss_tss_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeMetrics.ProtoReflect.Descriptor instead.
func (*NodeMetrics) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_tss_tss_proto_rawDescGZIP(), []int{2}
}

func (x *NodeMetrics) GetAccountID() string {
	if x != nil {
		return x.AccountID
	}
	return ""
}

func (x *NodeMetrics) GetTenantID() string {
	if x != nil {
		return x.TenantID
	}
	return ""
}

func (x *NodeMetrics) GetNetID() string {
	if x != nil {
		return x.NetID
	}
	return ""
}

func (x *NodeMetrics) GetSubnetID() string {
	if x != nil {
		return x.SubnetID
	}
	return ""
}

func (x *NodeMetrics) GetNodeID() string {
	if x != nil {
		return x.NodeID
	}
	return ""
}

func (x *NodeMetrics) GetQueryID() string {
	if x != nil {
		return x.QueryID
	}
	return ""
}

func (x *NodeMetrics) GetMetrics() map[string]*HostMetrics {
	if x != nil {
		return x.Metrics
	}
	return nil
}

func (x *NodeMetrics) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type HostMetrics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LoadAvg   []*TimeValue `protobuf:"bytes,11,rep,name=loadAvg,proto3" json:"loadAvg,omitempty"`
	CpuUsage  []*TimeValue `protobuf:"bytes,21,rep,name=cpuUsage,proto3" json:"cpuUsage,omitempty"`
	MemUsage  []*TimeValue `protobuf:"bytes,31,rep,name=memUsage,proto3" json:"memUsage,omitempty"`
	DiskUsage []*TimeValue `protobuf:"bytes,41,rep,name=diskUsage,proto3" json:"diskUsage,omitempty"`
	NetUsage  *NetMetrics  `protobuf:"bytes,101,opt,name=netUsage,proto3" json:"netUsage,omitempty"`
}

func (x *HostMetrics) Reset() {
	*x = HostMetrics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_tss_tss_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HostMetrics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HostMetrics) ProtoMessage() {}

func (x *HostMetrics) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_tss_tss_proto_msgTypes[3]
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
	return file_mmesh_protobuf_resources_v1_tss_tss_proto_rawDescGZIP(), []int{3}
}

func (x *HostMetrics) GetLoadAvg() []*TimeValue {
	if x != nil {
		return x.LoadAvg
	}
	return nil
}

func (x *HostMetrics) GetCpuUsage() []*TimeValue {
	if x != nil {
		return x.CpuUsage
	}
	return nil
}

func (x *HostMetrics) GetMemUsage() []*TimeValue {
	if x != nil {
		return x.MemUsage
	}
	return nil
}

func (x *HostMetrics) GetDiskUsage() []*TimeValue {
	if x != nil {
		return x.DiskUsage
	}
	return nil
}

func (x *HostMetrics) GetNetUsage() *NetMetrics {
	if x != nil {
		return x.NetUsage
	}
	return nil
}

type NetMetrics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NetRxBytes []*TimeValue `protobuf:"bytes,11,rep,name=netRxBytes,proto3" json:"netRxBytes,omitempty"`
	NetTxBytes []*TimeValue `protobuf:"bytes,21,rep,name=netTxBytes,proto3" json:"netTxBytes,omitempty"`
}

func (x *NetMetrics) Reset() {
	*x = NetMetrics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_tss_tss_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetMetrics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetMetrics) ProtoMessage() {}

func (x *NetMetrics) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_tss_tss_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetMetrics.ProtoReflect.Descriptor instead.
func (*NetMetrics) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_tss_tss_proto_rawDescGZIP(), []int{4}
}

func (x *NetMetrics) GetNetRxBytes() []*TimeValue {
	if x != nil {
		return x.NetRxBytes
	}
	return nil
}

func (x *NetMetrics) GetNetTxBytes() []*TimeValue {
	if x != nil {
		return x.NetTxBytes
	}
	return nil
}

type TimeValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp int64   `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Value     float64 `protobuf:"fixed64,101,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *TimeValue) Reset() {
	*x = TimeValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_tss_tss_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimeValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeValue) ProtoMessage() {}

func (x *TimeValue) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_tss_tss_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeValue.ProtoReflect.Descriptor instead.
func (*TimeValue) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_tss_tss_proto_rawDescGZIP(), []int{5}
}

func (x *TimeValue) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *TimeValue) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

var File_mmesh_protobuf_resources_v1_tss_tss_proto protoreflect.FileDescriptor

var file_mmesh_protobuf_resources_v1_tss_tss_proto_rawDesc = []byte{
	0x0a, 0x29, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x73,
	0x73, 0x2f, 0x74, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x74, 0x73, 0x73,
	0x22, 0x96, 0x01, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x61, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x2c, 0x0a, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x0e, 0x2e, 0x74, 0x73, 0x73, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x27, 0x0a, 0x06, 0x6d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x74, 0x73, 0x73,
	0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x52, 0x06, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x65, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xd2, 0x01, 0x0a, 0x0e, 0x4d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x65, 0x74, 0x49, 0x44, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x65, 0x74, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08,
	0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x64, 0x65,
	0x49, 0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x44,
	0x12, 0x22, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x49, 0x44,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c,
	0x65, 0x72, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x71, 0x75, 0x65, 0x72, 0x79, 0x49, 0x44, 0x18,
	0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x71, 0x75, 0x65, 0x72, 0x79, 0x49, 0x44, 0x22, 0xd1,
	0x02, 0x0a, 0x0b, 0x4e, 0x6f, 0x64, 0x65, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x1c,
	0x0a, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x65, 0x74, 0x49,
	0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x65, 0x74, 0x49, 0x44, 0x12, 0x1a,
	0x0a, 0x08, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f,
	0x64, 0x65, 0x49, 0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65,
	0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x71, 0x75, 0x65, 0x72, 0x79, 0x49, 0x44, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x71, 0x75, 0x65, 0x72, 0x79, 0x49, 0x44, 0x12, 0x37, 0x0a, 0x07,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18, 0x65, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x74, 0x73, 0x73, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e,
	0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x6d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x1d, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x18, 0xe9, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x1a, 0x4c, 0x0a, 0x0c, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x26, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x74, 0x73, 0x73, 0x2e, 0x48, 0x6f, 0x73, 0x74,
	0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0xea, 0x01, 0x0a, 0x0b, 0x48, 0x6f, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x73, 0x12, 0x28, 0x0a, 0x07, 0x6c, 0x6f, 0x61, 0x64, 0x41, 0x76, 0x67, 0x18, 0x0b, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x74, 0x73, 0x73, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x07, 0x6c, 0x6f, 0x61, 0x64, 0x41, 0x76, 0x67, 0x12, 0x2a, 0x0a, 0x08,
	0x63, 0x70, 0x75, 0x55, 0x73, 0x61, 0x67, 0x65, 0x18, 0x15, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x74, 0x73, 0x73, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08,
	0x63, 0x70, 0x75, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2a, 0x0a, 0x08, 0x6d, 0x65, 0x6d, 0x55,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x1f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x74, 0x73, 0x73,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x55,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x2c, 0x0a, 0x09, 0x64, 0x69, 0x73, 0x6b, 0x55, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x29, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x74, 0x73, 0x73, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x09, 0x64, 0x69, 0x73, 0x6b, 0x55, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x2b, 0x0a, 0x08, 0x6e, 0x65, 0x74, 0x55, 0x73, 0x61, 0x67, 0x65, 0x18, 0x65,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x74, 0x73, 0x73, 0x2e, 0x4e, 0x65, 0x74, 0x4d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x08, 0x6e, 0x65, 0x74, 0x55, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x6c, 0x0a, 0x0a, 0x4e, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x2e, 0x0a,
	0x0a, 0x6e, 0x65, 0x74, 0x52, 0x78, 0x42, 0x79, 0x74, 0x65, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x74, 0x73, 0x73, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x0a, 0x6e, 0x65, 0x74, 0x52, 0x78, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x2e, 0x0a,
	0x0a, 0x6e, 0x65, 0x74, 0x54, 0x78, 0x42, 0x79, 0x74, 0x65, 0x73, 0x18, 0x15, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x74, 0x73, 0x73, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x0a, 0x6e, 0x65, 0x74, 0x54, 0x78, 0x42, 0x79, 0x74, 0x65, 0x73, 0x22, 0x3f, 0x0a,
	0x09, 0x54, 0x69, 0x6d, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x65, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2a, 0x84,
	0x01, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x11, 0x0a, 0x0d,
	0x54, 0x54, 0x4c, 0x5f, 0x55, 0x4e, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x0a, 0x0a, 0x06, 0x54, 0x54, 0x4c, 0x5f, 0x31, 0x48, 0x10, 0x0b, 0x12, 0x0a, 0x0a, 0x06, 0x54,
	0x54, 0x4c, 0x5f, 0x36, 0x48, 0x10, 0x15, 0x12, 0x0b, 0x0a, 0x07, 0x54, 0x54, 0x4c, 0x5f, 0x31,
	0x32, 0x48, 0x10, 0x1f, 0x12, 0x0b, 0x0a, 0x07, 0x54, 0x54, 0x4c, 0x5f, 0x32, 0x34, 0x48, 0x10,
	0x29, 0x12, 0x0a, 0x0a, 0x06, 0x54, 0x54, 0x4c, 0x5f, 0x37, 0x44, 0x10, 0x33, 0x12, 0x0b, 0x0a,
	0x07, 0x54, 0x54, 0x4c, 0x5f, 0x31, 0x34, 0x44, 0x10, 0x3d, 0x12, 0x0b, 0x0a, 0x07, 0x54, 0x54,
	0x4c, 0x5f, 0x33, 0x30, 0x44, 0x10, 0x47, 0x12, 0x0c, 0x0a, 0x08, 0x54, 0x54, 0x4c, 0x5f, 0x33,
	0x36, 0x35, 0x44, 0x10, 0x51, 0x2a, 0x91, 0x01, 0x0a, 0x0a, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x0d, 0x0a, 0x09, 0x55, 0x4e, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x4e, 0x45, 0x54, 0x5f, 0x52, 0x58, 0x5f, 0x42, 0x59,
	0x54, 0x45, 0x53, 0x10, 0x0b, 0x12, 0x10, 0x0a, 0x0c, 0x4e, 0x45, 0x54, 0x5f, 0x54, 0x58, 0x5f,
	0x42, 0x59, 0x54, 0x45, 0x53, 0x10, 0x0c, 0x12, 0x11, 0x0a, 0x0d, 0x48, 0x4f, 0x53, 0x54, 0x5f,
	0x4c, 0x4f, 0x41, 0x44, 0x5f, 0x41, 0x56, 0x47, 0x10, 0x6f, 0x12, 0x12, 0x0a, 0x0e, 0x48, 0x4f,
	0x53, 0x54, 0x5f, 0x43, 0x50, 0x55, 0x5f, 0x55, 0x53, 0x41, 0x47, 0x45, 0x10, 0x79, 0x12, 0x13,
	0x0a, 0x0e, 0x48, 0x4f, 0x53, 0x54, 0x5f, 0x4d, 0x45, 0x4d, 0x5f, 0x55, 0x53, 0x41, 0x47, 0x45,
	0x10, 0x83, 0x01, 0x12, 0x14, 0x0a, 0x0f, 0x48, 0x4f, 0x53, 0x54, 0x5f, 0x44, 0x49, 0x53, 0x4b,
	0x5f, 0x55, 0x53, 0x41, 0x47, 0x45, 0x10, 0x8d, 0x01, 0x42, 0x27, 0x5a, 0x25, 0x6d, 0x6d, 0x65,
	0x73, 0x68, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x6d, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x74,
	0x73, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mmesh_protobuf_resources_v1_tss_tss_proto_rawDescOnce sync.Once
	file_mmesh_protobuf_resources_v1_tss_tss_proto_rawDescData = file_mmesh_protobuf_resources_v1_tss_tss_proto_rawDesc
)

func file_mmesh_protobuf_resources_v1_tss_tss_proto_rawDescGZIP() []byte {
	file_mmesh_protobuf_resources_v1_tss_tss_proto_rawDescOnce.Do(func() {
		file_mmesh_protobuf_resources_v1_tss_tss_proto_rawDescData = protoimpl.X.CompressGZIP(file_mmesh_protobuf_resources_v1_tss_tss_proto_rawDescData)
	})
	return file_mmesh_protobuf_resources_v1_tss_tss_proto_rawDescData
}

var file_mmesh_protobuf_resources_v1_tss_tss_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_mmesh_protobuf_resources_v1_tss_tss_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_mmesh_protobuf_resources_v1_tss_tss_proto_goTypes = []interface{}{
	(TimeRange)(0),         // 0: tss.TimeRange
	(MetricType)(0),        // 1: tss.MetricType
	(*DataPoint)(nil),      // 2: tss.DataPoint
	(*MetricsRequest)(nil), // 3: tss.MetricsRequest
	(*NodeMetrics)(nil),    // 4: tss.NodeMetrics
	(*HostMetrics)(nil),    // 5: tss.HostMetrics
	(*NetMetrics)(nil),     // 6: tss.NetMetrics
	(*TimeValue)(nil),      // 7: tss.TimeValue
	nil,                    // 8: tss.NodeMetrics.MetricsEntry
}
var file_mmesh_protobuf_resources_v1_tss_tss_proto_depIdxs = []int32{
	0,  // 0: tss.DataPoint.timeRange:type_name -> tss.TimeRange
	1,  // 1: tss.DataPoint.metric:type_name -> tss.MetricType
	8,  // 2: tss.NodeMetrics.metrics:type_name -> tss.NodeMetrics.MetricsEntry
	7,  // 3: tss.HostMetrics.loadAvg:type_name -> tss.TimeValue
	7,  // 4: tss.HostMetrics.cpuUsage:type_name -> tss.TimeValue
	7,  // 5: tss.HostMetrics.memUsage:type_name -> tss.TimeValue
	7,  // 6: tss.HostMetrics.diskUsage:type_name -> tss.TimeValue
	6,  // 7: tss.HostMetrics.netUsage:type_name -> tss.NetMetrics
	7,  // 8: tss.NetMetrics.netRxBytes:type_name -> tss.TimeValue
	7,  // 9: tss.NetMetrics.netTxBytes:type_name -> tss.TimeValue
	5,  // 10: tss.NodeMetrics.MetricsEntry.value:type_name -> tss.HostMetrics
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_mmesh_protobuf_resources_v1_tss_tss_proto_init() }
func file_mmesh_protobuf_resources_v1_tss_tss_proto_init() {
	if File_mmesh_protobuf_resources_v1_tss_tss_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mmesh_protobuf_resources_v1_tss_tss_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataPoint); i {
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
		file_mmesh_protobuf_resources_v1_tss_tss_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetricsRequest); i {
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
		file_mmesh_protobuf_resources_v1_tss_tss_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeMetrics); i {
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
		file_mmesh_protobuf_resources_v1_tss_tss_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_mmesh_protobuf_resources_v1_tss_tss_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetMetrics); i {
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
		file_mmesh_protobuf_resources_v1_tss_tss_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimeValue); i {
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
			RawDescriptor: file_mmesh_protobuf_resources_v1_tss_tss_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mmesh_protobuf_resources_v1_tss_tss_proto_goTypes,
		DependencyIndexes: file_mmesh_protobuf_resources_v1_tss_tss_proto_depIdxs,
		EnumInfos:         file_mmesh_protobuf_resources_v1_tss_tss_proto_enumTypes,
		MessageInfos:      file_mmesh_protobuf_resources_v1_tss_tss_proto_msgTypes,
	}.Build()
	File_mmesh_protobuf_resources_v1_tss_tss_proto = out.File
	file_mmesh_protobuf_resources_v1_tss_tss_proto_rawDesc = nil
	file_mmesh_protobuf_resources_v1_tss_tss_proto_goTypes = nil
	file_mmesh_protobuf_resources_v1_tss_tss_proto_depIdxs = nil
}
