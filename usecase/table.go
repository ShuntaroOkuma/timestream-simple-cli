package usecase

import (
	"context"
	"timestream-simple-cli/usecase/input"
)

type TableInteractor interface {
	DescribeTable(
		ctx context.Context,
		param *input.DescribeTable,
	) (string, error)
}
