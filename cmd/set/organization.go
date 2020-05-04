package set

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"gitlab.platformer.com/project-x/platformer-cli/internal/auth"
	"gitlab.platformer.com/project-x/platformer-cli/internal/cli"
	"gitlab.platformer.com/project-x/platformer-cli/internal/config"
)

var organizationSetCmd = &cobra.Command{
	Use:     "organization",
	Aliases: []string{"org"},
	Args:    cobra.ExactArgs(1), // exactly 1 argument required (orgName)
	Run: func(cmd *cobra.Command, args []string) {
		orgName := args[0]
		cli.HandleErrorAndExit(validateAndSetOrg(orgName))
	},
}

func validateAndSetOrg(orgName string) error {
	orgList, err := auth.LoadOrganizationList()
	if err != nil {
		return &cli.InternalError{Err: err, Message: "failed to load organization data"}
	}

	if _, ok := orgList[orgName]; !ok {
		return &cli.UserError{"invalid Organization name: " + orgName}
	}

	config.SetDefaultOrg(orgName)

	green := color.New(color.FgHiGreen).SprintfFunc()
	fmt.Printf("%s has been set as your default organization\n", green(orgName))

	return nil
}