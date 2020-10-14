package generate

import (
	"fmt"

	"github.com/platformercloud/platformer-cli/internal/cli"

	"github.com/spf13/cobra"
)

var (
	providerFlag string

	CICDCmd = &cobra.Command{
		Use:     "cicd",
		Aliases: []string{"ci"},
		Run: func(cmd *cobra.Command, args []string) {
			cli.HandleErrorAndExit(generateFile())
		},
	}
)

func init() {
	CICDCmd.Flags().StringVarP(&providerFlag, "provider", "p", "", "--provider=<PROVIDER_NAME> or -p <PROVIDER_NAME")
}

func generateFile() error {
	fmt.Println("Generate CI/CD")
	fmt.Println(providerFlag)
	return nil
}
