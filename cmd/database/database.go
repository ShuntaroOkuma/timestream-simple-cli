package database

import (
	"context"
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
