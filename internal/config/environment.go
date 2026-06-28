package config

type Environment string

const (
	EnvDevelopment Environment = "development"
	EnvProduction  Environment = "production"
)

func (e Environment) IsProduction() bool {
	return e == EnvProduction
}
