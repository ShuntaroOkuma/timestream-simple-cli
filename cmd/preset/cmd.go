package preset

import (
	"fmt"
	"timestream-simple-cli/dependency"
	"timestream-simple-cli/environment"

	"github.com/caarlos0/env/v10"
	"github.com/spf13/cobra"
)

func NewPresetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "preset",
		Short: "preset commands",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.HelpFunc()(cmd, args)
			}
		},
	}
	cmd.AddCommand(NewGenerateSampleDataCmd())
	return cmd
}

func NewGenerateSampleDataCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "data",
		Short: "generate sample data",
		Run: func(cmd *cobra.Command, args []string) {
			e := &environment.Environment{}
			if err := env.Parse(e); err != nil {
				panic(err)
			}

			ctx := cmd.Context()
			d := &dependency.Dependency{}
			d.Inject(ctx, e)

			h := &presetHandler{
				presetInteractor: d.PresetInteractor,
			}
			result, err := h.GenerateSampleData(ctx, cmd)
			if err != nil {
				fmt.Printf("generate sample data error: %v\n", err)
				return
			}
			fmt.Println(result)
		},
	}
	cmd.Flags().StringP("type", "t", "", "sample type, home or building")
	return cmd
}

// func NewGenerateSampleDatabaseCmd() *cobra.Command {}
// func NewGenerateSampleTableCmd() *cobra.Command {}
