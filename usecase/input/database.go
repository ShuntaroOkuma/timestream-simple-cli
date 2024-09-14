package input

import (
	"timestream-simple-cli/pkg/validation"

	"github.com/volatiletech/null/v8"
)

type DescribeDatabase struct {
	DatabaseName string `validate:"required"`
}

func NewDescribeDatabase(databaseName string) *DescribeDatabase {
	return &DescribeDatabase{
		DatabaseName: databaseName,
	}
}

func (p *DescribeDatabase) Validate() error {
	if err := validation.Validate(p); err != nil {
		return err
	}
	return nil
}

type CreateDatabase struct {
	DatabaseName string `validate:"required"`
	KmsKeyId     null.String
}

func NewCreateDatabase(
	databaseName string,
	kmsKeyId null.String,
) *CreateDatabase {

	return &CreateDatabase{
		DatabaseName: databaseName,
		KmsKeyId:     kmsKeyId,
	}
}

func (p *CreateDatabase) Validate() error {
	if err := validation.Validate(p); err != nil {
		return err
	}
	return nil
}
