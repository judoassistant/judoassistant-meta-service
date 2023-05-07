package config

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type Environment string

const (
	EnvironmentProduction  Environment = "production"
	EnvironmentDevelopment Environment = "development"

	_environmentVarName string      = "ENV"
	_environmentDefault Environment = EnvironmentProduction
)

type Config struct {
	Environment Environment

	AdminEmail            string `mapstructure:"admin_email"`
	AdminDefaultPassword  string `mapstructure:"admin_default_password"`
	AdminDefaultFirstName string `mapstructure:"admin_default_first_name"`
	AdminDefaultLastName  string `mapstructure:"admin_default_last_name"`

	Port int      `mapstructore:"port"`
	URL  []string `mapstructore:"url"`

	DatabasePath string `mapstructure:"database_path"`
}

func getEnvironment() (Environment, error) {
	viper.SetDefault(_environmentVarName, _environmentDefault)
	viper.BindEnv(_environmentVarName)

	environment := viper.GetString(_environmentVarName)
	if environment == string(EnvironmentDevelopment) {
		return EnvironmentDevelopment, nil
	}
	if environment == string(EnvironmentProduction) {
		return EnvironmentProduction, nil
	}
	return "", errors.Errorf("The environment variable %q must be one of %q but was %q", _environmentVarName, []Environment{EnvironmentProduction, EnvironmentDevelopment}, environment)
}

func NewConfig() (*Config, error) {
	environment, err := getEnvironment()
	if err != nil {
		return nil, err
	}
	viper.SetConfigName(fmt.Sprintf("config.%s.yaml", environment))
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")

	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.Wrap(err, "unable to read config")
	}
	config := &Config{}
	if err := viper.Unmarshal(config); err != nil {
		return nil, errors.Wrap(err, "unable to unmarshal config")
	}
	config.Environment = environment

	return config, nil
}
