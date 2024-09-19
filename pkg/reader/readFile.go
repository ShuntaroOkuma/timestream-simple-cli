package reader

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"timestream-simple-cli/types"
)

func ReadFileToValues(filePath string) ([]map[string]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var values []map[string]string
	if err := json.Unmarshal(data, &values); err != nil {
		return nil, fmt.Errorf("failed to unmarshal %s to []map[string]string: %w", filePath, err)
	}

	return values, nil
}

func ReadFileToSchema(filePath string) (types.Schema, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return types.Schema{}, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return types.Schema{}, fmt.Errorf("failed to read file: %w", err)
	}

	var schema types.Schema
	if err := json.Unmarshal(data, &schema); err != nil {
		return types.Schema{}, fmt.Errorf("failed to unmarshal %s to types.Schema: %w", filePath, err)
	}

	return schema, nil
}
