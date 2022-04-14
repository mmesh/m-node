// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: mmesh/protobuf/resources/v1/iam/auth/auth.proto

package auth

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	jwt "mmesh.dev/m-api-go/grpc/common/jwt"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AuthKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *AuthKey) Reset() {
	*x = AuthKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_iam_auth_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthKey) ProtoMessage() {}

func (x *AuthKey) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_iam_auth_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthKey.ProtoReflect.Descriptor instead.
func (*AuthKey) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_iam_auth_auth_proto_rawDescGZIP(), []int{0}
}

func (x *AuthKey) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type AuthToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subkey1       string   `protobuf:"bytes,1,opt,name=subkey1,proto3" json:"subkey1,omitempty"` // email
	Subkey2       string   `protobuf:"bytes,2,opt,name=subkey2,proto3" json:"subkey2,omitempty"` // realm
	ValidationKey string   `protobuf:"bytes,3,opt,name=validationKey,proto3" json:"validationKey,omitempty"`
	JWT           *jwt.JWT `protobuf:"bytes,4,opt,name=JWT,json=token,proto3" json:"JWT,omitempty"`
}

func (x *AuthToken) Reset() {
	*x = AuthToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_iam_auth_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthToken) ProtoMessage() {}

func (x *AuthToken) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_iam_auth_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthToken.ProtoReflect.Descriptor instead.
func (*AuthToken) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_iam_auth_auth_proto_rawDescGZIP(), []int{1}
}

func (x *AuthToken) GetSubkey1() string {
	if x != nil {
		return x.Subkey1
	}
	return ""
}

func (x *AuthToken) GetSubkey2() string {
	if x != nil {
		return x.Subkey2
	}
	return ""
}

func (x *AuthToken) GetValidationKey() string {
	if x != nil {
		return x.ValidationKey
	}
	return ""
}

func (x *AuthToken) GetJWT() *jwt.JWT {
	if x != nil {
		return x.JWT
	}
	return nil
}

type AccessRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserRealm  string `protobuf:"bytes,1,opt,name=userRealm,proto3" json:"userRealm,omitempty"`
	UserEmail  string `protobuf:"bytes,2,opt,name=userEmail,proto3" json:"userEmail,omitempty"`
	Realm      string `protobuf:"bytes,11,opt,name=realm,proto3" json:"realm,omitempty"`
	Tenant     string `protobuf:"bytes,12,opt,name=tenant,proto3" json:"tenant,omitempty"`
	Network    string `protobuf:"bytes,13,opt,name=network,proto3" json:"network,omitempty"`
	Subnet     string `protobuf:"bytes,14,opt,name=subnet,proto3" json:"subnet,omitempty"`
	Node       string `protobuf:"bytes,15,opt,name=node,proto3" json:"node,omitempty"`
	MMID       string `protobuf:"bytes,16,opt,name=MMID,proto3" json:"MMID,omitempty"`
	Permission int32  `protobuf:"varint,21,opt,name=permission,proto3" json:"permission,omitempty"`
	Endpoint   string `protobuf:"bytes,31,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
}

func (x *AccessRequest) Reset() {
	*x = AccessRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_iam_auth_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessRequest) ProtoMessage() {}

func (x *AccessRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_iam_auth_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessRequest.ProtoReflect.Descriptor instead.
func (*AccessRequest) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_iam_auth_auth_proto_rawDescGZIP(), []int{2}
}

func (x *AccessRequest) GetUserRealm() string {
	if x != nil {
		return x.UserRealm
	}
	return ""
}

func (x *AccessRequest) GetUserEmail() string {
	if x != nil {
		return x.UserEmail
	}
	return ""
}

func (x *AccessRequest) GetRealm() string {
	if x != nil {
		return x.Realm
	}
	return ""
}

func (x *AccessRequest) GetTenant() string {
	if x != nil {
		return x.Tenant
	}
	return ""
}

func (x *AccessRequest) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *AccessRequest) GetSubnet() string {
	if x != nil {
		return x.Subnet
	}
	return ""
}

func (x *AccessRequest) GetNode() string {
	if x != nil {
		return x.Node
	}
	return ""
}

func (x *AccessRequest) GetMMID() string {
	if x != nil {
		return x.MMID
	}
	return ""
}

func (x *AccessRequest) GetPermission() int32 {
	if x != nil {
		return x.Permission
	}
	return 0
}

func (x *AccessRequest) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

type AdminToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email     string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Realm     string `protobuf:"bytes,2,opt,name=realm,proto3" json:"realm,omitempty"`
	Target    string `protobuf:"bytes,3,opt,name=target,proto3" json:"target,omitempty"`
	Timestamp int64  `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	ExpTime   int64  `protobuf:"varint,5,opt,name=expTime,proto3" json:"expTime,omitempty"`
}

func (x *AdminToken) Reset() {
	*x = AdminToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_iam_auth_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminToken) ProtoMessage() {}

func (x *AdminToken) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_iam_auth_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminToken.ProtoReflect.Descriptor instead.
func (*AdminToken) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_iam_auth_auth_proto_rawDescGZIP(), []int{3}
}

func (x *AdminToken) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *AdminToken) GetRealm() string {
	if x != nil {
		return x.Realm
	}
	return ""
}

func (x *AdminToken) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *AdminToken) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *AdminToken) GetExpTime() int64 {
	if x != nil {
		return x.ExpTime
	}
	return 0
}

var File_mmesh_protobuf_resources_v1_iam_auth_auth_proto protoreflect.FileDescriptor

var file_mmesh_protobuf_resources_v1_iam_auth_auth_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x61,
	0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x61, 0x75, 0x74, 0x68, 0x1a, 0x26, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76,
	0x31, 0x2f, 0x6a, 0x77, 0x74, 0x2f, 0x6a, 0x77, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x1b, 0x0a, 0x07, 0x41, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x83, 0x01, 0x0a,
	0x09, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x62, 0x6b, 0x65, 0x79, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62,
	0x6b, 0x65, 0x79, 0x31, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6b, 0x65, 0x79, 0x32, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6b, 0x65, 0x79, 0x32, 0x12, 0x24,
	0x0a, 0x0d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4b, 0x65, 0x79, 0x12, 0x1c, 0x0a, 0x03, 0x4a, 0x57, 0x54, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x08, 0x2e, 0x6a, 0x77, 0x74, 0x2e, 0x4a, 0x57, 0x54, 0x52, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x22, 0x8f, 0x02, 0x0a, 0x0d, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x52, 0x65, 0x61, 0x6c,
	0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x52, 0x65, 0x61,
	0x6c, 0x6d, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x61, 0x6c, 0x6d, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x72, 0x65, 0x61, 0x6c, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x75, 0x62, 0x6e,
	0x65, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x4d, 0x4d, 0x49, 0x44, 0x18, 0x10, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x4d, 0x4d, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x15, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x70, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x22, 0x88, 0x01, 0x0a, 0x0a, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x61,
	0x6c, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x65, 0x61, 0x6c, 0x6d, 0x12,
	0x16, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x78, 0x70, 0x54, 0x69, 0x6d, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x65, 0x78, 0x70, 0x54, 0x69, 0x6d, 0x65, 0x42,
	0x2c, 0x5a, 0x2a, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x6d, 0x2d, 0x61,
	0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mmesh_protobuf_resources_v1_iam_auth_auth_proto_rawDescOnce sync.Once
	file_mmesh_protobuf_resources_v1_iam_auth_auth_proto_rawDescData = file_mmesh_protobuf_resources_v1_iam_auth_auth_proto_rawDesc
)

func file_mmesh_protobuf_resources_v1_iam_auth_auth_proto_rawDescGZIP() []byte {
	file_mmesh_protobuf_resources_v1_iam_auth_auth_proto_rawDescOnce.Do(func() {
		file_mmesh_protobuf_resources_v1_iam_auth_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_mmesh_protobuf_resources_v1_iam_auth_auth_proto_rawDescData)
	})
	return file_mmesh_protobuf_resources_v1_iam_auth_auth_proto_rawDescData
}

var file_mmesh_protobuf_resources_v1_iam_auth_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_mmesh_protobuf_resources_v1_iam_auth_auth_proto_goTypes = []interface{}{
	(*AuthKey)(nil),       // 0: auth.AuthKey
	(*AuthToken)(nil),     // 1: auth.AuthToken
	(*AccessRequest)(nil), // 2: auth.AccessRequest
	(*AdminToken)(nil),    // 3: auth.AdminToken
	(*jwt.JWT)(nil),       // 4: jwt.JWT
}
var file_mmesh_protobuf_resources_v1_iam_auth_auth_proto_depIdxs = []int32{
	4, // 0: auth.AuthToken.JWT:type_name -> jwt.JWT
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_mmesh_protobuf_resources_v1_iam_auth_auth_proto_init() }
func file_mmesh_protobuf_resources_v1_iam_auth_auth_proto_init() {
	if File_mmesh_protobuf_resources_v1_iam_auth_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mmesh_protobuf_resources_v1_iam_auth_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthKey); i {
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
		file_mmesh_protobuf_resources_v1_iam_auth_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthToken); i {
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
		file_mmesh_protobuf_resources_v1_iam_auth_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccessRequest); i {
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
		file_mmesh_protobuf_resources_v1_iam_auth_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminToken); i {
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
			RawDescriptor: file_mmesh_protobuf_resources_v1_iam_auth_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mmesh_protobuf_resources_v1_iam_auth_auth_proto_goTypes,
		DependencyIndexes: file_mmesh_protobuf_resources_v1_iam_auth_auth_proto_depIdxs,
		MessageInfos:      file_mmesh_protobuf_resources_v1_iam_auth_auth_proto_msgTypes,
	}.Build()
	File_mmesh_protobuf_resources_v1_iam_auth_auth_proto = out.File
	file_mmesh_protobuf_resources_v1_iam_auth_auth_proto_rawDesc = nil
	file_mmesh_protobuf_resources_v1_iam_auth_auth_proto_goTypes = nil
	file_mmesh_protobuf_resources_v1_iam_auth_auth_proto_depIdxs = nil
}
