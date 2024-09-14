package usecase

import (
	"context"
	"timestream-simple-cli/marshaller"
	"timestream-simple-cli/usecase/input"

	"github.com/aws/aws-sdk-go-v2/service/timestreamquery"
	"github.com/aws/aws-sdk-go-v2/service/timestreamwrite"
	"github.com/aws/aws-sdk-go/aws"
)

type databaseInteractor struct {
	WriteSvc *timestreamwrite.Client
	QuerySvc *timestreamquery.Client
}

func NewDatabaseInteractor(
	writeSvc *timestreamwrite.Client,
	querySvc *timestreamquery.Client,
) DatabaseInteractor {
	return &databaseInteractor{
		WriteSvc: writeSvc,
		QuerySvc: querySvc,
	}
}

func (i *databaseInteractor) DescribeDatabase(
	ctx context.Context,
	param *input.DescribeDatabase,
) (string, error) {
	if err := param.Validate(); err != nil {
		return "", err
	}

	input := &timestreamwrite.DescribeDatabaseInput{
		DatabaseName: aws.String(param.DatabaseName),
	}
	output, err := i.WriteSvc.DescribeDatabase(ctx, input)

	if err != nil {
		return "", err
	}

	return marshaller.JsonMarshal(output), nil
}

func (i *databaseInteractor) CreateDatabase(
	ctx context.Context,
	param *input.CreateDatabase,
) (string, error) {
	if err := param.Validate(); err != nil {
		return "", err
	}

	var input *timestreamwrite.CreateDatabaseInput
	if param.KmsKeyId.Valid {
		input = &timestreamwrite.CreateDatabaseInput{
			DatabaseName: aws.String(param.DatabaseName),
			KmsKeyId:     aws.String(param.KmsKeyId.String),
		}
	} else {
		input = &timestreamwrite.CreateDatabaseInput{
			DatabaseName: aws.String(param.DatabaseName),
		}
	}

	output, err := i.WriteSvc.CreateDatabase(ctx, input)

	if err != nil {
		return "", err
	}

	return marshaller.JsonMarshal(output), nil
}
