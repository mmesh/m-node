package utils

import (
	"os/exec"
	"strings"
)

func execCommand(name, sargs string) error {
	args := strings.Split(sargs, " ")
	cmd := exec.Command(name, args...)

	return cmd.Run()
}
