package database

import (
	"context"
	"encoding/json"
	"timestream-simple-cli/pkg/nullable"
	"timestream-simple-cli/types"
	"timestream-simple-cli/usecase"
	"timestream-simple-cli/usecase/input"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
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
	if databaseName == "" {
		return "", errors.Errorf("-name or -n param is required")
	}

	tags, err := cmd.Flags().GetString("tags")
	if err != nil {
		return "", err
	}

	// parse --tags input
	var tagsSlice types.Tags
	if tags != "" {
		if err := json.Unmarshal([]byte(tags), &tagsSlice); err != nil {
			return "", errors.Wrapf(err, "failed to parse tags")
		}
	}

	result, err := h.databaseInteractor.CreateDatabase(ctx, input.NewCreateDatabase(
		databaseName,
		nullable.TypeFrom(tagsSlice),
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
