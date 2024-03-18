package hsec

import (
	"time"

	dbTypes "github.com/aquasecurity/trivy-db/pkg/types"
	ftypes "github.com/aquasecurity/trivy/pkg/fanal/types"
	"github.com/aquasecurity/trivy/pkg/types"
	"mmesh.dev/m-api-go/grpc/resources/hsec"
)

func getHostSecurityReport(r *types.Report) *hsec.Report {
	return &hsec.Report{
		SchemaVersion: int32(r.SchemaVersion),
		CreatedAt:     r.CreatedAt.UnixMilli(),
		ArtifactName:  r.ArtifactName,
		ArtifactType:  string(r.ArtifactType),
		Metadata: &hsec.Metadata{
			Size:        r.Metadata.Size,
			OS:          getReportMetadataOS(r.Metadata),
			ImageID:     r.Metadata.ImageID,
			DiffIDs:     r.Metadata.DiffIDs,
			RepoTags:    r.Metadata.RepoTags,
			RepoDigests: r.Metadata.RepoDigests,
			ImageConfig: &hsec.ConfigFile{
				Architecture:  r.Metadata.ImageConfig.Architecture,
				Author:        r.Metadata.ImageConfig.Author,
				Container:     r.Metadata.ImageConfig.Container,
				Created:       r.Metadata.ImageConfig.Created.UnixMilli(),
				DockerVersion: r.Metadata.ImageConfig.DockerVersion,
				History:       getReportMetadataImageConfigHistory(r.Metadata),
				OS:            r.Metadata.ImageConfig.OS,
				RootFS:        getReportMetadataImageConfigRootFS(r.Metadata),
				Config: &hsec.Config{
					AttachStderr:    r.Metadata.ImageConfig.Config.AttachStderr,
					AttachStdin:     r.Metadata.ImageConfig.Config.AttachStdin,
					AttachStdout:    r.Metadata.ImageConfig.Config.AttachStdout,
					Cmd:             r.Metadata.ImageConfig.Config.Cmd,
					Healthcheck:     getReportMetadataImageConfigConfigHealthcheck(r.Metadata),
					DomainName:      r.Metadata.ImageConfig.Config.Domainname,
					Entrypoint:      r.Metadata.ImageConfig.Config.Entrypoint,
					Env:             r.Metadata.ImageConfig.Config.Env,
					Hostname:        r.Metadata.ImageConfig.Config.Hostname,
					Image:           r.Metadata.ImageConfig.Config.Image,
					Labels:          r.Metadata.ImageConfig.Config.Labels,
					OnBuild:         r.Metadata.ImageConfig.Config.OnBuild,
					OpenStdin:       r.Metadata.ImageConfig.Config.OpenStdin,
					StdinOnce:       r.Metadata.ImageConfig.Config.StdinOnce,
					Tty:             r.Metadata.ImageConfig.Config.Tty,
					User:            r.Metadata.ImageConfig.Config.User,
					Volumes:         getMapStructsToStrings(r.Metadata.ImageConfig.Config.Volumes),
					WorkingDir:      r.Metadata.ImageConfig.Config.WorkingDir,
					ExposedPorts:    getMapStructsToStrings(r.Metadata.ImageConfig.Config.ExposedPorts),
					ArgsEscaped:     r.Metadata.ImageConfig.Config.ArgsEscaped,
					NetworkDisabled: r.Metadata.ImageConfig.Config.NetworkDisabled,
					MacAddress:      r.Metadata.ImageConfig.Config.MacAddress,
					StopSignal:      r.Metadata.ImageConfig.Config.StopSignal,
					Shell:           r.Metadata.ImageConfig.Config.Shell,
				},
				OSVersion:  r.Metadata.ImageConfig.OSVersion,
				Variant:    r.Metadata.ImageConfig.Variant,
				OSFeatures: r.Metadata.ImageConfig.OSFeatures,
			}, // ImageConfig
		}, // Metadata
		Results:   getReportResults(r),
		CycloneDX: getReportCycloneDX(r.CycloneDX),
	}
}

func getReportMetadataOS(metadata types.Metadata) *hsec.OS {
	if metadata.OS == nil {
		return nil
	}

	return &hsec.OS{
		Family:   string(metadata.OS.Family),
		Name:     metadata.OS.Name,
		EOSL:     metadata.OS.Eosl,
		Extended: metadata.OS.Extended,
	}
}

func getReportMetadataImageConfigHistory(metadata types.Metadata) []*hsec.History {
	history := make([]*hsec.History, 0)

	for _, h := range metadata.ImageConfig.History {
		history = append(history, &hsec.History{
			Author:     h.Author,
			Created:    h.Created.UnixMilli(),
			CreatedBy:  h.CreatedBy,
			Comment:    h.Comment,
			EmptyLayer: h.EmptyLayer,
		})
	}

	return history
}

func getReportMetadataImageConfigRootFS(metadata types.Metadata) *hsec.RootFS {
	rootFS := &hsec.RootFS{
		Type:    metadata.ImageConfig.RootFS.Type,
		DiffIDs: make([]*hsec.Hash, 0),
	}

	for _, d := range metadata.ImageConfig.RootFS.DiffIDs {
		rootFS.DiffIDs = append(rootFS.DiffIDs, &hsec.Hash{
			Algorithm: d.Algorithm,
			Hex:       d.Hex,
		})
	}

	return rootFS
}

func getReportMetadataImageConfigConfigHealthcheck(metadata types.Metadata) *hsec.HealthConfig {
	if metadata.ImageConfig.Config.Healthcheck == nil {
		return nil
	}

	return &hsec.HealthConfig{
		Test:        metadata.ImageConfig.Config.Healthcheck.Test,
		Interval:    metadata.ImageConfig.Config.Healthcheck.Interval.Nanoseconds(),
		Timeout:     metadata.ImageConfig.Config.Healthcheck.Timeout.Nanoseconds(),
		StartPeriod: metadata.ImageConfig.Config.Healthcheck.StartPeriod.Nanoseconds(),
		Retries:     int32(metadata.ImageConfig.Config.Healthcheck.Retries),
	}
}

func getMapStructsToStrings(m map[string]struct{}) map[string]string {
	ms := make(map[string]string, 0)

	for id, _ := range m {
		ms[id] = ""
	}

	return ms
}

// Results

func getReportResults(r *types.Report) []*hsec.Result {
	if r == nil {
		return nil
	}

	results := make([]*hsec.Result, 0)

	for _, result := range r.Results {
		results = append(results, &hsec.Result{
			Target:            result.Target,
			Class:             string(result.Class),
			Type:              string(result.Type),
			ScannedPackages:   int32(len(result.Packages)),
			Packages:          getReportResultPackages(result),
			Vulnerabilities:   getReportResultVulnerabilities(result),
			MisconfSummary:    getReportResultMisconfSummary(result.MisconfSummary),
			Misconfigurations: getReportResultMisconfigurations(result.Misconfigurations),
			Secrets:           getReportResultSecrets(result.Secrets),
			Licenses:          getReportResultLicenses(result.Licenses),
			CustomResources:   getReportResultCustomResources(result.CustomResources),
		})
	}

	return results
}

// Packages

func getReportResultPackages(r types.Result) []*hsec.Package {
	pkgs := make([]*hsec.Package, 0)

	for _, p := range r.Packages {
		pkgs = append(pkgs, &hsec.Package{
			ID:              p.ID,
			Name:            p.Name,
			Identifier:      getPackageIdentifier(p.Identifier),
			Version:         p.Version,
			Release:         p.Release,
			Epoch:           int32(p.Epoch),
			Arch:            p.Arch,
			Dev:             p.Dev,
			SrcName:         p.SrcName,
			SrcVersion:      p.SrcVersion,
			SrcRelease:      p.SrcRelease,
			SrcEpoch:        int32(p.SrcEpoch),
			Licenses:        p.Licenses,
			Maintainer:      p.Maintainer,
			ModularityLabel: p.Modularitylabel,
			BuildInfo:       getPackageBuildInfo(p.BuildInfo),
			// Ref:             p.Ref,
			Indirect:  p.Indirect,
			DependsOn: p.DependsOn,
			Layer: &hsec.Layer{
				Digest:    p.Layer.Digest,
				DiffID:    p.Layer.DiffID,
				CreatedBy: p.Layer.CreatedBy,
			},
			FilePath:       p.FilePath,
			Digest:         string(p.Digest),
			Locations:      getLocations(p.Locations),
			InstalledFiles: p.InstalledFiles,
		})
	}

	return pkgs
}

func getPackageIdentifier(pkgIdentifier ftypes.PkgIdentifier) *hsec.PkgIdentifier {
	// if pkgIdentifier == nil {
	// 	return nil
	// }

	pkgURL := &hsec.PackageURL{}

	if pkgIdentifier.PURL != nil {
		pkgURL = &hsec.PackageURL{
			Type:       pkgIdentifier.PURL.Type,
			Namespace:  pkgIdentifier.PURL.Namespace,
			Name:       pkgIdentifier.PURL.Name,
			Version:    pkgIdentifier.PURL.Version,
			Qualifiers: make([]*hsec.Qualifier, 0),
			Subpath:    pkgIdentifier.PURL.Subpath,
		}

		for _, q := range pkgIdentifier.PURL.Qualifiers {
			pkgURL.Qualifiers = append(pkgURL.Qualifiers, &hsec.Qualifier{
				Key:   q.Key,
				Value: q.Value,
			})
		}
	}

	pkgID := &hsec.PkgIdentifier{
		PURL:   pkgURL,
		BOMRef: pkgIdentifier.BOMRef,
	}

	return pkgID
}

func getPackageBuildInfo(buildInfo *ftypes.BuildInfo) *hsec.BuildInfo {
	if buildInfo == nil {
		return nil
	}

	return &hsec.BuildInfo{
		ContentSets: buildInfo.ContentSets,
		Nvr:         buildInfo.Nvr,
		Arch:        buildInfo.Arch,
	}
}

func getLocations(locations []ftypes.Location) []*hsec.Location {
	hslocations := make([]*hsec.Location, 0)

	for _, l := range locations {
		hslocations = append(hslocations, &hsec.Location{
			StartLine: int32(l.StartLine),
			EndLine:   int32(l.EndLine),
		})
	}

	return hslocations
}

// Vulnerabilities

func getReportResultVulnerabilities(r types.Result) []*hsec.DetectedVulnerability {
	vulns := make([]*hsec.DetectedVulnerability, 0)

	for _, v := range r.Vulnerabilities {
		vulns = append(vulns, &hsec.DetectedVulnerability{
			VulnerabilityID:  v.VulnerabilityID,
			VendorIDs:        v.VendorIDs,
			PkgID:            v.PkgID,
			PkgName:          v.PkgName,
			PkgPath:          v.PkgPath,
			PkgIdentifier:    getPackageIdentifier(v.PkgIdentifier),
			InstalledVersion: v.InstalledVersion,
			FixedVersion:     v.FixedVersion,
			// Status:           getVulnerabilityStatus(v.Status),
			Status: v.Status.String(),
			Layer: &hsec.Layer{
				Digest:    v.Layer.Digest,
				DiffID:    v.Layer.DiffID,
				CreatedBy: v.Layer.CreatedBy,
			},
			SeveritySource: string(v.SeveritySource),
			PrimaryURL:     v.PrimaryURL,
			// PkgRef:         v.PkgRef,
			DataSource: getVulnerabilityDataSource(v.DataSource),
			Vulnerability: &hsec.Vulnerability{
				Title:            v.Title,
				Description:      v.Description,
				Severity:         v.Severity,
				CweIDs:           v.CweIDs,
				VendorSeverity:   getVulnerabilityVendorSeverity(v.VendorSeverity),
				CVSS:             getVulnerabilityVendorCVSS(v.CVSS),
				References:       v.References,
				PublishedDate:    getTime(v.PublishedDate),
				LastModifiedDate: getTime(v.LastModifiedDate),
			},
		})
	}

	return vulns
}

/*
func getVulnerabilityStatus(status dbTypes.Status) hsec.VulnerabilityStatus {
	switch status {
	case dbTypes.StatusUnknown:
		return hsec.VulnerabilityStatus_STATUS_UNKNOWN
	case dbTypes.StatusNotAffected:
		return hsec.VulnerabilityStatus_STATUS_NOT_AFFECTED
	case dbTypes.StatusAffected:
		return hsec.VulnerabilityStatus_STATUS_AFFECTED
	case dbTypes.StatusFixed:
		return hsec.VulnerabilityStatus_STATUS_FIXED
	case dbTypes.StatusUnderInvestigation:
		return hsec.VulnerabilityStatus_STATUS_UNDER_INVESTIGATION
	case dbTypes.StatusWillNotFix:
		return hsec.VulnerabilityStatus_STATUS_WILL_NOT_FIX
	case dbTypes.StatusFixDeferred:
		return hsec.VulnerabilityStatus_STATUS_FIX_DEFERRED
	case dbTypes.StatusEndOfLife:
		return hsec.VulnerabilityStatus_STATUS_END_OF_LIFE
	}

	return hsec.VulnerabilityStatus_STATUS_UNKNOWN
}
*/

func getVulnerabilityDataSource(ds *dbTypes.DataSource) *hsec.DataSource {
	if ds == nil {
		return nil
	}

	return &hsec.DataSource{
		ID:   string(ds.ID),
		Name: ds.Name,
		URL:  ds.URL,
	}
}

func getVulnerabilityVendorSeverity(vendorSeverity dbTypes.VendorSeverity) map[string]hsec.Severity {
	vs := make(map[string]hsec.Severity, 0)

	for sourceID, severity := range vendorSeverity {
		switch severity {
		case dbTypes.SeverityUnknown:
			vs[string(sourceID)] = hsec.Severity_SEVERITY_UNKNOWN
		case dbTypes.SeverityLow:
			vs[string(sourceID)] = hsec.Severity_SEVERITY_LOW
		case dbTypes.SeverityMedium:
			vs[string(sourceID)] = hsec.Severity_SEVERITY_MEDIUM
		case dbTypes.SeverityHigh:
			vs[string(sourceID)] = hsec.Severity_SEVERITY_HIGH
		case dbTypes.SeverityCritical:
			vs[string(sourceID)] = hsec.Severity_SEVERITY_CRITICAL
		}

		vs[string(sourceID)] = hsec.Severity_SEVERITY_UNKNOWN
	}

	return vs
}

func getVulnerabilityVendorCVSS(vendorCVSS dbTypes.VendorCVSS) map[string]*hsec.CVSS {
	vcvss := make(map[string]*hsec.CVSS, 0)

	for sourceID, cvss := range vendorCVSS {
		vcvss[string(sourceID)] = &hsec.CVSS{
			V2Vector: cvss.V2Vector,
			V3Vector: cvss.V3Vector,
			V2Score:  cvss.V2Score,
			V3Score:  cvss.V3Score,
		}
	}

	return vcvss
}

func getTime(tm *time.Time) int64 {
	if tm == nil {
		return 0
	}

	return tm.UnixMilli()
}

// MisconfSummary

func getReportResultMisconfSummary(ms *types.MisconfSummary) *hsec.MisconfSummary {
	if ms == nil {
		return nil
	}

	return &hsec.MisconfSummary{
		Successes:  int32(ms.Successes),
		Failures:   int32(ms.Failures),
		Exceptions: int32(ms.Exceptions),
	}
}

func getReportResultMisconfigurations(dmcfgs []types.DetectedMisconfiguration) []*hsec.DetectedMisconfiguration {
	misconfigs := make([]*hsec.DetectedMisconfiguration, 0)

	for _, m := range dmcfgs {
		misconfigs = append(misconfigs, &hsec.DetectedMisconfiguration{
			Type:        m.Type,
			ID:          m.ID,
			AVDID:       m.AVDID,
			Title:       m.Title,
			Description: m.Description,
			Message:     m.Message,
			Namespace:   m.Namespace,
			Query:       m.Query,
			Resolution:  m.Resolution,
			Severity:    m.Severity,
			PrimaryURL:  m.PrimaryURL,
			References:  m.References,
			Status:      getMisconfigStatus(m.Status),
			Layer: &hsec.Layer{
				Digest:    m.Layer.Digest,
				DiffID:    m.Layer.DiffID,
				CreatedBy: m.Layer.CreatedBy,
			},
			CauseMetadata: getCauseMetadata(m.CauseMetadata),
			Traces:        m.Traces,
		})
	}

	return misconfigs
}

func getMisconfigStatus(mstatus types.MisconfStatus) hsec.MisconfStatus {
	switch mstatus {
	case types.StatusPassed:
		return hsec.MisconfStatus_MISCONF_STATUS_PASSED
	case types.StatusFailure:
		return hsec.MisconfStatus_MISCONF_STATUS_FAILURE
	case types.StatusException:
		return hsec.MisconfStatus_MISCONF_STATUS_EXCEPTION
	}

	return hsec.MisconfStatus_MISCONF_STATUS_UNKNOWN
}

func getCauseMetadata(causeMeta ftypes.CauseMetadata) *hsec.CauseMetadata {
	return &hsec.CauseMetadata{
		Resource:    causeMeta.Resource,
		Provider:    causeMeta.Provider,
		Service:     causeMeta.Service,
		StartLine:   int32(causeMeta.StartLine),
		EndLine:     int32(causeMeta.EndLine),
		Code:        getCode(causeMeta.Code),
		Occurrences: getOccurrences(causeMeta.Occurrences),
	}
}

func getCode(c ftypes.Code) *hsec.Code {
	return &hsec.Code{
		Lines: getLines(c.Lines),
	}
}

func getLines(lines []ftypes.Line) []*hsec.Line {
	ls := make([]*hsec.Line, 0)

	for _, l := range lines {
		ls = append(ls, &hsec.Line{
			Number:      int32(l.Number),
			Content:     l.Content,
			IsCause:     l.IsCause,
			Annotation:  l.Annotation,
			Truncated:   l.Truncated,
			Highlighted: l.Highlighted,
			FirstCause:  l.FirstCause,
			LastCause:   l.LastCause,
		})
	}

	return ls
}

func getOccurrences(occurrences []ftypes.Occurrence) []*hsec.Occurrence {
	ocs := make([]*hsec.Occurrence, 0)

	for _, o := range occurrences {
		ocs = append(ocs, &hsec.Occurrence{
			Resource: o.Resource,
			Filename: o.Filename,
			Location: &hsec.Location{
				StartLine: int32(o.Location.StartLine),
				EndLine:   int32(o.Location.EndLine),
			},
		})
	}

	return ocs
}

// Secrets

func getReportResultSecrets(secrets []ftypes.SecretFinding) []*hsec.SecretFinding {
	ss := make([]*hsec.SecretFinding, 0)

	for _, s := range secrets {
		ss = append(ss, &hsec.SecretFinding{
			RuleID:    s.RuleID,
			Category:  string(s.Category),
			Severity:  s.Severity,
			Title:     s.Title,
			StartLine: int32(s.StartLine),
			EndLine:   int32(s.EndLine),
			Code:      getCode(s.Code),
			Match:     s.Match,
			Layer: &hsec.Layer{
				Digest:    s.Layer.Digest,
				DiffID:    s.Layer.DiffID,
				CreatedBy: s.Layer.CreatedBy,
			},
		})
	}

	return ss
}

// Licenses

func getReportResultLicenses(licenses []types.DetectedLicense) []*hsec.DetectedLicense {
	dls := make([]*hsec.DetectedLicense, 0)

	for _, l := range licenses {
		dls = append(dls, &hsec.DetectedLicense{
			Severity:   l.Severity,
			Category:   string(l.Category),
			PkgName:    l.PkgName,
			FilePath:   l.FilePath,
			Name:       l.Name,
			Confidence: l.Confidence,
			Link:       l.Link,
		})
	}

	return dls
}

// Custom Resources

func getReportResultCustomResources(customResources []ftypes.CustomResource) []*hsec.CustomResource {
	crs := make([]*hsec.CustomResource, 0)

	for _, cr := range customResources {
		crs = append(crs, &hsec.CustomResource{
			Type:     cr.Type,
			FilePath: cr.FilePath,
			Layer: &hsec.Layer{
				Digest:    cr.Layer.Digest,
				DiffID:    cr.Layer.DiffID,
				CreatedBy: cr.Layer.CreatedBy,
			},
		})
	}

	return crs
}

// CycloneDX

func getReportCycloneDX(cdx *ftypes.CycloneDX) *hsec.CycloneDX {
	if cdx == nil {
		return nil
	}

	return &hsec.CycloneDX{
		BOMFormat:    cdx.BOMFormat,
		SpecVersion:  int32(cdx.SpecVersion),
		SerialNumber: cdx.SerialNumber,
		Version:      int32(cdx.Version),
		Metadata: &hsec.BOMMetadata{
			Timestamp: cdx.Metadata.Timestamp,
			Component: &hsec.Component{
				BOMRef:     cdx.Metadata.Component.BOMRef,
				MIMEType:   cdx.Metadata.Component.MIMEType,
				Type:       string(cdx.Metadata.Component.Type),
				Name:       cdx.Metadata.Component.Name,
				Version:    cdx.Metadata.Component.Version,
				PackageURL: cdx.Metadata.Component.PackageURL,
			},
		},
		Components: getCycloneDXComponents(cdx.Components),
	}
}

func getCycloneDXComponents(components []ftypes.Component) []*hsec.Component {
	cs := make([]*hsec.Component, 0)

	for _, c := range components {
		cs = append(cs, &hsec.Component{
			BOMRef:     c.BOMRef,
			MIMEType:   c.MIMEType,
			Type:       string(c.Type),
			Name:       c.Name,
			Version:    c.Version,
			PackageURL: c.PackageURL,
		})
	}

	return cs
}
