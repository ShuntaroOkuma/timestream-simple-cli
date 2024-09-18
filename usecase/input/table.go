package input

import "timestream-simple-cli/pkg/validation"

type DescribeTable struct {
	DatabaseName string `validate:"required"`
	TableName    string `validate:"required"`
}

func NewDescribeTable(
	databaseName string,
	tableName string,
) *DescribeTable {
	return &DescribeTable{
		DatabaseName: databaseName,
		TableName:    tableName,
	}
}

func (p *DescribeTable) Validate() error {
	if err := validation.Validate(p); err != nil {
		return err
	}
	return nil
}
