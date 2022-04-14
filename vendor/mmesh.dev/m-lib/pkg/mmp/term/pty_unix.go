//go:build linux || darwin
// +build linux darwin

package term

import (
	"io"
	"os/exec"

	"github.com/creack/pty"
	// "github.com/donorp/pty"
	// "github.com/photostorm/pty"
)

func PTYRun(c *exec.Cmd, inrp *io.PipeReader, outwp *io.PipeWriter) error {
	// Start the command with a pty.
	ptmx, err := pty.Start(c)
	if err != nil {
		return err
	}
	// Make sure to close the pty at the end.
	defer func() { _ = ptmx.Close() }() // Best effort.

	// Handle pty size.
	// ch := make(chan os.Signal, 1)
	// signal.Notify(ch, syscall.SIGWINCH)
	// go func() {
	// 	for range ch {
	// 		if err := pty.InheritSize(os.Stdin, ptmx); err != nil {
	// 			log.Printf("error resizing pty: %s", err)
	// 		}
	// 	}
	// }()
	// ch <- syscall.SIGWINCH // Initial resize.

	// Copy stdin to the pty and the pty to stdout.
	go io.Copy(ptmx, inrp)
	go io.Copy(outwp, ptmx)

	c.Wait()

	return nil
}
