package hsec

import (
	"time"

	dbTypes "github.com/aquasecurity/trivy-db/pkg/types"
	ftypes "github.com/aquasecurity/trivy/pkg/fanal/types"
	"github.com/aquasecurity/trivy/pkg/types"
	"mmesh.dev/m-api-go/grpc/resources/hsec"
	"mmesh.dev/m-api-go/grpc/resources/nstore/hsecdb"
)

func query(r *hsecdb.HostSecurityReportRequest, hsr *hsec.Report) *hsecdb.HostSecurityReportResponse {
	hsrr := &hsecdb.HostSecurityReportResponse{
		AccountID: r.Request.AccountID,
		TenantID:  r.Request.TenantID,
		NodeID:    r.Request.NodeID,
		QueryID:   r.Request.QueryID,
		// Report:    hsr,
		Metadata:     nil,
		VulnReport:   nil,
		OsPkgsReport: nil,
		Timestamp:    time.Now().UnixMilli(),
	}

	if hsr != nil {
		// remove unnecessary metadata from report
		hsr.Metadata.ImageConfig = nil

		switch r.Type {
		case hsecdb.ReportQueryType_REPORT_UNSPECIFIED:
			filterVulnerabilities(hsr)
		case hsecdb.ReportQueryType_REPORT_OS_PKGS:
			filterOSPkgs(hsr)
			hsrr.OsPkgsReport = getOSPkgsReport(hsr)
		case hsecdb.ReportQueryType_REPORT_VULNERABILITIES:
			filterVulnerabilities(hsr)
			hsrr.VulnReport = getVulnReport(hsr)
		case hsecdb.ReportQueryType_REPORT_MISCONFIGS:
			filterMisconfigs(hsr)
		case hsecdb.ReportQueryType_REPORT_SECRETS:
			filterSecrets(hsr)
		case hsecdb.ReportQueryType_REPORT_LICENSES:
			filterLicenses(hsr)
		default:
			filterVulnerabilities(hsr)
		}

		hsrr.Metadata = getReportMetadata(hsr)
		hsrr.Timestamp = hsr.CreatedAt
	}

	return hsrr
}

func getReportMetadata(hsr *hsec.Report) *hsecdb.ReportMetadata {
	if hsr == nil || hsr.Metadata == nil || hsr.Metadata.OS == nil {
		return &hsecdb.ReportMetadata{}
	}

	return &hsecdb.ReportMetadata{
		OsName:   hsr.Metadata.OS.Name,
		OsFamily: hsr.Metadata.OS.Family,
	}
}

func getVulnReport(hsr *hsec.Report) []*hsecdb.VulnerabilityReport {
	vrl := make([]*hsecdb.VulnerabilityReport, 0)

	for _, r := range hsr.Results {
		if len(r.Vulnerabilities) == 0 {
			continue
		}

		vrl = append(vrl, &hsecdb.VulnerabilityReport{
			Target:          r.Target,
			Type:            r.Type,
			Totals:          getVulnTotals(r),
			Vulnerabilities: getVulns(r),
		})
	}

	return vrl
}

func getVulns(r *hsec.Result) []*hsecdb.Vulnerability {
	vl := make([]*hsecdb.Vulnerability, 0)

	for _, v := range r.Vulnerabilities {
		if v.Vulnerability == nil {
			continue
		}

		vl = append(vl, &hsecdb.Vulnerability{
			VulnerabilityID:  v.VulnerabilityID,
			PkgName:          v.PkgName,
			InstalledVersion: v.InstalledVersion,
			FixedVersion:     v.FixedVersion,
			Status:           v.Status,
			PrimaryURL:       v.PrimaryURL,
			Title:            v.Vulnerability.Title,
			Description:      v.Vulnerability.Description,
			Severity:         v.Vulnerability.Severity,
			PublishedDate:    v.Vulnerability.PublishedDate,
		})
	}

	return vl
}

func getVulnTotals(r *hsec.Result) *hsecdb.VulnTotals {
	var total, unknown, low, medium, high, critical int32

	for _, v := range r.Vulnerabilities {
		if v.Vulnerability == nil {
			continue
		}

		total++

		// See: https://pkg.go.dev/github.com/aquasecurity/trivy-db/pkg/types#pkg-constants

		switch v.Vulnerability.Severity {
		case dbTypes.SeverityUnknown.String():
			unknown++
		case dbTypes.SeverityLow.String():
			low++
		case dbTypes.SeverityMedium.String():
			medium++
		case dbTypes.SeverityHigh.String():
			high++
		case dbTypes.SeverityCritical.String():
			critical++
		}
	}

	return &hsecdb.VulnTotals{
		Total:    total,
		Unknown:  unknown,
		Low:      low,
		Medium:   medium,
		High:     high,
		Critical: critical,
	}
}

func getOSPkgsReport(hsr *hsec.Report) *hsecdb.OSPkgsReport {
	if len(hsr.Results) != 1 {
		return &hsecdb.OSPkgsReport{}
	}

	r := hsr.Results[0]

	if r.Class != "os-pkgs" {
		return &hsecdb.OSPkgsReport{}
	}

	osr := &hsecdb.OSPkgsReport{
		Target:    r.Target,
		Type:      r.Type,
		TotalPkgs: r.ScannedPackages,
		Pkgs:      make([]*hsecdb.Pkg, 0),
	}

	for _, pkg := range r.Packages {
		osr.Pkgs = append(osr.Pkgs, &hsecdb.Pkg{
			// PkgID:      pkg.ID,
			Name:           pkg.Name,
			Version:        pkg.Version,
			Release:        pkg.Release,
			Epoch:          pkg.Epoch,
			Arch:           pkg.Arch,
			Licenses:       pkg.Licenses,
			Maintainer:     pkg.Maintainer,
			InstalledFiles: int32(len(pkg.InstalledFiles)),
		})
	}

	return osr
}

func filterOSPkgs(hsr *hsec.Report) {
	results := make([]*hsec.Result, 0)

	for _, r := range hsr.Results {
		if r.Class != string(types.ClassOSPkg) {
			continue
		}

		results = append(results, r)
	}

	hsr.Results = results
}

func filterVulnerabilities(hsr *hsec.Report) {
	results := make([]*hsec.Result, 0)

	for _, r := range hsr.Results {
		if r.Class == string(types.ClassLangPkg) {
			if r.Type == string(ftypes.NodePkg) {
				continue
			}
			if r.Type == string(ftypes.PythonPkg) {
				continue
			}
			if r.Type == string(ftypes.CondaPkg) {
				continue
			}
		}

		// remove pkgs data from report
		r.Packages = nil

		results = append(results, r)
	}

	hsr.Results = results
}

func filterMisconfigs(hsr *hsec.Report) {
	results := make([]*hsec.Result, 0)

	for _, r := range hsr.Results {
		if r.Class != string(types.ClassConfig) {
			continue
		}

		// remove pkgs data from report
		r.Packages = nil

		results = append(results, r)
	}

	hsr.Results = results
}

func filterSecrets(hsr *hsec.Report) {
	results := make([]*hsec.Result, 0)

	for _, r := range hsr.Results {
		if r.Class != string(types.ClassSecret) {
			continue
		}

		// remove pkgs data from report
		r.Packages = nil

		results = append(results, r)
	}

	hsr.Results = results
}

func filterLicenses(hsr *hsec.Report) {
	results := make([]*hsec.Result, 0)

	for _, r := range hsr.Results {
		if r.Class != string(types.ClassLicense) {
			continue
		}

		// remove pkgs data from report
		r.Packages = nil

		results = append(results, r)
	}

	hsr.Results = results
}
