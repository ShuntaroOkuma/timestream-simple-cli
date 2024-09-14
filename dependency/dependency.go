package dependency

import (
	"context"
	"timestream-simple-cli/environment"
	"timestream-simple-cli/http"
	"timestream-simple-cli/usecase"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/timestreamquery"
	"github.com/aws/aws-sdk-go-v2/service/timestreamwrite"
)

type Dependency struct {
	DatabaseInteractor usecase.DatabaseInteractor
}

func (d *Dependency) Inject(
	ctx context.Context,
	e *environment.Environment,
) {
	httpClient := http.NewHTTPClient()

	cfg, _ := config.LoadDefaultConfig(ctx)

	writeSvc := timestreamwrite.NewFromConfig(cfg, func(o *timestreamwrite.Options) {
		o.Region = e.DBRegion
		o.HTTPClient = httpClient
	})
	querySvc := timestreamquery.NewFromConfig(cfg, func(o *timestreamquery.Options) {
		o.Region = e.DBRegion
		o.HTTPClient = httpClient
	})

	d.DatabaseInteractor = usecase.NewDatabaseInteractor(
		writeSvc,
		querySvc,
	)
}
