package cmd

import (
	"code.cloudfoundry.org/cli/plugin"
	"github.com/gmllt-samples/cf-sample-plugin/utils"
)

type ListAppsCommand struct{}

// Run list all deployed apps
func (c ListAppsCommand) Run(cliConnection plugin.CliConnection, args []string) {
	apps, err := cliConnection.GetApps()
	if err != nil {
		utils.HandleError(err)
	}

	headers := []string{"App Name", "State"}
	var rows [][]string
	for _, app := range apps {
		rows = append(rows, []string{app.Name, app.State})
	}

	utils.PrintTable(headers, rows)
}
