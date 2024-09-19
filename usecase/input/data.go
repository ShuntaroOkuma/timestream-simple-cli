package input

import (
	"timestream-simple-cli/pkg/validation"
	"timestream-simple-cli/types"
)

type WriteData struct {
	DatabaseName string              `validate:"required"`
	TableName    string              `validate:"required"`
	Schema       types.Schema        `validate:"required"`
	Values       []map[string]string `validate:"required"`
}

func NewWriteData(
	databaseName string,
	tableName string,
	schema types.Schema,
	values []map[string]string,
) *WriteData {
	return &WriteData{
		DatabaseName: databaseName,
		TableName:    tableName,
		Schema:       schema,
		Values:       values,
	}
}

func (p *WriteData) Validate() error {
	if err := validation.Validate(p); err != nil {
		return err
	}
	return nil
}
