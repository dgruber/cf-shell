package cfcli

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var extensions = [][]string{
	{"quit", "cf-shell: Exits the shell"},
	{"exit", "cf-shell: Exits the shell"},
	{"ls", "cf-shell: ls ..."},
	{"dir", "cf-shell: dir ..‚"},
	{"pwd", "cf-shell: pwd"},
}

func isPluginCommand(s string) bool {
	parts := strings.Split(s, " ")
	for _, cmd := range extensions {
		if parts[0] == cmd[0] {
			return true
		}
	}
	return false
}

func executePluginCommand(s string) {
	parts := strings.Split(s, " ")
	switch parts[0] {
	case "quit":
		fmt.Println("exiting cf-shell")
		os.Exit(0)
	case "exit":
		fmt.Println("exiting cf-shell")
		os.Exit(0)
	case "ls":
		execute(s)
	case "dir":
		execute(s)
	case "pwd":
		execute(s)
	default:
		fmt.Printf("Unrecognized command (%s)\n", s)
	}
}

func execute(s string) {
	//args := append([]string{"-c"}, strings.Split(s, " ")...)
	args := strings.Split(s, " ")
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error during executing %s: %s\n", cmd.Path, err.Error())
	}
}
