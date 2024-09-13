package usecase

import (
	"context"
	"timestream-simple-cli/usecase/input"
	"timestream-simple-cli/utils"

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

	return utils.JsonMarshal(output), nil

	// return output, nil
}

// func (i *databaseInteractor) ListDatabases(maxResultCount int32) error {
// 	listDatabasesMaxResult := maxResultCount
// 	var nextToken *string = nil

// 	for ok := true; ok; ok = nextToken != nil {
// 		listDatabasesInput := &timestreamwrite.ListDatabasesInput{
// 			MaxResults: &listDatabasesMaxResult,
// 		}
// 		if nextToken != nil {
// 			listDatabasesInput.NextToken = aws.String(*nextToken)
// 		}

// 		listDatabasesOutput, err := i.WriteSvc.ListDatabases(ctx, listDatabasesInput)

// 		if err != nil {
// 			fmt.Println("Error:")
// 			fmt.Println(err)
// 			return err
// 		} else {
// 			fmt.Println("List databases is successful, below is the output:")
// 			for _, database := range listDatabasesOutput.Databases {
// 				printDatabaseDetails(database)
// 			}
// 		}
// 		nextToken = listDatabasesOutput.NextToken
// 	}
// 	return nil
// }

// func (i *databaseInteractor) CreateDatabase(databaseName string) error {
// 	createDatabaseInput := &timestreamwrite.CreateDatabaseInput{
// 		DatabaseName: aws.String(databaseName),
// 	}
// 	createDatabaseOutput, err := i.WriteSvc.CreateDatabase(ctx, createDatabaseInput)

// 	if err != nil {
// 		var apiErr smithy.APIError
// 		if errors.As(err, &apiErr) {
// 			switch apiErr.ErrorCode() {
// 			case "ResourceNotFoundException":
// 				fmt.Println("ResourceNotFoundException", apiErr.Error())
// 			default:
// 				fmt.Printf("Error: %s", err.Error())
// 			}
// 		} else {
// 			fmt.Printf("Error: %s", err.Error())
// 		}
// 	} else {
// 		fmt.Printf("Database with name %s successfully created : %s\n", databaseName, JsonMarshalIgnoreError(createDatabaseOutput))
// 	}
// 	return err
// }

// func (i *databaseInteractor) UpdateDatabase(databaseName *string, kmsKeyId *string) error {
// 	updateDatabaseInput := &timestreamwrite.UpdateDatabaseInput{
// 		DatabaseName: aws.String(*databaseName),
// 		KmsKeyId:     aws.String(*kmsKeyId),
// 	}

// 	updateDatabaseOutput, err := i.WriteSvc.UpdateDatabase(ctx, updateDatabaseInput)

// 	if err != nil {
// 		fmt.Printf("Error: %s", err.Error())
// 	} else {
// 		fmt.Printf("Update database is successful")
// 		printDatabaseDetails(*updateDatabaseOutput.Database)
// 	}
// 	return err
// }
