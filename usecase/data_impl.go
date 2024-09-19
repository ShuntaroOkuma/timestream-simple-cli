package usecase

import (
	"context"
	"timestream-simple-cli/marshaller"
	"timestream-simple-cli/pkg/generator"
	"timestream-simple-cli/usecase/input"

	"github.com/aws/aws-sdk-go-v2/service/timestreamwrite"
)

type dataInteractor struct {
	WriteSvc *timestreamwrite.Client
}

func NewDataInteractor(
	writeSvc *timestreamwrite.Client,
) DataInteractor {
	return &dataInteractor{
		WriteSvc: writeSvc,
	}
}

func (i *dataInteractor) WriteData(
	ctx context.Context,
	param *input.WriteData,
) (string, error) {
	if err := param.Validate(); err != nil {
		return "", err
	}

	writeRecordsInput, err := generator.GenerateWriteRecordsInput(
		param.DatabaseName,
		param.TableName,
		param.Schema,
		param.Values,
	)
	if err != nil {
		return "", err
	}

	output, err := i.WriteSvc.WriteRecords(ctx, writeRecordsInput)
	if err != nil {
		return "", err
	}

	return marshaller.JsonMarshal(output), nil
}
