package dependency

import (
	"context"
	"timestream-simple-cli/environment"
	"timestream-simple-cli/http"
	"timestream-simple-cli/usecase"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/aws/aws-sdk-go-v2/service/timestreamquery"
	"github.com/aws/aws-sdk-go-v2/service/timestreamwrite"
)

type Dependency struct {
	DatabaseInteractor usecase.DatabaseInteractor
	TableInteractor    usecase.TableInteractor
	DataInteractor     usecase.DataInteractor
	PresetInteractor   usecase.PresetInteractor
}

func (d *Dependency) Inject(
	ctx context.Context,
	e *environment.Environment,
) {
	httpClient := http.NewHTTPClient()

	cfg, _ := config.LoadDefaultConfig(ctx)

	stsSvc := sts.NewFromConfig(cfg, func(o *sts.Options) {
		o.Region = e.DBRegion
		o.HTTPClient = httpClient
	})

	kmsSvc := kms.NewFromConfig(cfg, func(o *kms.Options) {
		o.Region = e.DBRegion
		o.HTTPClient = httpClient
	})

	writeSvc := timestreamwrite.NewFromConfig(cfg, func(o *timestreamwrite.Options) {
		o.Region = e.DBRegion
		o.HTTPClient = httpClient
	})
	querySvc := timestreamquery.NewFromConfig(cfg, func(o *timestreamquery.Options) {
		o.Region = e.DBRegion
		o.HTTPClient = httpClient
	})

	d.DatabaseInteractor = usecase.NewDatabaseInteractor(
		stsSvc,
		kmsSvc,
		writeSvc,
		querySvc,
	)

	d.TableInteractor = usecase.NewTableInteractor(
		writeSvc,
	)

	d.DataInteractor = usecase.NewDataInteractor(
		writeSvc,
	)

	d.PresetInteractor = usecase.NewPresetInteractor()
}
