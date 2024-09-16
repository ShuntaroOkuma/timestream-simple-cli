package usecase

import (
	"context"
	"timestream-simple-cli/marshaller"
	"timestream-simple-cli/usecase/input"

	"github.com/pkg/errors"

	"github.com/aws/aws-sdk-go-v2/service/kms"
	kmsTypes "github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/aws/aws-sdk-go-v2/service/timestreamquery"
	"github.com/aws/aws-sdk-go-v2/service/timestreamwrite"
	"github.com/aws/aws-sdk-go-v2/service/timestreamwrite/types"
	"github.com/aws/aws-sdk-go/aws"
)

type databaseInteractor struct {
	stsSvc   *sts.Client
	kmsSvc   *kms.Client
	WriteSvc *timestreamwrite.Client
	QuerySvc *timestreamquery.Client
}

func NewDatabaseInteractor(
	stsSvc *sts.Client,
	kmsSvc *kms.Client,
	writeSvc *timestreamwrite.Client,
	querySvc *timestreamquery.Client,
) DatabaseInteractor {
	return &databaseInteractor{
		stsSvc:   stsSvc,
		kmsSvc:   kmsSvc,
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
	tagKeyAutoAdded := "created-by"
	tagValueAutoAdded := "timestream-simple-cli"
	region := i.kmsSvc.Options().Region

	if err := param.Validate(); err != nil {
		return "", err
	}

	// Check if KMS key alias already exists
	aliasName := "alias/timestream-" + param.DatabaseName
	_, err := i.kmsSvc.DescribeKey(ctx, &kms.DescribeKeyInput{
		KeyId: aws.String(aliasName),
	})
	if err == nil {
		return "", errors.Errorf("KMS key alias %s already exists", aliasName)
	}

	// Check if database already exists
	_, err = i.WriteSvc.DescribeDatabase(ctx, &timestreamwrite.DescribeDatabaseInput{
		DatabaseName: aws.String(param.DatabaseName),
	})
	if err == nil {
		return "", errors.Errorf("Database %s already exists", param.DatabaseName)
	}

	// Get caller identity
	callerIdentity, err := i.stsSvc.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})
	if err != nil {
		return "", err
	}
	accountId := aws.StringValue(callerIdentity.Account)
	arn := aws.StringValue(callerIdentity.Arn)

	// Generate KMS key
	kmsTags := []kmsTypes.Tag{
		{TagKey: aws.String(tagKeyAutoAdded), TagValue: aws.String(tagValueAutoAdded)},
	}
	for _, tag := range param.Tags.Value() {
		kmsTags = append(kmsTags, kmsTypes.Tag{
			TagKey:   tag.Key,
			TagValue: tag.Value,
		})
	}

	// KMS Access Policy
	// This policy is based on the policy of the default kms key created by Timestream
	kmsAccessPolicy := `{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "Allow access through Amazon Timestream for all principals in the account that are authorized to use Amazon Timestream",
            "Effect": "Allow",
            "Principal": {
                "AWS": "*"
            },
            "Action": [
                "kms:Encrypt",
                "kms:Decrypt",
                "kms:ReEncrypt*",
                "kms:GenerateDataKey*",
                "kms:CreateGrant",
                "kms:DescribeKey"
            ],
            "Resource": "*",
            "Condition": {
                "StringEquals": {
                    "kms:ViaService": "timestream.` + region + `.amazonaws.com",
                    "kms:CallerAccount": "` + accountId + `"
                }
            }
        },
        {
            "Sid": "Allow direct access to key metadata to the account",
            "Effect": "Allow",
            "Principal": {
                "AWS": "` + arn + `"
            },
            "Action": [
                "kms:Describe*",
                "kms:Get*",
                "kms:List*",
                "kms:RevokeGrant",
								"kms:PutKeyPolicy",
								"kms:ScheduleKeyDeletion",
								"kms:CreateAlias"
            ],
            "Resource": "*"
        },
        {
            "Sid": "Allow the Timestream Service to describe the key directly",
            "Effect": "Allow",
            "Principal": {
                "Service": "timestream.amazonaws.com"
            },
            "Action": [
                "kms:Describe*",
                "kms:Get*",
                "kms:List*"
            ],
            "Resource": "*"
        }
    ]
  }`

	createKeyInput := &kms.CreateKeyInput{
		Description: aws.String("Timestream database encryption key for " + param.DatabaseName),
		KeySpec:     kmsTypes.KeySpecSymmetricDefault,
		Origin:      kmsTypes.OriginTypeAwsKms,
		MultiRegion: aws.Bool(false),
		Policy:      aws.String(kmsAccessPolicy),
		Tags:        kmsTags,
	}

	createKeyOutput, err := i.kmsSvc.CreateKey(ctx, createKeyInput)
	if err != nil {
		return "", err
	}

	keyId := createKeyOutput.KeyMetadata.KeyId

	// Add Key Alias
	createAliasInput := &kms.CreateAliasInput{
		AliasName:   aws.String(aliasName),
		TargetKeyId: keyId,
	}
	_, err = i.kmsSvc.CreateAlias(ctx, createAliasInput)
	if err != nil {
		_, errDeleteKey := i.kmsSvc.ScheduleKeyDeletion(ctx, &kms.ScheduleKeyDeletionInput{
			KeyId:               keyId,
			PendingWindowInDays: aws.Int32(7), // 7 days is the minimum
		})
		if errDeleteKey != nil {
			err = errors.Wrapf(err, "failed to delete KMS key: %v", errDeleteKey)
		}
		return "", err
	}

	// Create Database
	tsTags := []types.Tag{
		{Key: aws.String(tagKeyAutoAdded), Value: aws.String(tagValueAutoAdded)},
	}
	for _, tag := range param.Tags.Value() {
		tsTags = append(tsTags, types.Tag{
			Key:   tag.Key,
			Value: tag.Value,
		})
	}

	input := &timestreamwrite.CreateDatabaseInput{
		DatabaseName: aws.String(param.DatabaseName),
		KmsKeyId:     keyId,
		Tags:         tsTags,
	}

	output, err := i.WriteSvc.CreateDatabase(ctx, input)
	if err != nil {
		_, errDeleteKey := i.kmsSvc.ScheduleKeyDeletion(ctx, &kms.ScheduleKeyDeletionInput{
			KeyId:               keyId,
			PendingWindowInDays: aws.Int32(7), // 7 days is the minimum
		})
		if errDeleteKey != nil {
			err = errors.Wrapf(err, "failed to delete KMS key: %v", errDeleteKey)
		}
		return "", err
	}

	return marshaller.JsonMarshal(output), nil
}

func (i *databaseInteractor) UpdateKMS(
	ctx context.Context,
	param *input.UpdateKMS,
) (string, error) {
	if err := param.Validate(); err != nil {
		return "", err
	}

	input := &timestreamwrite.UpdateDatabaseInput{
		DatabaseName: aws.String(param.DatabaseName),
		KmsKeyId:     aws.String(param.KmsKeyId),
	}
	output, err := i.WriteSvc.UpdateDatabase(ctx, input)

	if err != nil {
		return "", err
	}

	return marshaller.JsonMarshal(output), nil
}
