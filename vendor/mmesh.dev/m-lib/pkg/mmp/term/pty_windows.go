//go:build windows
// +build windows

package term

import (
	"io"
	"os/exec"

	"mmesh.dev/m-lib/pkg/mmp/term/pty"
)

func PTYRun(c *exec.Cmd, inrp *io.PipeReader, outwp *io.PipeWriter) error {
	// Start the command with a pty.
	ptmx, err := pty.Start(c)
	if err != nil {
		return err
	}
	// Make sure to close the pty at the end.
	defer func() { _ = ptmx.Close() }() // Best effort.

	// Copy stdin to the pty and the pty to stdout.
	go io.Copy(ptmx, inrp)
	go io.Copy(outwp, ptmx)

	c.Wait()

	return nil
}
