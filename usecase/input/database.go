package input

import "timestream-simple-cli/utils"

type DescribeDatabase struct {
	DatabaseName string `validate:"required"`
}

func NewDescribeDatabase(databaseName string) *DescribeDatabase {
	return &DescribeDatabase{
		DatabaseName: databaseName,
	}
}

func (p *DescribeDatabase) Validate() error {
	if err := utils.Validate(p); err != nil {
		return err
	}
	return nil
}
