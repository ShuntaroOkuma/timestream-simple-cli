package data

import (
	"context"
	"fmt"
	"timestream-simple-cli/pkg/reader"
	"timestream-simple-cli/pkg/validation"
	"timestream-simple-cli/usecase"
	"timestream-simple-cli/usecase/input"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

type dataHandler struct {
	dataInteractor usecase.DataInteractor
}

func (h *dataHandler) WriteData(
	ctx context.Context,
	cmd *cobra.Command,
) (string, error) {

	// Get cmd flags
	databaseName, err := cmd.Flags().GetString("database")
	if err != nil {
		return "", err
	}
	if databaseName == "" {
		return "", errors.Errorf("-database or -d param is required")
	}

	tableName, err := cmd.Flags().GetString("table")
	if err != nil {
		return "", err
	}
	if tableName == "" {
		return "", errors.Errorf("-table or -t param is required")
	}

	schemaFilePath, err := cmd.Flags().GetString("schema-file")
	if err != nil {
		return "", err
	}
	if schemaFilePath == "" {
		return "", errors.Errorf("-schema-file or -s param is required")
	}

	valueFilePath, err := cmd.Flags().GetString("value-file")
	if err != nil {
		return "", err
	}
	if valueFilePath == "" {
		return "", errors.Errorf("-value-file or -v param is required")
	}

	// Read and convert, validate
	schema, err := reader.ReadFileToSchema(schemaFilePath)
	if err != nil {
		return "", err
	}
	if err := validation.Validate(schema); err != nil {
		return "", fmt.Errorf("failed to validation: %w", err)
	}
	if schema.Dimensions == nil {
		return "", errors.Errorf("schema dimension is empty")
	}
	if schema.Measures == nil {
		return "", errors.Errorf("schema measure is empty")
	}

	values, err := reader.ReadFileToValues(valueFilePath)
	if err != nil {
		return "", err
	}
	if values == nil {
		return "", errors.Errorf("values is empty")
	}

	// Call interactor
	result, err := h.dataInteractor.WriteData(ctx, input.NewWriteData(databaseName, tableName, schema, values))
	if err != nil {
		return "", err
	}
	return result, nil
}
