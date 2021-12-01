//go:build windows
// +build windows

package colors

import (
	"golang.org/x/sys/windows"
	"mmesh.dev/m-lib/pkg/errors"
)

// EnableWindowsVirtualTerminalProcessing Enable Console VT Processing
func EnableWindowsVirtualTerminalProcessing() error {
	console, err := windows.GetStdHandle(windows.STD_OUTPUT_HANDLE)
	if err != nil {
		return errors.Wrapf(err, "[%v] function windows.GetStdHandle()", errors.Trace())
	}

	var consoleMode uint32
	if err := windows.GetConsoleMode(console, &consoleMode); err != nil {
		return errors.Wrapf(err, "[%v] function windows.GetConsoleMode()", errors.Trace())
	}

	if err := windows.SetConsoleMode(console, consoleMode|windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING); err != nil {
		return errors.Wrapf(err, "[%v] function windows.SetConsoleMode()", errors.Trace())
	}

	return nil
}
