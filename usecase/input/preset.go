package input

import (
	"timestream-simple-cli/pkg/validation"
	"timestream-simple-cli/types"
)

type GenerateSampleData struct {
	SampleType types.SampleDataType
}

func NewGenerateSampleData(
	sampleType types.SampleDataType,
) *GenerateSampleData {
	return &GenerateSampleData{
		SampleType: sampleType,
	}
}

func (p *GenerateSampleData) Validate() error {
	if err := validation.Validate(p); err != nil {
		return err
	}
	return nil
}
