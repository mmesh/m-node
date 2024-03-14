// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.0--rc3
// source: mmesh/protobuf/resources/v1/nstore/hsecdb/hsecdb.proto

package hsecdb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	nstore "mmesh.dev/m-api-go/grpc/resources/nstore"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ReportQueryType int32

const (
	ReportQueryType_REPORT_UNSPECIFIED     ReportQueryType = 0
	ReportQueryType_REPORT_OS_PKGS         ReportQueryType = 11
	ReportQueryType_REPORT_VULNERABILITIES ReportQueryType = 21
	ReportQueryType_REPORT_MISCONFIGS      ReportQueryType = 31
	ReportQueryType_REPORT_SECRETS         ReportQueryType = 41
	ReportQueryType_REPORT_LICENSES        ReportQueryType = 51
)

// Enum value maps for ReportQueryType.
var (
	ReportQueryType_name = map[int32]string{
		0:  "REPORT_UNSPECIFIED",
		11: "REPORT_OS_PKGS",
		21: "REPORT_VULNERABILITIES",
		31: "REPORT_MISCONFIGS",
		41: "REPORT_SECRETS",
		51: "REPORT_LICENSES",
	}
	ReportQueryType_value = map[string]int32{
		"REPORT_UNSPECIFIED":     0,
		"REPORT_OS_PKGS":         11,
		"REPORT_VULNERABILITIES": 21,
		"REPORT_MISCONFIGS":      31,
		"REPORT_SECRETS":         41,
		"REPORT_LICENSES":        51,
	}
)

func (x ReportQueryType) Enum() *ReportQueryType {
	p := new(ReportQueryType)
	*p = x
	return p
}

func (x ReportQueryType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ReportQueryType) Descriptor() protoreflect.EnumDescriptor {
	return file_mmesh_protobuf_resources_v1_nstore_hsecdb_hsecdb_proto_enumTypes[0].Descriptor()
}

func (ReportQueryType) Type() protoreflect.EnumType {
	return &file_mmesh_protobuf_resources_v1_nstore_hsecdb_hsecdb_proto_enumTypes[0]
}

func (x ReportQueryType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ReportQueryType.Descriptor instead.
func (ReportQueryType) EnumDescriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_nstore_hsecdb_hsecdb_proto_rawDescGZIP(), []int{0}
}

type HostSecurityReportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request *nstore.DataRequest `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	Type    ReportQueryType     `protobuf:"varint,11,opt,name=type,proto3,enum=hsecdb.ReportQueryType" json:"type,omitempty"`
}

func (x *HostSecurityReportRequest) Reset() {
	*x = HostSecurityReportRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_nstore_hsecdb_hsecdb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HostSecurityReportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HostSecurityReportRequest) ProtoMessage() {}

func (x *HostSecurityReportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_nstore_hsecdb_hsecdb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HostSecurityReportRequest.ProtoReflect.Descriptor instead.
func (*HostSecurityReportRequest) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_nstore_hsecdb_hsecdb_proto_rawDescGZIP(), []int{0}
}

func (x *HostSecurityReportRequest) GetRequest() *nstore.DataRequest {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *HostSecurityReportRequest) GetType() ReportQueryType {
	if x != nil {
		return x.Type
	}
	return ReportQueryType_REPORT_UNSPECIFIED
}

type HostSecurityReportResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountID    string                 `protobuf:"bytes,1,opt,name=accountID,proto3" json:"accountID,omitempty"`
	TenantID     string                 `protobuf:"bytes,2,opt,name=tenantID,proto3" json:"tenantID,omitempty"`
	NodeID       string                 `protobuf:"bytes,5,opt,name=nodeID,proto3" json:"nodeID,omitempty"`
	QueryID      string                 `protobuf:"bytes,11,opt,name=queryID,proto3" json:"queryID,omitempty"`
	Metadata     *ReportMetadata        `protobuf:"bytes,101,opt,name=metadata,proto3" json:"metadata,omitempty"`
	VulnReport   []*VulnerabilityReport `protobuf:"bytes,201,rep,name=vulnReport,proto3" json:"vulnReport,omitempty"`
	OsPkgsReport *OSPkgsReport          `protobuf:"bytes,211,opt,name=osPkgsReport,proto3" json:"osPkgsReport,omitempty"`
	Timestamp    int64                  `protobuf:"varint,1001,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *HostSecurityReportResponse) Reset() {
	*x = HostSecurityReportResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_nstore_hsecdb_hsecdb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HostSecurityReportResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HostSecurityReportResponse) ProtoMessage() {}

func (x *HostSecurityReportResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_nstore_hsecdb_hsecdb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HostSecurityReportResponse.ProtoReflect.Descriptor instead.
func (*HostSecurityReportResponse) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_nstore_hsecdb_hsecdb_proto_rawDescGZIP(), []int{1}
}

func (x *HostSecurityReportResponse) GetAccountID() string {
	if x != nil {
		return x.AccountID
	}
	return ""
}

func (x *HostSecurityReportResponse) GetTenantID() string {
	if x != nil {
		return x.TenantID
	}
	return ""
}

func (x *HostSecurityReportResponse) GetNodeID() string {
	if x != nil {
		return x.NodeID
	}
	return ""
}

func (x *HostSecurityReportResponse) GetQueryID() string {
	if x != nil {
		return x.QueryID
	}
	return ""
}

func (x *HostSecurityReportResponse) GetMetadata() *ReportMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *HostSecurityReportResponse) GetVulnReport() []*VulnerabilityReport {
	if x != nil {
		return x.VulnReport
	}
	return nil
}

func (x *HostSecurityReportResponse) GetOsPkgsReport() *OSPkgsReport {
	if x != nil {
		return x.OsPkgsReport
	}
	return nil
}

func (x *HostSecurityReportResponse) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type ReportMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OsName   string `protobuf:"bytes,11,opt,name=osName,proto3" json:"osName,omitempty"`
	OsFamily string `protobuf:"bytes,12,opt,name=osFamily,proto3" json:"osFamily,omitempty"`
}

func (x *ReportMetadata) Reset() {
	*x = ReportMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_nstore_hsecdb_hsecdb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportMetadata) ProtoMessage() {}

func (x *ReportMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_nstore_hsecdb_hsecdb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportMetadata.ProtoReflect.Descriptor instead.
func (*ReportMetadata) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_nstore_hsecdb_hsecdb_proto_rawDescGZIP(), []int{2}
}

func (x *ReportMetadata) GetOsName() string {
	if x != nil {
		return x.OsName
	}
	return ""
}

func (x *ReportMetadata) GetOsFamily() string {
	if x != nil {
		return x.OsFamily
	}
	return ""
}

type VulnerabilityReport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Target          string           `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	Type            string           `protobuf:"bytes,11,opt,name=type,proto3" json:"type,omitempty"`
	Totals          *VulnTotals      `protobuf:"bytes,51,opt,name=totals,proto3" json:"totals,omitempty"`
	Vulnerabilities []*Vulnerability `protobuf:"bytes,101,rep,name=vulnerabilities,proto3" json:"vulnerabilities,omitempty"`
}

func (x *VulnerabilityReport) Reset() {
	*x = VulnerabilityReport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_nstore_hsecdb_hsecdb_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VulnerabilityReport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VulnerabilityReport) ProtoMessage() {}

func (x *VulnerabilityReport) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_nstore_hsecdb_hsecdb_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VulnerabilityReport.ProtoReflect.Descriptor instead.
func (*VulnerabilityReport) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_nstore_hsecdb_hsecdb_proto_rawDescGZIP(), []int{3}
}

func (x *VulnerabilityReport) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *VulnerabilityReport) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *VulnerabilityReport) GetTotals() *VulnTotals {
	if x != nil {
		return x.Totals
	}
	return nil
}

func (x *VulnerabilityReport) GetVulnerabilities() []*Vulnerability {
	if x != nil {
		return x.Vulnerabilities
	}
	return nil
}

type Vulnerability struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VulnerabilityID  string `protobuf:"bytes,1,opt,name=vulnerabilityID,proto3" json:"vulnerabilityID,omitempty"`
	PkgName          string `protobuf:"bytes,11,opt,name=pkgName,proto3" json:"pkgName,omitempty"`
	InstalledVersion string `protobuf:"bytes,21,opt,name=installedVersion,proto3" json:"installedVersion,omitempty"`
	FixedVersion     string `protobuf:"bytes,22,opt,name=fixedVersion,proto3" json:"fixedVersion,omitempty"`
	Status           string `protobuf:"bytes,31,opt,name=status,proto3" json:"status,omitempty"`
	PrimaryURL       string `protobuf:"bytes,41,opt,name=primaryURL,proto3" json:"primaryURL,omitempty"`
	Title            string `protobuf:"bytes,101,opt,name=title,proto3" json:"title,omitempty"`
	Description      string `protobuf:"bytes,111,opt,name=description,proto3" json:"description,omitempty"`
	Severity         string `protobuf:"bytes,121,opt,name=severity,proto3" json:"severity,omitempty"`
	PublishedDate    int64  `protobuf:"varint,131,opt,name=publishedDate,proto3" json:"publishedDate,omitempty"`
}

func (x *Vulnerability) Reset() {
	*x = Vulnerability{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_nstore_hsecdb_hsecdb_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Vulnerability) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vulnerability) ProtoMessage() {}

func (x *Vulnerability) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_nstore_hsecdb_hsecdb_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vulnerability.ProtoReflect.Descriptor instead.
func (*Vulnerability) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_nstore_hsecdb_hsecdb_proto_rawDescGZIP(), []int{4}
}

func (x *Vulnerability) GetVulnerabilityID() string {
	if x != nil {
		return x.VulnerabilityID
	}
	return ""
}

func (x *Vulnerability) GetPkgName() string {
	if x != nil {
		return x.PkgName
	}
	return ""
}

func (x *Vulnerability) GetInstalledVersion() string {
	if x != nil {
		return x.InstalledVersion
	}
	return ""
}

func (x *Vulnerability) GetFixedVersion() string {
	if x != nil {
		return x.FixedVersion
	}
	return ""
}

func (x *Vulnerability) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Vulnerability) GetPrimaryURL() string {
	if x != nil {
		return x.PrimaryURL
	}
	return ""
}

func (x *Vulnerability) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Vulnerability) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Vulnerability) GetSeverity() string {
	if x != nil {
		return x.Severity
	}
	return ""
}

func (x *Vulnerability) GetPublishedDate() int64 {
	if x != nil {
		return x.PublishedDate
	}
	return 0
}

type VulnTotals struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total    int32 `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Unknown  int32 `protobuf:"varint,11,opt,name=unknown,proto3" json:"unknown,omitempty"`
	Low      int32 `protobuf:"varint,21,opt,name=low,proto3" json:"low,omitempty"`
	Medium   int32 `protobuf:"varint,31,opt,name=medium,proto3" json:"medium,omitempty"`
	High     int32 `protobuf:"varint,41,opt,name=high,proto3" json:"high,omitempty"`
	Critical int32 `protobuf:"varint,51,opt,name=critical,proto3" json:"critical,omitempty"`
}

func (x *VulnTotals) Reset() {
	*x = VulnTotals{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_nstore_hsecdb_hsecdb_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VulnTotals) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VulnTotals) ProtoMessage() {}

func (x *VulnTotals) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_nstore_hsecdb_hsecdb_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VulnTotals.ProtoReflect.Descriptor instead.
func (*VulnTotals) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_nstore_hsecdb_hsecdb_proto_rawDescGZIP(), []int{5}
}

func (x *VulnTotals) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *VulnTotals) GetUnknown() int32 {
	if x != nil {
		return x.Unknown
	}
	return 0
}

func (x *VulnTotals) GetLow() int32 {
	if x != nil {
		return x.Low
	}
	return 0
}

func (x *VulnTotals) GetMedium() int32 {
	if x != nil {
		return x.Medium
	}
	return 0
}

func (x *VulnTotals) GetHigh() int32 {
	if x != nil {
		return x.High
	}
	return 0
}

func (x *VulnTotals) GetCritical() int32 {
	if x != nil {
		return x.Critical
	}
	return 0
}

type OSPkgsReport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Target    string `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	Type      string `protobuf:"bytes,11,opt,name=type,proto3" json:"type,omitempty"`
	TotalPkgs int32  `protobuf:"varint,51,opt,name=totalPkgs,proto3" json:"totalPkgs,omitempty"`
	Pkgs      []*Pkg `protobuf:"bytes,101,rep,name=pkgs,proto3" json:"pkgs,omitempty"`
}

func (x *OSPkgsReport) Reset() {
	*x = OSPkgsReport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_nstore_hsecdb_hsecdb_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OSPkgsReport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OSPkgsReport) ProtoMessage() {}

func (x *OSPkgsReport) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_nstore_hsecdb_hsecdb_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OSPkgsReport.ProtoReflect.Descriptor instead.
func (*OSPkgsReport) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_nstore_hsecdb_hsecdb_proto_rawDescGZIP(), []int{6}
}

func (x *OSPkgsReport) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *OSPkgsReport) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *OSPkgsReport) GetTotalPkgs() int32 {
	if x != nil {
		return x.TotalPkgs
	}
	return 0
}

func (x *OSPkgsReport) GetPkgs() []*Pkg {
	if x != nil {
		return x.Pkgs
	}
	return nil
}

type Pkg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// string pkgID = 1;
	Name           string   `protobuf:"bytes,11,opt,name=name,proto3" json:"name,omitempty"`
	Version        string   `protobuf:"bytes,21,opt,name=version,proto3" json:"version,omitempty"`
	Release        string   `protobuf:"bytes,31,opt,name=release,proto3" json:"release,omitempty"`
	Epoch          int32    `protobuf:"varint,41,opt,name=epoch,proto3" json:"epoch,omitempty"`
	Arch           string   `protobuf:"bytes,51,opt,name=arch,proto3" json:"arch,omitempty"`
	Licenses       []string `protobuf:"bytes,61,rep,name=licenses,proto3" json:"licenses,omitempty"`
	Maintainer     string   `protobuf:"bytes,71,opt,name=maintainer,proto3" json:"maintainer,omitempty"`
	InstalledFiles int32    `protobuf:"varint,81,opt,name=installedFiles,proto3" json:"installedFiles,omitempty"`
}

func (x *Pkg) Reset() {
	*x = Pkg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mmesh_protobuf_resources_v1_nstore_hsecdb_hsecdb_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pkg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pkg) ProtoMessage() {}

func (x *Pkg) ProtoReflect() protoreflect.Message {
	mi := &file_mmesh_protobuf_resources_v1_nstore_hsecdb_hsecdb_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pkg.ProtoReflect.Descriptor instead.
func (*Pkg) Descriptor() ([]byte, []int) {
	return file_mmesh_protobuf_resources_v1_nstore_hsecdb_hsecdb_proto_rawDescGZIP(), []int{7}
}

func (x *Pkg) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Pkg) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Pkg) GetRelease() string {
	if x != nil {
		return x.Release
	}
	return ""
}

func (x *Pkg) GetEpoch() int32 {
	if x != nil {
		return x.Epoch
	}
	return 0
}

func (x *Pkg) GetArch() string {
	if x != nil {
		return x.Arch
	}
	return ""
}

func (x *Pkg) GetLicenses() []string {
	if x != nil {
		return x.Licenses
	}
	return nil
}

func (x *Pkg) GetMaintainer() string {
	if x != nil {
		return x.Maintainer
	}
	return ""
}

func (x *Pkg) GetInstalledFiles() int32 {
	if x != nil {
		return x.InstalledFiles
	}
	return 0
}

var File_mmesh_protobuf_resources_v1_nstore_hsecdb_hsecdb_proto protoreflect.FileDescriptor

var file_mmesh_protobuf_resources_v1_nstore_hsecdb_hsecdb_proto_rawDesc = []byte{
	0x0a, 0x36, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2f, 0x68, 0x73, 0x65, 0x63, 0x64, 0x62, 0x2f, 0x68, 0x73, 0x65, 0x63,
	0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x68, 0x73, 0x65, 0x63, 0x64, 0x62,
	0x1a, 0x2f, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2f, 0x6e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x77, 0x0a, 0x19, 0x48, 0x6f, 0x73, 0x74, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74,
	0x79, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d,
	0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x6e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x68, 0x73,
	0x65, 0x63, 0x64, 0x62, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0xd4, 0x02, 0x0a, 0x1a, 0x48,
	0x6f, 0x73, 0x74, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x44, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x49, 0x44, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x49, 0x44, 0x12, 0x32, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x68, 0x73, 0x65, 0x63, 0x64, 0x62,
	0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x3c, 0x0a, 0x0a, 0x76, 0x75, 0x6c,
	0x6e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x18, 0xc9, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x68, 0x73, 0x65, 0x63, 0x64, 0x62, 0x2e, 0x56, 0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x0a, 0x76, 0x75, 0x6c,
	0x6e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x39, 0x0a, 0x0c, 0x6f, 0x73, 0x50, 0x6b, 0x67,
	0x73, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x18, 0xd3, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x68, 0x73, 0x65, 0x63, 0x64, 0x62, 0x2e, 0x4f, 0x53, 0x50, 0x6b, 0x67, 0x73, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x52, 0x0c, 0x6f, 0x73, 0x50, 0x6b, 0x67, 0x73, 0x52, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x12, 0x1d, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18,
	0xe9, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x22, 0x44, 0x0a, 0x0e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6f,
	0x73, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f,
	0x73, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x22, 0xae, 0x01, 0x0a, 0x13, 0x56, 0x75, 0x6c, 0x6e,
	0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x73, 0x18, 0x33, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x68, 0x73,
	0x65, 0x63, 0x64, 0x62, 0x2e, 0x56, 0x75, 0x6c, 0x6e, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x73, 0x52,
	0x06, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x73, 0x12, 0x3f, 0x0a, 0x0f, 0x76, 0x75, 0x6c, 0x6e, 0x65,
	0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x65, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x68, 0x73, 0x65, 0x63, 0x64, 0x62, 0x2e, 0x56, 0x75, 0x6c, 0x6e, 0x65, 0x72,
	0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x0f, 0x76, 0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x22, 0xd6, 0x02, 0x0a, 0x0d, 0x56, 0x75, 0x6c,
	0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x28, 0x0a, 0x0f, 0x76, 0x75,
	0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x76, 0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x6b, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x6b, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2a,
	0x0a, 0x10, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x65, 0x64, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c,
	0x6c, 0x65, 0x64, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x66, 0x69,
	0x78, 0x65, 0x64, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x66, 0x69, 0x78, 0x65, 0x64, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72,
	0x79, 0x55, 0x52, 0x4c, 0x18, 0x29, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x69, 0x6d,
	0x61, 0x72, 0x79, 0x55, 0x52, 0x4c, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x65, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x6f, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a,
	0x0a, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x18, 0x79, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x12, 0x25, 0x0a, 0x0d, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x44, 0x61, 0x74, 0x65, 0x18, 0x83, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x44, 0x61, 0x74,
	0x65, 0x22, 0x96, 0x01, 0x0a, 0x0a, 0x56, 0x75, 0x6c, 0x6e, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77,
	0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e,
	0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6f, 0x77, 0x18, 0x15, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6c,
	0x6f, 0x77, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x64, 0x69, 0x75, 0x6d, 0x18, 0x1f, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x6d, 0x65, 0x64, 0x69, 0x75, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x69,
	0x67, 0x68, 0x18, 0x29, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x68, 0x69, 0x67, 0x68, 0x12, 0x1a,
	0x0a, 0x08, 0x63, 0x72, 0x69, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x18, 0x33, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x63, 0x72, 0x69, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x22, 0x79, 0x0a, 0x0c, 0x4f, 0x53,
	0x50, 0x6b, 0x67, 0x73, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50,
	0x6b, 0x67, 0x73, 0x18, 0x33, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x50, 0x6b, 0x67, 0x73, 0x12, 0x1f, 0x0a, 0x04, 0x70, 0x6b, 0x67, 0x73, 0x18, 0x65, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x68, 0x73, 0x65, 0x63, 0x64, 0x62, 0x2e, 0x50, 0x6b, 0x67, 0x52,
	0x04, 0x70, 0x6b, 0x67, 0x73, 0x22, 0xdb, 0x01, 0x0a, 0x03, 0x50, 0x6b, 0x67, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x15, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x72,
	0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65,
	0x6c, 0x65, 0x61, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x18, 0x29,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x61,
	0x72, 0x63, 0x68, 0x18, 0x33, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x72, 0x63, 0x68, 0x12,
	0x1a, 0x0a, 0x08, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x73, 0x18, 0x3d, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x08, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x6d,
	0x61, 0x69, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x18, 0x47, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x6d, 0x61, 0x69, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x0e, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x65, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x51, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x65, 0x64, 0x46, 0x69,
	0x6c, 0x65, 0x73, 0x2a, 0x99, 0x01, 0x0a, 0x0f, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x12, 0x52, 0x45, 0x50, 0x4f, 0x52,
	0x54, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x12, 0x0a, 0x0e, 0x52, 0x45, 0x50, 0x4f, 0x52, 0x54, 0x5f, 0x4f, 0x53, 0x5f, 0x50, 0x4b, 0x47,
	0x53, 0x10, 0x0b, 0x12, 0x1a, 0x0a, 0x16, 0x52, 0x45, 0x50, 0x4f, 0x52, 0x54, 0x5f, 0x56, 0x55,
	0x4c, 0x4e, 0x45, 0x52, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x49, 0x45, 0x53, 0x10, 0x15, 0x12,
	0x15, 0x0a, 0x11, 0x52, 0x45, 0x50, 0x4f, 0x52, 0x54, 0x5f, 0x4d, 0x49, 0x53, 0x43, 0x4f, 0x4e,
	0x46, 0x49, 0x47, 0x53, 0x10, 0x1f, 0x12, 0x12, 0x0a, 0x0e, 0x52, 0x45, 0x50, 0x4f, 0x52, 0x54,
	0x5f, 0x53, 0x45, 0x43, 0x52, 0x45, 0x54, 0x53, 0x10, 0x29, 0x12, 0x13, 0x0a, 0x0f, 0x52, 0x45,
	0x50, 0x4f, 0x52, 0x54, 0x5f, 0x4c, 0x49, 0x43, 0x45, 0x4e, 0x53, 0x45, 0x53, 0x10, 0x33, 0x42,
	0x31, 0x5a, 0x2f, 0x6d, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x6d, 0x2d, 0x61,
	0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x6e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x68, 0x73, 0x65, 0x63,
	0x64, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mmesh_protobuf_resources_v1_nstore_hsecdb_hsecdb_proto_rawDescOnce sync.Once
	file_mmesh_protobuf_resources_v1_nstore_hsecdb_hsecdb_proto_rawDescData = file_mmesh_protobuf_resources_v1_nstore_hsecdb_hsecdb_proto_rawDesc
)

func file_mmesh_protobuf_resources_v1_nstore_hsecdb_hsecdb_proto_rawDescGZIP() []byte {
	file_mmesh_protobuf_resources_v1_nstore_hsecdb_hsecdb_proto_rawDescOnce.Do(func() {
		file_mmesh_protobuf_resources_v1_nstore_hsecdb_hsecdb_proto_rawDescData = protoimpl.X.CompressGZIP(file_mmesh_protobuf_resources_v1_nstore_hsecdb_hsecdb_proto_rawDescData)
	})
	return file_mmesh_protobuf_resources_v1_nstore_hsecdb_hsecdb_proto_rawDescData
}

var file_mmesh_protobuf_resources_v1_nstore_hsecdb_hsecdb_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_mmesh_protobuf_resources_v1_nstore_hsecdb_hsecdb_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_mmesh_protobuf_resources_v1_nstore_hsecdb_hsecdb_proto_goTypes = []interface{}{
	(ReportQueryType)(0),               // 0: hsecdb.ReportQueryType
	(*HostSecurityReportRequest)(nil),  // 1: hsecdb.HostSecurityReportRequest
	(*HostSecurityReportResponse)(nil), // 2: hsecdb.HostSecurityReportResponse
	(*ReportMetadata)(nil),             // 3: hsecdb.ReportMetadata
	(*VulnerabilityReport)(nil),        // 4: hsecdb.VulnerabilityReport
	(*Vulnerability)(nil),              // 5: hsecdb.Vulnerability
	(*VulnTotals)(nil),                 // 6: hsecdb.VulnTotals
	(*OSPkgsReport)(nil),               // 7: hsecdb.OSPkgsReport
	(*Pkg)(nil),                        // 8: hsecdb.Pkg
	(*nstore.DataRequest)(nil),         // 9: nstore.DataRequest
}
var file_mmesh_protobuf_resources_v1_nstore_hsecdb_hsecdb_proto_depIdxs = []int32{
	9, // 0: hsecdb.HostSecurityReportRequest.request:type_name -> nstore.DataRequest
	0, // 1: hsecdb.HostSecurityReportRequest.type:type_name -> hsecdb.ReportQueryType
	3, // 2: hsecdb.HostSecurityReportResponse.metadata:type_name -> hsecdb.ReportMetadata
	4, // 3: hsecdb.HostSecurityReportResponse.vulnReport:type_name -> hsecdb.VulnerabilityReport
	7, // 4: hsecdb.HostSecurityReportResponse.osPkgsReport:type_name -> hsecdb.OSPkgsReport
	6, // 5: hsecdb.VulnerabilityReport.totals:type_name -> hsecdb.VulnTotals
	5, // 6: hsecdb.VulnerabilityReport.vulnerabilities:type_name -> hsecdb.Vulnerability
	8, // 7: hsecdb.OSPkgsReport.pkgs:type_name -> hsecdb.Pkg
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_mmesh_protobuf_resources_v1_nstore_hsecdb_hsecdb_proto_init() }
func file_mmesh_protobuf_resources_v1_nstore_hsecdb_hsecdb_proto_init() {
	if File_mmesh_protobuf_resources_v1_nstore_hsecdb_hsecdb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mmesh_protobuf_resources_v1_nstore_hsecdb_hsecdb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HostSecurityReportRequest); i {
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
		file_mmesh_protobuf_resources_v1_nstore_hsecdb_hsecdb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HostSecurityReportResponse); i {
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
		file_mmesh_protobuf_resources_v1_nstore_hsecdb_hsecdb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportMetadata); i {
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
		file_mmesh_protobuf_resources_v1_nstore_hsecdb_hsecdb_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VulnerabilityReport); i {
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
		file_mmesh_protobuf_resources_v1_nstore_hsecdb_hsecdb_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Vulnerability); i {
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
		file_mmesh_protobuf_resources_v1_nstore_hsecdb_hsecdb_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VulnTotals); i {
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
		file_mmesh_protobuf_resources_v1_nstore_hsecdb_hsecdb_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OSPkgsReport); i {
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
		file_mmesh_protobuf_resources_v1_nstore_hsecdb_hsecdb_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pkg); i {
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
			RawDescriptor: file_mmesh_protobuf_resources_v1_nstore_hsecdb_hsecdb_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mmesh_protobuf_resources_v1_nstore_hsecdb_hsecdb_proto_goTypes,
		DependencyIndexes: file_mmesh_protobuf_resources_v1_nstore_hsecdb_hsecdb_proto_depIdxs,
		EnumInfos:         file_mmesh_protobuf_resources_v1_nstore_hsecdb_hsecdb_proto_enumTypes,
		MessageInfos:      file_mmesh_protobuf_resources_v1_nstore_hsecdb_hsecdb_proto_msgTypes,
	}.Build()
	File_mmesh_protobuf_resources_v1_nstore_hsecdb_hsecdb_proto = out.File
	file_mmesh_protobuf_resources_v1_nstore_hsecdb_hsecdb_proto_rawDesc = nil
	file_mmesh_protobuf_resources_v1_nstore_hsecdb_hsecdb_proto_goTypes = nil
	file_mmesh_protobuf_resources_v1_nstore_hsecdb_hsecdb_proto_depIdxs = nil
}
