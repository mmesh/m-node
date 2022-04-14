package netp2p

import (
	// "os"
	"os/exec"
	"strings"
)

func execCommand(name, sargs string) error {
	args := strings.Split(sargs, " ")
	cmd := exec.Command(name, args...)
	// cmd.Stderr = os.Stderr
	// cmd.Stdout = os.Stdout

	// log.Printf("[command] %s %s", name, sargs)

	return cmd.Run()
}
