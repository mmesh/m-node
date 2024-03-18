//go:build linux
// +build linux

package hsec

func reportFile() string {
	return "/var/lib/mmesh/report.hsr"
}

func rootTargetDir() string {
	return "/"
}

func skipDirs() []string {
	return []string{
		"/proc/",
		"/run/",
		"/srv/",
		"/mnt/",
		"/var/lib/docker/",
	}
}

func globalCacheDir() string {
	return "/var/cache/mmesh"
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
