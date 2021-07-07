// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.1
// source: mmesh/protobuf/resources/v1/iam/user.proto

package iam

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	metrics "mmesh.dev/m-api-go/grpc/resources/metrics"
	resource "mmesh.dev/m-api-go/grpc/resources/resource"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Realm          string                `protobuf:"bytes,1,opt,name=realm,proto3" json:"realm,omitempty"` // accountID
	UserID         string                `protobuf:"bytes,2,opt,name=userID,proto3" json:"userID,omitempty"`
	Username       string                `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	Name           string                `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	Email          string                `protobuf:"bytes,11,opt,name=email,proto3" json:"email,omitempty"`
	EmailVerified  bool                  `protobuf:"varint,12,opt,name=emailVerified,proto3" json:"emailVerified,omitempty"`
	NewEmail       string                `protobuf:"bytes,13,opt,name=newEmail,proto3" json:"newEmail,omitempty"`
	Credentials    *UserCredentials      `protobuf:"bytes,21,opt,name=credentials,proto3" json:"credentials,omitempty"`
	Status         *UserStatus           `protobuf:"bytes,31,opt,name=status,proto3" json:"status,omitempty"`
	SecurityGroups []string              `protobuf:"bytes,41,rep,name=securityGroups,proto3" json:"securityGroups,omitempty"`
	Roles          []string              `protobuf:"bytes,42,rep,name=roles,proto3" json:"roles,omitempty"`
	ACLs           []string              `protobuf:"bytes,43,rep,name=ACLs,proto3" json:"ACLs,omitempty"`
	Stats          *UserStats            `protobuf:"bytes,51,opt,name=stats,proto3" json:"stats,omitempty"`
	EventMetrics   *metrics.EventMetrics `protobuf:"bytes,61,opt,name=eventMetrics,proto3" json:"eventMetrics,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_iam_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_iam_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_iam_user_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetRealm() string {
	if x != nil {
		return x.Realm
	}
	return ""
}

func (x *User) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetEmailVerified() bool {
	if x != nil {
		return x.EmailVerified
	}
	return false
}

func (x *User) GetNewEmail() string {
	if x != nil {
		return x.NewEmail
	}
	return ""
}

func (x *User) GetCredentials() *UserCredentials {
	if x != nil {
		return x.Credentials
	}
	return nil
}

func (x *User) GetStatus() *UserStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *User) GetSecurityGroups() []string {
	if x != nil {
		return x.SecurityGroups
	}
	return nil
}

func (x *User) GetRoles() []string {
	if x != nil {
		return x.Roles
	}
	return nil
}

func (x *User) GetACLs() []string {
	if x != nil {
		return x.ACLs
	}
	return nil
}

func (x *User) GetStats() *UserStats {
	if x != nil {
		return x.Stats
	}
	return nil
}

func (x *User) GetEventMetrics() *metrics.EventMetrics {
	if x != nil {
		return x.EventMetrics
	}
	return nil
}

type Users struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta  *resource.ListResponse `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	Users []*User                `protobuf:"bytes,2,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *Users) Reset() {
	*x = Users{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_iam_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Users) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Users) ProtoMessage() {}

func (x *Users) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_iam_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Users.ProtoReflect.Descriptor instead.
func (*Users) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_iam_user_proto_rawDescGZIP(), []int{1}
}

func (x *Users) GetMeta() *resource.ListResponse {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *Users) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

type ListUsersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta      *resource.ListRequest `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	AccountID string                `protobuf:"bytes,2,opt,name=accountID,proto3" json:"accountID,omitempty"`
}

func (x *ListUsersRequest) Reset() {
	*x = ListUsersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_iam_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUsersRequest) ProtoMessage() {}

func (x *ListUsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_iam_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUsersRequest.ProtoReflect.Descriptor instead.
func (*ListUsersRequest) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_iam_user_proto_rawDescGZIP(), []int{2}
}

func (x *ListUsersRequest) GetMeta() *resource.ListRequest {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *ListUsersRequest) GetAccountID() string {
	if x != nil {
		return x.AccountID
	}
	return ""
}

type UserCredentials struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Password      string `protobuf:"bytes,1,opt,name=password,proto3" json:"password,omitempty"`
	PasswordReset string `protobuf:"bytes,2,opt,name=passwordReset,proto3" json:"passwordReset,omitempty"`
	SSH           *SSH   `protobuf:"bytes,11,opt,name=SSH,proto3" json:"SSH,omitempty"`
	TOTP          *TOTP  `protobuf:"bytes,21,opt,name=TOTP,proto3" json:"TOTP,omitempty"`
}

func (x *UserCredentials) Reset() {
	*x = UserCredentials{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_iam_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserCredentials) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserCredentials) ProtoMessage() {}

func (x *UserCredentials) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_iam_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserCredentials.ProtoReflect.Descriptor instead.
func (*UserCredentials) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_iam_user_proto_rawDescGZIP(), []int{3}
}

func (x *UserCredentials) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UserCredentials) GetPasswordReset() string {
	if x != nil {
		return x.PasswordReset
	}
	return ""
}

func (x *UserCredentials) GetSSH() *SSH {
	if x != nil {
		return x.SSH
	}
	return nil
}

func (x *UserCredentials) GetTOTP() *TOTP {
	if x != nil {
		return x.TOTP
	}
	return nil
}

type SSH struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enabled   bool   `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	PublicKey string `protobuf:"bytes,11,opt,name=publicKey,proto3" json:"publicKey,omitempty"` // string signature = 21;
}

func (x *SSH) Reset() {
	*x = SSH{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_iam_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SSH) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SSH) ProtoMessage() {}

func (x *SSH) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_iam_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SSH.ProtoReflect.Descriptor instead.
func (*SSH) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_iam_user_proto_rawDescGZIP(), []int{4}
}

func (x *SSH) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *SSH) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

type TOTP struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enabled bool   `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Secret  string `protobuf:"bytes,2,opt,name=secret,proto3" json:"secret,omitempty"`
}

func (x *TOTP) Reset() {
	*x = TOTP{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_iam_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TOTP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TOTP) ProtoMessage() {}

func (x *TOTP) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_iam_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TOTP.ProtoReflect.Descriptor instead.
func (*TOTP) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_iam_user_proto_rawDescGZIP(), []int{5}
}

func (x *TOTP) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *TOTP) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

type UserStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
}

func (x *UserStatus) Reset() {
	*x = UserStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_iam_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserStatus) ProtoMessage() {}

func (x *UserStatus) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_iam_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserStatus.ProtoReflect.Descriptor instead.
func (*UserStatus) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_iam_user_proto_rawDescGZIP(), []int{6}
}

func (x *UserStatus) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

type UserStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SecurityIssues *Metric            `protobuf:"bytes,1,opt,name=securityIssues,proto3" json:"securityIssues,omitempty"`
	Activity       map[string]*Metric `protobuf:"bytes,2,rep,name=activity,proto3" json:"activity,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // map[userActivity]*Metric
	FirstSeen      int64              `protobuf:"varint,3,opt,name=firstSeen,proto3" json:"firstSeen,omitempty"`
	LastSeen       int64              `protobuf:"varint,4,opt,name=lastSeen,proto3" json:"lastSeen,omitempty"`
	Timespan       string             `protobuf:"bytes,5,opt,name=timespan,proto3" json:"timespan,omitempty"`
	LastActivity   string             `protobuf:"bytes,6,opt,name=lastActivity,proto3" json:"lastActivity,omitempty"`
}

func (x *UserStats) Reset() {
	*x = UserStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_iam_user_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserStats) ProtoMessage() {}

func (x *UserStats) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_iam_user_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserStats.ProtoReflect.Descriptor instead.
func (*UserStats) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_iam_user_proto_rawDescGZIP(), []int{7}
}

func (x *UserStats) GetSecurityIssues() *Metric {
	if x != nil {
		return x.SecurityIssues
	}
	return nil
}

func (x *UserStats) GetActivity() map[string]*Metric {
	if x != nil {
		return x.Activity
	}
	return nil
}

func (x *UserStats) GetFirstSeen() int64 {
	if x != nil {
		return x.FirstSeen
	}
	return 0
}

func (x *UserStats) GetLastSeen() int64 {
	if x != nil {
		return x.LastSeen
	}
	return 0
}

func (x *UserStats) GetTimespan() string {
	if x != nil {
		return x.Timespan
	}
	return ""
}

func (x *UserStats) GetLastActivity() string {
	if x != nil {
		return x.LastActivity
	}
	return ""
}

type Metric struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total   uint32  `protobuf:"fixed32,1,opt,name=total,proto3" json:"total,omitempty"`
	Average float32 `protobuf:"fixed32,2,opt,name=average,proto3" json:"average,omitempty"`
}

func (x *Metric) Reset() {
	*x = Metric{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_iam_user_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metric) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metric) ProtoMessage() {}

func (x *Metric) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_iam_user_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metric.ProtoReflect.Descriptor instead.
func (*Metric) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_iam_user_proto_rawDescGZIP(), []int{8}
}

func (x *Metric) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *Metric) GetAverage() float32 {
	if x != nil {
		return x.Average
	}
	return 0
}

var File_mmesh_protobuf_resources_v1_iam_user_proto protoreflect.FileDescriptor

var file_mmesh_protobuf_resources_v1_iam_user_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x61,
	0x6d, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x69, 0x61,
	0x6d, 0x1a, 0x2f, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x2f, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xd0, 0x03, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05,
	0x72, 0x65, 0x61, 0x6c, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x65, 0x61,
	0x6c, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x24, 0x0a, 0x0d, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65,
	0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x56, 0x65,
	0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x65, 0x77, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x65, 0x77, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x36, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c,
	0x73, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x52, 0x0b, 0x63,
	0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x27, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x69, 0x61, 0x6d,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x29, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x65, 0x63,
	0x75, 0x72, 0x69, 0x74, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x72,
	0x6f, 0x6c, 0x65, 0x73, 0x18, 0x2a, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x41, 0x43, 0x4c, 0x73, 0x18, 0x2b, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x04, 0x41, 0x43, 0x4c, 0x73, 0x12, 0x24, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x33,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53,
	0x74, 0x61, 0x74, 0x73, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x12, 0x39, 0x0a, 0x0c, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18, 0x3d, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x0c, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x4d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x22, 0x54, 0x0a, 0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12,
	0x2a, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x1f, 0x0a, 0x05, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x69, 0x61, 0x6d,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x22, 0x5b, 0x0a, 0x10,
	0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x29, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x22, 0x8e, 0x01, 0x0a, 0x0f, 0x55, 0x73,
	0x65, 0x72, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x65, 0x74, 0x12,
	0x1a, 0x0a, 0x03, 0x53, 0x53, 0x48, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x69,
	0x61, 0x6d, 0x2e, 0x53, 0x53, 0x48, 0x52, 0x03, 0x53, 0x53, 0x48, 0x12, 0x1d, 0x0a, 0x04, 0x54,
	0x4f, 0x54, 0x50, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x69, 0x61, 0x6d, 0x2e,
	0x54, 0x4f, 0x54, 0x50, 0x52, 0x04, 0x54, 0x4f, 0x54, 0x50, 0x22, 0x3d, 0x0a, 0x03, 0x53, 0x53,
	0x48, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x22, 0x38, 0x0a, 0x04, 0x54, 0x4f, 0x54,
	0x50, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x22, 0x26, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x22, 0xbe, 0x02, 0x0a, 0x09,
	0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x33, 0x0a, 0x0e, 0x73, 0x65, 0x63,
	0x75, 0x72, 0x69, 0x74, 0x79, 0x49, 0x73, 0x73, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0b, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x52, 0x0e,
	0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x49, 0x73, 0x73, 0x75, 0x65, 0x73, 0x12, 0x38,
	0x0a, 0x08, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x73,
	0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73,
	0x74, 0x53, 0x65, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x66, 0x69, 0x72,
	0x73, 0x74, 0x53, 0x65, 0x65, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x53, 0x65,
	0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x53, 0x65,
	0x65, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x6e, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x6e, 0x12, 0x22,
	0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x1a, 0x48, 0x0a, 0x0d, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x21, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x38, 0x0a, 0x06,
	0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x07, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x18, 0x0a, 0x07,
	0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x61,
	0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x42, 0x27, 0x5a, 0x25, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2e,
	0x64, 0x65, 0x76, 0x2f, 0x6d, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x69, 0x61, 0x6d, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mmesh_protobuf_resources_v1_iam_user_proto_rawDescOnce sync.Once
	file_mmesh_protobuf_resources_v1_iam_user_proto_rawDescData = file_mmesh_protobuf_resources_v1_iam_user_proto_rawDesc
)

func file_mmesh_protobuf_resources_v1_iam_user_proto_rawDescGZIP() []byte {
	file_mmesh_protobuf_resources_v1_iam_user_proto_rawDescOnce.Do(func() {
		file_mmesh_protobuf_resources_v1_iam_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_mmesh_protobuf_resources_v1_iam_user_proto_rawDescData)
	})
	return file_mmesh_protobuf_resources_v1_iam_user_proto_rawDescData
}

var file_mmesh_protobuf_resources_v1_iam_user_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_mmesh_protobuf_resources_v1_iam_user_proto_goTypes = []interface{}{
	(*User)(nil),                  // 0: iam.User
	(*Users)(nil),                 // 1: iam.Users
	(*ListUsersRequest)(nil),      // 2: iam.ListUsersRequest
	(*UserCredentials)(nil),       // 3: iam.UserCredentials
	(*SSH)(nil),                   // 4: iam.SSH
	(*TOTP)(nil),                  // 5: iam.TOTP
	(*UserStatus)(nil),            // 6: iam.UserStatus
	(*UserStats)(nil),             // 7: iam.UserStats
	(*Metric)(nil),                // 8: iam.Metric
	nil,                           // 9: iam.UserStats.ActivityEntry
	(*metrics.EventMetrics)(nil),  // 10: metrics.EventMetrics
	(*resource.ListResponse)(nil), // 11: resource.ListResponse
	(*resource.ListRequest)(nil),  // 12: resource.ListRequest
}
var file_mmesh_protobuf_resources_v1_iam_user_proto_depIdxs = []int32{
	3,  // 0: iam.User.credentials:type_name -> iam.UserCredentials
	6,  // 1: iam.User.status:type_name -> iam.UserStatus
	7,  // 2: iam.User.stats:type_name -> iam.UserStats
	10, // 3: iam.User.eventMetrics:type_name -> metrics.EventMetrics
	11, // 4: iam.Users.meta:type_name -> resource.ListResponse
	0,  // 5: iam.Users.users:type_name -> iam.User
	12, // 6: iam.ListUsersRequest.meta:type_name -> resource.ListRequest
	4,  // 7: iam.UserCredentials.SSH:type_name -> iam.SSH
	5,  // 8: iam.UserCredentials.TOTP:type_name -> iam.TOTP
	8,  // 9: iam.UserStats.securityIssues:type_name -> iam.Metric
	9,  // 10: iam.UserStats.activity:type_name -> iam.UserStats.ActivityEntry
	8,  // 11: iam.UserStats.ActivityEntry.value:type_name -> iam.Metric
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_mmesh_protobuf_resources_v1_iam_user_proto_init() }
func file_mmesh_protobuf_resources_v1_iam_user_proto_init() {
	if File_mmesh_protobuf_resources_v1_iam_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mmesh_protobuf_resources_v1_iam_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_mmesh_protobuf_resources_v1_iam_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Users); i {
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
		file_mmesh_protobuf_resources_v1_iam_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUsersRequest); i {
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
		file_mmesh_protobuf_resources_v1_iam_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserCredentials); i {
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
		file_mmesh_protobuf_resources_v1_iam_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SSH); i {
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
		file_mmesh_protobuf_resources_v1_iam_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TOTP); i {
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
		file_mmesh_protobuf_resources_v1_iam_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserStatus); i {
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
		file_mmesh_protobuf_resources_v1_iam_user_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserStats); i {
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
		file_mmesh_protobuf_resources_v1_iam_user_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metric); i {
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
			RawDescriptor: file_mmesh_protobuf_resources_v1_iam_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mmesh_protobuf_resources_v1_iam_user_proto_goTypes,
		DependencyIndexes: file_mmesh_protobuf_resources_v1_iam_user_proto_depIdxs,
		MessageInfos:      file_mmesh_protobuf_resources_v1_iam_user_proto_msgTypes,
	}.Build()
	File_mmesh_protobuf_resources_v1_iam_user_proto = out.File
	file_mmesh_protobuf_resources_v1_iam_user_proto_rawDesc = nil
	file_mmesh_protobuf_resources_v1_iam_user_proto_goTypes = nil
	file_mmesh_protobuf_resources_v1_iam_user_proto_depIdxs = nil
}
