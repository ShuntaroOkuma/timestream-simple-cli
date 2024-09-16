package types

type Tag struct {
	Key   *string `json:"key"`
	Value *string `json:"value"`
}

type Tags []Tag
