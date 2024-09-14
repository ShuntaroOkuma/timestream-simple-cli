package environment

type Environment struct {
	DBRegion string `env:"DB_REGION,required"`
}
