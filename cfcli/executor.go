package cfcli

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func Executor(s string) {
	s = strings.TrimSpace(s)
	if s == "" {
		return
	}
	if isPluginCommand(s) {
		executePluginCommand(s)
	} else {
		cmd := exec.Command("/bin/sh", "-c", "cf "+s)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Printf("Error during executing %s: %s\n", cmd.Path, err.Error())
		}
	}
}
