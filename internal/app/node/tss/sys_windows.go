//go:build windows
// +build windows

package tss

import (
	"fmt"
	"os"
)

func tssDir() string {
	programFiles := os.Getenv("ProgramFiles")

	if len(programFiles) == 0 {
		programFiles = `C:\Program Files`
	}

	return fmt.Sprintf(`%s\mmesh\tss`, programFiles)
}
