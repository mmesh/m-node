// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.0--rc3
// source: mmesh/protobuf/resources/v1/events/alert.proto

package events

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	resource "mmesh.dev/m-api-go/grpc/resources/resource"
	tenant "mmesh.dev/m-api-go/grpc/resources/tenant"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AlertListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountID string `protobuf:"bytes,1,opt,name=accountID,proto3" json:"accountID,omitempty"`
	TenantID  string `protobuf:"bytes,2,opt,name=tenantID,proto3" json:"tenantID,omitempty"`
}

func (x *AlertListRequest) Reset() {
	*x = AlertListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_events_alert_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlertListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlertListRequest) ProtoMessage() {}

func (x *AlertListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_events_alert_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlertListRequest.ProtoReflect.Descriptor instead.
func (*AlertListRequest) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_events_alert_proto_rawDescGZIP(), []int{0}
}

func (x *AlertListRequest) GetAccountID() string {
	if x != nil {
		return x.AccountID
	}
	return ""
}

func (x *AlertListRequest) GetTenantID() string {
	if x != nil {
		return x.TenantID
	}
	return ""
}

type AlertList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Alerts []*Alert `protobuf:"bytes,1,rep,name=alerts,proto3" json:"alerts,omitempty"`
}

func (x *AlertList) Reset() {
	*x = AlertList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_events_alert_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlertList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlertList) ProtoMessage() {}

func (x *AlertList) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_events_alert_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlertList.ProtoReflect.Descriptor instead.
func (*AlertList) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_events_alert_proto_rawDescGZIP(), []int{1}
}

func (x *AlertList) GetAlerts() []*Alert {
	if x != nil {
		return x.Alerts
	}
	return nil
}

type Alert struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountID string `protobuf:"bytes,1,opt,name=accountID,proto3" json:"accountID,omitempty"`
	TenantID  string `protobuf:"bytes,2,opt,name=tenantID,proto3" json:"tenantID,omitempty"`
	AlertID   string `protobuf:"bytes,3,opt,name=alertID,proto3" json:"alertID,omitempty"` // == eventID
	Timestamp int64  `protobuf:"varint,11,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// string netID = 22;
	// string subnetID = 23;
	NodeID        string            `protobuf:"bytes,24,opt,name=nodeID,proto3" json:"nodeID,omitempty"`
	NodeName      string            `protobuf:"bytes,25,opt,name=nodeName,proto3" json:"nodeName,omitempty"`
	Class         Class             `protobuf:"varint,31,opt,name=class,proto3,enum=events.Class" json:"class,omitempty"`
	Group         Group             `protobuf:"varint,41,opt,name=group,proto3,enum=events.Group" json:"group,omitempty"`
	Component     string            `protobuf:"bytes,51,opt,name=component,proto3" json:"component,omitempty"`
	Severity      Severity          `protobuf:"varint,61,opt,name=severity,proto3,enum=events.Severity" json:"severity,omitempty"`
	Status        Status            `protobuf:"varint,71,opt,name=status,proto3,enum=events.Status" json:"status,omitempty"`
	Summary       string            `protobuf:"bytes,101,opt,name=summary,proto3" json:"summary,omitempty"`
	CustomDetails map[string]string `protobuf:"bytes,111,rep,name=customDetails,proto3" json:"customDetails,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// map<int64, AlertComment> comments = 201;  // map[timestamp]*AlertComment
	Comments    []*AlertComment `protobuf:"bytes,202,rep,name=comments,proto3" json:"comments,omitempty"`
	LastUpdated int64           `protobuf:"varint,1001,opt,name=lastUpdated,proto3" json:"lastUpdated,omitempty"`
}

func (x *Alert) Reset() {
	*x = Alert{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_events_alert_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Alert) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Alert) ProtoMessage() {}

func (x *Alert) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_events_alert_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Alert.ProtoReflect.Descriptor instead.
func (*Alert) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_events_alert_proto_rawDescGZIP(), []int{2}
}

func (x *Alert) GetAccountID() string {
	if x != nil {
		return x.AccountID
	}
	return ""
}

func (x *Alert) GetTenantID() string {
	if x != nil {
		return x.TenantID
	}
	return ""
}

func (x *Alert) GetAlertID() string {
	if x != nil {
		return x.AlertID
	}
	return ""
}

func (x *Alert) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Alert) GetNodeID() string {
	if x != nil {
		return x.NodeID
	}
	return ""
}

func (x *Alert) GetNodeName() string {
	if x != nil {
		return x.NodeName
	}
	return ""
}

func (x *Alert) GetClass() Class {
	if x != nil {
		return x.Class
	}
	return Class_UNDEFINED_CLASS
}

func (x *Alert) GetGroup() Group {
	if x != nil {
		return x.Group
	}
	return Group_UNDEFINED_GROUP
}

func (x *Alert) GetComponent() string {
	if x != nil {
		return x.Component
	}
	return ""
}

func (x *Alert) GetSeverity() Severity {
	if x != nil {
		return x.Severity
	}
	return Severity_UNDEFINED_SEVERITY
}

func (x *Alert) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_UNDEFINED_STATUS
}

func (x *Alert) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

func (x *Alert) GetCustomDetails() map[string]string {
	if x != nil {
		return x.CustomDetails
	}
	return nil
}

func (x *Alert) GetComments() []*AlertComment {
	if x != nil {
		return x.Comments
	}
	return nil
}

func (x *Alert) GetLastUpdated() int64 {
	if x != nil {
		return x.LastUpdated
	}
	return 0
}

type Alerts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta   *resource.ListResponse `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	Alerts []*Alert               `protobuf:"bytes,2,rep,name=alerts,proto3" json:"alerts,omitempty"`
}

func (x *Alerts) Reset() {
	*x = Alerts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_events_alert_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Alerts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Alerts) ProtoMessage() {}

func (x *Alerts) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_events_alert_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Alerts.ProtoReflect.Descriptor instead.
func (*Alerts) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_events_alert_proto_rawDescGZIP(), []int{3}
}

func (x *Alerts) GetMeta() *resource.ListResponse {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *Alerts) GetAlerts() []*Alert {
	if x != nil {
		return x.Alerts
	}
	return nil
}

type ListAlertsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta   *resource.ListRequest `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	Tenant *tenant.TenantReq     `protobuf:"bytes,2,opt,name=tenant,proto3" json:"tenant,omitempty"`
}

func (x *ListAlertsRequest) Reset() {
	*x = ListAlertsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_events_alert_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAlertsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAlertsRequest) ProtoMessage() {}

func (x *ListAlertsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_events_alert_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAlertsRequest.ProtoReflect.Descriptor instead.
func (*ListAlertsRequest) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_events_alert_proto_rawDescGZIP(), []int{4}
}

func (x *ListAlertsRequest) GetMeta() *resource.ListRequest {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *ListAlertsRequest) GetTenant() *tenant.TenantReq {
	if x != nil {
		return x.Tenant
	}
	return nil
}

type AlertComment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp int64  `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	UserID    string `protobuf:"bytes,11,opt,name=userID,proto3" json:"userID,omitempty"`
	UserEmail string `protobuf:"bytes,12,opt,name=userEmail,proto3" json:"userEmail,omitempty"`
	Text      string `protobuf:"bytes,21,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *AlertComment) Reset() {
	*x = AlertComment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_events_alert_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlertComment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlertComment) ProtoMessage() {}

func (x *AlertComment) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_events_alert_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlertComment.ProtoReflect.Descriptor instead.
func (*AlertComment) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_events_alert_proto_rawDescGZIP(), []int{5}
}

func (x *AlertComment) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *AlertComment) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *AlertComment) GetUserEmail() string {
	if x != nil {
		return x.UserEmail
	}
	return ""
}

func (x *AlertComment) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type AlertReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountID string `protobuf:"bytes,1,opt,name=accountID,proto3" json:"accountID,omitempty"`
	TenantID  string `protobuf:"bytes,2,opt,name=tenantID,proto3" json:"tenantID,omitempty"`
	AlertID   string `protobuf:"bytes,3,opt,name=alertID,proto3" json:"alertID,omitempty"`
	// string netID = 12;
	// string SubnetID = 13;
	NodeID string `protobuf:"bytes,14,opt,name=nodeID,proto3" json:"nodeID,omitempty"`
}

func (x *AlertReq) Reset() {
	*x = AlertReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_events_alert_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlertReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlertReq) ProtoMessage() {}

func (x *AlertReq) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_events_alert_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlertReq.ProtoReflect.Descriptor instead.
func (*AlertReq) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_events_alert_proto_rawDescGZIP(), []int{6}
}

func (x *AlertReq) GetAccountID() string {
	if x != nil {
		return x.AccountID
	}
	return ""
}

func (x *AlertReq) GetTenantID() string {
	if x != nil {
		return x.TenantID
	}
	return ""
}

func (x *AlertReq) GetAlertID() string {
	if x != nil {
		return x.AlertID
	}
	return ""
}

func (x *AlertReq) GetNodeID() string {
	if x != nil {
		return x.NodeID
	}
	return ""
}

type AlertNewCommentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountID string        `protobuf:"bytes,1,opt,name=accountID,proto3" json:"accountID,omitempty"`
	TenantID  string        `protobuf:"bytes,2,opt,name=tenantID,proto3" json:"tenantID,omitempty"`
	AlertID   string        `protobuf:"bytes,3,opt,name=alertID,proto3" json:"alertID,omitempty"`
	Comment   *AlertComment `protobuf:"bytes,11,opt,name=comment,proto3" json:"comment,omitempty"`
}

func (x *AlertNewCommentRequest) Reset() {
	*x = AlertNewCommentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_events_alert_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlertNewCommentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlertNewCommentRequest) ProtoMessage() {}

func (x *AlertNewCommentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_events_alert_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlertNewCommentRequest.ProtoReflect.Descriptor instead.
func (*AlertNewCommentRequest) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_events_alert_proto_rawDescGZIP(), []int{7}
}

func (x *AlertNewCommentRequest) GetAccountID() string {
	if x != nil {
		return x.AccountID
	}
	return ""
}

func (x *AlertNewCommentRequest) GetTenantID() string {
	if x != nil {
		return x.TenantID
	}
	return ""
}

func (x *AlertNewCommentRequest) GetAlertID() string {
	if x != nil {
		return x.AlertID
	}
	return ""
}

func (x *AlertNewCommentRequest) GetComment() *AlertComment {
	if x != nil {
		return x.Comment
	}
	return nil
}

var File_mmesh_protobuf_resources_v1_events_alert_proto protoreflect.FileDescriptor

var file_mmesh_protobuf_resources_v1_events_alert_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0x2f, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x6c,
	0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2f, 0x6d, 0x6d, 0x65, 0x73, 0x68,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2f, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x6d, 0x6d, 0x65, 0x73,
	0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4c, 0x0a, 0x10, 0x41, 0x6c,
	0x65, 0x72, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x22, 0x32, 0x0a, 0x09, 0x41, 0x6c, 0x65, 0x72,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x06, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x41,
	0x6c, 0x65, 0x72, 0x74, 0x52, 0x06, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x22, 0xe5, 0x04, 0x0a,
	0x05, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44,
	0x12, 0x18, 0x0a, 0x07, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x64, 0x65,
	0x49, 0x44, 0x18, 0x18, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x44,
	0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x19, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x05,
	0x63, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x05, 0x63, 0x6c, 0x61, 0x73,
	0x73, 0x12, 0x23, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x29, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0d, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52,
	0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e,
	0x65, 0x6e, 0x74, 0x18, 0x33, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6f,
	0x6e, 0x65, 0x6e, 0x74, 0x12, 0x2c, 0x0a, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79,
	0x18, 0x3d, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e,
	0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x52, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69,
	0x74, 0x79, 0x12, 0x26, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x47, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x65, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x6d,
	0x6d, 0x61, 0x72, 0x79, 0x12, 0x46, 0x0a, 0x0d, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x6f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0d, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x31, 0x0a, 0x08,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0xca, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12,
	0x21, 0x0a, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0xe9,
	0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x1a, 0x40, 0x0a, 0x12, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0x5b, 0x0a, 0x06, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x12, 0x2a,
	0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x25, 0x0a, 0x06, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x06, 0x61, 0x6c, 0x65, 0x72, 0x74,
	0x73, 0x22, 0x69, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x6d, 0x65, 0x74,
	0x61, 0x12, 0x29, 0x0a, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x54, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x52, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x22, 0x76, 0x0a, 0x0c,
	0x41, 0x6c, 0x65, 0x72, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x65, 0x78, 0x74, 0x22, 0x76, 0x0a, 0x08, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71,
	0x12, 0x1c, 0x0a, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x1a,
	0x0a, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x44, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x44, 0x22, 0x9c, 0x01, 0x0a,
	0x16, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x4e, 0x65, 0x77, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49,
	0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49,
	0x44, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x49, 0x44, 0x12, 0x2e, 0x0a, 0x07, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x2a, 0x5a, 0x28, 0x6d,
	0x6d, 0x65, 0x73, 0x68, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x6d, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x67,
	0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mmesh_protobuf_resources_v1_events_alert_proto_rawDescOnce sync.Once
	file_mmesh_protobuf_resources_v1_events_alert_proto_rawDescData = file_mmesh_protobuf_resources_v1_events_alert_proto_rawDesc
)

func file_mmesh_protobuf_resources_v1_events_alert_proto_rawDescGZIP() []byte {
	file_mmesh_protobuf_resources_v1_events_alert_proto_rawDescOnce.Do(func() {
		file_mmesh_protobuf_resources_v1_events_alert_proto_rawDescData = protoimpl.X.CompressGZIP(file_mmesh_protobuf_resources_v1_events_alert_proto_rawDescData)
	})
	return file_mmesh_protobuf_resources_v1_events_alert_proto_rawDescData
}

var file_mmesh_protobuf_resources_v1_events_alert_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_mmesh_protobuf_resources_v1_events_alert_proto_goTypes = []interface{}{
	(*AlertListRequest)(nil),       // 0: events.AlertListRequest
	(*AlertList)(nil),              // 1: events.AlertList
	(*Alert)(nil),                  // 2: events.Alert
	(*Alerts)(nil),                 // 3: events.Alerts
	(*ListAlertsRequest)(nil),      // 4: events.ListAlertsRequest
	(*AlertComment)(nil),           // 5: events.AlertComment
	(*AlertReq)(nil),               // 6: events.AlertReq
	(*AlertNewCommentRequest)(nil), // 7: events.AlertNewCommentRequest
	nil,                            // 8: events.Alert.CustomDetailsEntry
	(Class)(0),                     // 9: events.Class
	(Group)(0),                     // 10: events.Group
	(Severity)(0),                  // 11: events.Severity
	(Status)(0),                    // 12: events.Status
	(*resource.ListResponse)(nil),  // 13: resource.ListResponse
	(*resource.ListRequest)(nil),   // 14: resource.ListRequest
	(*tenant.TenantReq)(nil),       // 15: tenant.TenantReq
}
var file_mmesh_protobuf_resources_v1_events_alert_proto_depIdxs = []int32{
	2,  // 0: events.AlertList.alerts:type_name -> events.Alert
	9,  // 1: events.Alert.class:type_name -> events.Class
	10, // 2: events.Alert.group:type_name -> events.Group
	11, // 3: events.Alert.severity:type_name -> events.Severity
	12, // 4: events.Alert.status:type_name -> events.Status
	8,  // 5: events.Alert.customDetails:type_name -> events.Alert.CustomDetailsEntry
	5,  // 6: events.Alert.comments:type_name -> events.AlertComment
	13, // 7: events.Alerts.meta:type_name -> resource.ListResponse
	2,  // 8: events.Alerts.alerts:type_name -> events.Alert
	14, // 9: events.ListAlertsRequest.meta:type_name -> resource.ListRequest
	15, // 10: events.ListAlertsRequest.tenant:type_name -> tenant.TenantReq
	5,  // 11: events.AlertNewCommentRequest.comment:type_name -> events.AlertComment
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_mmesh_protobuf_resources_v1_events_alert_proto_init() }
func file_mmesh_protobuf_resources_v1_events_alert_proto_init() {
	if File_mmesh_protobuf_resources_v1_events_alert_proto != nil {
		return
	}
	file_mmesh_protobuf_resources_v1_events_event_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_mmesh_protobuf_resources_v1_events_alert_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlertListRequest); i {
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
		file_mmesh_protobuf_resources_v1_events_alert_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlertList); i {
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
		file_mmesh_protobuf_resources_v1_events_alert_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Alert); i {
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
		file_mmesh_protobuf_resources_v1_events_alert_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Alerts); i {
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
		file_mmesh_protobuf_resources_v1_events_alert_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAlertsRequest); i {
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
		file_mmesh_protobuf_resources_v1_events_alert_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlertComment); i {
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
		file_mmesh_protobuf_resources_v1_events_alert_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlertReq); i {
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
		file_mmesh_protobuf_resources_v1_events_alert_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlertNewCommentRequest); i {
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
			RawDescriptor: file_mmesh_protobuf_resources_v1_events_alert_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mmesh_protobuf_resources_v1_events_alert_proto_goTypes,
		DependencyIndexes: file_mmesh_protobuf_resources_v1_events_alert_proto_depIdxs,
		MessageInfos:      file_mmesh_protobuf_resources_v1_events_alert_proto_msgTypes,
	}.Build()
	File_mmesh_protobuf_resources_v1_events_alert_proto = out.File
	file_mmesh_protobuf_resources_v1_events_alert_proto_rawDesc = nil
	file_mmesh_protobuf_resources_v1_events_alert_proto_goTypes = nil
	file_mmesh_protobuf_resources_v1_events_alert_proto_depIdxs = nil
}
