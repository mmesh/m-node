// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.2
// source: mmesh/protobuf/rpc/v1/billingAPI.proto

package rpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status1 "google.golang.org/grpc/status"
	status "mmesh.dev/m-api-go/grpc/common/status"
	billing "mmesh.dev/m-api-go/grpc/resources/billing"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	BillingAPI_SearchCustomer_FullMethodName    = "/api.BillingAPI/SearchCustomer"
	BillingAPI_GetCustomerPortal_FullMethodName = "/api.BillingAPI/GetCustomerPortal"
	BillingAPI_ListInvoices_FullMethodName      = "/api.BillingAPI/ListInvoices"
	BillingAPI_GetInvoice_FullMethodName        = "/api.BillingAPI/GetInvoice"
	BillingAPI_DeleteInvoice_FullMethodName     = "/api.BillingAPI/DeleteInvoice"
	BillingAPI_ListBilledItems_FullMethodName   = "/api.BillingAPI/ListBilledItems"
	BillingAPI_GetBilledItem_FullMethodName     = "/api.BillingAPI/GetBilledItem"
	BillingAPI_ApplyPromotion_FullMethodName    = "/api.BillingAPI/ApplyPromotion"
)

// BillingAPIClient is the client API for BillingAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BillingAPIClient interface {
	// customer
	SearchCustomer(ctx context.Context, in *billing.Customer, opts ...grpc.CallOption) (*status.SearchResponse, error)
	GetCustomerPortal(ctx context.Context, in *billing.CustomerPortalRequest, opts ...grpc.CallOption) (*billing.CustomerPortalResponse, error)
	// invoice
	//
	//	rpc NewInvoice(billing.Invoice) returns (billing.Invoice) {
	//	  option (google.api.http) = {
	//	    post: "/api/v1/accounts/{accountID}/billing/invoices"
	//	    body: "*"
	//	  };
	//	}
	ListInvoices(ctx context.Context, in *billing.ListInvoicesRequest, opts ...grpc.CallOption) (*billing.Invoices, error)
	GetInvoice(ctx context.Context, in *billing.Invoice, opts ...grpc.CallOption) (*billing.Invoice, error)
	//	rpc SetInvoice(billing.Invoice) returns (billing.Invoice) {
	//	  option (google.api.http) = {
	//	    post: "/api/v1/accounts/{accountID}/billing/invoices/{invoiceID}"
	//	    body: "*"
	//	  };
	//	}
	DeleteInvoice(ctx context.Context, in *billing.Invoice, opts ...grpc.CallOption) (*status.StatusResponse, error)
	// billed-item
	//
	//	rpc NewBilledItem(billing.InvoiceItem) returns (billing.InvoiceItem) {
	//	  option (google.api.http) = {
	//	    post: "/api/v1/accounts/{accountID}/billing/items"
	//	    body: "*"
	//	  };
	//	}
	ListBilledItems(ctx context.Context, in *billing.ListInvoiceItemsRequest, opts ...grpc.CallOption) (*billing.InvoiceItems, error)
	GetBilledItem(ctx context.Context, in *billing.InvoiceItem, opts ...grpc.CallOption) (*billing.InvoiceItem, error)
	// subscription
	ApplyPromotion(ctx context.Context, in *billing.ApplyPromotionRequest, opts ...grpc.CallOption) (*status.StatusResponse, error)
}

type billingAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewBillingAPIClient(cc grpc.ClientConnInterface) BillingAPIClient {
	return &billingAPIClient{cc}
}

func (c *billingAPIClient) SearchCustomer(ctx context.Context, in *billing.Customer, opts ...grpc.CallOption) (*status.SearchResponse, error) {
	out := new(status.SearchResponse)
	err := c.cc.Invoke(ctx, BillingAPI_SearchCustomer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *billingAPIClient) GetCustomerPortal(ctx context.Context, in *billing.CustomerPortalRequest, opts ...grpc.CallOption) (*billing.CustomerPortalResponse, error) {
	out := new(billing.CustomerPortalResponse)
	err := c.cc.Invoke(ctx, BillingAPI_GetCustomerPortal_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *billingAPIClient) ListInvoices(ctx context.Context, in *billing.ListInvoicesRequest, opts ...grpc.CallOption) (*billing.Invoices, error) {
	out := new(billing.Invoices)
	err := c.cc.Invoke(ctx, BillingAPI_ListInvoices_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *billingAPIClient) GetInvoice(ctx context.Context, in *billing.Invoice, opts ...grpc.CallOption) (*billing.Invoice, error) {
	out := new(billing.Invoice)
	err := c.cc.Invoke(ctx, BillingAPI_GetInvoice_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *billingAPIClient) DeleteInvoice(ctx context.Context, in *billing.Invoice, opts ...grpc.CallOption) (*status.StatusResponse, error) {
	out := new(status.StatusResponse)
	err := c.cc.Invoke(ctx, BillingAPI_DeleteInvoice_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *billingAPIClient) ListBilledItems(ctx context.Context, in *billing.ListInvoiceItemsRequest, opts ...grpc.CallOption) (*billing.InvoiceItems, error) {
	out := new(billing.InvoiceItems)
	err := c.cc.Invoke(ctx, BillingAPI_ListBilledItems_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *billingAPIClient) GetBilledItem(ctx context.Context, in *billing.InvoiceItem, opts ...grpc.CallOption) (*billing.InvoiceItem, error) {
	out := new(billing.InvoiceItem)
	err := c.cc.Invoke(ctx, BillingAPI_GetBilledItem_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *billingAPIClient) ApplyPromotion(ctx context.Context, in *billing.ApplyPromotionRequest, opts ...grpc.CallOption) (*status.StatusResponse, error) {
	out := new(status.StatusResponse)
	err := c.cc.Invoke(ctx, BillingAPI_ApplyPromotion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BillingAPIServer is the server API for BillingAPI service.
// All implementations must embed UnimplementedBillingAPIServer
// for forward compatibility
type BillingAPIServer interface {
	// customer
	SearchCustomer(context.Context, *billing.Customer) (*status.SearchResponse, error)
	GetCustomerPortal(context.Context, *billing.CustomerPortalRequest) (*billing.CustomerPortalResponse, error)
	// invoice
	//
	//	rpc NewInvoice(billing.Invoice) returns (billing.Invoice) {
	//	  option (google.api.http) = {
	//	    post: "/api/v1/accounts/{accountID}/billing/invoices"
	//	    body: "*"
	//	  };
	//	}
	ListInvoices(context.Context, *billing.ListInvoicesRequest) (*billing.Invoices, error)
	GetInvoice(context.Context, *billing.Invoice) (*billing.Invoice, error)
	//	rpc SetInvoice(billing.Invoice) returns (billing.Invoice) {
	//	  option (google.api.http) = {
	//	    post: "/api/v1/accounts/{accountID}/billing/invoices/{invoiceID}"
	//	    body: "*"
	//	  };
	//	}
	DeleteInvoice(context.Context, *billing.Invoice) (*status.StatusResponse, error)
	// billed-item
	//
	//	rpc NewBilledItem(billing.InvoiceItem) returns (billing.InvoiceItem) {
	//	  option (google.api.http) = {
	//	    post: "/api/v1/accounts/{accountID}/billing/items"
	//	    body: "*"
	//	  };
	//	}
	ListBilledItems(context.Context, *billing.ListInvoiceItemsRequest) (*billing.InvoiceItems, error)
	GetBilledItem(context.Context, *billing.InvoiceItem) (*billing.InvoiceItem, error)
	// subscription
	ApplyPromotion(context.Context, *billing.ApplyPromotionRequest) (*status.StatusResponse, error)
	mustEmbedUnimplementedBillingAPIServer()
}

// UnimplementedBillingAPIServer must be embedded to have forward compatible implementations.
type UnimplementedBillingAPIServer struct {
}

func (UnimplementedBillingAPIServer) SearchCustomer(context.Context, *billing.Customer) (*status.SearchResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method SearchCustomer not implemented")
}
func (UnimplementedBillingAPIServer) GetCustomerPortal(context.Context, *billing.CustomerPortalRequest) (*billing.CustomerPortalResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method GetCustomerPortal not implemented")
}
func (UnimplementedBillingAPIServer) ListInvoices(context.Context, *billing.ListInvoicesRequest) (*billing.Invoices, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method ListInvoices not implemented")
}
func (UnimplementedBillingAPIServer) GetInvoice(context.Context, *billing.Invoice) (*billing.Invoice, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method GetInvoice not implemented")
}
func (UnimplementedBillingAPIServer) DeleteInvoice(context.Context, *billing.Invoice) (*status.StatusResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method DeleteInvoice not implemented")
}
func (UnimplementedBillingAPIServer) ListBilledItems(context.Context, *billing.ListInvoiceItemsRequest) (*billing.InvoiceItems, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method ListBilledItems not implemented")
}
func (UnimplementedBillingAPIServer) GetBilledItem(context.Context, *billing.InvoiceItem) (*billing.InvoiceItem, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method GetBilledItem not implemented")
}
func (UnimplementedBillingAPIServer) ApplyPromotion(context.Context, *billing.ApplyPromotionRequest) (*status.StatusResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method ApplyPromotion not implemented")
}
func (UnimplementedBillingAPIServer) mustEmbedUnimplementedBillingAPIServer() {}

// UnsafeBillingAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BillingAPIServer will
// result in compilation errors.
type UnsafeBillingAPIServer interface {
	mustEmbedUnimplementedBillingAPIServer()
}

func RegisterBillingAPIServer(s grpc.ServiceRegistrar, srv BillingAPIServer) {
	s.RegisterService(&BillingAPI_ServiceDesc, srv)
}

func _BillingAPI_SearchCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(billing.Customer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillingAPIServer).SearchCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BillingAPI_SearchCustomer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillingAPIServer).SearchCustomer(ctx, req.(*billing.Customer))
	}
	return interceptor(ctx, in, info, handler)
}

func _BillingAPI_GetCustomerPortal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(billing.CustomerPortalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillingAPIServer).GetCustomerPortal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BillingAPI_GetCustomerPortal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillingAPIServer).GetCustomerPortal(ctx, req.(*billing.CustomerPortalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BillingAPI_ListInvoices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(billing.ListInvoicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillingAPIServer).ListInvoices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BillingAPI_ListInvoices_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillingAPIServer).ListInvoices(ctx, req.(*billing.ListInvoicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BillingAPI_GetInvoice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(billing.Invoice)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillingAPIServer).GetInvoice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BillingAPI_GetInvoice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillingAPIServer).GetInvoice(ctx, req.(*billing.Invoice))
	}
	return interceptor(ctx, in, info, handler)
}

func _BillingAPI_DeleteInvoice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(billing.Invoice)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillingAPIServer).DeleteInvoice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BillingAPI_DeleteInvoice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillingAPIServer).DeleteInvoice(ctx, req.(*billing.Invoice))
	}
	return interceptor(ctx, in, info, handler)
}

func _BillingAPI_ListBilledItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(billing.ListInvoiceItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillingAPIServer).ListBilledItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BillingAPI_ListBilledItems_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillingAPIServer).ListBilledItems(ctx, req.(*billing.ListInvoiceItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BillingAPI_GetBilledItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(billing.InvoiceItem)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillingAPIServer).GetBilledItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BillingAPI_GetBilledItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillingAPIServer).GetBilledItem(ctx, req.(*billing.InvoiceItem))
	}
	return interceptor(ctx, in, info, handler)
}

func _BillingAPI_ApplyPromotion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(billing.ApplyPromotionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillingAPIServer).ApplyPromotion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BillingAPI_ApplyPromotion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillingAPIServer).ApplyPromotion(ctx, req.(*billing.ApplyPromotionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BillingAPI_ServiceDesc is the grpc.ServiceDesc for BillingAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BillingAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.BillingAPI",
	HandlerType: (*BillingAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SearchCustomer",
			Handler:    _BillingAPI_SearchCustomer_Handler,
		},
		{
			MethodName: "GetCustomerPortal",
			Handler:    _BillingAPI_GetCustomerPortal_Handler,
		},
		{
			MethodName: "ListInvoices",
			Handler:    _BillingAPI_ListInvoices_Handler,
		},
		{
			MethodName: "GetInvoice",
			Handler:    _BillingAPI_GetInvoice_Handler,
		},
		{
			MethodName: "DeleteInvoice",
			Handler:    _BillingAPI_DeleteInvoice_Handler,
		},
		{
			MethodName: "ListBilledItems",
			Handler:    _BillingAPI_ListBilledItems_Handler,
		},
		{
			MethodName: "GetBilledItem",
			Handler:    _BillingAPI_GetBilledItem_Handler,
		},
		{
			MethodName: "ApplyPromotion",
			Handler:    _BillingAPI_ApplyPromotion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mmesh/protobuf/rpc/v1/billingAPI.proto",
}
