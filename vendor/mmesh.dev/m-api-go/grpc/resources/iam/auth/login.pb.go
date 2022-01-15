// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: mmesh/protobuf/resources/v1/iam/auth/login.proto

package auth

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

type AuthMethod int32

const (
	AuthMethod_PASSWORD AuthMethod = 0
	AuthMethod_SSH_KEY  AuthMethod = 1
)

// Enum value maps for AuthMethod.
var (
	AuthMethod_name = map[int32]string{
		0: "PASSWORD",
		1: "SSH_KEY",
	}
	AuthMethod_value = map[string]int32{
		"PASSWORD": 0,
		"SSH_KEY":  1,
	}
)

func (x AuthMethod) Enum() *AuthMethod {
	p := new(AuthMethod)
	*p = x
	return p
}

func (x AuthMethod) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AuthMethod) Descriptor() protoreflect.EnumDescriptor {
	return file_mmesh_protobuf_resources_v1_iam_auth_login_proto_enumTypes[0].Descriptor()
}

func (AuthMethod) Type() protoreflect.EnumType {
	return &file_mmesh_protobuf_resources_v1_iam_auth_login_proto_enumTypes[0]
}

func (x AuthMethod) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AuthMethod.Descriptor instead.
func (AuthMethod) EnumDescriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_iam_auth_login_proto_rawDescGZIP(), []int{0}
}

type LoginResult int32

const (
	LoginResult_LOGIN_SUCCESSFUL   LoginResult = 0
	LoginResult_LOGIN_FAILED       LoginResult = 1
	LoginResult_LOGIN_2FA_REQUIRED LoginResult = 2
)

// Enum value maps for LoginResult.
var (
	LoginResult_name = map[int32]string{
		0: "LOGIN_SUCCESSFUL",
		1: "LOGIN_FAILED",
		2: "LOGIN_2FA_REQUIRED",
	}
	LoginResult_value = map[string]int32{
		"LOGIN_SUCCESSFUL":   0,
		"LOGIN_FAILED":       1,
		"LOGIN_2FA_REQUIRED": 2,
	}
)

func (x LoginResult) Enum() *LoginResult {
	p := new(LoginResult)
	*p = x
	return p
}

func (x LoginResult) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LoginResult) Descriptor() protoreflect.EnumDescriptor {
	return file_mmesh_protobuf_resources_v1_iam_auth_login_proto_enumTypes[1].Descriptor()
}

func (LoginResult) Type() protoreflect.EnumType {
	return &file_mmesh_protobuf_resources_v1_iam_auth_login_proto_enumTypes[1]
}

func (x LoginResult) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LoginResult.Descriptor instead.
func (LoginResult) EnumDescriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_iam_auth_login_proto_rawDescGZIP(), []int{1}
}

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Realm string `protobuf:"bytes,1,opt,name=realm,proto3" json:"realm,omitempty"` // accountID
	Email string `protobuf:"bytes,11,opt,name=email,proto3" json:"email,omitempty"`
	// bool admin = 21;
	AuthMethod AuthMethod `protobuf:"varint,31,opt,name=authMethod,proto3,enum=auth.AuthMethod" json:"authMethod,omitempty"`
	Password   string     `protobuf:"bytes,41,opt,name=password,proto3" json:"password,omitempty"`
	TotpCode   string     `protobuf:"bytes,51,opt,name=totpCode,proto3" json:"totpCode,omitempty"`
	LocationID string     `protobuf:"bytes,101,opt,name=locationID,proto3" json:"locationID,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_iam_auth_login_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_iam_auth_login_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_iam_auth_login_proto_rawDescGZIP(), []int{0}
}

func (x *LoginRequest) GetRealm() string {
	if x != nil {
		return x.Realm
	}
	return ""
}

func (x *LoginRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *LoginRequest) GetAuthMethod() AuthMethod {
	if x != nil {
		return x.AuthMethod
	}
	return AuthMethod_PASSWORD
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *LoginRequest) GetTotpCode() string {
	if x != nil {
		return x.TotpCode
	}
	return ""
}

func (x *LoginRequest) GetLocationID() string {
	if x != nil {
		return x.LocationID
	}
	return ""
}

type LoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result       LoginResult `protobuf:"varint,1,opt,name=result,proto3,enum=auth.LoginResult" json:"result,omitempty"`
	AuthKey      *AuthKey    `protobuf:"bytes,11,opt,name=authKey,proto3" json:"authKey,omitempty"`
	FederationID string      `protobuf:"bytes,21,opt,name=federationID,proto3" json:"federationID,omitempty"`
	IsAdmin      bool        `protobuf:"varint,31,opt,name=isAdmin,proto3" json:"isAdmin,omitempty"`
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_iam_auth_login_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_iam_auth_login_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_iam_auth_login_proto_rawDescGZIP(), []int{1}
}

func (x *LoginResponse) GetResult() LoginResult {
	if x != nil {
		return x.Result
	}
	return LoginResult_LOGIN_SUCCESSFUL
}

func (x *LoginResponse) GetAuthKey() *AuthKey {
	if x != nil {
		return x.AuthKey
	}
	return nil
}

func (x *LoginResponse) GetFederationID() string {
	if x != nil {
		return x.FederationID
	}
	return ""
}

func (x *LoginResponse) GetIsAdmin() bool {
	if x != nil {
		return x.IsAdmin
	}
	return false
}

var File_mmesh_protobuf_resources_v1_iam_auth_login_proto protoreflect.FileDescriptor

var file_mmesh_protobuf_resources_v1_iam_auth_login_proto_rawDesc = []byte{
	0x0a, 0x30, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x61,
	0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x61, 0x75, 0x74, 0x68, 0x1a, 0x2f, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc4, 0x01, 0x0a, 0x0c, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65,
	0x61, 0x6c, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x65, 0x61, 0x6c, 0x6d,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x30, 0x0a, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x0a, 0x61, 0x75,
	0x74, 0x68, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x29, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x6f, 0x74, 0x70, 0x43, 0x6f, 0x64, 0x65,
	0x18, 0x33, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x6f, 0x74, 0x70, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x1e, 0x0a, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x65,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44,
	0x22, 0xa1, 0x01, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x29, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x11, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x27, 0x0a,
	0x07, 0x61, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x52, 0x07, 0x61,
	0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x12, 0x22, 0x0a, 0x0c, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x65,
	0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x73,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0x2a, 0x27, 0x0a, 0x0a, 0x41, 0x75, 0x74, 0x68, 0x4d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x12, 0x0c, 0x0a, 0x08, 0x50, 0x41, 0x53, 0x53, 0x57, 0x4f, 0x52, 0x44, 0x10, 0x00,
	0x12, 0x0b, 0x0a, 0x07, 0x53, 0x53, 0x48, 0x5f, 0x4b, 0x45, 0x59, 0x10, 0x01, 0x2a, 0x4d, 0x0a,
	0x0b, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x14, 0x0a, 0x10,
	0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x46, 0x55, 0x4c,
	0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x5f, 0x46, 0x41, 0x49, 0x4c,
	0x45, 0x44, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x5f, 0x32, 0x46,
	0x41, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x49, 0x52, 0x45, 0x44, 0x10, 0x02, 0x42, 0x2c, 0x5a, 0x2a,
	0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x6d, 0x2d, 0x61, 0x70, 0x69, 0x2d,
	0x67, 0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_mmesh_protobuf_resources_v1_iam_auth_login_proto_rawDescOnce sync.Once
	file_mmesh_protobuf_resources_v1_iam_auth_login_proto_rawDescData = file_mmesh_protobuf_resources_v1_iam_auth_login_proto_rawDesc
)

func file_mmesh_protobuf_resources_v1_iam_auth_login_proto_rawDescGZIP() []byte {
	file_mmesh_protobuf_resources_v1_iam_auth_login_proto_rawDescOnce.Do(func() {
		file_mmesh_protobuf_resources_v1_iam_auth_login_proto_rawDescData = protoimpl.X.CompressGZIP(file_mmesh_protobuf_resources_v1_iam_auth_login_proto_rawDescData)
	})
	return file_mmesh_protobuf_resources_v1_iam_auth_login_proto_rawDescData
}

var file_mmesh_protobuf_resources_v1_iam_auth_login_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_mmesh_protobuf_resources_v1_iam_auth_login_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_mmesh_protobuf_resources_v1_iam_auth_login_proto_goTypes = []interface{}{
	(AuthMethod)(0),       // 0: auth.AuthMethod
	(LoginResult)(0),      // 1: auth.LoginResult
	(*LoginRequest)(nil),  // 2: auth.LoginRequest
	(*LoginResponse)(nil), // 3: auth.LoginResponse
	(*AuthKey)(nil),       // 4: auth.AuthKey
}
var file_mmesh_protobuf_resources_v1_iam_auth_login_proto_depIdxs = []int32{
	0, // 0: auth.LoginRequest.authMethod:type_name -> auth.AuthMethod
	1, // 1: auth.LoginResponse.result:type_name -> auth.LoginResult
	4, // 2: auth.LoginResponse.authKey:type_name -> auth.AuthKey
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_mmesh_protobuf_resources_v1_iam_auth_login_proto_init() }
func file_mmesh_protobuf_resources_v1_iam_auth_login_proto_init() {
	if File_mmesh_protobuf_resources_v1_iam_auth_login_proto != nil {
		return
	}
	file_mmesh_protobuf_resources_v1_iam_auth_auth_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_mmesh_protobuf_resources_v1_iam_auth_login_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest); i {
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
		file_mmesh_protobuf_resources_v1_iam_auth_login_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResponse); i {
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
			RawDescriptor: file_mmesh_protobuf_resources_v1_iam_auth_login_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mmesh_protobuf_resources_v1_iam_auth_login_proto_goTypes,
		DependencyIndexes: file_mmesh_protobuf_resources_v1_iam_auth_login_proto_depIdxs,
		EnumInfos:         file_mmesh_protobuf_resources_v1_iam_auth_login_proto_enumTypes,
		MessageInfos:      file_mmesh_protobuf_resources_v1_iam_auth_login_proto_msgTypes,
	}.Build()
	File_mmesh_protobuf_resources_v1_iam_auth_login_proto = out.File
	file_mmesh_protobuf_resources_v1_iam_auth_login_proto_rawDesc = nil
	file_mmesh_protobuf_resources_v1_iam_auth_login_proto_goTypes = nil
	file_mmesh_protobuf_resources_v1_iam_auth_login_proto_depIdxs = nil
}
