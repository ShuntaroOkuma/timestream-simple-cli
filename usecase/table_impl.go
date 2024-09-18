package usecase

import (
	"context"
	"timestream-simple-cli/marshaller"
	"timestream-simple-cli/usecase/input"

	"github.com/aws/aws-sdk-go-v2/service/timestreamwrite"
	"github.com/aws/aws-sdk-go/aws"
)

type tableInteractor struct {
	WriteSvc *timestreamwrite.Client
}

func NewTableInteractor(
	writeSvc *timestreamwrite.Client,
) TableInteractor {
	return &tableInteractor{
		WriteSvc: writeSvc,
	}
}

func (i *tableInteractor) DescribeTable(
	ctx context.Context,
	param *input.DescribeTable,
) (string, error) {
	if err := param.Validate(); err != nil {
		return "", err
	}

	input := &timestreamwrite.DescribeTableInput{
		DatabaseName: aws.String(param.DatabaseName),
		TableName:    aws.String(param.TableName),
	}
	output, err := i.WriteSvc.DescribeTable(ctx, input)

	if err != nil {
		return "", err
	}

	return marshaller.JsonMarshal(output), nil
}
