package data

import (
	"fmt"
	"timestream-simple-cli/dependency"
	"timestream-simple-cli/environment"

	"github.com/caarlos0/env/v10"
	"github.com/spf13/cobra"
)

func NewDataCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "data",
		Short: "data commands",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.HelpFunc()(cmd, args)
			}
		},
	}
	cmd.AddCommand(NewWriteDataCmd())
	return cmd
}

func NewWriteDataCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "write",
		Short: "write data",
		Run: func(cmd *cobra.Command, args []string) {
			e := &environment.Environment{}
			if err := env.Parse(e); err != nil {
				panic(err)
			}

			ctx := cmd.Context()
			d := &dependency.Dependency{}
			d.Inject(ctx, e)

			h := &dataHandler{
				dataInteractor: d.DataInteractor,
			}
			result, err := h.WriteData(ctx, cmd)
			if err != nil {
				fmt.Printf("data error: %v\n", err)
				return
			}
			fmt.Println("success:", result)
		},
	}
	cmd.Flags().StringP("database", "d", "", "database name")
	cmd.Flags().StringP("table", "t", "", "table name")
	cmd.Flags().StringP("schema-file", "s", "", "schema file path")
	cmd.Flags().StringP("value-file", "v", "", "value file path")
	return cmd
}
