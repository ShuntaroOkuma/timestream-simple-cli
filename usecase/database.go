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
	// ListDatabases(
	// 	ctx context.Context,
	// 	param *input.ListDatabases,
	// ) (*timestreamwrite.ListDatabases, error)
	// CreateDatabase(
	// 	ctx context.Context,
	// 	param *input.CreateDatabase,
	// ) (*timestreamwrite.CreateDatabase, error)
	// UpdateDatabase(
	// 	ctx context.Context,
	// 	param *input.UpdateDatabase,
	// ) (*timestreamwrite.UpdateDatabase, error)
	// DeleteDatabase(
	// 	ctx context.Context,
	// 	param *input.DeleteDatabase,
	// ) (*timestreamwrite.DeleteDatabase, error)
}
