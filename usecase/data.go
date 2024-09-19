package usecase

import (
	"context"
	"timestream-simple-cli/usecase/input"
)

type DataInteractor interface {
	WriteData(
		ctx context.Context,
		param *input.WriteData,
	) (string, error)
}
