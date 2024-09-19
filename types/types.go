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
	Dimensions []Dimension `json:"dimensions"`
	Measures   []Measure   `json:"measures"`
}

type Dimension struct {
	Name string `json:"name"`
	Type string `json:"type" validate:"eq=VARCHAR"`
}

type Measure struct {
	Name string `json:"name"`
	Type string `json:"type" validate:"eq=DOUBLE|eq=BIGINT|eq=BOOLEAN|eq=VARCHAR|eq=TIMESTAMP"`
}
