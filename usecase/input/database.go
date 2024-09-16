package input

import (
	"timestream-simple-cli/pkg/nullable"
	"timestream-simple-cli/pkg/validation"
	"timestream-simple-cli/types"
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
	Tags         nullable.Type[types.Tags]
}

func NewCreateDatabase(
	databaseName string,
	tags nullable.Type[types.Tags],
) *CreateDatabase {
	return &CreateDatabase{
		DatabaseName: databaseName,
		Tags:         tags,
	}
}

func (p *CreateDatabase) Validate() error {
	if err := validation.Validate(p); err != nil {
		return err
	}
	return nil
}

type UpdateKMS struct {
	DatabaseName string `validate:"required"`
	KmsKeyId     string `validate:"required"`
}

func NewUpdateKMS(
	databaseName string,
	kmsKeyId string,
) *UpdateKMS {

	return &UpdateKMS{
		DatabaseName: databaseName,
		KmsKeyId:     kmsKeyId,
	}
}

func (p *UpdateKMS) Validate() error {
	if err := validation.Validate(p); err != nil {
		return err
	}
	return nil
}
