package table

import (
	"fmt"
	"timestream-simple-cli/dependency"
	"timestream-simple-cli/environment"

	"github.com/caarlos0/env/v10"
	"github.com/spf13/cobra"
)

func NewTableCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "table",
		Short: "table commands",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.HelpFunc()(cmd, args)
			}
		},
	}
	cmd.AddCommand(NewDescribeTableCmd())
	return cmd
}

func NewDescribeTableCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "describe",
		Short: "describe table",
		Run: func(cmd *cobra.Command, args []string) {
			e := &environment.Environment{}
			if err := env.Parse(e); err != nil {
				panic(err)
			}

			ctx := cmd.Context()
			d := &dependency.Dependency{}
			d.Inject(ctx, e)

			h := &tableHandler{
				tableInteractor: d.TableInteractor,
			}
			result, err := h.DescribeTable(ctx, cmd)
			if err != nil {
				fmt.Printf("table error: %v\n", err)
				return
			}
			fmt.Println("success:", result)
		},
	}
	cmd.Flags().StringP("database", "d", "", "database name")
	cmd.Flags().StringP("table", "t", "", "table name")
	return cmd
}
