package hsec

import (
	"time"

	ftypes "github.com/aquasecurity/trivy/pkg/fanal/types"
	"github.com/aquasecurity/trivy/pkg/types"
	"mmesh.dev/m-api-go/grpc/resources/hsec"
	"mmesh.dev/m-api-go/grpc/resources/nstore/hsecdb"
)

func query(r *hsecdb.HostSecurityReportRequest, hsr *hsec.Report) *hsecdb.HostSecurityReportResponse {
	if hsr == nil {
		hsr = &hsec.Report{}
	} else {
		// remove unnecessary metadata from report
		hsr.Metadata.ImageConfig = nil

		switch r.Type {
		case hsecdb.ReportQueryType_REPORT_UNSPECIFIED:
			filterVulnerabilities(hsr)
		case hsecdb.ReportQueryType_REPORT_OS_PKGS:
			filterOSPkgs(hsr)
		case hsecdb.ReportQueryType_REPORT_VULNERABILITIES:
			filterVulnerabilities(hsr)
		case hsecdb.ReportQueryType_REPORT_MISCONFIGS:
			filterMisconfigs(hsr)
		case hsecdb.ReportQueryType_REPORT_SECRETS:
			filterSecrets(hsr)
		case hsecdb.ReportQueryType_REPORT_LICENSES:
			filterLicenses(hsr)
		default:
			filterVulnerabilities(hsr)
		}
	}

	return &hsecdb.HostSecurityReportResponse{
		AccountID: r.Request.AccountID,
		TenantID:  r.Request.TenantID,
		NodeID:    r.Request.NodeID,
		QueryID:   r.Request.QueryID,
		Report:    hsr,
		Timestamp: time.Now().UnixMilli(),
	}
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
