package cmd

import (
	"os/exec"
	"strings"
)

// CmdExecute runs the external commands
func CmdExecute(command string) ([]byte, error) {
	cmdArgs := strings.Fields(command)
	cmd := exec.Command(cmdArgs[0], cmdArgs[1:len(cmdArgs)]...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return []byte{}, err
	}
	return out, nil
}
