// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.0--rc2
// source: mmesh/protobuf/rpc/v1/accountAPI.proto

package rpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status1 "google.golang.org/grpc/status"
	status "mmesh.dev/m-api-go/grpc/common/status"
	account "mmesh.dev/m-api-go/grpc/resources/account"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	AccountAPI_NewAccount_FullMethodName              = "/api.AccountAPI/NewAccount"
	AccountAPI_ListAccounts_FullMethodName            = "/api.AccountAPI/ListAccounts"
	AccountAPI_GetAccount_FullMethodName              = "/api.AccountAPI/GetAccount"
	AccountAPI_CancelAccount_FullMethodName           = "/api.AccountAPI/CancelAccount"
	AccountAPI_DeleteAccount_FullMethodName           = "/api.AccountAPI/DeleteAccount"
	AccountAPI_SetAccountMetadata_FullMethodName      = "/api.AccountAPI/SetAccountMetadata"
	AccountAPI_SetAccountIntegrations_FullMethodName  = "/api.AccountAPI/SetAccountIntegrations"
	AccountAPI_EnableAccount_FullMethodName           = "/api.AccountAPI/EnableAccount"
	AccountAPI_DisableAccount_FullMethodName          = "/api.AccountAPI/DisableAccount"
	AccountAPI_SuspendAccountService_FullMethodName   = "/api.AccountAPI/SuspendAccountService"
	AccountAPI_UnsuspendAccountService_FullMethodName = "/api.AccountAPI/UnsuspendAccountService"
	AccountAPI_CancelAccountService_FullMethodName    = "/api.AccountAPI/CancelAccountService"
)

// AccountAPIClient is the client API for AccountAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccountAPIClient interface {
	// account
	NewAccount(ctx context.Context, in *account.NewAccountRequest, opts ...grpc.CallOption) (*account.NewAccountResponse, error)
	ListAccounts(ctx context.Context, in *account.ListAccountsRequest, opts ...grpc.CallOption) (*account.Accounts, error)
	GetAccount(ctx context.Context, in *account.AccountReq, opts ...grpc.CallOption) (*account.Account, error)
	//	rpc UpdateAccount(account.Account) returns (account.Account) {
	//	  option (google.api.http) = {
	//	    post: "/api/v1/accounts/{accountID}"
	//	    body: "*"
	//	  };
	//	}
	CancelAccount(ctx context.Context, in *account.AccountReq, opts ...grpc.CallOption) (*status.StatusResponse, error)
	DeleteAccount(ctx context.Context, in *account.AccountReq, opts ...grpc.CallOption) (*status.StatusResponse, error)
	SetAccountMetadata(ctx context.Context, in *account.SetAccountMetadataRequest, opts ...grpc.CallOption) (*account.Account, error)
	SetAccountIntegrations(ctx context.Context, in *account.SetAccountIntegrationsRequest, opts ...grpc.CallOption) (*account.Account, error)
	EnableAccount(ctx context.Context, in *account.AccountReq, opts ...grpc.CallOption) (*account.Account, error)
	DisableAccount(ctx context.Context, in *account.AccountReq, opts ...grpc.CallOption) (*account.Account, error)
	// rpc SetAccountService(account.Account) returns (account.Account) {}
	SuspendAccountService(ctx context.Context, in *account.AccountReq, opts ...grpc.CallOption) (*account.Account, error)
	UnsuspendAccountService(ctx context.Context, in *account.AccountReq, opts ...grpc.CallOption) (*account.Account, error)
	CancelAccountService(ctx context.Context, in *account.AccountReq, opts ...grpc.CallOption) (*account.Account, error)
}

type accountAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountAPIClient(cc grpc.ClientConnInterface) AccountAPIClient {
	return &accountAPIClient{cc}
}

func (c *accountAPIClient) NewAccount(ctx context.Context, in *account.NewAccountRequest, opts ...grpc.CallOption) (*account.NewAccountResponse, error) {
	out := new(account.NewAccountResponse)
	err := c.cc.Invoke(ctx, AccountAPI_NewAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountAPIClient) ListAccounts(ctx context.Context, in *account.ListAccountsRequest, opts ...grpc.CallOption) (*account.Accounts, error) {
	out := new(account.Accounts)
	err := c.cc.Invoke(ctx, AccountAPI_ListAccounts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountAPIClient) GetAccount(ctx context.Context, in *account.AccountReq, opts ...grpc.CallOption) (*account.Account, error) {
	out := new(account.Account)
	err := c.cc.Invoke(ctx, AccountAPI_GetAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountAPIClient) CancelAccount(ctx context.Context, in *account.AccountReq, opts ...grpc.CallOption) (*status.StatusResponse, error) {
	out := new(status.StatusResponse)
	err := c.cc.Invoke(ctx, AccountAPI_CancelAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountAPIClient) DeleteAccount(ctx context.Context, in *account.AccountReq, opts ...grpc.CallOption) (*status.StatusResponse, error) {
	out := new(status.StatusResponse)
	err := c.cc.Invoke(ctx, AccountAPI_DeleteAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountAPIClient) SetAccountMetadata(ctx context.Context, in *account.SetAccountMetadataRequest, opts ...grpc.CallOption) (*account.Account, error) {
	out := new(account.Account)
	err := c.cc.Invoke(ctx, AccountAPI_SetAccountMetadata_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountAPIClient) SetAccountIntegrations(ctx context.Context, in *account.SetAccountIntegrationsRequest, opts ...grpc.CallOption) (*account.Account, error) {
	out := new(account.Account)
	err := c.cc.Invoke(ctx, AccountAPI_SetAccountIntegrations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountAPIClient) EnableAccount(ctx context.Context, in *account.AccountReq, opts ...grpc.CallOption) (*account.Account, error) {
	out := new(account.Account)
	err := c.cc.Invoke(ctx, AccountAPI_EnableAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountAPIClient) DisableAccount(ctx context.Context, in *account.AccountReq, opts ...grpc.CallOption) (*account.Account, error) {
	out := new(account.Account)
	err := c.cc.Invoke(ctx, AccountAPI_DisableAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountAPIClient) SuspendAccountService(ctx context.Context, in *account.AccountReq, opts ...grpc.CallOption) (*account.Account, error) {
	out := new(account.Account)
	err := c.cc.Invoke(ctx, AccountAPI_SuspendAccountService_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountAPIClient) UnsuspendAccountService(ctx context.Context, in *account.AccountReq, opts ...grpc.CallOption) (*account.Account, error) {
	out := new(account.Account)
	err := c.cc.Invoke(ctx, AccountAPI_UnsuspendAccountService_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountAPIClient) CancelAccountService(ctx context.Context, in *account.AccountReq, opts ...grpc.CallOption) (*account.Account, error) {
	out := new(account.Account)
	err := c.cc.Invoke(ctx, AccountAPI_CancelAccountService_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountAPIServer is the server API for AccountAPI service.
// All implementations must embed UnimplementedAccountAPIServer
// for forward compatibility
type AccountAPIServer interface {
	// account
	NewAccount(context.Context, *account.NewAccountRequest) (*account.NewAccountResponse, error)
	ListAccounts(context.Context, *account.ListAccountsRequest) (*account.Accounts, error)
	GetAccount(context.Context, *account.AccountReq) (*account.Account, error)
	//	rpc UpdateAccount(account.Account) returns (account.Account) {
	//	  option (google.api.http) = {
	//	    post: "/api/v1/accounts/{accountID}"
	//	    body: "*"
	//	  };
	//	}
	CancelAccount(context.Context, *account.AccountReq) (*status.StatusResponse, error)
	DeleteAccount(context.Context, *account.AccountReq) (*status.StatusResponse, error)
	SetAccountMetadata(context.Context, *account.SetAccountMetadataRequest) (*account.Account, error)
	SetAccountIntegrations(context.Context, *account.SetAccountIntegrationsRequest) (*account.Account, error)
	EnableAccount(context.Context, *account.AccountReq) (*account.Account, error)
	DisableAccount(context.Context, *account.AccountReq) (*account.Account, error)
	// rpc SetAccountService(account.Account) returns (account.Account) {}
	SuspendAccountService(context.Context, *account.AccountReq) (*account.Account, error)
	UnsuspendAccountService(context.Context, *account.AccountReq) (*account.Account, error)
	CancelAccountService(context.Context, *account.AccountReq) (*account.Account, error)
	mustEmbedUnimplementedAccountAPIServer()
}

// UnimplementedAccountAPIServer must be embedded to have forward compatible implementations.
type UnimplementedAccountAPIServer struct {
}

func (UnimplementedAccountAPIServer) NewAccount(context.Context, *account.NewAccountRequest) (*account.NewAccountResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method NewAccount not implemented")
}
func (UnimplementedAccountAPIServer) ListAccounts(context.Context, *account.ListAccountsRequest) (*account.Accounts, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method ListAccounts not implemented")
}
func (UnimplementedAccountAPIServer) GetAccount(context.Context, *account.AccountReq) (*account.Account, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method GetAccount not implemented")
}
func (UnimplementedAccountAPIServer) CancelAccount(context.Context, *account.AccountReq) (*status.StatusResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method CancelAccount not implemented")
}
func (UnimplementedAccountAPIServer) DeleteAccount(context.Context, *account.AccountReq) (*status.StatusResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method DeleteAccount not implemented")
}
func (UnimplementedAccountAPIServer) SetAccountMetadata(context.Context, *account.SetAccountMetadataRequest) (*account.Account, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method SetAccountMetadata not implemented")
}
func (UnimplementedAccountAPIServer) SetAccountIntegrations(context.Context, *account.SetAccountIntegrationsRequest) (*account.Account, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method SetAccountIntegrations not implemented")
}
func (UnimplementedAccountAPIServer) EnableAccount(context.Context, *account.AccountReq) (*account.Account, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method EnableAccount not implemented")
}
func (UnimplementedAccountAPIServer) DisableAccount(context.Context, *account.AccountReq) (*account.Account, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method DisableAccount not implemented")
}
func (UnimplementedAccountAPIServer) SuspendAccountService(context.Context, *account.AccountReq) (*account.Account, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method SuspendAccountService not implemented")
}
func (UnimplementedAccountAPIServer) UnsuspendAccountService(context.Context, *account.AccountReq) (*account.Account, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method UnsuspendAccountService not implemented")
}
func (UnimplementedAccountAPIServer) CancelAccountService(context.Context, *account.AccountReq) (*account.Account, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method CancelAccountService not implemented")
}
func (UnimplementedAccountAPIServer) mustEmbedUnimplementedAccountAPIServer() {}

// UnsafeAccountAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccountAPIServer will
// result in compilation errors.
type UnsafeAccountAPIServer interface {
	mustEmbedUnimplementedAccountAPIServer()
}

func RegisterAccountAPIServer(s grpc.ServiceRegistrar, srv AccountAPIServer) {
	s.RegisterService(&AccountAPI_ServiceDesc, srv)
}

func _AccountAPI_NewAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(account.NewAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountAPIServer).NewAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountAPI_NewAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountAPIServer).NewAccount(ctx, req.(*account.NewAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountAPI_ListAccounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(account.ListAccountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountAPIServer).ListAccounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountAPI_ListAccounts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountAPIServer).ListAccounts(ctx, req.(*account.ListAccountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountAPI_GetAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(account.AccountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountAPIServer).GetAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountAPI_GetAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountAPIServer).GetAccount(ctx, req.(*account.AccountReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountAPI_CancelAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(account.AccountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountAPIServer).CancelAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountAPI_CancelAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountAPIServer).CancelAccount(ctx, req.(*account.AccountReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountAPI_DeleteAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(account.AccountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountAPIServer).DeleteAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountAPI_DeleteAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountAPIServer).DeleteAccount(ctx, req.(*account.AccountReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountAPI_SetAccountMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(account.SetAccountMetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountAPIServer).SetAccountMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountAPI_SetAccountMetadata_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountAPIServer).SetAccountMetadata(ctx, req.(*account.SetAccountMetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountAPI_SetAccountIntegrations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(account.SetAccountIntegrationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountAPIServer).SetAccountIntegrations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountAPI_SetAccountIntegrations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountAPIServer).SetAccountIntegrations(ctx, req.(*account.SetAccountIntegrationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountAPI_EnableAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(account.AccountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountAPIServer).EnableAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountAPI_EnableAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountAPIServer).EnableAccount(ctx, req.(*account.AccountReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountAPI_DisableAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(account.AccountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountAPIServer).DisableAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountAPI_DisableAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountAPIServer).DisableAccount(ctx, req.(*account.AccountReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountAPI_SuspendAccountService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(account.AccountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountAPIServer).SuspendAccountService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountAPI_SuspendAccountService_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountAPIServer).SuspendAccountService(ctx, req.(*account.AccountReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountAPI_UnsuspendAccountService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(account.AccountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountAPIServer).UnsuspendAccountService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountAPI_UnsuspendAccountService_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountAPIServer).UnsuspendAccountService(ctx, req.(*account.AccountReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountAPI_CancelAccountService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(account.AccountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountAPIServer).CancelAccountService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountAPI_CancelAccountService_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountAPIServer).CancelAccountService(ctx, req.(*account.AccountReq))
	}
	return interceptor(ctx, in, info, handler)
}

// AccountAPI_ServiceDesc is the grpc.ServiceDesc for AccountAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AccountAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.AccountAPI",
	HandlerType: (*AccountAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewAccount",
			Handler:    _AccountAPI_NewAccount_Handler,
		},
		{
			MethodName: "ListAccounts",
			Handler:    _AccountAPI_ListAccounts_Handler,
		},
		{
			MethodName: "GetAccount",
			Handler:    _AccountAPI_GetAccount_Handler,
		},
		{
			MethodName: "CancelAccount",
			Handler:    _AccountAPI_CancelAccount_Handler,
		},
		{
			MethodName: "DeleteAccount",
			Handler:    _AccountAPI_DeleteAccount_Handler,
		},
		{
			MethodName: "SetAccountMetadata",
			Handler:    _AccountAPI_SetAccountMetadata_Handler,
		},
		{
			MethodName: "SetAccountIntegrations",
			Handler:    _AccountAPI_SetAccountIntegrations_Handler,
		},
		{
			MethodName: "EnableAccount",
			Handler:    _AccountAPI_EnableAccount_Handler,
		},
		{
			MethodName: "DisableAccount",
			Handler:    _AccountAPI_DisableAccount_Handler,
		},
		{
			MethodName: "SuspendAccountService",
			Handler:    _AccountAPI_SuspendAccountService_Handler,
		},
		{
			MethodName: "UnsuspendAccountService",
			Handler:    _AccountAPI_UnsuspendAccountService_Handler,
		},
		{
			MethodName: "CancelAccountService",
			Handler:    _AccountAPI_CancelAccountService_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mmesh/protobuf/rpc/v1/accountAPI.proto",
}
