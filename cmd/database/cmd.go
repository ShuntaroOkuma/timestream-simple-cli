package database

import (
	"fmt"
	"timestream-simple-cli/dependency"
	"timestream-simple-cli/utils"

	"github.com/caarlos0/env/v10"
	"github.com/spf13/cobra"
)

func NewDatabaseCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "database",
		Short: "crud database",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.HelpFunc()(cmd, args)
			}
		},
	}
	cmd.AddCommand(NewDescribeDatabaseCmd())
	return cmd
}

func NewDescribeDatabaseCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "describe",
		Short: "describe database",
		Run: func(cmd *cobra.Command, args []string) {
			e := &utils.Environment{}
			if err := env.Parse(e); err != nil {
				panic(err)
			}

			ctx := cmd.Context()
			d := &dependency.Dependency{}
			d.Inject(ctx, e)

			h := &databaseHandler{
				databaseInteractor: d.DatabaseInteractor,
			}
			result, err := h.DescribeDatabase(ctx, cmd)
			if err != nil {
				fmt.Printf("database error: %v\n", err)
				return
			}
			fmt.Println("success:", result)
		},
	}
	cmd.Flags().StringP("name", "n", "", "database name")
	return cmd
}
