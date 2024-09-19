package generator

import (
	"strconv"
	"time"
	"timestream-simple-cli/types"

	"github.com/aws/aws-sdk-go-v2/service/timestreamwrite"
	tsw "github.com/aws/aws-sdk-go-v2/service/timestreamwrite/types"
	"github.com/aws/aws-sdk-go/aws"
)

func GenerateWriteRecordsInput(
	databaseName string,
	tableName string,
	schema types.Schema,
	values []map[string]string,
) (*timestreamwrite.WriteRecordsInput, error) {

	currentTimeInSeconds := time.Now().Unix()

	var records []tsw.Record
	for _, value := range values {
		var dimensions []tsw.Dimension
		for _, dimension := range schema.Dimensions {
			dimensions = append(dimensions, tsw.Dimension{
				Name:               aws.String(dimension.Name),
				Value:              aws.String(value[dimension.Name]),
				DimensionValueType: tsw.DimensionValueType(dimension.Type),
			})
		}

		if value["Timestamp"] != "" {
			parsedTimestamp, err := time.Parse(time.RFC3339, value["Timestamp"])
			if err != nil {
				return nil, err
			}
			currentTimeInSeconds = parsedTimestamp.Unix()
		}

		if len(schema.Measures) == 1 {
			records = append(records, tsw.Record{
				Dimensions:       dimensions,
				MeasureName:      aws.String(schema.Measures[0].Name),
				MeasureValueType: tsw.MeasureValueType(schema.Measures[0].Type),
				MeasureValue:     aws.String(value[schema.Measures[0].Name]),
				Time:             aws.String(strconv.FormatInt(currentTimeInSeconds, 10)),
				TimeUnit:         tsw.TimeUnitSeconds,
			})
		} else {
			var measures []tsw.MeasureValue
			var multiMeasureName string
			for _, measure := range schema.Measures {
				measures = append(measures, tsw.MeasureValue{
					Name:  aws.String(measure.Name),
					Value: aws.String(value[measure.Name]),
					Type:  tsw.MeasureValueType(measure.Type),
				})
				multiMeasureName = multiMeasureName + measure.Name // multiMeasureName unites all measure names
			}

			records = append(records, tsw.Record{
				Dimensions:       dimensions,
				MeasureName:      aws.String(multiMeasureName),
				MeasureValueType: "MULTI",
				MeasureValues:    measures,
				Time:             aws.String(strconv.FormatInt(currentTimeInSeconds, 10)),
				TimeUnit:         tsw.TimeUnitSeconds,
			})
		}
	}

	return &timestreamwrite.WriteRecordsInput{
		DatabaseName: aws.String(databaseName),
		TableName:    aws.String(tableName),
		Records:      records,
	}, nil
}
