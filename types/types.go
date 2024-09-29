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

type Schema struct {
	Dimensions       []Dimension `json:"dimensions" validate:"required,min=1,max=128"`
	Measures         []Measure   `json:"measures" validate:"required,min=1"`
	MultiMeasureName string      `json:"multiMeasureName"`
}

type Dimension struct {
	Name string `json:"name" validate:"required,min=1,max=60"`
	Type string `json:"type" validate:"eq=VARCHAR"`
}

type Measure struct {
	Name string `json:"name" validate:"required,min=1,max=256"`
	Type string `json:"type" validate:"required,eq=DOUBLE|eq=BIGINT|eq=BOOLEAN|eq=VARCHAR|eq=TIMESTAMP|eq=MULTI"`
}
