package usecase

import (
	"context"
	"timestream-simple-cli/usecase/input"
)

type DatabaseInteractor interface {
	DescribeDatabase(
		ctx context.Context,
		param *input.DescribeDatabase,
	) (string, error)
	CreateDatabase(
		ctx context.Context,
		param *input.CreateDatabase,
	) (string, error)
}
