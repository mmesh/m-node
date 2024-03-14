package hsec

import (
	"mmesh.dev/m-api-go/grpc/resources/nstore"
	"mmesh.dev/m-api-go/grpc/resources/nstore/hsecdb"
	"mmesh.dev/m-lib/pkg/errors"
)

type SummaryReport struct {
	TotalOSPkgs int32
	Vulns       *hsecdb.VulnTotals
}

func GetSummary(r *nstore.DataRequest) (*SummaryReport, error) {
	s := &SummaryReport{
		Vulns: &hsecdb.VulnTotals{},
	}

	hsr, err := readReportFile()
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function os.ReadFile()", errors.Trace())
	}

	hsrr := query(&hsecdb.HostSecurityReportRequest{
		Request: r,
		Type:    hsecdb.ReportQueryType_REPORT_OS_PKGS,
	}, hsr)

	if hsrr != nil {
		if hsrr.OsPkgsReport != nil {
			s.TotalOSPkgs = hsrr.OsPkgsReport.TotalPkgs
		}
	}

	hsr, err = readReportFile()
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function os.ReadFile()", errors.Trace())
	}

	hsrr = query(&hsecdb.HostSecurityReportRequest{
		Request: r,
		Type:    hsecdb.ReportQueryType_REPORT_VULNERABILITIES,
	}, hsr)

	if hsrr != nil {
		for _, vr := range hsrr.VulnReport {
			if vr.Totals == nil {
				continue
			}

			s.Vulns.Total += vr.Totals.Total
			s.Vulns.Unknown += vr.Totals.Unknown
			s.Vulns.Low += vr.Totals.Low
			s.Vulns.Medium += vr.Totals.Medium
			s.Vulns.High += vr.Totals.High
			s.Vulns.Critical += vr.Totals.Critical
		}
	}

	return s, nil
}
