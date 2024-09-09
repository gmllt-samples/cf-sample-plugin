package main

import (
	"code.cloudfoundry.org/cli/plugin"
	"fmt"
	"github.com/gmllt-samples/cf-sample-plugin/cmd"
	"os"
	"strconv"
	"strings"
)

// Injected via GoReleaser at build time
var version = "v0.0.0" // default value

// PluginBase structure du plugin
type PluginBase struct{}

// Run est appelée lorsque le plugin est exécuté
func (p *PluginBase) Run(cliConnection plugin.CliConnection, args []string) {
	if len(args) == 0 {
		fmt.Println("No command provided. Use 'cf help' to see available commands.")
		os.Exit(1)
	}

	switch args[0] {
	case "hello":
		fmt.Println("Hello from the plugin!")
	case "list-apps":
		cmd.ListAppsCommand{}.Run(cliConnection, args)
	case "sample-command":
		cmd.SampleCommand{}.Run(cliConnection, args)
	case "sample-command-with-args":
		cmd.SampleCommandWithArgs{}.Run(cliConnection, args)
	default:
		fmt.Printf("Unknown command: %s\n", args[0])
	}
}

func (p *PluginBase) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "SamplePlugin",
		Version: plugin.VersionType{
			Major: parseVersionComponent(0),
			Minor: parseVersionComponent(1),
			Build: parseVersionComponent(2),
		},
		Commands: []plugin.Command{
			{
				Name:     "hello",
				HelpText: "Prints Hello message",
			},
			{
				Name:     "list-apps",
				HelpText: "Lists all deployed apps",
			},
		},
	}
}

// Helper function to extract version components
func parseVersionComponent(idx int) int {
	parts := strings.Split(version, ".")
	if idx >= len(parts) {
		return 0
	}
	part, _ := strconv.Atoi(parts[idx])
	return part
}
