//go:build windows
// +build windows

package kvstore

import (
	"fmt"
	"os"
)

func dbDir() string {
	programFiles := os.Getenv("ProgramFiles")

	if len(programFiles) == 0 {
		programFiles = `C:\Program Files`
	}

	return fmt.Sprintf(`%s\mmesh\db`, programFiles)
}
