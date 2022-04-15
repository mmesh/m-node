// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.0
// source: mmesh/protobuf/resources/v1/messaging/chat.proto

package messaging

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

type ChatBackendType int32

const (
	ChatBackendType_UNSPECIFIED ChatBackendType = 0
	ChatBackendType_CRISP       ChatBackendType = 5
	ChatBackendType_SLACK       ChatBackendType = 11
	ChatBackendType_DISCORD     ChatBackendType = 12 // WHATSAPP = 21;
)

// Enum value maps for ChatBackendType.
var (
	ChatBackendType_name = map[int32]string{
		0:  "UNSPECIFIED",
		5:  "CRISP",
		11: "SLACK",
		12: "DISCORD",
	}
	ChatBackendType_value = map[string]int32{
		"UNSPECIFIED": 0,
		"CRISP":       5,
		"SLACK":       11,
		"DISCORD":     12,
	}
)

func (x ChatBackendType) Enum() *ChatBackendType {
	p := new(ChatBackendType)
	*p = x
	return p
}

func (x ChatBackendType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChatBackendType) Descriptor() protoreflect.EnumDescriptor {
	return file_mmesh_protobuf_resources_v1_messaging_chat_proto_enumTypes[0].Descriptor()
}

func (ChatBackendType) Type() protoreflect.EnumType {
	return &file_mmesh_protobuf_resources_v1_messaging_chat_proto_enumTypes[0]
}

func (x ChatBackendType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChatBackendType.Descriptor instead.
func (ChatBackendType) EnumDescriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_messaging_chat_proto_rawDescGZIP(), []int{0}
}

type ChatMessageDirection int32

const (
	ChatMessageDirection_USER_TO_PROVIDER ChatMessageDirection = 0
	ChatMessageDirection_PROVIDER_TO_USER ChatMessageDirection = 1
)

// Enum value maps for ChatMessageDirection.
var (
	ChatMessageDirection_name = map[int32]string{
		0: "USER_TO_PROVIDER",
		1: "PROVIDER_TO_USER",
	}
	ChatMessageDirection_value = map[string]int32{
		"USER_TO_PROVIDER": 0,
		"PROVIDER_TO_USER": 1,
	}
)

func (x ChatMessageDirection) Enum() *ChatMessageDirection {
	p := new(ChatMessageDirection)
	*p = x
	return p
}

func (x ChatMessageDirection) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChatMessageDirection) Descriptor() protoreflect.EnumDescriptor {
	return file_mmesh_protobuf_resources_v1_messaging_chat_proto_enumTypes[1].Descriptor()
}

func (ChatMessageDirection) Type() protoreflect.EnumType {
	return &file_mmesh_protobuf_resources_v1_messaging_chat_proto_enumTypes[1]
}

func (x ChatMessageDirection) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChatMessageDirection.Descriptor instead.
func (ChatMessageDirection) EnumDescriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_messaging_chat_proto_rawDescGZIP(), []int{1}
}

type ChatMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceID        string               `protobuf:"bytes,1,opt,name=serviceID,proto3" json:"serviceID,omitempty"`
	ProviderID       string               `protobuf:"bytes,2,opt,name=providerID,proto3" json:"providerID,omitempty"`
	IssueID          string               `protobuf:"bytes,3,opt,name=issueID,proto3" json:"issueID,omitempty"`
	ThreadID         string               `protobuf:"bytes,4,opt,name=threadID,proto3" json:"threadID,omitempty"` // crisp conversation sessionID | slack conversation threadTS
	BackendType      ChatBackendType      `protobuf:"varint,11,opt,name=backendType,proto3,enum=messaging.ChatBackendType" json:"backendType,omitempty"`
	Direction        ChatMessageDirection `protobuf:"varint,12,opt,name=direction,proto3,enum=messaging.ChatMessageDirection" json:"direction,omitempty"`
	AccountID        string               `protobuf:"bytes,21,opt,name=accountID,proto3" json:"accountID,omitempty"`
	UserEmail        string               `protobuf:"bytes,22,opt,name=userEmail,proto3" json:"userEmail,omitempty"`
	UserNickname     string               `protobuf:"bytes,23,opt,name=userNickname,proto3" json:"userNickname,omitempty"`
	ProviderChannel  string               `protobuf:"bytes,31,opt,name=providerChannel,proto3" json:"providerChannel,omitempty"`
	OperatorEmail    string               `protobuf:"bytes,32,opt,name=operatorEmail,proto3" json:"operatorEmail,omitempty"`
	OperatorNickname string               `protobuf:"bytes,33,opt,name=operatorNickname,proto3" json:"operatorNickname,omitempty"`
	Payload          []byte               `protobuf:"bytes,41,opt,name=payload,proto3" json:"payload,omitempty"`
	Timestamp        int64                `protobuf:"varint,101,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *ChatMessage) Reset() {
	*x = ChatMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_messaging_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatMessage) ProtoMessage() {}

func (x *ChatMessage) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_messaging_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatMessage.ProtoReflect.Descriptor instead.
func (*ChatMessage) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_messaging_chat_proto_rawDescGZIP(), []int{0}
}

func (x *ChatMessage) GetServiceID() string {
	if x != nil {
		return x.ServiceID
	}
	return ""
}

func (x *ChatMessage) GetProviderID() string {
	if x != nil {
		return x.ProviderID
	}
	return ""
}

func (x *ChatMessage) GetIssueID() string {
	if x != nil {
		return x.IssueID
	}
	return ""
}

func (x *ChatMessage) GetThreadID() string {
	if x != nil {
		return x.ThreadID
	}
	return ""
}

func (x *ChatMessage) GetBackendType() ChatBackendType {
	if x != nil {
		return x.BackendType
	}
	return ChatBackendType_UNSPECIFIED
}

func (x *ChatMessage) GetDirection() ChatMessageDirection {
	if x != nil {
		return x.Direction
	}
	return ChatMessageDirection_USER_TO_PROVIDER
}

func (x *ChatMessage) GetAccountID() string {
	if x != nil {
		return x.AccountID
	}
	return ""
}

func (x *ChatMessage) GetUserEmail() string {
	if x != nil {
		return x.UserEmail
	}
	return ""
}

func (x *ChatMessage) GetUserNickname() string {
	if x != nil {
		return x.UserNickname
	}
	return ""
}

func (x *ChatMessage) GetProviderChannel() string {
	if x != nil {
		return x.ProviderChannel
	}
	return ""
}

func (x *ChatMessage) GetOperatorEmail() string {
	if x != nil {
		return x.OperatorEmail
	}
	return ""
}

func (x *ChatMessage) GetOperatorNickname() string {
	if x != nil {
		return x.OperatorNickname
	}
	return ""
}

func (x *ChatMessage) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *ChatMessage) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type ChatThread struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ThreadID    string                 `protobuf:"bytes,1,opt,name=threadID,proto3" json:"threadID,omitempty"`
	Comments    map[int64]*ChatComment `protobuf:"bytes,11,rep,name=comments,proto3" json:"comments,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // map[timestamp]*ChatComment
	LastUpdated int64                  `protobuf:"varint,101,opt,name=lastUpdated,proto3" json:"lastUpdated,omitempty"`
}

func (x *ChatThread) Reset() {
	*x = ChatThread{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_messaging_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatThread) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatThread) ProtoMessage() {}

func (x *ChatThread) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_messaging_chat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatThread.ProtoReflect.Descriptor instead.
func (*ChatThread) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_messaging_chat_proto_rawDescGZIP(), []int{1}
}

func (x *ChatThread) GetThreadID() string {
	if x != nil {
		return x.ThreadID
	}
	return ""
}

func (x *ChatThread) GetComments() map[int64]*ChatComment {
	if x != nil {
		return x.Comments
	}
	return nil
}

func (x *ChatThread) GetLastUpdated() int64 {
	if x != nil {
		return x.LastUpdated
	}
	return 0
}

type ChatComment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fingerprint int64  `protobuf:"varint,5,opt,name=fingerprint,proto3" json:"fingerprint,omitempty"`
	Nickname    string `protobuf:"bytes,11,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Text        []byte `protobuf:"bytes,21,opt,name=text,proto3" json:"text,omitempty"`
	Timestamp   int64  `protobuf:"varint,31,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	IsCustomer  bool   `protobuf:"varint,41,opt,name=isCustomer,proto3" json:"isCustomer,omitempty"`
	Edited      bool   `protobuf:"varint,51,opt,name=edited,proto3" json:"edited,omitempty"`
}

func (x *ChatComment) Reset() {
	*x = ChatComment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_messaging_chat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatComment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatComment) ProtoMessage() {}

func (x *ChatComment) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_messaging_chat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatComment.ProtoReflect.Descriptor instead.
func (*ChatComment) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_messaging_chat_proto_rawDescGZIP(), []int{2}
}

func (x *ChatComment) GetFingerprint() int64 {
	if x != nil {
		return x.Fingerprint
	}
	return 0
}

func (x *ChatComment) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *ChatComment) GetText() []byte {
	if x != nil {
		return x.Text
	}
	return nil
}

func (x *ChatComment) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *ChatComment) GetIsCustomer() bool {
	if x != nil {
		return x.IsCustomer
	}
	return false
}

func (x *ChatComment) GetEdited() bool {
	if x != nil {
		return x.Edited
	}
	return false
}

type ChatParticipant struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email      string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Nickname   string `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Company    string `protobuf:"bytes,11,opt,name=company,proto3" json:"company,omitempty"`
	IsCustomer bool   `protobuf:"varint,21,opt,name=isCustomer,proto3" json:"isCustomer,omitempty"`
}

func (x *ChatParticipant) Reset() {
	*x = ChatParticipant{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_messaging_chat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatParticipant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatParticipant) ProtoMessage() {}

func (x *ChatParticipant) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_messaging_chat_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatParticipant.ProtoReflect.Descriptor instead.
func (*ChatParticipant) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_messaging_chat_proto_rawDescGZIP(), []int{3}
}

func (x *ChatParticipant) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *ChatParticipant) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *ChatParticipant) GetCompany() string {
	if x != nil {
		return x.Company
	}
	return ""
}

func (x *ChatParticipant) GetIsCustomer() bool {
	if x != nil {
		return x.IsCustomer
	}
	return false
}

var File_mmesh_protobuf_resources_v1_messaging_chat_proto protoreflect.FileDescriptor

var file_mmesh_protobuf_resources_v1_messaging_chat_proto_rawDesc = []byte{
	0x0a, 0x30, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x22, 0x92, 0x04,
	0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x69,
	0x73, 0x73, 0x75, 0x65, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x73,
	0x73, 0x75, 0x65, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x49,
	0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x49,
	0x44, 0x12, 0x3c, 0x0a, 0x0b, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69,
	0x6e, 0x67, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x0b, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x3d, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x43,
	0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c,
	0x0a, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x15, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09,
	0x75, 0x73, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x75, 0x73, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x22, 0x0a, 0x0c, 0x75, 0x73,
	0x65, 0x72, 0x4e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x28,
	0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x24, 0x0a, 0x0d, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x20, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x2a,
	0x0a, 0x10, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x4e, 0x69, 0x63, 0x6b, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x21, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x4e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x29, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x18, 0x65, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x22, 0xe0, 0x01, 0x0a, 0x0a, 0x43, 0x68, 0x61, 0x74, 0x54, 0x68, 0x72, 0x65, 0x61,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x49, 0x44, 0x12, 0x3f, 0x0a,
	0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x23, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x68, 0x61, 0x74,
	0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x20,
	0x0a, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x65, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x1a, 0x53, 0x0a, 0x0d, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x2c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x43,
	0x68, 0x61, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xb5, 0x01, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x66, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x70,
	0x72, 0x69, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x66, 0x69, 0x6e, 0x67,
	0x65, 0x72, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x15, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x73, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x18, 0x29, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x64, 0x69, 0x74, 0x65, 0x64, 0x18,
	0x33, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x65, 0x64, 0x69, 0x74, 0x65, 0x64, 0x22, 0x7d, 0x0a,
	0x0f, 0x43, 0x68, 0x61, 0x74, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x1e, 0x0a, 0x0a,
	0x69, 0x73, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x18, 0x15, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0a, 0x69, 0x73, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2a, 0x45, 0x0a, 0x0f,
	0x43, 0x68, 0x61, 0x74, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x09, 0x0a, 0x05, 0x43, 0x52, 0x49, 0x53, 0x50, 0x10, 0x05, 0x12, 0x09, 0x0a, 0x05, 0x53,
	0x4c, 0x41, 0x43, 0x4b, 0x10, 0x0b, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x49, 0x53, 0x43, 0x4f, 0x52,
	0x44, 0x10, 0x0c, 0x2a, 0x42, 0x0a, 0x14, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x10, 0x55,
	0x53, 0x45, 0x52, 0x5f, 0x54, 0x4f, 0x5f, 0x50, 0x52, 0x4f, 0x56, 0x49, 0x44, 0x45, 0x52, 0x10,
	0x00, 0x12, 0x14, 0x0a, 0x10, 0x50, 0x52, 0x4f, 0x56, 0x49, 0x44, 0x45, 0x52, 0x5f, 0x54, 0x4f,
	0x5f, 0x55, 0x53, 0x45, 0x52, 0x10, 0x01, 0x42, 0x2d, 0x5a, 0x2b, 0x6d, 0x6d, 0x65, 0x73, 0x68,
	0x2e, 0x64, 0x65, 0x76, 0x2f, 0x6d, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x72,
	0x70, 0x63, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mmesh_protobuf_resources_v1_messaging_chat_proto_rawDescOnce sync.Once
	file_mmesh_protobuf_resources_v1_messaging_chat_proto_rawDescData = file_mmesh_protobuf_resources_v1_messaging_chat_proto_rawDesc
)

func file_mmesh_protobuf_resources_v1_messaging_chat_proto_rawDescGZIP() []byte {
	file_mmesh_protobuf_resources_v1_messaging_chat_proto_rawDescOnce.Do(func() {
		file_mmesh_protobuf_resources_v1_messaging_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_mmesh_protobuf_resources_v1_messaging_chat_proto_rawDescData)
	})
	return file_mmesh_protobuf_resources_v1_messaging_chat_proto_rawDescData
}

var file_mmesh_protobuf_resources_v1_messaging_chat_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_mmesh_protobuf_resources_v1_messaging_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_mmesh_protobuf_resources_v1_messaging_chat_proto_goTypes = []interface{}{
	(ChatBackendType)(0),      // 0: messaging.ChatBackendType
	(ChatMessageDirection)(0), // 1: messaging.ChatMessageDirection
	(*ChatMessage)(nil),       // 2: messaging.ChatMessage
	(*ChatThread)(nil),        // 3: messaging.ChatThread
	(*ChatComment)(nil),       // 4: messaging.ChatComment
	(*ChatParticipant)(nil),   // 5: messaging.ChatParticipant
	nil,                       // 6: messaging.ChatThread.CommentsEntry
}
var file_mmesh_protobuf_resources_v1_messaging_chat_proto_depIdxs = []int32{
	0, // 0: messaging.ChatMessage.backendType:type_name -> messaging.ChatBackendType
	1, // 1: messaging.ChatMessage.direction:type_name -> messaging.ChatMessageDirection
	6, // 2: messaging.ChatThread.comments:type_name -> messaging.ChatThread.CommentsEntry
	4, // 3: messaging.ChatThread.CommentsEntry.value:type_name -> messaging.ChatComment
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_mmesh_protobuf_resources_v1_messaging_chat_proto_init() }
func file_mmesh_protobuf_resources_v1_messaging_chat_proto_init() {
	if File_mmesh_protobuf_resources_v1_messaging_chat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mmesh_protobuf_resources_v1_messaging_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatMessage); i {
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
		file_mmesh_protobuf_resources_v1_messaging_chat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatThread); i {
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
		file_mmesh_protobuf_resources_v1_messaging_chat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatComment); i {
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
		file_mmesh_protobuf_resources_v1_messaging_chat_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatParticipant); i {
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
			RawDescriptor: file_mmesh_protobuf_resources_v1_messaging_chat_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mmesh_protobuf_resources_v1_messaging_chat_proto_goTypes,
		DependencyIndexes: file_mmesh_protobuf_resources_v1_messaging_chat_proto_depIdxs,
		EnumInfos:         file_mmesh_protobuf_resources_v1_messaging_chat_proto_enumTypes,
		MessageInfos:      file_mmesh_protobuf_resources_v1_messaging_chat_proto_msgTypes,
	}.Build()
	File_mmesh_protobuf_resources_v1_messaging_chat_proto = out.File
	file_mmesh_protobuf_resources_v1_messaging_chat_proto_rawDesc = nil
	file_mmesh_protobuf_resources_v1_messaging_chat_proto_goTypes = nil
	file_mmesh_protobuf_resources_v1_messaging_chat_proto_depIdxs = nil
}
