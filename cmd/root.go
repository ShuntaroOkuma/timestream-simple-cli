package cmd

import (
	"timestream-simple-cli/cmd/data"
	"timestream-simple-cli/cmd/database"
	"timestream-simple-cli/cmd/preset"
	"timestream-simple-cli/cmd/table"

	"github.com/spf13/cobra"
)

func NewCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ts",
		Short: "timestream simple cli",
		Long:  "timestream simple cli",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.HelpFunc()(cmd, args)
			}
		},
	}
	cmd.AddCommand(database.NewDatabaseCmd())
	cmd.AddCommand(table.NewTableCmd())
	cmd.AddCommand(data.NewDataCmd())
	cmd.AddCommand(preset.NewPresetCmd())
	return cmd
}
