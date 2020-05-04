package list

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"gitlab.platformer.com/project-x/platformer-cli/internal/auth"
	"gitlab.platformer.com/project-x/platformer-cli/internal/cli"
)

var organizationListCmd = &cobra.Command{
	Use:     "organizations",
	Aliases: []string{"organization", "org", "orgs"},
	Run: func(cmd *cobra.Command, args []string) {
		cli.HandleErrorAndExit(printOrgList())
	},
}

func printOrgList() error {
	orgList, err := auth.LoadOrganizationList()
	if err != nil {
		return &cli.InternalError{Err: err, Message: "failed to load organization data"}
	}

	fmt.Println(strings.Join(orgList.Names(), "\n"))
	return nil
}