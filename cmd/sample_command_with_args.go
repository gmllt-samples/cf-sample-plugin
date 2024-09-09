package cmd

import (
	"code.cloudfoundry.org/cli/plugin"
	"fmt"
	"github.com/gmllt-samples/cf-sample-plugin/utils"
)

type SampleCommandWithArgs struct{}

// Run execute a sample command that takes arguments
func (c SampleCommandWithArgs) Run(cliConnection plugin.CliConnection, args []string) {
	utils.ValidateArgs(args, 2) // Verify that the number of arguments is correct

	fmt.Printf("Argument 1: %s, Argument 2: %s\n", args[0], args[1])
}
