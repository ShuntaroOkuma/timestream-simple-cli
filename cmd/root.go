package cmd

import (
	"timestream-simple-cli/cmd/database"
	"timestream-simple-cli/cmd/table"

	"github.com/spf13/cobra"
)

func NewCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tt",
		Short: "timestream trial",
		Long:  "timestream trial",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.HelpFunc()(cmd, args)
			}
		},
	}
	cmd.AddCommand(database.NewDatabaseCmd())
	cmd.AddCommand(table.NewTableCmd())
	return cmd
}
