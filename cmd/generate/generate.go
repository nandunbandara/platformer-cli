package generate

import (
	"github.com/platformercloud/platformer-cli/internal/cli"
	"github.com/spf13/cobra"
)

var GenerateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate resource files",
	ValidArgs: []string{
		CICDCmd.Use,
	},
	ArgAliases: append(
		CICDCmd.Aliases,
	),
	Args: cobra.ExactValidArgs(1),
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		cli.HandleErrorAndExit(func() error {
			// if !auth.isLoggedIn() {
			// 	return &cli.NotLoggedInError{}
			// }
			return nil
		}())
	},
}

func init() {
	GenerateCmd.AddCommand(CICDCmd)
}
