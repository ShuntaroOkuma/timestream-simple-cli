package input

import (
	"timestream-simple-cli/pkg/errors"
	"timestream-simple-cli/pkg/validation"
	"timestream-simple-cli/types"
)

type WriteData struct {
	DatabaseName string              `validate:"required"`
	TableName    string              `validate:"required"`
	Schema       types.Schema        `validate:"required"`
	Values       []map[string]string `validate:"required"`
}

func NewWriteData(
	databaseName string,
	tableName string,
	schema types.Schema,
	values []map[string]string,
) *WriteData {
	return &WriteData{
		DatabaseName: databaseName,
		TableName:    tableName,
		Schema:       schema,
		Values:       values,
	}
}

func (p *WriteData) Validate() error {
	if err := validation.Validate(p); err != nil {
		return err
	}

	errorList := []error{}
	dimensions := p.Schema.Dimensions
	measures := p.Schema.Measures

	// Check if multiMeasureName is included in the schema if there are two or more measures
	if len(measures) > 1 {
		if p.Schema.MultiMeasureName == "" {
			errorList = append(errorList, errors.New("multiMeasureName should be included"))
		}
	}

	// Check if the columns included in Values are only included in the columns included in Schema (dimensions and measures)
	schemaColumns := []string{}
	for _, v := range dimensions {
		schemaColumns = append(schemaColumns, v.Name)
	}
	for _, v := range measures {
		schemaColumns = append(schemaColumns, v.Name)
	}
	falseKeys := []string{}
	for _, value := range p.Values {
		for key := range value {
			// skip if key is Timestamp
			if key == "Timestamp" {
				continue
			}
			// If the key is included in falseKeys, do not check for the same key
			if isIncluded(key, falseKeys) {
				continue
			}

			if !isIncluded(key, schemaColumns) {
				falseKeys = append(falseKeys, key)
			}
		}
	}
	// Convert falseKeys to a string
	falseKeysStr := ""
	for i, v := range falseKeys {
		if i == 0 {
			falseKeysStr = v
		} else {
			falseKeysStr = falseKeysStr + ", " + v
		}
	}

	if falseKeysStr != "" {
		errorList = append(errorList, errors.Errorf("column %s in value is not included in schema", falseKeysStr))
	}

	// Check that the value of the key in Values is less than 2048 characters for the same measure
	for _, value := range p.Values {
		for _, measure := range measures {
			if len(value[measure.Name]) > 2048 {
				errorList = append(errorList, errors.Errorf("value of %s is too long", measure.Name))
			}
		}
	}

	if len(errorList) > 0 {
		return errors.Wrap(errors.AggregateErrors(errorList), "failed to validation")
	}

	return nil
}

func isIncluded(target string, list []string) bool {
	for _, v := range list {
		if target == v {
			return true
		}
	}
	return false
}
