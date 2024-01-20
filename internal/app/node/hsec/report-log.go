package hsec

import (
	"github.com/aquasecurity/trivy/pkg/types"
	"mmesh.dev/m-lib/pkg/xlog"
)

func logSecurityReport(result types.Result, idx int) {
	xlog.Info("-----------------------------------------------------")
	xlog.Infof("[HSEC-%05d] Target: %s", idx, result.Target)
	xlog.Infof("[HSEC-%05d] Class: %s", idx, result.Class)
	xlog.Infof("[HSEC-%05d] Type: %s", idx, result.Type)
	xlog.Infof("[HSEC-%05d] Scanned Packages: %d", idx, len(result.Packages))
	xlog.Infof("[HSEC-%05d] Vulnerabilities: %d", idx, len(result.Vulnerabilities))
	xlog.Infof("[HSEC-%05d] Misconfigurations: %d", idx, len(result.Misconfigurations))
	if result.MisconfSummary != nil {
		xlog.Infof("    Successes: %d", result.MisconfSummary.Successes)
		xlog.Infof("    Failures: %d", result.MisconfSummary.Failures)
		xlog.Infof("    Exceptions: %d", result.MisconfSummary.Exceptions)
	}
	xlog.Infof("[HSEC-%05d] Secrets: %d", idx, len(result.Secrets))
	xlog.Infof("[HSEC-%05d] Licenses: %d", idx, len(result.Licenses))
	xlog.Infof("[HSEC-%05d] Custom Resources: %d", idx, len(result.CustomResources))
	xlog.Info("-----------------------------------------------------")
}

/*
func scanReport(ctx context.Context, runner artifact.Runner, opts flag.Options, r types.Report) error {
	var err error

	r, err = runner.Filter(ctx, opts, r)
	if err != nil {
		return errors.Wrapf(err, "[%v] function runner.Filter()", errors.Trace())
	}

	if err = runner.Report(ctx, opts, r); err != nil {
		return errors.Wrapf(err, "[%v] function runner.Report()", errors.Trace())
	}

	return nil
}
*/
