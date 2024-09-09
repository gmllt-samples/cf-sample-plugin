package cmd

import (
	"code.cloudfoundry.org/cli/plugin"
	"fmt"
)

type SampleCommand struct{}

// Run execute a sample command without arguments
func (c SampleCommand) Run(cliConnection plugin.CliConnection, args []string) {
	fmt.Println("This is a sample command without arguments.")
}
