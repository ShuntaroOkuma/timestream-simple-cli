package preset

import (
	"context"
	"fmt"
	"timestream-simple-cli/types"
	"timestream-simple-cli/usecase"
	"timestream-simple-cli/usecase/input"

	"github.com/spf13/cobra"
)

type presetHandler struct {
	presetInteractor usecase.PresetInteractor
}

func (h *presetHandler) GenerateSampleData(
	ctx context.Context,
	cmd *cobra.Command,
) (string, error) {
	sampleDataTypeString, err := cmd.Flags().GetString("type")
	if err != nil {
		return "", err
	}
	if sampleDataTypeString == "" {
		// use default type which is home
		sampleDataTypeString = "home"
	}

	var sampleDataType types.SampleDataType
	switch sampleDataTypeString {
	case "home":
		sampleDataType = types.SampleDataTypeHome
	case "building":
		sampleDataType = types.SampleDataTypeBuilding
	default:
		return "", fmt.Errorf("invalid preset type: %s", sampleDataTypeString)
	}

	result, err := h.presetInteractor.GenerateSampleData(ctx, input.NewGenerateSampleData(sampleDataType))
	if err != nil {
		return "", err
	}
	return result, nil
}
