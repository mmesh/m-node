//go:build darwin
// +build darwin

package hsec

func reportFile() string {
	return "/opt/mmesh/var/lib/report.hsr"
}

func rootTargetDir() string {
	return "/"
}

func skipDirs() []string {
	return []string{}
}

func globalCacheDir() string {
	return "/opt/mmesh/var/cache"
}

/*
func globalCacheDir() string {
	tmpDir, err := os.UserCacheDir()
	if err != nil {
		tmpDir = os.TempDir()
	}

	return filepath.Join(tmpDir, "mmesh", "cache")
}
*/
