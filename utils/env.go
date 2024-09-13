package utils

type Environment struct {
	DBRegion string `env:"DB_REGION,required"`
}
