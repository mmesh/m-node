package hsec

import (
	"context"

	// "github.com/aquasecurity/trivy/pkg/utils/fsutils"
	"github.com/aquasecurity/trivy/pkg/commands/artifact"
	"github.com/aquasecurity/trivy/pkg/log"
	"github.com/aquasecurity/trivy/pkg/types"
	"mmesh.dev/m-lib/pkg/errors"
	"mmesh.dev/m-lib/pkg/xlog"
	_ "modernc.org/sqlite"
)

func scan() error {
	// Generate config options
	opts := newOptions(&optsConfig{
		globalCacheDir:     globalCacheDir(), // Default: fsutils.CacheDir()
		reportFormat:       types.FormatTable,
		reportReportFormat: "all", // "all", "summary"
		scanTarget:         rootTargetDir(),
		scanSkipDirs:       skipDirs(),
		scanSkipFiles:      []string{},
		scanFilePatterns:   []string{},
		scanParallel:       scanParallel, // Default: 5
	})

	// Initialize logger
	if err := log.InitLogger(opts.Debug, opts.Quiet); err != nil {
		return errors.Wrapf(err, "[%v] function log.InitLogger()", errors.Trace())
	}

	var err error

	ctx, cancel := context.WithTimeout(context.Background(), opts.Timeout)
	defer cancel()

	defer func() {
		if errors.Is(err, context.DeadlineExceeded) {
			xlog.Error("[host-security] Security scan timeout exceeded")
		}
	}()

	r, err := artifact.NewRunner(ctx, opts)
	if err != nil {
		if errors.Is(err, artifact.SkipScan) {
			xlog.Warnf("[host-security] Security scan skipped: %v", err)
			return nil
		}
		return errors.Wrapf(err, "[%v] function artifact.NewRunner()", errors.Trace())
	}
	defer r.Close(ctx)

	xlog.Info("[host-security] Scanning rootFS...")

	report, err := r.ScanRootfs(ctx, opts)
	if err != nil {
		return errors.Wrapf(err, "[%v] function r.ScanRootfs()", errors.Trace())
	}

	hsr := getHostSecurityReport(&report)

	if err := writeReportFile(hsr); err != nil {
		return errors.Wrapf(err, "[%v] function writeReportFile()", errors.Trace())
	}

	// if err := scanReport(ctx, r, opts, report); err != nil {
	// 	return errors.Wrapf(err, "[%v] function scanRport()", errors.Trace())
	// }

	for idx, result := range report.Results {
		logSecurityReport(result, idx)
	}

	xlog.Infof("[host-security] Security scan completed: %d results", len(report.Results))

	return nil
}
