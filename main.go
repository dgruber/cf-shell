package main

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"github.com/dgruber/cf-shell/cfcli"

	"code.cloudfoundry.org/cli/plugin"
)

func getTarget(cliConnection plugin.CliConnection) (string, bool) {
	org, err := cliConnection.GetCurrentOrg()
	if err != nil {
		return "", false
	}
	space, err := cliConnection.GetCurrentSpace()
	if err != nil {
		return "", false
	}
	if space.Name == "" && org.Name == "" {
		return "", false
	}
	return fmt.Sprintf("%s|%s", org.Name, space.Name), true
}

func shell(target string, hasTarget bool) {
	prefix := fmt.Sprintf("cf %c ", 8594)
	if hasTarget {
		prefix = fmt.Sprintf("[%s] %c ", target, 8594)
	}
	p := prompt.New(
		cfcli.Executor,
		cfcli.Completer,
		prompt.OptionTitle("cf shell: interactive Cloud Foundry command shell"),
		prompt.OptionPrefix(prefix),
		prompt.OptionInputTextColor(prompt.Blue),
	)
	p.Run()
	fmt.Println("exiting cf-shell")
}

type Shell struct{}

func (c *Shell) Run(cliConnection plugin.CliConnection, args []string) {
	if args[0] == "shell" {
		fmt.Printf("Starting interactive shell...\n")
		cfcli.SetCFContext(cliConnection)
		shell(getTarget(cliConnection))
	}
}

func (c *Shell) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "Shell",
		Version: plugin.VersionType{
			Major: 0,
			Minor: 1,
			Build: 0,
		},
		Commands: []plugin.Command{
			{
				Name:     "shell",
				HelpText: "shell starts an interactive Cloud Foundry shell with auto-completion",
				UsageDetails: plugin.Usage{
					Usage: "shell - interactive shell for cf cli\n   cf shell",
				},
			},
		},
	}
}

func main() {
	plugin.Start(new(Shell))
}
