// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: mmesh/protobuf/resources/v1/billing/invoice.proto

package billing

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type Invoice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountID            string                  `protobuf:"bytes,1,opt,name=accountID,proto3" json:"accountID,omitempty"`
	InvoiceID            string                  `protobuf:"bytes,2,opt,name=invoiceID,proto3" json:"invoiceID,omitempty"`
	Description          string                  `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	ServiceID            string                  `protobuf:"bytes,5,opt,name=serviceID,proto3" json:"serviceID,omitempty"`
	PricingPlanID        string                  `protobuf:"bytes,6,opt,name=pricingPlanID,proto3" json:"pricingPlanID,omitempty"`
	CustomerEmail        string                  `protobuf:"bytes,8,opt,name=customerEmail,proto3" json:"customerEmail,omitempty"`
	StripeCustomerID     string                  `protobuf:"bytes,11,opt,name=stripeCustomerID,proto3" json:"stripeCustomerID,omitempty"`
	StripeInvoiceID      string                  `protobuf:"bytes,12,opt,name=stripeInvoiceID,proto3" json:"stripeInvoiceID,omitempty"`
	StripeSubscriptionID string                  `protobuf:"bytes,13,opt,name=stripeSubscriptionID,proto3" json:"stripeSubscriptionID,omitempty"`
	AutoAdvance          bool                    `protobuf:"varint,14,opt,name=autoAdvance,proto3" json:"autoAdvance,omitempty"`
	AmountDue            int64                   `protobuf:"varint,21,opt,name=amountDue,proto3" json:"amountDue,omitempty"`
	AmountPaid           int64                   `protobuf:"varint,22,opt,name=amountPaid,proto3" json:"amountPaid,omitempty"`
	AmountRemaining      int64                   `protobuf:"varint,23,opt,name=amountRemaining,proto3" json:"amountRemaining,omitempty"`
	StartingBalance      int64                   `protobuf:"varint,25,opt,name=startingBalance,proto3" json:"startingBalance,omitempty"`
	EndingBalance        int64                   `protobuf:"varint,26,opt,name=endingBalance,proto3" json:"endingBalance,omitempty"`
	Attempted            bool                    `protobuf:"varint,31,opt,name=attempted,proto3" json:"attempted,omitempty"`
	Paid                 bool                    `protobuf:"varint,32,opt,name=paid,proto3" json:"paid,omitempty"`
	Deleted              bool                    `protobuf:"varint,33,opt,name=deleted,proto3" json:"deleted,omitempty"`
	AttemptCount         int64                   `protobuf:"varint,36,opt,name=attemptCount,proto3" json:"attemptCount,omitempty"`
	NextPaymentAttempt   int64                   `protobuf:"varint,37,opt,name=nextPaymentAttempt,proto3" json:"nextPaymentAttempt,omitempty"`
	Subtotal             int64                   `protobuf:"varint,41,opt,name=subtotal,proto3" json:"subtotal,omitempty"`
	Tax                  int64                   `protobuf:"varint,42,opt,name=tax,proto3" json:"tax,omitempty"`
	Total                int64                   `protobuf:"varint,43,opt,name=total,proto3" json:"total,omitempty"`
	TotalDiscount        int64                   `protobuf:"varint,48,opt,name=totalDiscount,proto3" json:"totalDiscount,omitempty"`
	Currency             string                  `protobuf:"bytes,51,opt,name=currency,proto3" json:"currency,omitempty"`
	InvoiceNumber        string                  `protobuf:"bytes,52,opt,name=invoiceNumber,proto3" json:"invoiceNumber,omitempty"`
	ReceiptNumber        string                  `protobuf:"bytes,53,opt,name=receiptNumber,proto3" json:"receiptNumber,omitempty"`
	StatementDescriptor  string                  `protobuf:"bytes,54,opt,name=statementDescriptor,proto3" json:"statementDescriptor,omitempty"`
	HostedInvoiceURL     string                  `protobuf:"bytes,55,opt,name=hostedInvoiceURL,proto3" json:"hostedInvoiceURL,omitempty"`
	InvoicePDF           string                  `protobuf:"bytes,56,opt,name=invoicePDF,proto3" json:"invoicePDF,omitempty"`
	CreationDate         int64                   `protobuf:"varint,61,opt,name=creationDate,proto3" json:"creationDate,omitempty"`
	LastModified         int64                   `protobuf:"varint,62,opt,name=lastModified,proto3" json:"lastModified,omitempty"`
	Period               string                  `protobuf:"bytes,71,opt,name=period,proto3" json:"period,omitempty"`
	PeriodStart          int64                   `protobuf:"varint,72,opt,name=periodStart,proto3" json:"periodStart,omitempty"`
	PeriodEnd            int64                   `protobuf:"varint,73,opt,name=periodEnd,proto3" json:"periodEnd,omitempty"`
	DueDate              int64                   `protobuf:"varint,74,opt,name=dueDate,proto3" json:"dueDate,omitempty"`
	TrialEndDate         int64                   `protobuf:"varint,78,opt,name=trialEndDate,proto3" json:"trialEndDate,omitempty"`
	TrialPeriod          bool                    `protobuf:"varint,79,opt,name=trialPeriod,proto3" json:"trialPeriod,omitempty"`
	Status               string                  `protobuf:"bytes,81,opt,name=status,proto3" json:"status,omitempty"`                           // stripe invoice.status
	PaymentIntentStatus  string                  `protobuf:"bytes,85,opt,name=paymentIntentStatus,proto3" json:"paymentIntentStatus,omitempty"` // stripe payment_intent.status
	SubscriptionInvoice  bool                    `protobuf:"varint,88,opt,name=subscriptionInvoice,proto3" json:"subscriptionInvoice,omitempty"`
	SummaryItems         []*InvoiceItem          `protobuf:"bytes,91,rep,name=summaryItems,proto3" json:"summaryItems,omitempty"`
	Items                map[string]*InvoiceItem `protobuf:"bytes,92,rep,name=items,proto3" json:"items,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // map[itemID]*InvoiceItem
}

func (x *Invoice) Reset() {
	*x = Invoice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_billing_invoice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Invoice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Invoice) ProtoMessage() {}

func (x *Invoice) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_billing_invoice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Invoice.ProtoReflect.Descriptor instead.
func (*Invoice) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_billing_invoice_proto_rawDescGZIP(), []int{0}
}

func (x *Invoice) GetAccountID() string {
	if x != nil {
		return x.AccountID
	}
	return ""
}

func (x *Invoice) GetInvoiceID() string {
	if x != nil {
		return x.InvoiceID
	}
	return ""
}

func (x *Invoice) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Invoice) GetServiceID() string {
	if x != nil {
		return x.ServiceID
	}
	return ""
}

func (x *Invoice) GetPricingPlanID() string {
	if x != nil {
		return x.PricingPlanID
	}
	return ""
}

func (x *Invoice) GetCustomerEmail() string {
	if x != nil {
		return x.CustomerEmail
	}
	return ""
}

func (x *Invoice) GetStripeCustomerID() string {
	if x != nil {
		return x.StripeCustomerID
	}
	return ""
}

func (x *Invoice) GetStripeInvoiceID() string {
	if x != nil {
		return x.StripeInvoiceID
	}
	return ""
}

func (x *Invoice) GetStripeSubscriptionID() string {
	if x != nil {
		return x.StripeSubscriptionID
	}
	return ""
}

func (x *Invoice) GetAutoAdvance() bool {
	if x != nil {
		return x.AutoAdvance
	}
	return false
}

func (x *Invoice) GetAmountDue() int64 {
	if x != nil {
		return x.AmountDue
	}
	return 0
}

func (x *Invoice) GetAmountPaid() int64 {
	if x != nil {
		return x.AmountPaid
	}
	return 0
}

func (x *Invoice) GetAmountRemaining() int64 {
	if x != nil {
		return x.AmountRemaining
	}
	return 0
}

func (x *Invoice) GetStartingBalance() int64 {
	if x != nil {
		return x.StartingBalance
	}
	return 0
}

func (x *Invoice) GetEndingBalance() int64 {
	if x != nil {
		return x.EndingBalance
	}
	return 0
}

func (x *Invoice) GetAttempted() bool {
	if x != nil {
		return x.Attempted
	}
	return false
}

func (x *Invoice) GetPaid() bool {
	if x != nil {
		return x.Paid
	}
	return false
}

func (x *Invoice) GetDeleted() bool {
	if x != nil {
		return x.Deleted
	}
	return false
}

func (x *Invoice) GetAttemptCount() int64 {
	if x != nil {
		return x.AttemptCount
	}
	return 0
}

func (x *Invoice) GetNextPaymentAttempt() int64 {
	if x != nil {
		return x.NextPaymentAttempt
	}
	return 0
}

func (x *Invoice) GetSubtotal() int64 {
	if x != nil {
		return x.Subtotal
	}
	return 0
}

func (x *Invoice) GetTax() int64 {
	if x != nil {
		return x.Tax
	}
	return 0
}

func (x *Invoice) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *Invoice) GetTotalDiscount() int64 {
	if x != nil {
		return x.TotalDiscount
	}
	return 0
}

func (x *Invoice) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *Invoice) GetInvoiceNumber() string {
	if x != nil {
		return x.InvoiceNumber
	}
	return ""
}

func (x *Invoice) GetReceiptNumber() string {
	if x != nil {
		return x.ReceiptNumber
	}
	return ""
}

func (x *Invoice) GetStatementDescriptor() string {
	if x != nil {
		return x.StatementDescriptor
	}
	return ""
}

func (x *Invoice) GetHostedInvoiceURL() string {
	if x != nil {
		return x.HostedInvoiceURL
	}
	return ""
}

func (x *Invoice) GetInvoicePDF() string {
	if x != nil {
		return x.InvoicePDF
	}
	return ""
}

func (x *Invoice) GetCreationDate() int64 {
	if x != nil {
		return x.CreationDate
	}
	return 0
}

func (x *Invoice) GetLastModified() int64 {
	if x != nil {
		return x.LastModified
	}
	return 0
}

func (x *Invoice) GetPeriod() string {
	if x != nil {
		return x.Period
	}
	return ""
}

func (x *Invoice) GetPeriodStart() int64 {
	if x != nil {
		return x.PeriodStart
	}
	return 0
}

func (x *Invoice) GetPeriodEnd() int64 {
	if x != nil {
		return x.PeriodEnd
	}
	return 0
}

func (x *Invoice) GetDueDate() int64 {
	if x != nil {
		return x.DueDate
	}
	return 0
}

func (x *Invoice) GetTrialEndDate() int64 {
	if x != nil {
		return x.TrialEndDate
	}
	return 0
}

func (x *Invoice) GetTrialPeriod() bool {
	if x != nil {
		return x.TrialPeriod
	}
	return false
}

func (x *Invoice) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Invoice) GetPaymentIntentStatus() string {
	if x != nil {
		return x.PaymentIntentStatus
	}
	return ""
}

func (x *Invoice) GetSubscriptionInvoice() bool {
	if x != nil {
		return x.SubscriptionInvoice
	}
	return false
}

func (x *Invoice) GetSummaryItems() []*InvoiceItem {
	if x != nil {
		return x.SummaryItems
	}
	return nil
}

func (x *Invoice) GetItems() map[string]*InvoiceItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type Invoices struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta     *resource.ListResponse `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	Invoices []*Invoice             `protobuf:"bytes,2,rep,name=invoices,proto3" json:"invoices,omitempty"`
}

func (x *Invoices) Reset() {
	*x = Invoices{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_billing_invoice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Invoices) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Invoices) ProtoMessage() {}

func (x *Invoices) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_billing_invoice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Invoices.ProtoReflect.Descriptor instead.
func (*Invoices) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_billing_invoice_proto_rawDescGZIP(), []int{1}
}

func (x *Invoices) GetMeta() *resource.ListResponse {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *Invoices) GetInvoices() []*Invoice {
	if x != nil {
		return x.Invoices
	}
	return nil
}

type ListInvoicesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta      *resource.ListRequest `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	AccountID string                `protobuf:"bytes,2,opt,name=accountID,proto3" json:"accountID,omitempty"`
}

func (x *ListInvoicesRequest) Reset() {
	*x = ListInvoicesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_billing_invoice_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListInvoicesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListInvoicesRequest) ProtoMessage() {}

func (x *ListInvoicesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_billing_invoice_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListInvoicesRequest.ProtoReflect.Descriptor instead.
func (*ListInvoicesRequest) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_billing_invoice_proto_rawDescGZIP(), []int{2}
}

func (x *ListInvoicesRequest) GetMeta() *resource.ListRequest {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *ListInvoicesRequest) GetAccountID() string {
	if x != nil {
		return x.AccountID
	}
	return ""
}

var File_mmesh_protobuf_resources_v1_billing_invoice_proto protoreflect.FileDescriptor

var file_mmesh_protobuf_resources_v1_billing_invoice_proto_rawDesc = []byte{
	0x0a, 0x31, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x69,
	0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x07, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x1a, 0x2f, 0x6d, 0x6d,
	0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x6d,
	0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69,
	0x6e, 0x67, 0x2f, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xde, 0x0c,
	0x0a, 0x07, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6e, 0x76, 0x6f, 0x69,
	0x63, 0x65, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6e, 0x76, 0x6f,
	0x69, 0x63, 0x65, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x49, 0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x49, 0x44, 0x12, 0x24, 0x0a, 0x0d, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67,
	0x50, 0x6c, 0x61, 0x6e, 0x49, 0x44, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x72,
	0x69, 0x63, 0x69, 0x6e, 0x67, 0x50, 0x6c, 0x61, 0x6e, 0x49, 0x44, 0x12, 0x24, 0x0a, 0x0d, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x12, 0x2a, 0x0a, 0x10, 0x73, 0x74, 0x72, 0x69, 0x70, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x49, 0x44, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x73, 0x74, 0x72,
	0x69, 0x70, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x44, 0x12, 0x28, 0x0a,
	0x0f, 0x73, 0x74, 0x72, 0x69, 0x70, 0x65, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x49, 0x44,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x74, 0x72, 0x69, 0x70, 0x65, 0x49, 0x6e,
	0x76, 0x6f, 0x69, 0x63, 0x65, 0x49, 0x44, 0x12, 0x32, 0x0a, 0x14, 0x73, 0x74, 0x72, 0x69, 0x70,
	0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x73, 0x74, 0x72, 0x69, 0x70, 0x65, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x61,
	0x75, 0x74, 0x6f, 0x41, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0b, 0x61, 0x75, 0x74, 0x6f, 0x41, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x75, 0x65, 0x18, 0x15, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x75, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x61, 0x69, 0x64, 0x18, 0x16, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x61, 0x69, 0x64, 0x12, 0x28, 0x0a, 0x0f, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x18, 0x17,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x6d, 0x61,
	0x69, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x28, 0x0a, 0x0f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x69, 0x6e,
	0x67, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x19, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12,
	0x24, 0x0a, 0x0d, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x18, 0x1a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x42, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74,
	0x65, 0x64, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x61, 0x74, 0x74, 0x65, 0x6d, 0x70,
	0x74, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x69, 0x64, 0x18, 0x20, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x04, 0x70, 0x61, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x64, 0x18, 0x21, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x24, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x61, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2e, 0x0a, 0x12, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x41, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x18, 0x25, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x12, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x74,
	0x74, 0x65, 0x6d, 0x70, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x75, 0x62, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x18, 0x29, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x75, 0x62, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x78, 0x18, 0x2a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03,
	0x74, 0x61, 0x78, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x2b, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x24, 0x0a, 0x0d, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x30, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x33, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x24, 0x0a, 0x0d, 0x69,
	0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x34, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x24, 0x0a, 0x0d, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x35, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70,
	0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x13, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x18, 0x36,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x12, 0x2a, 0x0a, 0x10, 0x68, 0x6f, 0x73,
	0x74, 0x65, 0x64, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x55, 0x52, 0x4c, 0x18, 0x37, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x10, 0x68, 0x6f, 0x73, 0x74, 0x65, 0x64, 0x49, 0x6e, 0x76, 0x6f, 0x69,
	0x63, 0x65, 0x55, 0x52, 0x4c, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65,
	0x50, 0x44, 0x46, 0x18, 0x38, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e, 0x76, 0x6f, 0x69,
	0x63, 0x65, 0x50, 0x44, 0x46, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x44, 0x61, 0x74, 0x65, 0x18, 0x3d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x6c, 0x61, 0x73,
	0x74, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x3e, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0c, 0x6c, 0x61, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x47, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70,
	0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x18, 0x48, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x69,
	0x6f, 0x64, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x65, 0x72, 0x69, 0x6f,
	0x64, 0x45, 0x6e, 0x64, 0x18, 0x49, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x65, 0x72, 0x69,
	0x6f, 0x64, 0x45, 0x6e, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x75, 0x65, 0x44, 0x61, 0x74, 0x65,
	0x18, 0x4a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x64, 0x75, 0x65, 0x44, 0x61, 0x74, 0x65, 0x12,
	0x22, 0x0a, 0x0c, 0x74, 0x72, 0x69, 0x61, 0x6c, 0x45, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x18,
	0x4e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x74, 0x72, 0x69, 0x61, 0x6c, 0x45, 0x6e, 0x64, 0x44,
	0x61, 0x74, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x72, 0x69, 0x61, 0x6c, 0x50, 0x65, 0x72, 0x69,
	0x6f, 0x64, 0x18, 0x4f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x74, 0x72, 0x69, 0x61, 0x6c, 0x50,
	0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x51, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x30, 0x0a,
	0x13, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x55, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x30, 0x0a, 0x13, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x18, 0x58, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x73, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63,
	0x65, 0x12, 0x38, 0x0a, 0x0c, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x49, 0x74, 0x65, 0x6d,
	0x73, 0x18, 0x5b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e,
	0x67, 0x2e, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x0c, 0x73,
	0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x31, 0x0a, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x18, 0x5c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x62, 0x69, 0x6c,
	0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x74, 0x65,
	0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x1a, 0x4e,
	0x0a, 0x0a, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x49,
	0x74, 0x65, 0x6d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x64,
	0x0a, 0x08, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x12, 0x2a, 0x0a, 0x04, 0x6d, 0x65,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x2c, 0x0a, 0x08, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69,
	0x6e, 0x67, 0x2e, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x08, 0x69, 0x6e, 0x76, 0x6f,
	0x69, 0x63, 0x65, 0x73, 0x22, 0x5e, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x76, 0x6f,
	0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x04, 0x6d,
	0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x49, 0x44, 0x42, 0x2b, 0x5a, 0x29, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x64, 0x65,
	0x76, 0x2f, 0x6d, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e,
	0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mmesh_protobuf_resources_v1_billing_invoice_proto_rawDescOnce sync.Once
	file_mmesh_protobuf_resources_v1_billing_invoice_proto_rawDescData = file_mmesh_protobuf_resources_v1_billing_invoice_proto_rawDesc
)

func file_mmesh_protobuf_resources_v1_billing_invoice_proto_rawDescGZIP() []byte {
	file_mmesh_protobuf_resources_v1_billing_invoice_proto_rawDescOnce.Do(func() {
		file_mmesh_protobuf_resources_v1_billing_invoice_proto_rawDescData = protoimpl.X.CompressGZIP(file_mmesh_protobuf_resources_v1_billing_invoice_proto_rawDescData)
	})
	return file_mmesh_protobuf_resources_v1_billing_invoice_proto_rawDescData
}

var file_mmesh_protobuf_resources_v1_billing_invoice_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_mmesh_protobuf_resources_v1_billing_invoice_proto_goTypes = []interface{}{
	(*Invoice)(nil),               // 0: billing.Invoice
	(*Invoices)(nil),              // 1: billing.Invoices
	(*ListInvoicesRequest)(nil),   // 2: billing.ListInvoicesRequest
	nil,                           // 3: billing.Invoice.ItemsEntry
	(*InvoiceItem)(nil),           // 4: billing.InvoiceItem
	(*resource.ListResponse)(nil), // 5: resource.ListResponse
	(*resource.ListRequest)(nil),  // 6: resource.ListRequest
}
var file_mmesh_protobuf_resources_v1_billing_invoice_proto_depIdxs = []int32{
	4, // 0: billing.Invoice.summaryItems:type_name -> billing.InvoiceItem
	3, // 1: billing.Invoice.items:type_name -> billing.Invoice.ItemsEntry
	5, // 2: billing.Invoices.meta:type_name -> resource.ListResponse
	0, // 3: billing.Invoices.invoices:type_name -> billing.Invoice
	6, // 4: billing.ListInvoicesRequest.meta:type_name -> resource.ListRequest
	4, // 5: billing.Invoice.ItemsEntry.value:type_name -> billing.InvoiceItem
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_mmesh_protobuf_resources_v1_billing_invoice_proto_init() }
func file_mmesh_protobuf_resources_v1_billing_invoice_proto_init() {
	if File_mmesh_protobuf_resources_v1_billing_invoice_proto != nil {
		return
	}
	file_mmesh_protobuf_resources_v1_billing_item_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_mmesh_protobuf_resources_v1_billing_invoice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Invoice); i {
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
		file_mmesh_protobuf_resources_v1_billing_invoice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Invoices); i {
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
		file_mmesh_protobuf_resources_v1_billing_invoice_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListInvoicesRequest); i {
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
			RawDescriptor: file_mmesh_protobuf_resources_v1_billing_invoice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mmesh_protobuf_resources_v1_billing_invoice_proto_goTypes,
		DependencyIndexes: file_mmesh_protobuf_resources_v1_billing_invoice_proto_depIdxs,
		MessageInfos:      file_mmesh_protobuf_resources_v1_billing_invoice_proto_msgTypes,
	}.Build()
	File_mmesh_protobuf_resources_v1_billing_invoice_proto = out.File
	file_mmesh_protobuf_resources_v1_billing_invoice_proto_rawDesc = nil
	file_mmesh_protobuf_resources_v1_billing_invoice_proto_goTypes = nil
	file_mmesh_protobuf_resources_v1_billing_invoice_proto_depIdxs = nil
}
