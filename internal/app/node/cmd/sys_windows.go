//go:build windows
// +build windows

package cmd

import (
	"fmt"
	"os"
	"runtime"

	"mmesh.dev/m-lib/pkg/utils/colors"
)

func ConsoleInit() error {
	if runtime.GOOS == "windows" {
		if err := colors.EnableWindowsVirtualTerminalProcessing(); err != nil {
			return err
		}
	}

	return nil
}

func defaultConfigFile() string {
	programFiles := os.Getenv("ProgramFiles")

	if len(programFiles) == 0 {
		programFiles = `C:\Program Files`
	}

	return fmt.Sprintf(`%s\mmesh\mmesh-node.yml`, programFiles)
}
