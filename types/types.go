package types

type Tag struct {
	Key   *string `json:"key"`
	Value *string `json:"value"`
}

type Tags []Tag

type SampleDataType string

const (
	SampleDataTypeHome     SampleDataType = "home"
	SampleDataTypeBuilding SampleDataType = "building"
)
