package database

import (
	"context"
	"timestream-simple-cli/usecase"
	"timestream-simple-cli/usecase/input"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/volatiletech/null/v8"
)

type databaseHandler struct {
	databaseInteractor usecase.DatabaseInteractor
}

func (h *databaseHandler) DescribeDatabase(ctx context.Context, cmd *cobra.Command) (string, error) {
	databaseName, err := cmd.Flags().GetString("name")
	if err != nil {
		return "", err
	}

	if databaseName == "" {
		return "", errors.Errorf("-name or -n param is required")
	}

	result, err := h.databaseInteractor.DescribeDatabase(ctx, input.NewDescribeDatabase(databaseName))
	if err != nil {
		return "", err
	}
	return result, nil
}

func (h *databaseHandler) CreateDatabase(ctx context.Context, cmd *cobra.Command) (string, error) {
	databaseName, err := cmd.Flags().GetString("name")
	if err != nil {
		return "", err
	}

	KmsKeyId, err := cmd.Flags().GetString("kms-key-id")
	if err != nil {
		return "", err
	}

	if databaseName == "" {
		return "", errors.Errorf("-name or -n param is required")
	}

	result, err := h.databaseInteractor.CreateDatabase(ctx, input.NewCreateDatabase(
		databaseName,
		null.StringFrom(KmsKeyId),
	))
	if err != nil {
		return "", err
	}
	return result, nil
}

func (h *databaseHandler) UpdateKMS(ctx context.Context, cmd *cobra.Command) (string, error) {
	databaseName, err := cmd.Flags().GetString("name")
	if err != nil {
		return "", err
	}

	KmsKeyId, err := cmd.Flags().GetString("kms-key-id")
	if err != nil {
		return "", err
	}

	if databaseName == "" && KmsKeyId == "" {
		return "", errors.Errorf("at least one of params is required")
	}

	result, err := h.databaseInteractor.UpdateKMS(ctx, input.NewUpdateKMS(
		databaseName,
		KmsKeyId,
	))
	if err != nil {
		return "", err
	}
	return result, nil
}
