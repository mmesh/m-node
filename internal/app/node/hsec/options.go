package hsec

import (
	"net/http"
	"os"
	"path/filepath"
	"time"

	defsecTypes "github.com/aquasecurity/defsec/pkg/types"
	dbTypes "github.com/aquasecurity/trivy-db/pkg/types"
	"github.com/aquasecurity/trivy/pkg/compliance/spec"
	ftypes "github.com/aquasecurity/trivy/pkg/fanal/types"
	"github.com/aquasecurity/trivy/pkg/flag"
	"github.com/aquasecurity/trivy/pkg/licensing"
	"github.com/aquasecurity/trivy/pkg/types"
)

const (
	defaultDBRepository     string = "ghcr.io/aquasecurity/trivy-db"
	defaultJavaDBRepository string = "ghcr.io/aquasecurity/trivy-java-db"
	defaultRekorURL         string = "https://rekor.sigstore.dev"

	debug bool = false
	quiet bool = true

	timeout = 14400 // 4 hours

	// Number of goroutines enabled for parallel scanning, set 0 to auto-detect parallelism
	scanParallel int = 1 // Default: 5
)

type optsConfig struct {
	globalCacheDir     string
	reportFormat       types.Format
	reportReportFormat string
	scanTarget         string
	scanSkipDirs       []string
	scanSkipFiles      []string
	scanFilePatterns   []string
	scanParallel       int
}

func globalCacheDir() string {
	tmpDir, err := os.UserCacheDir()
	if err != nil {
		tmpDir = os.TempDir()
	}
	// Default: "/tmp/mmesh/cache",
	return filepath.Join(tmpDir, "mmesh", "cache")
}

func newOptions(cfg *optsConfig) flag.Options {
	return flag.Options{
		GlobalOptions: flag.GlobalOptions{
			ConfigFile:            "",
			ShowVersion:           false,
			Quiet:                 quiet,
			Debug:                 debug,
			Insecure:              false,
			Timeout:               timeout * time.Second,
			CacheDir:              cfg.globalCacheDir,
			GenerateDefaultConfig: false,
		},
		AWSOptions: flag.AWSOptions{
			Region:       "",
			Endpoint:     "",
			Services:     nil,
			SkipServices: nil,
			Account:      "",
			ARN:          "",
		},
		CacheOptions: flag.CacheOptions{
			ClearCache:   false,
			CacheBackend: "fs",
			CacheTTL:     0,
			RedisTLS:     false,
			RedisOptions: flag.RedisOptions{
				RedisCACert: "",
				RedisCert:   "",
				RedisKey:    "",
			},
		},
		CloudOptions: flag.CloudOptions{
			MaxCacheAge: 0,
			UpdateCache: false,
		},
		DBOptions: flag.DBOptions{
			Reset:              false,
			DownloadDBOnly:     false,
			SkipDBUpdate:       false,
			DownloadJavaDBOnly: false,
			SkipJavaDBUpdate:   false,
			NoProgress:         true,
			DBRepository:       defaultDBRepository,
			JavaDBRepository:   defaultJavaDBRepository,
			Light:              false,
		},
		ImageOptions: flag.ImageOptions{
			Input:               "",
			ImageConfigScanners: nil,
			ScanRemovedPkgs:     false,
			Platform: ftypes.Platform{
				Platform: nil,
				Force:    false,
			},
			DockerHost:   "",
			ImageSources: nil,
		},
		K8sOptions: flag.K8sOptions{
			ClusterContext:         "",
			Namespace:              "",
			KubeConfig:             "",
			Components:             nil,
			K8sVersion:             "",
			Tolerations:            nil,
			NodeCollectorImageRef:  "",
			AllNamespaces:          false,
			NodeCollectorNamespace: "",
			ExcludeOwned:           false,
			ExcludeNodes:           map[string]string(nil),
		},
		LicenseOptions: flag.LicenseOptions{
			LicenseFull:            false,
			IgnoredLicenses:        nil,
			LicenseConfidenceLevel: 0.9,
			LicenseRiskThreshold:   0,
			LicenseCategories: map[ftypes.LicenseCategory][]string{
				ftypes.CategoryForbidden:    licensing.ForbiddenLicenses,
				ftypes.CategoryNotice:       licensing.NoticeLicenses,
				ftypes.CategoryPermissive:   licensing.PermissiveLicenses,
				ftypes.CategoryReciprocal:   licensing.ReciprocalLicenses,
				ftypes.CategoryRestricted:   licensing.RestrictedLicenses,
				ftypes.CategoryUnencumbered: licensing.UnencumberedLicenses,
			},
		},
		MisconfOptions: flag.MisconfOptions{
			IncludeNonFailures:      false,
			ResetPolicyBundle:       false,
			PolicyBundleRepository:  "",
			HelmValues:              nil,
			HelmValueFiles:          nil,
			HelmFileValues:          nil,
			HelmStringValues:        nil,
			TerraformTFVars:         nil,
			CloudFormationParamVars: nil,
			TfExcludeDownloaded:     false,
			MisconfigScanners:       nil,
		},
		ModuleOptions: flag.ModuleOptions{
			ModuleDir:      "",
			EnabledModules: nil,
		},
		RegistryOptions: flag.RegistryOptions{
			Credentials:   nil,
			RegistryToken: "",
		},
		RegoOptions: flag.RegoOptions{
			SkipPolicyUpdate: false,
			Trace:            false,
			PolicyPaths:      nil,
			DataPaths:        nil,
			PolicyNamespaces: nil,
		},
		RemoteOptions: flag.RemoteOptions{
			Token:         "",
			TokenHeader:   "",
			ServerAddr:    "",
			Listen:        "",
			CustomHeaders: http.Header{},
		},
		RepoOptions: flag.RepoOptions{
			RepoBranch: "",
			RepoCommit: "",
			RepoTag:    "",
		},
		ReportOptions: flag.ReportOptions{
			Format:           cfg.reportFormat,
			ReportFormat:     cfg.reportReportFormat,
			Template:         "",
			DependencyTree:   true,
			ListAllPkgs:      true,
			IgnoreFile:       "",
			ExitCode:         0,
			ExitOnEOL:        0,
			IgnorePolicy:     "",
			Output:           "",
			OutputPluginArgs: nil,
			Severities: []dbTypes.Severity{
				dbTypes.SeverityCritical,
				dbTypes.SeverityHigh,
				dbTypes.SeverityMedium,
				dbTypes.SeverityLow,
				dbTypes.SeverityUnknown,
			},
			Compliance: spec.ComplianceSpec{
				Spec: defsecTypes.Spec{
					ID:               "",
					Title:            "",
					Description:      "",
					Version:          "",
					RelatedResources: nil,
					Controls:         nil,
				},
			},
		},
		SBOMOptions: flag.SBOMOptions{},
		ScanOptions: flag.ScanOptions{
			Target:      cfg.scanTarget,
			SkipDirs:    cfg.scanSkipDirs,
			SkipFiles:   cfg.scanSkipFiles,
			OfflineScan: false,
			Scanners: types.Scanners{
				types.VulnerabilityScanner,
				types.MisconfigScanner,
				// types.SecretScanner,
				// types.LicenseScanner,
			},
			FilePatterns: cfg.scanFilePatterns,
			// Number of goroutines enabled for parallel scanning, set 0 to auto-detect parallelism
			Parallel:       cfg.scanParallel, // Default: 5
			SBOMSources:    nil,
			RekorURL:       defaultRekorURL,
			IncludeDevDeps: false,
		},
		SecretOptions: flag.SecretOptions{
			SecretConfigPath: "",
		},
		VulnerabilityOptions: flag.VulnerabilityOptions{
			VulnType: []string{
				types.VulnTypeOS,
				types.VulnTypeLibrary,
			},
			IgnoreStatuses: nil,
			VEXPath:        "",
		},
		AppVersion:        "dev",
		DisabledAnalyzers: nil,
	}
}
