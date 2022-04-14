package version

import (
	"fmt"
	"strings"

	"github.com/coreos/go-semver/semver"
)

var (
	// GitRepository is the git repository that was compiled
	GitRepository string
	// GitCommit is the git commit that was compiled
	GitCommit string
	// VersionMajor is for an API incompatible changes
	VersionMajor int64
	// VersionMinor is for functionality in a backwards-compatible manner
	VersionMinor int64
	// VersionPatch is for backwards-compatible bug fixes
	VersionPatch int64
	// VersionPre indicates prerelease
	VersionPre = ""
	// VersionDev indicates development branch. Releases will be empty string.
	VersionDev string
	// VersionDate is the compilation date
	VersionDate string
	// VersionNumber is a string with the format n.n.n
	VersionNumber string
	// GoVersion is a string with go-version
	GoVersion string
)

// Version is the specification version that the package types support.
var Version = semver.Version{
	Major:      VersionMajor,
	Minor:      VersionMinor,
	Patch:      VersionPatch,
	PreRelease: semver.PreRelease(VersionPre),
	Metadata:   VersionDev,
}

func init() {
	if strings.HasPrefix(VersionNumber, "v") {
		VersionNumber = strings.TrimPrefix(VersionNumber, "v")
	}

	if err := Version.Set(VersionNumber); err != nil {
		fmt.Println(err)
	}
}

func GetVersion() string {
	return "v" + Version.String() + "-" + VersionDate + "+" + GitCommit + "--" + GoVersion
}
