//go:build windows
// +build windows

package hsec

import (
	"fmt"
	"os"
)

func reportFile() string {
	programFiles := os.Getenv("ProgramFiles")

	if len(programFiles) == 0 {
		programFiles = `C:\Program Files`
	}

	return fmt.Sprintf(`%s\mmesh\report.hsr`, programFiles)
}

func rootTargetDir() string {
	return `C:\`
}

func skipDirs() []string {
	return []string{}
}

func globalCacheDir() string {
	programFiles := os.Getenv("ProgramFiles")

	if len(programFiles) == 0 {
		programFiles = `C:\Program Files`
	}

	return fmt.Sprintf(`%s\mmesh\cache`, programFiles)
}
