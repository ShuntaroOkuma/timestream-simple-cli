package usecase

import (
	"context"
	"fmt"
	"time"
	"timestream-simple-cli/marshaller"
	"timestream-simple-cli/pkg/generator"
	"timestream-simple-cli/types"
	"timestream-simple-cli/usecase/input"
)

type presetInteractor struct{}

func NewPresetInteractor() PresetInteractor {
	return &presetInteractor{}
}

func (i *presetInteractor) GenerateSampleData(
	ctx context.Context,
	param *input.GenerateSampleData,
) (string, error) {
	if err := param.Validate(); err != nil {
		return "", err
	}

	var output string
	var err error
	switch param.SampleType {
	case types.SampleDataTypeHome:
		// startTime is 1 hour ago from now
		// data length is 10
		res := generator.GenerateHomeData("device-001", "Tokyo", time.Now().Add(-time.Hour), 10)
		output = marshaller.JsonMarshal(res)
		err = nil
	case types.SampleDataTypeBuilding:
		// startTime is 1 hour ago from now
		// data length is 10
		res := generator.GenerateSmartBuildingData("Building-A", "5", time.Now().Add(-time.Hour), 10)
		output = marshaller.JsonMarshal(res)
		err = nil
	default:
		output = ""
		err = fmt.Errorf("invalid sample type: %s", param.SampleType)
	}

	return output, err
}
