package usecase

import (
	"context"
	"timestream-simple-cli/usecase/input"
)

type PresetInteractor interface {
	GenerateSampleData(
		ctx context.Context,
		param *input.GenerateSampleData,
	) (string, error)
}
