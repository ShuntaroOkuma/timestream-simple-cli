package table

import (
	"context"
	"timestream-simple-cli/usecase"
	"timestream-simple-cli/usecase/input"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

type tableHandler struct {
	tableInteractor usecase.TableInteractor
}

func (h *tableHandler) DescribeTable(
	ctx context.Context,
	cmd *cobra.Command,
) (string, error) {
	databaseName, err := cmd.Flags().GetString("database")
	if err != nil {
		return "", err
	}
	tableName, err := cmd.Flags().GetString("table")
	if err != nil {
		return "", err
	}

	if databaseName == "" {
		return "", errors.Errorf("-database or -d param is required")
	}
	if tableName == "" {
		return "", errors.Errorf("-table or -t param is required")
	}

	result, err := h.tableInteractor.DescribeTable(ctx, input.NewDescribeTable(databaseName, tableName))
	if err != nil {
		return "", err
	}
	return result, nil
}
